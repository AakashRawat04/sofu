package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	method  string
	target  string
	version string
	headers map[string]string
	body    string
}

func readRequest(rawRequest string) *Request {
	fmt.Println("reading raw request: ", rawRequest)
	lines := strings.Split(rawRequest, "\r\n")
	if len(lines) < 1 {
		fmt.Println("Empty request")
		return &Request{}
	}

	requestLine := strings.Fields(lines[0])
	if len(requestLine) != 3 {
		fmt.Println("Bad request line")
		return &Request{}
	}

	headerMap := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			break
		}
		pair := strings.SplitN(line, ":", 2)
		if len(pair) != 2 {
			fmt.Println("Malformed header: ", line)
			continue
		}
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])
		headerMap[key] = value
	}

	return &Request{
		method:  requestLine[0],
		target:  requestLine[1],
		version: requestLine[2],
		headers: headerMap,
		body:    "",
	}
}

type Response struct {
	status  string
	headers map[string]string
	body    string
}

func NewResponse(status string) *Response {
	return &Response{
		status:  status,
		headers: make(map[string]string),
		body:    "",
	}
}

func (r *Response) SetHeader(key, value string) {
	r.headers[key] = value
}

func (r *Response) SetBody(body string) {
	r.body = body
	r.SetHeader("Content-Length", strconv.Itoa(len(body)))
}

func (r *Response) Build() string {
	var builder strings.Builder
	builder.WriteString(r.status)
	builder.WriteString("\r\n")
	for key, value := range r.headers {
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(value)
		builder.WriteString("\r\n")
	}
	builder.WriteString("\r\n")
	builder.WriteString(r.body)
	return builder.String()
}

func handleConnection(conn net.Conn) {
	fmt.Println("yoo got the connection.. or not ?: ", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	var rawRequest strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading request: ", err)
			return
		}
		rawRequest.WriteString(line)
		if strings.TrimSpace(line) == "" {
			break
		}
	}

	request := readRequest(rawRequest.String())
	fmt.Println("request constructed: ", request)

	if strings.HasPrefix(request.target, "/echo/") {
		echoStr := strings.TrimPrefix(request.target, "/echo/")
		resp := NewResponse("HTTP/1.1 200 OK")
		resp.SetHeader("Content-Type", "text/plain")
		resp.SetBody(echoStr)
		conn.Write([]byte(resp.Build()))
	} else if request.target == "/" {
		resp := NewResponse("HTTP/1.1 200 OK")
		conn.Write([]byte(resp.Build()))
	} else {
		resp := NewResponse("HTTP/1.1 404 Not Found")
		conn.Write([]byte(resp.Build()))
	}
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		fmt.Println("waiting for a connection")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("got connection: ", conn.RemoteAddr())
		fmt.Println("handling go function")
		go handleConnection(conn)
	}
}

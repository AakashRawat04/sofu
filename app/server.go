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

	// Parse request line
	requestLine := strings.Fields(lines[0])
	if len(requestLine) != 3 {
		fmt.Println("Bad request line")
		return &Request{}
	}

	// Parse headers
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
		body:    "", // No body yet
	}
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
		if strings.TrimSpace(line) == "" { // End of headers
			break
		}
	}

	request := readRequest(rawRequest.String())
	fmt.Println("request constructed: ", request)

	if request.target == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.HasPrefix(request.target, "/echo/") {
		fmt.Println("Echo endpoint hit.. returning whatever is passed!")
		echoStr := strings.TrimPrefix(request.target, "/echo/")
		contentLength := strconv.Itoa(len(echoStr))
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: " + contentLength + "\r\n" +
			"\r\n" +
			echoStr
		conn.Write([]byte(response))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
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
		// accept incoming requests
		conn, err := l.Accept()
		fmt.Println("got conenction: ", conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("handling go functoin")
		go handleConnection(conn)
	}
}

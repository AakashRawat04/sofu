// internal/handlers/handlers.go
package handlers

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/codecrafters-http-server-go/internal/models"
)

func readRequest(rawRequest string, reader *bufio.Reader) *models.Request {
	fmt.Println("reading raw request: ", rawRequest)
	lines := strings.Split(rawRequest, "\r\n")
	if len(lines) < 1 {
		fmt.Println("Empty request")
		return &models.Request{}
	}

	requestLine := strings.Fields(lines[0])
	if len(requestLine) != 3 {
		fmt.Println("Bad request line")
		return &models.Request{}
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

	body := ""
	if lengthStr, ok := headerMap["Content-Length"]; ok {
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length < 0 {
			fmt.Println("Invalid Content-Length: ", lengthStr)
			return &models.Request{} // Could return nil or a special error request
		}
		if length > 0 {
			bodyBytes := make([]byte, length)
			n, err := io.ReadFull(reader, bodyBytes)
			if err != nil {
				fmt.Println("Error reading body: ", err)
				return &models.Request{}
			}
			if n != length {
				fmt.Println("Incomplete body read: expected", length, "got", n)
				return &models.Request{}
			}
			body = string(bodyBytes)
		}
	}

	return &models.Request{
		Method:  requestLine[0],
		Target:  requestLine[1],
		Version: requestLine[2],
		Headers: headerMap,
		Body:    body,
	}
}

func Handle(conn net.Conn, directory string) {
	defer conn.Close()
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

	request := readRequest(rawRequest.String(), reader)
	fmt.Println("request constructed: ", request)

	resp := HandleRequest(request, directory)
	conn.Write([]byte(resp.Build()))
}

func HandleRequest(request *models.Request, directory string) *models.Response {
	if strings.HasPrefix(request.Target, "/echo/") {
		echoStr := strings.TrimPrefix(request.Target, "/echo/")
		resp := models.NewResponse("HTTP/1.1 200 OK")
		resp.SetHeader("Content-Type", "text/plain")
		resp.SetBody(echoStr)
		return resp
	} else if request.Target == "/user-agent" {
		resp := models.NewResponse("HTTP/1.1 200 OK")
		respBody := request.Headers["User-Agent"]
		resp.SetHeader("Content-Type", "text/plain")
		resp.SetBody(respBody)
		return resp
	} else if request.Method == "POST" && strings.HasPrefix(request.Target, "/files/") {
		filename := strings.TrimPrefix(request.Target, "/files/")
		filePath := directory + "/" + filename
		err := os.WriteFile(filePath, []byte(request.Body), 0644)
		if err != nil {
			fmt.Println("Error writing file: ", err)
			return models.NewResponse("HTTP/1.1 404 Not Found")
		}
		return models.NewResponse("HTTP/1.1 201 Created")
	} else if strings.HasPrefix(request.Target, "/files/") {
		filename := strings.TrimPrefix(request.Target, "/files/")
		filePath := directory + "/" + filename
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("File not found: ", err)
			return models.NewResponse("HTTP/1.1 404 Not Found")
		}
		resp := models.NewResponse("HTTP/1.1 200 OK")
		contentType := "application/octet-stream"
		resp.SetHeader("Content-Type", contentType)
		resp.SetBody(string(data))
		return resp
	} else if request.Target == "/" {
		return models.NewResponse("HTTP/1.1 200 OK")
	}
	return models.NewResponse("HTTP/1.1 404 Not Found")
}

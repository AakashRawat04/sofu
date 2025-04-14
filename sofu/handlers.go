package sofu

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HandlerFunc func(*Context)

func readRequest(c *Context, reader *bufio.Reader) {
	var rawRequest strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading request:", err)
			return
		}
		rawRequest.WriteString(line)
		if strings.TrimSpace(line) == "" {
			break
		}
	}

	lines := strings.Split(rawRequest.String(), "\r\n")
	if len(lines) < 1 {
		return
	}
	requestLine := strings.Fields(lines[0])
	if len(requestLine) != 3 {
		return
	}

	c.Request.Method = requestLine[0]
	c.Request.Path = requestLine[1]
	c.Request.Version = requestLine[2]

	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			break
		}
		pair := strings.SplitN(line, ":", 2)
		if len(pair) == 2 {
			c.Request.Headers[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
		}
	}

	if lengthStr, ok := c.Request.Headers["Content-Length"]; ok {
		length, err := strconv.Atoi(lengthStr)
		if err == nil && length > 0 {
			bodyBytes := make([]byte, length)
			_, err := io.ReadFull(reader, bodyBytes)
			if err == nil {
				c.Request.Body = string(bodyBytes)
			}
		}
	}
}

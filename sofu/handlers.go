package sofu

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HandlerFunc func(*Context)

func readRequest(c *Context, reader *bufio.Reader) error {
	// Read the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return err // Connection closed
		}
		fmt.Println("Error reading request line:", err)
		return err
	}

	// Parse the request line
	parts := strings.Fields(strings.TrimSpace(requestLine))
	if len(parts) != 3 {
		return errors.New("invalid request line")
	}

	c.Request.Method = parts[0]
	c.Request.Path = parts[1]
	c.Request.Version = parts[2]

	// Read headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break // End of headers
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue // Invalid header
		}

		headerName := strings.TrimSpace(parts[0])
		headerValue := strings.TrimSpace(parts[1])
		c.Request.Headers[headerName] = headerValue
	}

	// Read body if Content-Length is present
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

	return nil
}

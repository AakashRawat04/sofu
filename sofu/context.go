package sofu

import (
	"bufio"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/codecrafters-http-server-go/sofu/compressions"
)

type Context struct {
	Conn    net.Conn
	Request *Request
	writer  *bufio.Writer
	headers map[string]string // Store custom headers
}

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
	Params  map[string]string
}

func NewContext(conn net.Conn) *Context {
	return &Context{
		Conn:    conn,
		Request: &Request{Headers: make(map[string]string), Params: make(map[string]string)},
		writer:  bufio.NewWriter(conn),
		headers: make(map[string]string),
	}
}

// SetHeader adds or updates a header in the response
func (c *Context) SetHeader(key, value string) {
	c.headers[key] = value
}

// String sends a response with the given status code and body
func (c *Context) String(status int, body string) {
	statusLine := "HTTP/1.1 " + strconv.Itoa(status) + " " + statusText(status) + "\r\n"

	// Set default headers if not already set
	if _, ok := c.headers[HeaderContentType]; !ok {
		c.headers[HeaderContentType] = ContentTypeTextPlain
	}

	if _, ok := c.headers[HeaderContentLength]; !ok {
		c.headers[HeaderContentLength] = strconv.Itoa(len(body))
	}

	var responseBody = body
	if acceptEncoding, ok := c.Request.Headers[HeaderAcceptEncoding]; ok {
		compressedBody, encoding := compressions.HandleCompression(acceptEncoding, body)

		// Only set Content-Encoding if compression was applied
		if encoding != "" {
			responseBody = compressedBody
			c.SetHeader(HeaderContentLength, strconv.Itoa(len(compressedBody)))
			c.SetHeader(HeaderContentEncoding, encoding)
		}
	}

	// Build headers
	var headerStr strings.Builder
	for key, value := range c.headers {
		headerStr.WriteString(key)
		headerStr.WriteString(": ")
		headerStr.WriteString(value)
		headerStr.WriteString("\r\n")
	}
	headerStr.WriteString("\r\n")

	// Write response
	c.writer.WriteString(statusLine)
	c.writer.WriteString(headerStr.String())
	c.writer.WriteString(responseBody)
	c.writer.Flush()
}

func (c *Context) Param(key string) string {
	return c.Request.Params[key]
}

func statusText(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 404:
		return "Not Found"
	case 400:
		return "Bad Request"
	default:
		return "Unknown"
	}
}

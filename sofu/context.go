package sofu

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/AakashRawat04/sofu/sofu/compressions"
)

type Context struct {
	Conn    net.Conn
	Request *Request
	headers map[string]string // Store response headers
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
		headers: make(map[string]string),
	}
}

// SetHeader adds or updates a header in the response
func (c *Context) SetHeader(key, value string) {
	c.headers[key] = value
}

// WriteResponse sends an HTTP response with the given status code and body
func (c *Context) WriteResponse(statusCode int, body string) {
	statusMessage, ok := StatusText[statusCode]
	if !ok {
		statusMessage = "Unknown"
	}

	if _, ok := c.headers[HeaderContentLength]; !ok {
		c.headers[HeaderContentLength] = strconv.Itoa(len(body))
	}

	if _, ok := c.headers[HeaderContentType]; !ok {
		c.headers[HeaderContentType] = ContentTypeTextPlain
	}

	// handle compression
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

	// Check if client requested connection close
	if closeConn, ok := c.Request.Headers["Connection"]; ok && strings.ToLower(closeConn) == "close" {
		c.headers["Connection"] = "close"
	}

	response := fmt.Sprintf("%s %d %s\r\n", HTTPVersion1_1, statusCode, statusMessage)

	for key, value := range c.headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	response += "\r\n" + responseBody

	c.Conn.Write([]byte(response))
}

// Helper method for common status OK responses
func (c *Context) String(statusCode int, body string) {
	c.WriteResponse(statusCode, body)
}

func (c *Context) ShouldCloseConnection() bool {
	// Check HTTP version - HTTP/1.0 defaults to non-persistent
	if c.Request.Version == HTTPVersion1_0 {
		// But Connection: keep-alive can override this
		if keepAlive, ok := c.Request.Headers["Connection"]; ok && strings.ToLower(keepAlive) == "keep-alive" {
			return false
		}
		return true
	}

	// HTTP/1.1 defaults to persistent connections
	// But Connection: close can override this
	if closeConn, ok := c.Request.Headers["Connection"]; ok && strings.ToLower(closeConn) == "close" {
		return true
	}

	return false
}

// Helper method to get a parameter from the request
func (c *Context) Param(key string) string {
	return c.Request.Params[key]
}

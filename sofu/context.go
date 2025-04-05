package sofu

import (
	"bufio"
	"net"
	"strconv"
)

type Context struct {
	Conn    net.Conn
	Request *Request
	writer  *bufio.Writer
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
	}
}

func (c *Context) String(status int, body string) {
	statusLine := "HTTP/1.1 " + strconv.Itoa(status) + " " + statusText(status) + "\r\n"
	headers := "Content-Type: text/plain\r\nContent-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n"
	c.writer.WriteString(statusLine + headers + body)
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

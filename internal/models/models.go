// internal/models/models.go
package models

import (
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Target  string
	Version string
	Headers map[string]string
	Body    string
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

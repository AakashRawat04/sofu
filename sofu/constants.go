package sofu

// Constants for HTTP status codes
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
)

// Constants for HTTP methods
const (
	MethodGET     = "GET"
	MethodPOST    = "POST"
	MethodPUT     = "PUT"
	MethodDELETE  = "DELETE"
	MethodOPTIONS = "OPTIONS"
)

// Constants for common headers
const (
	HeaderContentType     = "Content-Type"
	HeaderContentLength   = "Content-Length"
	HeaderUserAgent       = "User-Agent"
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"
	HeaderConnection      = "Connection"
)

// Constants for common Content Types
const (
	ContentTypeTextPlain              = "text/plain"
	ContentTypeApplicationJSON        = "application/json"
	ContentTypeTextHTML               = "text/html"
	ContentTypeApplicationOctetStream = "application/octet-stream"
)

// Map of status codes to their respective messages
var StatusText = map[int]string{
	StatusOK:                  "OK",
	StatusCreated:             "Created",
	StatusNoContent:           "No Content",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not Found",
	StatusMethodNotAllowed:    "Method Not Allowed",
	StatusInternalServerError: "Internal Server Error",
	StatusNotImplemented:      "Not Implemented",
	StatusBadGateway:          "Bad Gateway",
	StatusServiceUnavailable:  "Service Unavailable",
}

// Constants for common HTTP version
const (
	HTTPVersion1_0 = "HTTP/1.0"
	HTTPVersion1_1 = "HTTP/1.1"
)

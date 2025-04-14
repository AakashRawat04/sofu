package sofu

// Constants for HTTP status codes
const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusInternalServerError = 500
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
)

// Constants for common Content Types
const (
	ContentTypeTextPlain              = "text/plain"
	ContentTypeApplicationJSON        = "application/json"
	ContentTypeTextHTML               = "text/html"
	ContentTypeApplicationOctetStream = "application/octet-stream"
)

// Constants for common compression schemes
const (
	CompressionGzip    = "gzip"
	CompressionDeflate = "deflate"
)

// Constants for common HTTP status messages
const (
	StatusMessageOK                  = "OK"
	StatusMessageBadRequest          = "Bad Request"
	StatusMessageNotFound            = "Not Found"
	StatusMessageInternalServerError = "Internal Server Error"
)

// Constants for common HTTP version
const (
	HTTPVersion1_0 = "HTTP/1.0"
	HTTPVersion1_1 = "HTTP/1.1"
)

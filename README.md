[![progress-banner](https://backend.codecrafters.io/progress/http-server/d149aba3-35f0-4522-b661-07f058fd7808)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

![image](https://github.com/user-attachments/assets/3710a962-564f-4208-8400-0bbed6b8f665)


# Sofu HTTP Server Framework

Sofu is a lightweight, flexible HTTP server framework for Go that makes it easy to build web applications and APIs. Built as part of the CodeCrafters HTTP server challenge, Sofu provides a clean, intuitive interface for handling HTTP requests.

## Features

- Simple and intuitive routing system with path parameter support
- HTTP method-based route handling (`GET`, `POST`, etc.)
- Context-based request handling
- Built-in response helpers
- Support for HTTP compression (content negotiation)
- File serving capabilities
- Persistent connection support (keep-alive)
- Easy setup with minimal configuration

## Installation

```bash
go get github.com/AakashRawat04/sofu
```

## Quick Start

```go
package main

import "github.com/AakashRawat04/sofu/sofu"

func main() {
    // Create a new server instance
    server := sofu.New()
    
    // Register routes
    server.GET("/", func(c *sofu.Context) {
        c.WriteResponse(sofu.StatusOK, "Hello, World!")
    })
    
    // Start the server
    server.Start(":4221")
}
```

## Usage Examples

### Basic Routes

```go
// Define routes and their handlers
server.GET("/", func(c *sofu.Context) {
    c.WriteResponse(sofu.StatusOK, "Welcome to Sofu!")
})

server.POST("/submit", func(c *sofu.Context) {
    // Handle POST request
    c.WriteResponse(sofu.StatusCreated, "Resource created")
})
```

### URL Parameters

```go
// Route with a parameter
server.GET("/echo/:message", func(c *sofu.Context) {
    // Extract the parameter value
    message := c.Param("message")
    c.WriteResponse(sofu.StatusOK, message)
})
```

### Access Request Headers

```go
server.GET("/user-agent", func(c *sofu.Context) {
    userAgent := c.Request.Headers["User-Agent"]
    c.WriteResponse(sofu.StatusOK, userAgent)
})
```

### Serving Files

```go
server.GET("/files/:filename", func(c *sofu.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join(server.Directory, filename)

    content, err := os.ReadFile(filePath)
    if err != nil {
        c.WriteResponse(sofu.StatusNotFound, "File not found")
        return
    }

    contentType := sofu.ContentTypeApplicationOctetStream
    if strings.HasSuffix(filename, ".txt") {
        contentType = sofu.ContentTypeTextPlain
    } else if strings.HasSuffix(filename, ".html") {
        contentType = sofu.ContentTypeTextHTML
    }

    c.SetHeader(sofu.HeaderContentType, contentType)
    c.WriteResponse(sofu.StatusOK, string(content))
})
```

### File Upload

```go
server.POST("/files/:filename", func(c *sofu.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join(server.Directory, filename)

    err := os.WriteFile(filePath, []byte(c.Request.Body), 0644)
    if err != nil {
        c.WriteResponse(sofu.StatusInternalServerError, "Failed to save file")
        return
    }

    c.WriteResponse(sofu.StatusCreated, "")
})
```

## Command Line Options

The server accepts command line flags:

- `-directory`: Specify the directory for serving files (default: "files")

Example:
```bash
./server -directory /path/to/files
```

## API Reference

### Server

- `New() *Server`: Creates a new server instance
- `Start(addr string)`: Starts the server on the given address
- `GET(path string, handler HandlerFunc)`: Registers a GET route
- `POST(path string, handler HandlerFunc)`: Registers a POST route

### Context

- `WriteResponse(statusCode int, body string)`: Writes a response with the given status code and body
- `String(statusCode int, body string)`: Shorthand for WriteResponse
- `SetHeader(key, value string)`: Sets a response header
- `Param(key string)`: Gets a URL parameter value

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

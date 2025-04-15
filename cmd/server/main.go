package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/codecrafters-http-server-go/sofu"
)

func main() {
	server := sofu.New()

	// Register routes
	server.GET("/", func(c *sofu.Context) {
		c.WriteResponse(sofu.StatusOK, "")
	})

	server.GET("/user-agent", func(c *sofu.Context) {
		userAgent := c.Request.Headers["User-Agent"]
		c.WriteResponse(sofu.StatusOK, userAgent)
	})

	server.GET("/echo/:message", func(c *sofu.Context) {
		message := c.Param("message")
		c.WriteResponse(sofu.StatusOK, message)
	})

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
		c.SetHeader(sofu.HeaderContentLength, fmt.Sprintf("%d", len(content)))
		c.WriteResponse(sofu.StatusOK, string(content))
	})

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

	// Start the server
	server.Start(":4221")
}

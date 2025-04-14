package main

import (
	"os"

	"github.com/codecrafters-io/codecrafters-http-server-go/sofu"
)

func main() {
	server := sofu.New()

	server.GET("/", func(c *sofu.Context) {
		c.String(200, "")
	})

	server.GET("/echo/:message", func(c *sofu.Context) {
		message := c.Param("message")
		c.SetHeader("Content-Type", "text/plain")
		c.String(200, message)
	})

	server.GET("/user-agent", func(c *sofu.Context) {
		userAgent := c.Request.Headers["User-Agent"]
		if userAgent == "" {
			c.String(400, "User-Agent header missing")
		} else {
			c.SetHeader("Content-Type", "text/plain")
			c.String(200, userAgent)
		}
	})

	server.GET("/files/:filename", func(c *sofu.Context) {
		filename := c.Param("filename")
		data, err := os.ReadFile(server.Directory + "/" + filename)
		if err != nil {
			c.String(404, "Not Found")
		} else {
			c.SetHeader("Content-Type", "application/octet-stream")
			c.String(200, string(data))
		}
	})

	server.POST("/files/:filename", func(c *sofu.Context) {
		filename := c.Param("filename")
		data := c.Request.Body
		err := os.WriteFile(server.Directory+"/"+filename, []byte(data), 0644)
		if err != nil {
			c.String(500, "Internal Server Error")
		} else {
			c.SetHeader("Content-Type", "text/plain")
			c.String(201, "Created")
		}
	})

	server.Start("0.0.0.0:4221")
}

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
		c.String(200, message)
	})

	server.GET("/user-agent", func(c *sofu.Context) {
		userAgent := c.Request.Headers["User-Agent"]
		if userAgent == "" {
			c.String(400, "User-Agent header missing")
		} else {
			c.String(200, userAgent)
		}
	})

	server.GET("/files/:filename", func(c *sofu.Context) {
		filename := c.Param("filename")
		data, err := os.ReadFile(server.Directory + "/" + filename)
		if err != nil {
			c.String(404, "Not Found")
		} else {
			c.String(200, string(data))
		}
	})

	server.Start("0.0.0.0:4221")
}

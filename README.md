[![progress-banner](https://backend.codecrafters.io/progress/http-server/d149aba3-35f0-4522-b661-07f058fd7808)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# HTTP Server Framework

A lightweight HTTP framework built for the CodeCrafters HTTP Server challenge.

## Usage
```go
package main

import "github.com/codecrafters-io/codecrafters-http-server-go/framework"

func main() {
    server := framework.New()
    server.GET("/", func(c *framework.Context) {
        c.String(200, "Hello, World!")
    })
    server.Start(":8080")
}
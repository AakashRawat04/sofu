// internal/server/server.go
package server

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/codecrafters-http-server-go/internal/handlers"
)

func Run() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	directory := flag.String("directory", "servefiles", "files to serve")
	flag.Parse()
	fmt.Println("directory for serving files set to: ", *directory)

	for {
		fmt.Println("waiting for a connection")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("got connection: ", conn.RemoteAddr())
		fmt.Println("handling go function")
		go handlers.Handle(conn, *directory)
	}
}

package sofu

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type Server struct {
	Router    *Router
	Directory string
}

func New() *Server {
	directory := flag.String("directory", "files", "directory for serving files")
	flag.Parse()
	return &Server{
		Router:    NewRouter(),
		Directory: *directory,
	}
}

func (s *Server) Start(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to bind to port:", addr)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Server started on", addr, "serving files from", s.Directory)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		c := NewContext(conn)

		// Read and parse the HTTP request
		if err := readRequest(c, reader); err != nil {
			// Connection closed or error reading request
			break
		}

		// Handle the request
		s.Router.Handle(c)

		// Check if we should close the connection
		if c.ShouldCloseConnection() {
			break
		}
	}
}

func (s *Server) GET(path string, handler HandlerFunc) {
	s.Router.GET(path, handler)
}

func (s *Server) POST(path string, handler HandlerFunc) {
	s.Router.POST(path, handler)
}

func (s *Server) PUT(path string, handler HandlerFunc) {
	s.Router.PUT(path, handler)
}

func (s *Server) DELETE(path string, handler HandlerFunc) {
	s.Router.DELETE(path, handler)
}

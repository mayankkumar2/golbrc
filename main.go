package main

import (
	"github.com/mayankkumar2/golbrc/pkg/proxy"
	"github.com/mayankkumar2/golbrc/pkg/server"
	"log"
	"os"
)



func main() {
	_logger := server.NewLogger(os.Stdout)

	s := server.NewServer(proxy.ProxyMiddleware, _logger)

	if err := s.StartServer(server.Options{
		TLSEnabled: false,
		Address: "0.0.0.0",
		Port: "8080",
	}); err != nil {
		log.Fatalln(err)
	}
}
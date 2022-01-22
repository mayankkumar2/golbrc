package server

import "github.com/valyala/fasthttp"

type Options struct {
	Port string
	Address string
	TLSEnabled bool
	CertFile string
	KeyFile string
}

type Server interface {
	StartServer(options Options) error
}


type HTTPServer struct {
	server *fasthttp.Server
}

func NewServer(handler fasthttp.RequestHandler, logger fasthttp.Logger) Server {
	return &HTTPServer{
		server: &fasthttp.Server{
			Handler:           handler,
			TCPKeepalive:      true,
			ReduceMemoryUsage: true,
			LogAllErrors:      true,
			StreamRequestBody: true,
			Logger:            logger,
		},
	}
}

func (h *HTTPServer) StartServer(options Options) error {
	if options.TLSEnabled {
		err := h.server.ListenAndServeTLS(options.Address+":"+options.Port, options.CertFile, options.KeyFile)
		return err
	} else {
		err := h.server.ListenAndServe(options.Address+ ":" +options.Port)
		return err
	}
}
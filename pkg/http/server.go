package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// Server is the HTTP Server
type Server struct {
	router *http.ServeMux
}

// NewServer
func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
	}
	s.registerHandlers()
	return s
}

// Serve
func (s *Server) Serve(ln net.Listener) error {

	srv := &http.Server{
		Handler: s.router,
	}

	if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) registerHandlers() {

	// TODO: Setup Proxy
	// TODO: Setup Handlers

	s.router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		url := req.URL.Query().Get("url")

		if url != "" {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Fprintf(w, string(body))
		}
	})
}

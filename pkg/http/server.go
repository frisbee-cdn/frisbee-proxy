package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/frisbee-cdn/frisbee-proxy/pkg/cache"
)

// Server is the HTTP Server
type Server struct {
	router    *http.ServeMux
	datastore *cache.MapDataStore
}

// NewServer
func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
		// TODO: Provide configuration structure
		datastore: cache.NewMapDataStore(50),
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

		body, err := s.getFromOrigin(url)
		if err != nil {
			log.Fatalln(err)
		}

		s.datastore.Put(url, body)

		r, _ := s.datastore.Get(url)
		fmt.Fprintf(w, string(r))
	})
}

func (s *Server) getFromOrigin(url string) ([]byte, error) {

	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	return nil, nil
}

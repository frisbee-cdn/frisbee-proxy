package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/frisbee-cdn/frisbee-proxy/pkg/proxy"
)

func (s *Server) handleURLSearch(w http.ResponseWriter, req *http.Request) {

	// TODO: Setup Proxy

	np := proxy.NewProxy()

	if err := np.Connect(context.Background(), ":5002"); err != nil {
		log.Fatalln(err)
	}
	defer np.Close()

	url := req.URL.Query().Get("url")

	r, err := s.datastore.Get(url)

	if err != nil {

	}

	// body, err := s.getFromOrigin(url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// s.datastore.Put(url, body)

	fmt.Fprintf(w, string(r))
}

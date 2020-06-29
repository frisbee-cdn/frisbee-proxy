package main

import (
	"fmt"
	"net"
	"os"

	"github.com/frisbee-cdn/frisbee-proxy/pkg/http"
)

// func serveReverseProy(targetAddress string, res http.ResponseWriter, req *http.Request) {

// 	url, _ := url.Parse(targetAddress)

// 	proxy := httputil.NewSingleHostReverseProxy(url)

// 	req.URL.Host = url.Host
// 	req.URL.Scheme = url.Scheme
// 	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
// 	req.Host = url.Host

// 	proxy.ServeHTTP(res, req)
// }

// func handleRequestAndRedirect(rest http.ResponseWriter, req *http.Request) {

// 	// requestPayload := pkg.ParseRequestBody(req)

// }

func main() {
	s := http.NewServer()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 6001))
	if err != nil {
		fmt.Printf("Error Failted to listen HTTP port %s\n", err)
		os.Exit(1)
	}
	s.Serve(ln)
}

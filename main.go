package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/frisbee-cdn/frisbee-proxy/pkg/parser"
)

func serveReverseProy(tarhetAddres string, res http.ResponseWriter, req *http.Request) {

	url, _ := url.Parse(targetAddress)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(rest http.ResponseWriter, req *http.Request) {

	requestPayload := parser.ParseRequestBody(req)
}

func main() {

	log.Info("This is an information to be displayed")

	router := mux.NewRouter()

	router.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":"+"9090", router); err != nil {
		panic(err)
	}
}

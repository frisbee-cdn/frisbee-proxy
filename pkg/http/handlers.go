package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/frisbee-cdn/frisbee-proxy/pkg/grpc"
	proxy "github.com/frisbee-cdn/frisbee-proxy/pkg/proxy"
	"log"
	"net/http"
)

func (s *Server) handleURLSearch(w http.ResponseWriter, req *http.Request) {


	np := proxy.NewProxy()

	if err := np.Connect(context.Background(), ":5002"); err != nil {
		log.Fatalln(err)
	}
	defer np.Close()

	url := req.URL.Query().Get("url")

	r, err := s.datastore.Get(url)
	if err != nil{

		grpcReq := grpc.FindValue{url}
		b, err := json.Marshal(grpcReq)
		if err != nil{
			log.Fatalf("Proxy Error: %s", err)
		}
		grpcRes, err := np.Call(context.Background(), "models.FrisbeeProtocol", "FindValueProxy", b)

		if err != nil {
			log.Fatal("Proxy Error: %s", err)
		}
		revealRes, _ := grpcRes.MarshalJSON()
		var reply grpc.FindValueReply
		_ = json.Unmarshal(revealRes, &reply)

		if reply.Value == nil{
			log.Printf("No content for key: %s obtain from origin", url)
			originContent, err := s.getFromOrigin(url)
			if err != nil{
				log.Fatalf("Error fetching from origin: %s", err)
			}
			_ = s.datastore.Put(url, originContent)
			storeReq := grpc.Store{
				Key: url,
				Value: originContent,
			}
			b, err := json.Marshal(storeReq)
			if err != nil{
				log.Fatalf("Proxy Error: %s", err)
			}
			storeRes, err := np.Call(context.Background(), "models.FrisbeeProtocol", "StoreProxy", b)
			if err != nil {
				log.Fatalf("Proxy Error: %s", err)
			}

			revealRes, _ := storeRes.MarshalJSON()
			var reply grpc.Error
			_ = json.Unmarshal(revealRes, &reply)
			log.Print("Stored content in cache")
			r = originContent
		}else{
			log.Print("Received value from peer, storing in cache")
			_ = s.datastore.Put(url, reply.Value)
			r = reply.Value
		}

		fmt.Fprintf(w, string(r))
	}


	fmt.Fprintf(w, string(r))
}



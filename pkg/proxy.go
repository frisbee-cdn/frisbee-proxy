package pkg

import (
	"context"
	"net/url"

	ref "github.com/frisbee-cdn/frisbee-proxy/pkg/grpc"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// Proxy
type Proxy struct {
	cc        *grpc.ClientConn
	reflector *ref.Reflector
}

// NewProxy
func NewProxy() *Proxy {
	return &Proxy{}
}

// Connect
func (p *Proxy) Connect(ctx context.Context, target *url.URL) error {

	cc, err := grpc.DialContext(ctx, target.String(), grpc.WithInsecure())

	if err != nil {
		return err
	}

	p.cc = cc

	rc := grpcreflect.NewClient(ctx, rpb.NewServerReflectionClient(p.cc))
	p.reflector = ref.NewReflector(rc)
	return nil
}

// Close
func (p *Proxy) Close() error {
	return p.cc.Close()
}

// Call
func (p *Proxy) Call(ctx context.Context, service, method string, message []byte) ([]byte, error) {

	return nil, nil
}

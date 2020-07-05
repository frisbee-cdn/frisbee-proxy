package pkg

import (
	"context"

	ref "github.com/frisbee-cdn/frisbee-proxy/pkg/grpc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// Proxy describes the Proxy
type Proxy struct {
	cc        *grpc.ClientConn
	reflector *ref.Reflector
	stub      grpcdynamic.Stub
}

// NewProxy creates a new Proxy
func NewProxy() *Proxy {
	return &Proxy{}
}

// Connect establishes an new connection with the RPC server
func (p *Proxy) Connect(ctx context.Context, addr string) error {

	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())

	if err != nil {
		return err
	}
	p.cc = cc
	rc := grpcreflect.NewClient(ctx, rpb.NewServerReflectionClient(p.cc))
	p.reflector = ref.NewReflector(rc)
	p.stub = grpcdynamic.NewStub(p.cc)
	return nil
}

// Close closes the RPC connection
func (p *Proxy) Close() error {
	return p.cc.Close()
}

// Call perfrms the gRPC call after reflection to obtain the descriptors
func (p *Proxy) Call(ctx context.Context, service, method string, message []byte) (ref.Message, error) {

	invocation, err := p.reflector.CreateInvocation(ctx, service, method, message)
	if err != nil {
		return nil, err
	}

	o, err := p.stub.InvokeRpc(ctx,
		invocation.MethodDescriptor.AsProtoreflectDescriptor(),
		invocation.Message.AsProtoreflectMessage())
	if err != nil {
		return nil, err
	}

	outputMsg := invocation.MethodDescriptor.GetOutputType().NewMessage()
	err = outputMsg.ConvertFrom(o)
	if err != nil {
		return nil, err
	}
	return outputMsg, nil
}

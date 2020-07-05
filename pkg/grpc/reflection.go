package grpc

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
)

// Reflector
type Reflector struct {
	reflectClient *grpcreflect.Client
}

// NewReflector
func NewReflector(rc *grpcreflect.Client) *Reflector {
	return &Reflector{reflectClient: rc}
}

// CreateInvocation
func (r *Reflector) CreateInvocation(ctx context.Context, serviceName, methodName string, input []byte) (*MethodInvocation, error) {

	s, err := r.reflectClient.ResolveService(serviceName)
	if err != nil {
		return nil, err
	}
	methodDesc, err := r.findMethodByName(methodName, s)
	if err != nil {
		return nil, err
	}

	inputMsg := methodDesc.GetInputType().NewMessage()
	err = inputMsg.UnmarshalJSON(input)
	if err != nil {
		return nil, err
	}

	return &MethodInvocation{
		MethodDescriptor: methodDesc,
		Message:          inputMsg,
	}, nil
}

func (r *Reflector) findMethodByName(name string, service *desc.ServiceDescriptor) (*MethodDescriptor, error) {

	md := service.FindMethodByName(name)

	if md == nil {
		return nil, nil
	}

	return &MethodDescriptor{
		MethodDescriptor: md,
	}, nil
}

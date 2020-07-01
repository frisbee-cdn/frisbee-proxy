package grpc

import (
	"context"

	"github.com/jhump/protoreflect/grpcreflect"
)

// Reflector
type Reflector struct {
	reflectClient *grpcreflect.Client
}

// NewReflector
func NewReflector(rc *grpcreflect.Client) Reflector {
	return Reflector{reflectClient: rc}
}

// CreateInvocation
func (r *Reflector) CreateInvocation(ctx context.Context,
	serviceName, methodName string, input []byte) (*MethodInvocation, error) {

	d, err := r.reflectClient.ResolveService(serviceName)
	if err != nil {
		return nil, err
	}
	methodDesc, err := d.FindMethodByName(methodName)
	if err != nil {
		return nil, err
	}

	inputMsg := methodDesc.GetInputType()
	err = inputMsg.UnmarshalJSON(input)
	if err != nil {
		return nil, err
	}

	return &MethodInvocation{
		MethodDescriptor: methodDesc,
		Message:          inputMsg,
	}, nil
}

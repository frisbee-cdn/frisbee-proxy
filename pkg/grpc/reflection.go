package grpc

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
)

// Reflector
type Reflector struct {
	rc *grpcreflect.Client
}

// NewReflector
func NewReflector(reflectClient *grpcreflect.Client) *Reflector {

	return &Reflector{
		rc: reflectClient,
	}
}

// CreateInvocation
func (r *Reflector) CreateInvocation(ctx context.Context,
	serviceName, methodName string, input []byte) (*MethodInvocation, error) {

	servDesc, err := r.resolveService(serviceName)
	if err != nil {
		return nil, err
	}

	methodDesc, err := servDesc.FindMethodByName(methodName)
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

func (r *Reflector) resolveService(serviceName string) (*ServiceDescriptor, error) {

	d, err := r.rc.ResolveService(serviceName)
	if err != nil {
		return nil, err
		// TODO: Build custom error ProxyError
	}

	return &ServiceDescriptor{
		ServiceDescriptor: d,
	}, nil
}

// ServiceDescriptor
type ServiceDescriptor struct {
	*desc.ServiceDescriptor
}

// FindMethodByName
func (s *ServiceDescriptor) FindMethodByName(methodName string) (*MethodDescriptor, error) {

	m := s.ServiceDescriptor.FindMethodByName(methodName)
	if m == nil {
		return nil, nil
	}
	return &MethodDescriptor{
		MethodDescriptor: m,
	}, nil

}

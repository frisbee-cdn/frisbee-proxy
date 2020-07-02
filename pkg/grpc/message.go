package grpc

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

// MethodInvocation contains the MethodDescriptor and the Message used to invoke an RPC
type MethodInvocation struct {
	*MethodDescriptor
	Message
}

// Message defines general methods to map JSON to grpc
type Message interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error

	ConvertFrom(target proto.Message) error
	AsProtoreflectMessage() *dynamic.Message
}

// MethodDescriptor represents the method type
type MethodDescriptor struct {
	*desc.MethodDescriptor
}

// GetInputType gets the MessageDescriptor for the method input type
func (m *MethodDescriptor) GetInputType() *MessageDescriptor {

	return &MessageDescriptor{
		desc: m.MethodDescriptor.GetInputType(),
	}
}

// GetOutputType  gets the MessageDescriptor for the method output type
func (m *MethodDescriptor) GetOutputType() *MessageDescriptor {
	return &MessageDescriptor{
		desc: m.MethodDescriptor.GetOutputType(),
	}
}

// AsProtoreflectDescriptor
func (m *MethodDescriptor) AsProtoreflectDescriptor() *desc.MethodDescriptor {
	return m.MethodDescriptor
}

// MessageDescriptor represents the message type
type MessageDescriptor struct {
	desc *desc.MessageDescriptor
}

// NewMessage
func (m *MessageDescriptor) NewMessage() *MessageImpl {

	return &MessageImpl{
		Message: dynamic.NewMessage(m.desc),
	}
}

type MessageImpl struct {
	*dynamic.Message
}

// MarshalJSON
func (m *MessageImpl) MarshalJSON() ([]byte, error) {

	b, err := m.Message.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// UnmarshalJSON
func (m *MessageImpl) UnmarshalJSON(b []byte) error {

	if err := m.Message.UnmarshalJSON(b); err != nil {
		return err
	}
	return nil
}

// ConvertFrom
func (m *MessageImpl) ConvertFrom(target proto.Message) error {
	return m.Message.ConvertFrom(target)
}

// AsProtoreflectMessage
func (m *MessageImpl) AsProtoreflectMessage() *dynamic.Message {
	return m.Message
}

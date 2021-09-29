package main

import "github.com/golang/protobuf/proto"

...
package helloworld
...

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset() { *m = HelloRequest{} }
func (m *HelloRequest) String() { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage() {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
...


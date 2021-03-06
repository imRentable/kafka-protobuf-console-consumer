// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	sample.proto

It has these top-level messages:
	SampleMessage
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SampleMessage struct {
	StringField  string `protobuf:"bytes,1,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	IntegerField int32  `protobuf:"varint,2,opt,name=integer_field,json=integerField" json:"integer_field,omitempty"`
}

func (m *SampleMessage) Reset()                    { *m = SampleMessage{} }
func (m *SampleMessage) String() string            { return proto.CompactTextString(m) }
func (*SampleMessage) ProtoMessage()               {}
func (*SampleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SampleMessage) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

func (m *SampleMessage) GetIntegerField() int32 {
	if m != nil {
		return m.IntegerField
	}
	return 0
}

func init() {
	proto.RegisterType((*SampleMessage)(nil), "sample_package.SampleMessage")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xf0, 0xe2, 0x0b, 0x12, 0x93,
	0xb3, 0x13, 0xd3, 0x53, 0x95, 0xc2, 0xb9, 0x78, 0x83, 0xc1, 0x22, 0xbe, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0x42, 0x8a, 0x5c, 0x3c, 0xc5, 0x25, 0x45, 0x99, 0x79, 0xe9, 0xf1, 0x69, 0x99, 0xa9,
	0x39, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x10, 0x31, 0x37, 0x90, 0x90, 0x90,
	0x32, 0x17, 0x6f, 0x66, 0x5e, 0x49, 0x6a, 0x7a, 0x6a, 0x11, 0x54, 0x0d, 0x93, 0x02, 0xa3, 0x06,
	0x6b, 0x10, 0x0f, 0x54, 0x10, 0xac, 0x28, 0x89, 0x0d, 0x6c, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x2d, 0xec, 0xf7, 0x7f, 0x00, 0x00, 0x00,
}

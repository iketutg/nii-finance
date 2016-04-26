// Code generated by protoc-gen-go.
// source: github.com/nii236/nii-finance/tickSubscriber/proto/tickSubscriber.proto
// DO NOT EDIT!

/*
Package tickSubscriber is a generated protocol buffer package.

It is generated from these files:
	github.com/nii236/nii-finance/tickSubscriber/proto/tickSubscriber.proto

It has these top-level messages:
	Tick
*/
package tickSubscriber

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Tick struct {
	Pair string `protobuf:"bytes,1,opt,name=pair" json:"pair,omitempty"`
	Time int64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Bid  int32  `protobuf:"varint,3,opt,name=bid" json:"bid,omitempty"`
	Ask  int32  `protobuf:"varint,4,opt,name=ask" json:"ask,omitempty"`
	Last int32  `protobuf:"varint,5,opt,name=last" json:"last,omitempty"`
}

func (m *Tick) Reset()                    { *m = Tick{} }
func (m *Tick) String() string            { return proto.CompactTextString(m) }
func (*Tick) ProtoMessage()               {}
func (*Tick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Tick)(nil), "Tick")
}

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcb, 0xcc, 0x34, 0x32, 0x36, 0x03, 0x51, 0xba,
	0x69, 0x99, 0x79, 0x89, 0x79, 0xc9, 0xa9, 0xfa, 0x25, 0x99, 0xc9, 0xd9, 0xc1, 0xa5, 0x49, 0xc5,
	0xc9, 0x45, 0x99, 0x49, 0xa9, 0x45, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x68, 0x82, 0x7a, 0x60,
	0x41, 0x25, 0x2f, 0x2e, 0x96, 0x10, 0xa0, 0xb8, 0x10, 0x0f, 0x17, 0x4b, 0x41, 0x62, 0x66, 0x91,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x27, 0x88, 0x57, 0x92, 0x99, 0x9b, 0x2a, 0xc1, 0x04, 0xe4, 0x31,
	0x0b, 0x71, 0x73, 0x31, 0x27, 0x65, 0xa6, 0x48, 0x30, 0x03, 0x39, 0xac, 0x20, 0x4e, 0x62, 0x71,
	0xb6, 0x04, 0x0b, 0x98, 0x03, 0x54, 0x97, 0x93, 0x58, 0x5c, 0x22, 0xc1, 0x0a, 0xe2, 0x25, 0xb1,
	0x81, 0x8d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xa6, 0x3c, 0x2f, 0x9d, 0x00, 0x00,
	0x00,
}
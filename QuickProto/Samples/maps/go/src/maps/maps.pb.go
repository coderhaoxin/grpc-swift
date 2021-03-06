// Code generated by protoc-gen-go.
// source: maps.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	maps.proto

It has these top-level messages:
	MapTest
*/
package test

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

type MapTest struct {
	Name              string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Properties        map[string]string `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	IntegerProperties map[int32]int32   `protobuf:"bytes,3,rep,name=integer_properties,json=integerProperties" json:"integer_properties,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *MapTest) Reset()                    { *m = MapTest{} }
func (m *MapTest) String() string            { return proto.CompactTextString(m) }
func (*MapTest) ProtoMessage()               {}
func (*MapTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapTest) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *MapTest) GetIntegerProperties() map[int32]int32 {
	if m != nil {
		return m.IntegerProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*MapTest)(nil), "test.MapTest")
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0x28,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e, 0x51, 0xda, 0xc6, 0xc4,
	0xc5, 0xee, 0x9b, 0x58, 0x10, 0x02, 0x64, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0xb6, 0x5c, 0x5c, 0x40, 0xe5, 0x05, 0xa9,
	0x45, 0x25, 0x99, 0xa9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xb2, 0x7a, 0x20, 0xad,
	0x7a, 0x50, 0x6d, 0x7a, 0x01, 0x70, 0x79, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x24, 0x0d, 0x42,
	0xc1, 0x5c, 0x42, 0x99, 0x79, 0x25, 0xa9, 0xe9, 0xa9, 0x45, 0xf1, 0x48, 0xc6, 0x30, 0x83, 0x8d,
	0x51, 0x41, 0x35, 0xc6, 0x13, 0xa2, 0x0e, 0xdd, 0x34, 0xc1, 0x4c, 0x74, 0x71, 0x29, 0x5b, 0x2e,
	0x7e, 0x34, 0x55, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x97, 0x83, 0x98, 0x42, 0x22,
	0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x40, 0x37, 0x83, 0xc4, 0x20, 0x1c, 0x2b, 0x26, 0x0b,
	0x46, 0x29, 0x17, 0x2e, 0x31, 0xec, 0x76, 0x21, 0x9b, 0xc2, 0x8a, 0xc5, 0x14, 0x56, 0x24, 0x53,
	0x92, 0xd8, 0xc0, 0xa1, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x76, 0x99, 0xbd, 0xf3, 0x53,
	0x01, 0x00, 0x00,
}

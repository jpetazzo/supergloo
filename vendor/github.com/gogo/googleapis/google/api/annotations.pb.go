// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/vendor/github.com/gogo/googleapis/google/api/annotations.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import annotations "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*annotations.HttpRule)(nil),
	Field:         72295728,
	Name:          "google.api.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "github.com/solo-io/supergloo/vendor/github.com/gogo/googleapis/google/api/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/vendor/github.com/gogo/googleapis/google/api/annotations.proto", fileDescriptor_annotations_e8792e9012e583ae)
}

var fileDescriptor_annotations_e8792e9012e583ae = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x60, 0xaa, 0xc5, 0x43, 0x04, 0x91, 0xa2, 0x20, 0x3d, 0x48, 0xf1, 0xe4, 0xc5, 0x1d,
	0xd1, 0x9b, 0xb7, 0xf6, 0xa2, 0x22, 0x62, 0xe8, 0x51, 0x4f, 0x9b, 0x64, 0xdd, 0x0c, 0xa4, 0xf9,
	0x87, 0xdd, 0x89, 0x0f, 0xe4, 0x13, 0xf8, 0x00, 0x3e, 0x81, 0x4f, 0x25, 0xdd, 0xa4, 0x34, 0xb7,
	0xdd, 0x99, 0x8f, 0xff, 0x9f, 0xec, 0xc3, 0xb3, 0xd6, 0x5d, 0x61, 0x4a, 0x6c, 0x28, 0xa2, 0xc1,
	0x0d, 0x83, 0x62, 0x27, 0x2e, 0xf8, 0x06, 0xa0, 0x2f, 0xd7, 0x56, 0x08, 0x34, 0x32, 0x1e, 0x1e,
	0xe4, 0x01, 0xdf, 0x38, 0x2b, 0x1c, 0x87, 0x27, 0x59, 0x61, 0xb2, 0x6d, 0x0b, 0xb5, 0xca, 0x68,
	0xa3, 0x91, 0x00, 0xc5, 0x2c, 0xeb, 0xb7, 0xc6, 0x0a, 0xcf, 0xcf, 0x47, 0xb2, 0x56, 0x95, 0x9e,
	0xcc, 0x17, 0xc3, 0x38, 0xfd, 0x8a, 0xee, 0x93, 0x2a, 0x17, 0xcb, 0xc0, 0xa2, 0x08, 0xbd, 0x78,
	0x78, 0xc9, 0xa6, 0x5b, 0x3f, 0xbb, 0x34, 0x43, 0xda, 0x8e, 0x9a, 0x57, 0xa7, 0x35, 0xaa, 0x37,
	0x49, 0x95, 0x17, 0x3f, 0x7f, 0xbf, 0x57, 0x8b, 0xc9, 0xf5, 0xf1, 0xdd, 0x99, 0xd9, 0xd7, 0x9a,
	0x27, 0x55, 0x59, 0x77, 0x8d, 0x5b, 0xa7, 0x90, 0xd5, 0x6d, 0x76, 0x52, 0x62, 0x33, 0x02, 0xab,
	0xd3, 0xe5, 0xfe, 0xec, 0x7c, 0x9b, 0x9c, 0x4f, 0xde, 0x0f, 0xad, 0xf0, 0xf7, 0xc1, 0xf4, 0x71,
	0x99, 0x3f, 0x17, 0x47, 0xa9, 0xee, 0xfe, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x2a, 0x5f, 0x93,
	0x29, 0x01, 0x00, 0x00,
}

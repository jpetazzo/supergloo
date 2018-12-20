// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/vendor/k8s.io/api/storage/v1alpha1/generated.proto

package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import _ "k8s.io/apimachinery/pkg/runtime"
import _ "k8s.io/apimachinery/pkg/runtime/schema"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// VolumeAttachment captures the intent to attach or detach the specified volume
// to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachment struct {
	// Standard object metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the desired attach/detach volume behavior.
	// Populated by the Kubernetes system.
	Spec *VolumeAttachmentSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status of the VolumeAttachment request.
	// Populated by the entity completing the attach or detach
	// operation, i.e. the external-attacher.
	// +optional
	Status               *VolumeAttachmentStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VolumeAttachment) Reset()         { *m = VolumeAttachment{} }
func (m *VolumeAttachment) String() string { return proto.CompactTextString(m) }
func (*VolumeAttachment) ProtoMessage()    {}
func (*VolumeAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{0}
}
func (m *VolumeAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttachment.Unmarshal(m, b)
}
func (m *VolumeAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttachment.Marshal(b, m, deterministic)
}
func (dst *VolumeAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttachment.Merge(dst, src)
}
func (m *VolumeAttachment) XXX_Size() int {
	return xxx_messageInfo_VolumeAttachment.Size(m)
}
func (m *VolumeAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttachment proto.InternalMessageInfo

func (m *VolumeAttachment) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VolumeAttachment) GetSpec() *VolumeAttachmentSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *VolumeAttachment) GetStatus() *VolumeAttachmentStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentList struct {
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is the list of VolumeAttachments
	Items                []*VolumeAttachment `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VolumeAttachmentList) Reset()         { *m = VolumeAttachmentList{} }
func (m *VolumeAttachmentList) String() string { return proto.CompactTextString(m) }
func (*VolumeAttachmentList) ProtoMessage()    {}
func (*VolumeAttachmentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{1}
}
func (m *VolumeAttachmentList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttachmentList.Unmarshal(m, b)
}
func (m *VolumeAttachmentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttachmentList.Marshal(b, m, deterministic)
}
func (dst *VolumeAttachmentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttachmentList.Merge(dst, src)
}
func (m *VolumeAttachmentList) XXX_Size() int {
	return xxx_messageInfo_VolumeAttachmentList.Size(m)
}
func (m *VolumeAttachmentList) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttachmentList.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttachmentList proto.InternalMessageInfo

func (m *VolumeAttachmentList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VolumeAttachmentList) GetItems() []*VolumeAttachment {
	if m != nil {
		return m.Items
	}
	return nil
}

// VolumeAttachmentSource represents a volume that should be attached.
// Right now only PersistenVolumes can be attached via external attacher,
// in future we may allow also inline volumes in pods.
// Exactly one member can be set.
type VolumeAttachmentSource struct {
	// Name of the persistent volume to attach.
	// +optional
	PersistentVolumeName *string  `protobuf:"bytes,1,opt,name=persistentVolumeName" json:"persistentVolumeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeAttachmentSource) Reset()         { *m = VolumeAttachmentSource{} }
func (m *VolumeAttachmentSource) String() string { return proto.CompactTextString(m) }
func (*VolumeAttachmentSource) ProtoMessage()    {}
func (*VolumeAttachmentSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{2}
}
func (m *VolumeAttachmentSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttachmentSource.Unmarshal(m, b)
}
func (m *VolumeAttachmentSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttachmentSource.Marshal(b, m, deterministic)
}
func (dst *VolumeAttachmentSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttachmentSource.Merge(dst, src)
}
func (m *VolumeAttachmentSource) XXX_Size() int {
	return xxx_messageInfo_VolumeAttachmentSource.Size(m)
}
func (m *VolumeAttachmentSource) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttachmentSource.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttachmentSource proto.InternalMessageInfo

func (m *VolumeAttachmentSource) GetPersistentVolumeName() string {
	if m != nil && m.PersistentVolumeName != nil {
		return *m.PersistentVolumeName
	}
	return ""
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this
	// request. This is the name returned by GetPluginName().
	Attacher *string `protobuf:"bytes,1,opt,name=attacher" json:"attacher,omitempty"`
	// Source represents the volume that should be attached.
	Source *VolumeAttachmentSource `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// The node that the volume should be attached to.
	NodeName             *string  `protobuf:"bytes,3,opt,name=nodeName" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeAttachmentSpec) Reset()         { *m = VolumeAttachmentSpec{} }
func (m *VolumeAttachmentSpec) String() string { return proto.CompactTextString(m) }
func (*VolumeAttachmentSpec) ProtoMessage()    {}
func (*VolumeAttachmentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{3}
}
func (m *VolumeAttachmentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttachmentSpec.Unmarshal(m, b)
}
func (m *VolumeAttachmentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttachmentSpec.Marshal(b, m, deterministic)
}
func (dst *VolumeAttachmentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttachmentSpec.Merge(dst, src)
}
func (m *VolumeAttachmentSpec) XXX_Size() int {
	return xxx_messageInfo_VolumeAttachmentSpec.Size(m)
}
func (m *VolumeAttachmentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttachmentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttachmentSpec proto.InternalMessageInfo

func (m *VolumeAttachmentSpec) GetAttacher() string {
	if m != nil && m.Attacher != nil {
		return *m.Attacher
	}
	return ""
}

func (m *VolumeAttachmentSpec) GetSource() *VolumeAttachmentSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *VolumeAttachmentSpec) GetNodeName() string {
	if m != nil && m.NodeName != nil {
		return *m.NodeName
	}
	return ""
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// Indicates the volume is successfully attached.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	Attached *bool `protobuf:"varint,1,opt,name=attached" json:"attached,omitempty"`
	// Upon successful attach, this field is populated with any
	// information returned by the attach operation that must be passed
	// into subsequent WaitForAttach or Mount calls.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
	AttachmentMetadata map[string]string `protobuf:"bytes,2,rep,name=attachmentMetadata" json:"attachmentMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The last error encountered during attach operation, if any.
	// This field must only be set by the entity completing the attach
	// operation, i.e. the external-attacher.
	// +optional
	AttachError *VolumeError `protobuf:"bytes,3,opt,name=attachError" json:"attachError,omitempty"`
	// The last error encountered during detach operation, if any.
	// This field must only be set by the entity completing the detach
	// operation, i.e. the external-attacher.
	// +optional
	DetachError          *VolumeError `protobuf:"bytes,4,opt,name=detachError" json:"detachError,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VolumeAttachmentStatus) Reset()         { *m = VolumeAttachmentStatus{} }
func (m *VolumeAttachmentStatus) String() string { return proto.CompactTextString(m) }
func (*VolumeAttachmentStatus) ProtoMessage()    {}
func (*VolumeAttachmentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{4}
}
func (m *VolumeAttachmentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttachmentStatus.Unmarshal(m, b)
}
func (m *VolumeAttachmentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttachmentStatus.Marshal(b, m, deterministic)
}
func (dst *VolumeAttachmentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttachmentStatus.Merge(dst, src)
}
func (m *VolumeAttachmentStatus) XXX_Size() int {
	return xxx_messageInfo_VolumeAttachmentStatus.Size(m)
}
func (m *VolumeAttachmentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttachmentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttachmentStatus proto.InternalMessageInfo

func (m *VolumeAttachmentStatus) GetAttached() bool {
	if m != nil && m.Attached != nil {
		return *m.Attached
	}
	return false
}

func (m *VolumeAttachmentStatus) GetAttachmentMetadata() map[string]string {
	if m != nil {
		return m.AttachmentMetadata
	}
	return nil
}

func (m *VolumeAttachmentStatus) GetAttachError() *VolumeError {
	if m != nil {
		return m.AttachError
	}
	return nil
}

func (m *VolumeAttachmentStatus) GetDetachError() *VolumeError {
	if m != nil {
		return m.DetachError
	}
	return nil
}

// VolumeError captures an error encountered during a volume operation.
type VolumeError struct {
	// Time the error was encountered.
	// +optional
	Time *v1.Time `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	// String detailing the error encountered during Attach or Detach operation.
	// This string maybe logged, so it should not contain sensitive
	// information.
	// +optional
	Message              *string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeError) Reset()         { *m = VolumeError{} }
func (m *VolumeError) String() string { return proto.CompactTextString(m) }
func (*VolumeError) ProtoMessage()    {}
func (*VolumeError) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_c8554ad24a6e43dd, []int{5}
}
func (m *VolumeError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeError.Unmarshal(m, b)
}
func (m *VolumeError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeError.Marshal(b, m, deterministic)
}
func (dst *VolumeError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeError.Merge(dst, src)
}
func (m *VolumeError) XXX_Size() int {
	return xxx_messageInfo_VolumeError.Size(m)
}
func (m *VolumeError) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeError.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeError proto.InternalMessageInfo

func (m *VolumeError) GetTime() *v1.Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *VolumeError) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VolumeAttachment)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachment")
	proto.RegisterType((*VolumeAttachmentList)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentList")
	proto.RegisterType((*VolumeAttachmentSource)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentSource")
	proto.RegisterType((*VolumeAttachmentSpec)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentSpec")
	proto.RegisterType((*VolumeAttachmentStatus)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentStatus")
	proto.RegisterMapType((map[string]string)(nil), "k8s.io.api.storage.v1alpha1.VolumeAttachmentStatus.AttachmentMetadataEntry")
	proto.RegisterType((*VolumeError)(nil), "k8s.io.api.storage.v1alpha1.VolumeError")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/vendor/k8s.io/api/storage/v1alpha1/generated.proto", fileDescriptor_generated_c8554ad24a6e43dd)
}

var fileDescriptor_generated_c8554ad24a6e43dd = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x3e, 0x0a, 0xa9, 0x73, 0xa9, 0xac, 0x08, 0xa2, 0x70, 0xa9, 0xf6, 0x14, 0x21, 0xd5,
	0x26, 0x01, 0xa1, 0x8a, 0x03, 0x52, 0x41, 0xb9, 0xb4, 0x29, 0x95, 0x0c, 0xe2, 0xc0, 0xcd, 0xdd,
	0x1d, 0x6d, 0x4c, 0xe2, 0xb5, 0x65, 0x7b, 0x23, 0x45, 0xfc, 0x09, 0xee, 0xfc, 0x00, 0xfe, 0x20,
	0x3f, 0x00, 0xd9, 0xbb, 0x9b, 0x44, 0xd9, 0x14, 0x92, 0xde, 0x32, 0x93, 0xf7, 0xde, 0xbc, 0x37,
	0xb3, 0xbb, 0xe8, 0x2e, 0x15, 0x6e, 0x96, 0xdf, 0x93, 0x58, 0x49, 0x6a, 0xd5, 0x42, 0x5d, 0x08,
	0x45, 0x6d, 0xae, 0xc1, 0xa4, 0x0b, 0xa5, 0xe8, 0x12, 0xb2, 0x44, 0x19, 0x3a, 0xbf, 0xb4, 0x44,
	0x28, 0xca, 0xb5, 0xa0, 0xd6, 0x29, 0xc3, 0x53, 0xa0, 0xcb, 0x11, 0x5f, 0xe8, 0x19, 0x1f, 0xd1,
	0x14, 0x32, 0x30, 0xdc, 0x41, 0x42, 0xb4, 0x51, 0x4e, 0xe1, 0x17, 0x05, 0x98, 0x70, 0x2d, 0x48,
	0x09, 0x26, 0x15, 0x78, 0xf0, 0x66, 0xa3, 0x24, 0x79, 0x3c, 0x13, 0x19, 0x98, 0x15, 0xd5, 0xf3,
	0xd4, 0x37, 0x2c, 0x95, 0xe0, 0x38, 0x5d, 0xd6, 0x24, 0x07, 0xf4, 0x21, 0x96, 0xc9, 0x33, 0x27,
	0x24, 0xd4, 0x08, 0x6f, 0xff, 0x47, 0xb0, 0xf1, 0x0c, 0x24, 0xdf, 0xe5, 0x45, 0x7f, 0x1a, 0xe8,
	0xec, 0xab, 0x5a, 0xe4, 0x12, 0xae, 0x9c, 0xe3, 0xf1, 0x4c, 0x42, 0xe6, 0xf0, 0x14, 0x75, 0xbc,
	0xb1, 0x84, 0x3b, 0xde, 0x6f, 0x9c, 0x37, 0x86, 0xdd, 0xf1, 0x2b, 0xb2, 0xc9, 0xb8, 0xd6, 0x27,
	0x7a, 0x9e, 0xfa, 0x86, 0x25, 0x1e, 0x4d, 0x96, 0x23, 0x72, 0x77, 0xff, 0x1d, 0x62, 0x77, 0x0b,
	0x8e, 0xb3, 0xb5, 0x02, 0x9e, 0xa0, 0xb6, 0xd5, 0x10, 0xf7, 0x9b, 0x41, 0x69, 0x44, 0xfe, 0xb1,
	0x2d, 0xb2, 0x6b, 0xe5, 0xb3, 0x86, 0x98, 0x05, 0x3a, 0xbe, 0x41, 0x4f, 0xac, 0xe3, 0x2e, 0xb7,
	0xfd, 0x56, 0x10, 0x7a, 0x7d, 0x9c, 0x50, 0xa0, 0xb2, 0x52, 0x22, 0xfa, 0xdd, 0x40, 0xbd, 0x5d,
	0xc8, 0x54, 0x58, 0x87, 0xaf, 0x6b, 0xd1, 0xc9, 0x61, 0xd1, 0x3d, 0x7b, 0x27, 0xf8, 0x47, 0x74,
	0x22, 0x1c, 0x48, 0xdb, 0x6f, 0x9e, 0xb7, 0x86, 0xdd, 0xf1, 0xc5, 0x51, 0x86, 0x59, 0xc1, 0x8d,
	0xa6, 0xe8, 0x59, 0x2d, 0x8b, 0xca, 0x4d, 0x0c, 0x78, 0x8c, 0x7a, 0x1a, 0x8c, 0x15, 0xd6, 0x41,
	0xe6, 0x0a, 0xcc, 0x27, 0x2e, 0x21, 0xd8, 0x3e, 0x65, 0x7b, 0xff, 0x8b, 0x7e, 0xed, 0xc9, 0xed,
	0x77, 0x8c, 0x07, 0xa8, 0xc3, 0x43, 0x07, 0x4c, 0x29, 0xb0, 0xae, 0xc3, 0xe6, 0xc3, 0xc8, 0xf2,
	0x84, 0x47, 0x6e, 0x3e, 0x50, 0x59, 0x29, 0xe1, 0x07, 0x65, 0x2a, 0x29, 0x9c, 0xb6, 0x8a, 0x41,
	0x55, 0x1d, 0xfd, 0x6c, 0xed, 0x09, 0x1b, 0x0e, 0xb6, 0xe5, 0x2f, 0x09, 0xfe, 0x3a, 0x6b, 0x7f,
	0x09, 0xfe, 0x81, 0x30, 0x5f, 0xe3, 0x6f, 0xab, 0xeb, 0x15, 0x4b, 0xbf, 0x79, 0xc4, 0x53, 0x42,
	0xae, 0x6a, 0x6a, 0x93, 0xcc, 0x99, 0x15, 0xdb, 0x33, 0x06, 0x5f, 0xa3, 0x6e, 0xd1, 0x9d, 0x18,
	0xa3, 0x4c, 0xf9, 0x6c, 0x0e, 0x0f, 0x98, 0x1a, 0xf0, 0x6c, 0x9b, 0xec, 0xb5, 0x12, 0xd8, 0x68,
	0xb5, 0x8f, 0xd5, 0xda, 0x22, 0x0f, 0x26, 0xe8, 0xf9, 0x03, 0x31, 0xf0, 0x19, 0x6a, 0xcd, 0x61,
	0x55, 0x9e, 0xd9, 0xff, 0xc4, 0x3d, 0x74, 0xb2, 0xe4, 0x8b, 0xbc, 0x38, 0xf0, 0x29, 0x2b, 0x8a,
	0x77, 0xcd, 0xcb, 0x46, 0x94, 0xa2, 0xee, 0xd6, 0x08, 0xfc, 0x1e, 0xb5, 0xfd, 0xd7, 0xa4, 0x7c,
	0x35, 0x5e, 0x1e, 0xf6, 0x6a, 0x7c, 0x11, 0x12, 0x58, 0xe0, 0xe1, 0x3e, 0x7a, 0x2a, 0xc1, 0x5a,
	0x9e, 0x56, 0xa3, 0xaa, 0xf2, 0x03, 0xfa, 0xd6, 0xa9, 0x42, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x10, 0xb9, 0x4e, 0xa2, 0x05, 0x00, 0x00,
}

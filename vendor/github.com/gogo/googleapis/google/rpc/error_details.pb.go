// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/vendor/github.com/gogo/googleapis/google/rpc/error_details.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retires have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay           *types.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RetryInfo) Reset()         { *m = RetryInfo{} }
func (m *RetryInfo) String() string { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()    {}
func (*RetryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{0}
}
func (m *RetryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryInfo.Unmarshal(m, b)
}
func (m *RetryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryInfo.Marshal(b, m, deterministic)
}
func (dst *RetryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryInfo.Merge(dst, src)
}
func (m *RetryInfo) XXX_Size() int {
	return xxx_messageInfo_RetryInfo.Size(m)
}
func (m *RetryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RetryInfo proto.InternalMessageInfo

func (m *RetryInfo) GetRetryDelay() *types.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugInfo) Reset()         { *m = DebugInfo{} }
func (m *DebugInfo) String() string { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()    {}
func (*DebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{1}
}
func (m *DebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugInfo.Unmarshal(m, b)
}
func (m *DebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugInfo.Marshal(b, m, deterministic)
}
func (dst *DebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugInfo.Merge(dst, src)
}
func (m *DebugInfo) XXX_Size() int {
	return xxx_messageInfo_DebugInfo.Size(m)
}
func (m *DebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DebugInfo proto.InternalMessageInfo

func (m *DebugInfo) GetStackEntries() []string {
	if m != nil {
		return m.StackEntries
	}
	return nil
}

func (m *DebugInfo) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryDetail and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations           []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *QuotaFailure) Reset()         { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()    {}
func (*QuotaFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{2}
}
func (m *QuotaFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaFailure.Unmarshal(m, b)
}
func (m *QuotaFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaFailure.Marshal(b, m, deterministic)
}
func (dst *QuotaFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaFailure.Merge(dst, src)
}
func (m *QuotaFailure) XXX_Size() int {
	return xxx_messageInfo_QuotaFailure.Size(m)
}
func (m *QuotaFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaFailure.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaFailure proto.InternalMessageInfo

func (m *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "project:<Google
	// developer project id>".
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaFailure_Violation) Reset()         { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()    {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{2, 0}
}
func (m *QuotaFailure_Violation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaFailure_Violation.Unmarshal(m, b)
}
func (m *QuotaFailure_Violation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaFailure_Violation.Marshal(b, m, deterministic)
}
func (dst *QuotaFailure_Violation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaFailure_Violation.Merge(dst, src)
}
func (m *QuotaFailure_Violation) XXX_Size() int {
	return xxx_messageInfo_QuotaFailure_Violation.Size(m)
}
func (m *QuotaFailure_Violation) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaFailure_Violation.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaFailure_Violation proto.InternalMessageInfo

func (m *QuotaFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *QuotaFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes what preconditions have failed.
//
// For example, if an RPC failed because it required the Terms of Service to be
// acknowledged, it could list the terms of service violation in the
// PreconditionFailure message.
type PreconditionFailure struct {
	// Describes all precondition violations.
	Violations           []*PreconditionFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PreconditionFailure) Reset()         { *m = PreconditionFailure{} }
func (m *PreconditionFailure) String() string { return proto.CompactTextString(m) }
func (*PreconditionFailure) ProtoMessage()    {}
func (*PreconditionFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{3}
}
func (m *PreconditionFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionFailure.Unmarshal(m, b)
}
func (m *PreconditionFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionFailure.Marshal(b, m, deterministic)
}
func (dst *PreconditionFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionFailure.Merge(dst, src)
}
func (m *PreconditionFailure) XXX_Size() int {
	return xxx_messageInfo_PreconditionFailure.Size(m)
}
func (m *PreconditionFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionFailure.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionFailure proto.InternalMessageInfo

func (m *PreconditionFailure) GetViolations() []*PreconditionFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single precondition failure.
type PreconditionFailure_Violation struct {
	// The type of PreconditionFailure. We recommend using a service-specific
	// enum type to define the supported precondition violation types. For
	// example, "TOS" for "Terms of Service violation".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The subject, relative to the type, that failed.
	// For example, "google.com/cloud" relative to the "TOS" type would
	// indicate which terms of service is being referenced.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the precondition failed. Developers can use this
	// description to understand how to fix the failure.
	//
	// For example: "Terms of service not accepted".
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreconditionFailure_Violation) Reset()         { *m = PreconditionFailure_Violation{} }
func (m *PreconditionFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*PreconditionFailure_Violation) ProtoMessage()    {}
func (*PreconditionFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{3, 0}
}
func (m *PreconditionFailure_Violation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionFailure_Violation.Unmarshal(m, b)
}
func (m *PreconditionFailure_Violation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionFailure_Violation.Marshal(b, m, deterministic)
}
func (dst *PreconditionFailure_Violation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionFailure_Violation.Merge(dst, src)
}
func (m *PreconditionFailure_Violation) XXX_Size() int {
	return xxx_messageInfo_PreconditionFailure_Violation.Size(m)
}
func (m *PreconditionFailure_Violation) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionFailure_Violation.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionFailure_Violation proto.InternalMessageInfo

func (m *PreconditionFailure_Violation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PreconditionFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *PreconditionFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	// Describes all violations in a client request.
	FieldViolations      []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations" json:"field_violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BadRequest) Reset()         { *m = BadRequest{} }
func (m *BadRequest) String() string { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()    {}
func (*BadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{4}
}
func (m *BadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequest.Unmarshal(m, b)
}
func (m *BadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequest.Marshal(b, m, deterministic)
}
func (dst *BadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequest.Merge(dst, src)
}
func (m *BadRequest) XXX_Size() int {
	return xxx_messageInfo_BadRequest.Size(m)
}
func (m *BadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequest proto.InternalMessageInfo

func (m *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if m != nil {
		return m.FieldViolations
	}
	return nil
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	// A path leading to a field in the request body. The value will be a
	// sequence of dot-separated identifiers that identify a protocol buffer
	// field. E.g., "field_violations.field" would identify this field.
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// A description of why the request element is bad.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BadRequest_FieldViolation) Reset()         { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()    {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{4, 0}
}
func (m *BadRequest_FieldViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequest_FieldViolation.Unmarshal(m, b)
}
func (m *BadRequest_FieldViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequest_FieldViolation.Marshal(b, m, deterministic)
}
func (dst *BadRequest_FieldViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequest_FieldViolation.Merge(dst, src)
}
func (m *BadRequest_FieldViolation) XXX_Size() int {
	return xxx_messageInfo_BadRequest_FieldViolation.Size(m)
}
func (m *BadRequest_FieldViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequest_FieldViolation.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequest_FieldViolation proto.InternalMessageInfo

func (m *BadRequest_FieldViolation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *BadRequest_FieldViolation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service generating
	// it. For example, it can be used to identify requests in the service's logs.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData          string   `protobuf:"bytes,2,opt,name=serving_data,json=servingData,proto3" json:"serving_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{5}
}
func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (dst *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(dst, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RequestInfo) GetServingData() string {
	if m != nil {
		return m.ServingData
	}
	return ""
}

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	// "cloud storage bucket", "file", "Google calendar"; or the type URL
	// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The name of the resource being accessed.  For example, a shared calendar
	// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The owner of the resource (optional).
	// For example, "user:<owner email>" or "project:<Google developer project
	// id>".
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Describes what error is encountered when accessing this resource.
	// For example, updating a cloud project may require the `writer` permission
	// on the developer console project.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceInfo) Reset()         { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()    {}
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{6}
}
func (m *ResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceInfo.Unmarshal(m, b)
}
func (m *ResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceInfo.Marshal(b, m, deterministic)
}
func (dst *ResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceInfo.Merge(dst, src)
}
func (m *ResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ResourceInfo.Size(m)
}
func (m *ResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceInfo proto.InternalMessageInfo

func (m *ResourceInfo) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceInfo) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ResourceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ResourceInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links                []*Help_Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Help) Reset()         { *m = Help{} }
func (m *Help) String() string { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()    {}
func (*Help) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{7}
}
func (m *Help) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Help.Unmarshal(m, b)
}
func (m *Help) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Help.Marshal(b, m, deterministic)
}
func (dst *Help) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Help.Merge(dst, src)
}
func (m *Help) XXX_Size() int {
	return xxx_messageInfo_Help.Size(m)
}
func (m *Help) XXX_DiscardUnknown() {
	xxx_messageInfo_Help.DiscardUnknown(m)
}

var xxx_messageInfo_Help proto.InternalMessageInfo

func (m *Help) GetLinks() []*Help_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// Describes a URL link.
type Help_Link struct {
	// Describes what the link offers.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The URL of the link.
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Help_Link) Reset()         { *m = Help_Link{} }
func (m *Help_Link) String() string { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()    {}
func (*Help_Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{7, 0}
}
func (m *Help_Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Help_Link.Unmarshal(m, b)
}
func (m *Help_Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Help_Link.Marshal(b, m, deterministic)
}
func (dst *Help_Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Help_Link.Merge(dst, src)
}
func (m *Help_Link) XXX_Size() int {
	return xxx_messageInfo_Help_Link.Size(m)
}
func (m *Help_Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Help_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Help_Link proto.InternalMessageInfo

func (m *Help_Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Help_Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Provides a localized error message that is safe to return to the user
// which can be attached to an RPC error.
type LocalizedMessage struct {
	// The locale used following the specification defined at
	// http://www.rfc-editor.org/rfc/bcp/bcp47.txt.
	// Examples are: "en-US", "fr-CH", "es-MX"
	Locale string `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	// The localized error message in the above locale.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalizedMessage) Reset()         { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()    {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_details_f8968d9248e86288, []int{8}
}
func (m *LocalizedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalizedMessage.Unmarshal(m, b)
}
func (m *LocalizedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalizedMessage.Marshal(b, m, deterministic)
}
func (dst *LocalizedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalizedMessage.Merge(dst, src)
}
func (m *LocalizedMessage) XXX_Size() int {
	return xxx_messageInfo_LocalizedMessage.Size(m)
}
func (m *LocalizedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalizedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LocalizedMessage proto.InternalMessageInfo

func (m *LocalizedMessage) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *LocalizedMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RetryInfo)(nil), "google.rpc.RetryInfo")
	proto.RegisterType((*DebugInfo)(nil), "google.rpc.DebugInfo")
	proto.RegisterType((*QuotaFailure)(nil), "google.rpc.QuotaFailure")
	proto.RegisterType((*QuotaFailure_Violation)(nil), "google.rpc.QuotaFailure.Violation")
	proto.RegisterType((*PreconditionFailure)(nil), "google.rpc.PreconditionFailure")
	proto.RegisterType((*PreconditionFailure_Violation)(nil), "google.rpc.PreconditionFailure.Violation")
	proto.RegisterType((*BadRequest)(nil), "google.rpc.BadRequest")
	proto.RegisterType((*BadRequest_FieldViolation)(nil), "google.rpc.BadRequest.FieldViolation")
	proto.RegisterType((*RequestInfo)(nil), "google.rpc.RequestInfo")
	proto.RegisterType((*ResourceInfo)(nil), "google.rpc.ResourceInfo")
	proto.RegisterType((*Help)(nil), "google.rpc.Help")
	proto.RegisterType((*Help_Link)(nil), "google.rpc.Help.Link")
	proto.RegisterType((*LocalizedMessage)(nil), "google.rpc.LocalizedMessage")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/vendor/github.com/gogo/googleapis/google/rpc/error_details.proto", fileDescriptor_error_details_f8968d9248e86288)
}

var fileDescriptor_error_details_f8968d9248e86288 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0xb5, 0xdb, 0x5f, 0x39, 0xed, 0x7f, 0x8c, 0xf0, 0xa2, 0x52, 0x09, 0x54, 0x82,
	0x90, 0x86, 0x10, 0x09, 0x1a, 0x77, 0xbb, 0x2c, 0xdd, 0x9b, 0x34, 0xa0, 0x44, 0x88, 0x0b, 0x10,
	0x8a, 0xdc, 0xe4, 0x34, 0x98, 0xa5, 0x39, 0xc1, 0x76, 0x86, 0xc6, 0xa7, 0xe0, 0x9e, 0x3b, 0xae,
	0xf8, 0x12, 0x7c, 0x37, 0xe4, 0xc4, 0x59, 0xbd, 0x75, 0x20, 0xee, 0xfc, 0x1c, 0xff, 0xfc, 0xe4,
	0x39, 0x47, 0x8e, 0xe1, 0x43, 0xc6, 0xd5, 0xc7, 0x6a, 0x16, 0x24, 0xb4, 0x08, 0x25, 0xe5, 0xf4,
	0x84, 0x53, 0x28, 0xab, 0x12, 0x45, 0x96, 0x13, 0x85, 0xa7, 0x58, 0xa4, 0x24, 0x42, 0x8b, 0xc9,
	0x28, 0xa3, 0x30, 0x23, 0xca, 0x72, 0x64, 0x25, 0x97, 0x66, 0x19, 0x8a, 0x32, 0x09, 0x51, 0x08,
	0x12, 0x71, 0x8a, 0x8a, 0xf1, 0x5c, 0x06, 0xa5, 0x20, 0x45, 0x1e, 0x34, 0xfb, 0x81, 0x28, 0x93,
	0xe1, 0x3d, 0xc3, 0xd6, 0x3b, 0xb3, 0x6a, 0x1e, 0xa6, 0x95, 0x60, 0x8a, 0x53, 0xd1, 0xb0, 0xfe,
	0x01, 0xb8, 0x11, 0x2a, 0x71, 0x76, 0x54, 0xcc, 0xc9, 0xdb, 0x85, 0x9e, 0xd0, 0x22, 0x4e, 0x31,
	0x67, 0x67, 0x03, 0x67, 0xe4, 0x6c, 0xf7, 0x76, 0xee, 0x04, 0xc6, 0xae, 0xb5, 0x08, 0x26, 0xc6,
	0x22, 0x82, 0x9a, 0x9e, 0x68, 0xd8, 0x3f, 0x04, 0x77, 0x82, 0xb3, 0x2a, 0xab, 0x8d, 0x1e, 0xc0,
	0xff, 0x52, 0xb1, 0xe4, 0x24, 0xc6, 0x42, 0x09, 0x8e, 0x72, 0xe0, 0x8c, 0x3a, 0xdb, 0x6e, 0xd4,
	0xaf, 0x8b, 0x7b, 0x4d, 0xcd, 0xbb, 0x0d, 0x1b, 0x4d, 0xee, 0xc1, 0xda, 0xc8, 0xd9, 0x76, 0x23,
	0xa3, 0xfc, 0xef, 0x0e, 0xf4, 0x5f, 0x57, 0xa4, 0xd8, 0x3e, 0xe3, 0x79, 0x25, 0xd0, 0x1b, 0x03,
	0x9c, 0x72, 0xca, 0xeb, 0x6f, 0x36, 0x56, 0xbd, 0x1d, 0x3f, 0x58, 0x36, 0x19, 0xd8, 0x74, 0xf0,
	0xb6, 0x45, 0x23, 0xeb, 0xd4, 0xf0, 0x00, 0xdc, 0xf3, 0x0d, 0x6f, 0x00, 0xff, 0xc9, 0x6a, 0xf6,
	0x09, 0x13, 0x55, 0xf7, 0xe8, 0x46, 0xad, 0xf4, 0x46, 0xd0, 0x4b, 0x51, 0x26, 0x82, 0x97, 0x1a,
	0x34, 0xc1, 0xec, 0x92, 0xff, 0xcb, 0x81, 0x1b, 0x53, 0x81, 0x09, 0x15, 0x29, 0xd7, 0x85, 0x36,
	0xe4, 0xd1, 0x15, 0x21, 0x1f, 0xd9, 0x21, 0xaf, 0x38, 0xf4, 0x87, 0xac, 0xef, 0xed, 0xac, 0x1e,
	0x74, 0xd5, 0x59, 0x89, 0x26, 0x68, 0xbd, 0xb6, 0xf3, 0xaf, 0xfd, 0x35, 0x7f, 0x67, 0x35, 0xff,
	0x4f, 0x07, 0x60, 0xcc, 0xd2, 0x08, 0x3f, 0x57, 0x28, 0x95, 0x37, 0x85, 0xad, 0x39, 0xc7, 0x3c,
	0x8d, 0x57, 0xc2, 0x3f, 0xb4, 0xc3, 0x2f, 0x4f, 0x04, 0xfb, 0x1a, 0x5f, 0x06, 0xbf, 0x36, 0xbf,
	0xa0, 0xe5, 0xf0, 0x10, 0x36, 0x2f, 0x22, 0xde, 0x4d, 0x58, 0xaf, 0x21, 0xd3, 0x43, 0x23, 0xfe,
	0x61, 0xd4, 0xaf, 0xa0, 0x67, 0x3e, 0x5a, 0x5f, 0xaa, 0xbb, 0x00, 0xa2, 0x91, 0x31, 0x6f, 0xbd,
	0x5c, 0x53, 0x39, 0x4a, 0xbd, 0xfb, 0xd0, 0x97, 0x28, 0x4e, 0x79, 0x91, 0xc5, 0x29, 0x53, 0xac,
	0x35, 0x34, 0xb5, 0x09, 0x53, 0xcc, 0xff, 0xe6, 0x40, 0x3f, 0x42, 0x49, 0x95, 0x48, 0xb0, 0xbd,
	0xa7, 0xc2, 0xe8, 0xd8, 0x9a, 0x72, 0xbf, 0x2d, 0xbe, 0xd1, 0xd3, 0xb6, 0xa1, 0x82, 0x2d, 0xd0,
	0x38, 0x9f, 0x43, 0x2f, 0xd9, 0x02, 0x75, 0x8f, 0xf4, 0xa5, 0x40, 0x61, 0x46, 0xde, 0x88, 0xcb,
	0x3d, 0x76, 0x57, 0x7b, 0x24, 0xe8, 0x1e, 0x62, 0x5e, 0x7a, 0x8f, 0x61, 0x3d, 0xe7, 0xc5, 0x49,
	0x3b, 0xfc, 0x5b, 0xf6, 0xf0, 0x35, 0x10, 0x1c, 0xf3, 0xe2, 0x24, 0x6a, 0x98, 0xe1, 0x2e, 0x74,
	0xb5, 0xbc, 0x6c, 0xef, 0xac, 0xd8, 0x7b, 0x5b, 0xd0, 0xa9, 0x44, 0xfb, 0x83, 0xe9, 0xa5, 0x3f,
	0x81, 0xad, 0x63, 0x4a, 0x58, 0xce, 0xbf, 0x62, 0xfa, 0x02, 0xa5, 0x64, 0x19, 0xea, 0x3f, 0x31,
	0xd7, 0xb5, 0xb6, 0x7f, 0xa3, 0xf4, 0x3d, 0x5b, 0x34, 0x48, 0x7b, 0xcf, 0x8c, 0x1c, 0x3f, 0x85,
	0xcd, 0x84, 0x16, 0x56, 0xc8, 0xf1, 0xf5, 0x3d, 0xfd, 0x12, 0x4d, 0x9a, 0x87, 0x68, 0xaa, 0x9f,
	0x8a, 0xa9, 0xf3, 0xae, 0x23, 0xca, 0xe4, 0xc7, 0x5a, 0x27, 0x9a, 0x3e, 0x9f, 0x6d, 0xd4, 0xcf,
	0xc7, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0x34, 0xc0, 0xd5, 0xfc, 0x04, 0x00, 0x00,
}

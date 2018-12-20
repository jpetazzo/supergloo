// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/vendor/k8s.io/api/batch/v1beta1/generated.proto

package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v12 "k8s.io/api/batch/v1"
import v11 "k8s.io/api/core/v1"
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

// CronJob represents the configuration of a single cron job.
type CronJob struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the desired behavior of a cron job, including the schedule.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec *CronJobSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Current status of a cron job.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status               *CronJobStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CronJob) Reset()         { *m = CronJob{} }
func (m *CronJob) String() string { return proto.CompactTextString(m) }
func (*CronJob) ProtoMessage()    {}
func (*CronJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{0}
}
func (m *CronJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJob.Unmarshal(m, b)
}
func (m *CronJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJob.Marshal(b, m, deterministic)
}
func (dst *CronJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJob.Merge(dst, src)
}
func (m *CronJob) XXX_Size() int {
	return xxx_messageInfo_CronJob.Size(m)
}
func (m *CronJob) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJob.DiscardUnknown(m)
}

var xxx_messageInfo_CronJob proto.InternalMessageInfo

func (m *CronJob) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CronJob) GetSpec() *CronJobSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CronJob) GetStatus() *CronJobStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// CronJobList is a collection of cron jobs.
type CronJobList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// items is the list of CronJobs.
	Items                []*CronJob `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CronJobList) Reset()         { *m = CronJobList{} }
func (m *CronJobList) String() string { return proto.CompactTextString(m) }
func (*CronJobList) ProtoMessage()    {}
func (*CronJobList) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{1}
}
func (m *CronJobList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJobList.Unmarshal(m, b)
}
func (m *CronJobList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJobList.Marshal(b, m, deterministic)
}
func (dst *CronJobList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJobList.Merge(dst, src)
}
func (m *CronJobList) XXX_Size() int {
	return xxx_messageInfo_CronJobList.Size(m)
}
func (m *CronJobList) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJobList.DiscardUnknown(m)
}

var xxx_messageInfo_CronJobList proto.InternalMessageInfo

func (m *CronJobList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CronJobList) GetItems() []*CronJob {
	if m != nil {
		return m.Items
	}
	return nil
}

// CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpec struct {
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule *string `protobuf:"bytes,1,opt,name=schedule" json:"schedule,omitempty"`
	// Optional deadline in seconds for starting the job if it misses scheduled
	// time for any reason.  Missed jobs executions will be counted as failed ones.
	// +optional
	StartingDeadlineSeconds *int64 `protobuf:"varint,2,opt,name=startingDeadlineSeconds" json:"startingDeadlineSeconds,omitempty"`
	// Specifies how to treat concurrent executions of a Job.
	// Valid values are:
	// - "Allow" (default): allows CronJobs to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy *string `protobuf:"bytes,3,opt,name=concurrencyPolicy" json:"concurrencyPolicy,omitempty"`
	// This flag tells the controller to suspend subsequent executions, it does
	// not apply to already started executions.  Defaults to false.
	// +optional
	Suspend *bool `protobuf:"varint,4,opt,name=suspend" json:"suspend,omitempty"`
	// Specifies the job that will be created when executing a CronJob.
	JobTemplate *JobTemplateSpec `protobuf:"bytes,5,opt,name=jobTemplate" json:"jobTemplate,omitempty"`
	// The number of successful finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 3.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `protobuf:"varint,6,opt,name=successfulJobsHistoryLimit" json:"successfulJobsHistoryLimit,omitempty"`
	// The number of failed finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 1.
	// +optional
	FailedJobsHistoryLimit *int32   `protobuf:"varint,7,opt,name=failedJobsHistoryLimit" json:"failedJobsHistoryLimit,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *CronJobSpec) Reset()         { *m = CronJobSpec{} }
func (m *CronJobSpec) String() string { return proto.CompactTextString(m) }
func (*CronJobSpec) ProtoMessage()    {}
func (*CronJobSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{2}
}
func (m *CronJobSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJobSpec.Unmarshal(m, b)
}
func (m *CronJobSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJobSpec.Marshal(b, m, deterministic)
}
func (dst *CronJobSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJobSpec.Merge(dst, src)
}
func (m *CronJobSpec) XXX_Size() int {
	return xxx_messageInfo_CronJobSpec.Size(m)
}
func (m *CronJobSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJobSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CronJobSpec proto.InternalMessageInfo

func (m *CronJobSpec) GetSchedule() string {
	if m != nil && m.Schedule != nil {
		return *m.Schedule
	}
	return ""
}

func (m *CronJobSpec) GetStartingDeadlineSeconds() int64 {
	if m != nil && m.StartingDeadlineSeconds != nil {
		return *m.StartingDeadlineSeconds
	}
	return 0
}

func (m *CronJobSpec) GetConcurrencyPolicy() string {
	if m != nil && m.ConcurrencyPolicy != nil {
		return *m.ConcurrencyPolicy
	}
	return ""
}

func (m *CronJobSpec) GetSuspend() bool {
	if m != nil && m.Suspend != nil {
		return *m.Suspend
	}
	return false
}

func (m *CronJobSpec) GetJobTemplate() *JobTemplateSpec {
	if m != nil {
		return m.JobTemplate
	}
	return nil
}

func (m *CronJobSpec) GetSuccessfulJobsHistoryLimit() int32 {
	if m != nil && m.SuccessfulJobsHistoryLimit != nil {
		return *m.SuccessfulJobsHistoryLimit
	}
	return 0
}

func (m *CronJobSpec) GetFailedJobsHistoryLimit() int32 {
	if m != nil && m.FailedJobsHistoryLimit != nil {
		return *m.FailedJobsHistoryLimit
	}
	return 0
}

// CronJobStatus represents the current state of a cron job.
type CronJobStatus struct {
	// A list of pointers to currently running jobs.
	// +optional
	Active []*v11.ObjectReference `protobuf:"bytes,1,rep,name=active" json:"active,omitempty"`
	// Information when was the last time the job was successfully scheduled.
	// +optional
	LastScheduleTime     *v1.Time `protobuf:"bytes,4,opt,name=lastScheduleTime" json:"lastScheduleTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CronJobStatus) Reset()         { *m = CronJobStatus{} }
func (m *CronJobStatus) String() string { return proto.CompactTextString(m) }
func (*CronJobStatus) ProtoMessage()    {}
func (*CronJobStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{3}
}
func (m *CronJobStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJobStatus.Unmarshal(m, b)
}
func (m *CronJobStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJobStatus.Marshal(b, m, deterministic)
}
func (dst *CronJobStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJobStatus.Merge(dst, src)
}
func (m *CronJobStatus) XXX_Size() int {
	return xxx_messageInfo_CronJobStatus.Size(m)
}
func (m *CronJobStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJobStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CronJobStatus proto.InternalMessageInfo

func (m *CronJobStatus) GetActive() []*v11.ObjectReference {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *CronJobStatus) GetLastScheduleTime() *v1.Time {
	if m != nil {
		return m.LastScheduleTime
	}
	return nil
}

// JobTemplate describes a template for creating copies of a predefined pod.
type JobTemplate struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Defines jobs that will be created from this template.
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Template             *JobTemplateSpec `protobuf:"bytes,2,opt,name=template" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JobTemplate) Reset()         { *m = JobTemplate{} }
func (m *JobTemplate) String() string { return proto.CompactTextString(m) }
func (*JobTemplate) ProtoMessage()    {}
func (*JobTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{4}
}
func (m *JobTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobTemplate.Unmarshal(m, b)
}
func (m *JobTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobTemplate.Marshal(b, m, deterministic)
}
func (dst *JobTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobTemplate.Merge(dst, src)
}
func (m *JobTemplate) XXX_Size() int {
	return xxx_messageInfo_JobTemplate.Size(m)
}
func (m *JobTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_JobTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_JobTemplate proto.InternalMessageInfo

func (m *JobTemplate) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *JobTemplate) GetTemplate() *JobTemplateSpec {
	if m != nil {
		return m.Template
	}
	return nil
}

// JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpec struct {
	// Standard object's metadata of the jobs created from this template.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the desired behavior of the job.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec                 *v12.JobSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *JobTemplateSpec) Reset()         { *m = JobTemplateSpec{} }
func (m *JobTemplateSpec) String() string { return proto.CompactTextString(m) }
func (*JobTemplateSpec) ProtoMessage()    {}
func (*JobTemplateSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_9718d467dfea2c0d, []int{5}
}
func (m *JobTemplateSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobTemplateSpec.Unmarshal(m, b)
}
func (m *JobTemplateSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobTemplateSpec.Marshal(b, m, deterministic)
}
func (dst *JobTemplateSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobTemplateSpec.Merge(dst, src)
}
func (m *JobTemplateSpec) XXX_Size() int {
	return xxx_messageInfo_JobTemplateSpec.Size(m)
}
func (m *JobTemplateSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_JobTemplateSpec.DiscardUnknown(m)
}

var xxx_messageInfo_JobTemplateSpec proto.InternalMessageInfo

func (m *JobTemplateSpec) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *JobTemplateSpec) GetSpec() *v12.JobSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*CronJob)(nil), "k8s.io.api.batch.v1beta1.CronJob")
	proto.RegisterType((*CronJobList)(nil), "k8s.io.api.batch.v1beta1.CronJobList")
	proto.RegisterType((*CronJobSpec)(nil), "k8s.io.api.batch.v1beta1.CronJobSpec")
	proto.RegisterType((*CronJobStatus)(nil), "k8s.io.api.batch.v1beta1.CronJobStatus")
	proto.RegisterType((*JobTemplate)(nil), "k8s.io.api.batch.v1beta1.JobTemplate")
	proto.RegisterType((*JobTemplateSpec)(nil), "k8s.io.api.batch.v1beta1.JobTemplateSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/vendor/k8s.io/api/batch/v1beta1/generated.proto", fileDescriptor_generated_9718d467dfea2c0d)
}

var fileDescriptor_generated_9718d467dfea2c0d = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xee, 0xa3, 0x9d, 0x23, 0x04, 0xf8, 0x02, 0xa2, 0x8a, 0x8b, 0x92, 0x09, 0x51,
	0x10, 0x38, 0xdb, 0x84, 0xc6, 0x10, 0x12, 0x48, 0x7c, 0x48, 0xa8, 0x14, 0x81, 0xbc, 0x89, 0x0b,
	0xee, 0x1c, 0xe7, 0x2c, 0xf5, 0x96, 0xc4, 0x91, 0xed, 0x54, 0xea, 0x63, 0xc0, 0x23, 0x20, 0xde,
	0x88, 0x17, 0xe1, 0x11, 0x90, 0xbd, 0xac, 0xdf, 0xa5, 0x43, 0xda, 0x65, 0x72, 0xfe, 0xbf, 0x93,
	0xff, 0xf9, 0x1f, 0xc7, 0xa8, 0x9f, 0x0a, 0x33, 0xa8, 0x62, 0xc2, 0x65, 0x1e, 0x69, 0x99, 0xc9,
	0xa7, 0x42, 0x46, 0xba, 0x2a, 0x41, 0xa5, 0x99, 0x94, 0xd1, 0x10, 0x8a, 0x44, 0xaa, 0xe8, 0xfc,
	0x48, 0x13, 0x21, 0x23, 0x56, 0x8a, 0x28, 0x66, 0x86, 0x0f, 0xa2, 0xe1, 0x7e, 0x0c, 0x86, 0xed,
	0x47, 0x29, 0x14, 0xa0, 0x98, 0x81, 0x84, 0x94, 0x4a, 0x1a, 0x89, 0x83, 0x0b, 0x25, 0x61, 0xa5,
	0x20, 0x4e, 0x49, 0x6a, 0x65, 0x7b, 0x77, 0x49, 0x8f, 0x79, 0xbc, 0x1d, 0x4e, 0x89, 0xb8, 0x54,
	0xb0, 0x4c, 0xf3, 0x6c, 0xa2, 0xc9, 0x19, 0x1f, 0x88, 0x02, 0xd4, 0x28, 0x2a, 0xcf, 0x53, 0xfb,
	0x42, 0x47, 0x39, 0x18, 0xb6, 0x8c, 0x8a, 0x56, 0x51, 0xaa, 0x2a, 0x8c, 0xc8, 0x61, 0x01, 0x38,
	0x5c, 0x07, 0x68, 0x3e, 0x80, 0x9c, 0xcd, 0x73, 0xe1, 0x6f, 0x0f, 0x35, 0xdf, 0x2a, 0x59, 0xf4,
	0x64, 0x8c, 0xfb, 0xa8, 0x65, 0xfd, 0x24, 0xcc, 0xb0, 0xc0, 0xeb, 0x78, 0x5d, 0xff, 0x60, 0x8f,
	0x4c, 0x02, 0x1a, 0xb7, 0x25, 0xe5, 0x79, 0x6a, 0x5f, 0x68, 0x62, 0xd5, 0x64, 0xb8, 0x4f, 0x3e,
	0xc7, 0x67, 0xc0, 0xcd, 0x27, 0x30, 0x8c, 0x8e, 0x3b, 0xe0, 0x17, 0x68, 0x53, 0x97, 0xc0, 0x83,
	0x86, 0xeb, 0xf4, 0x80, 0xac, 0x8a, 0x9a, 0xd4, 0x9f, 0x3f, 0x2e, 0x81, 0x53, 0x87, 0xe0, 0xd7,
	0x68, 0x5b, 0x1b, 0x66, 0x2a, 0x1d, 0x6c, 0x38, 0xf8, 0xe1, 0x7a, 0xd8, 0xc9, 0x69, 0x8d, 0x85,
	0x3f, 0x3c, 0xe4, 0xd7, 0x95, 0xbe, 0xd0, 0x06, 0xf7, 0x16, 0x26, 0x23, 0x57, 0x9b, 0xcc, 0xd2,
	0x73, 0x73, 0x3d, 0x47, 0x5b, 0xc2, 0x40, 0xae, 0x83, 0x46, 0x67, 0xa3, 0xeb, 0x1f, 0xdc, 0x5f,
	0xeb, 0x8d, 0x5e, 0xe8, 0xc3, 0x3f, 0x8d, 0xb1, 0x29, 0x3b, 0x2b, 0x6e, 0xa3, 0x96, 0x5d, 0x4a,
	0x52, 0x65, 0xe0, 0x4c, 0xed, 0xd0, 0xf1, 0x33, 0x3e, 0x42, 0x77, 0xb5, 0x61, 0xca, 0x88, 0x22,
	0x7d, 0x07, 0x2c, 0xc9, 0x44, 0x01, 0xc7, 0xc0, 0x65, 0x91, 0x68, 0x97, 0xe7, 0x06, 0x5d, 0x55,
	0xc6, 0x4f, 0xd0, 0x6d, 0x2e, 0x0b, 0x5e, 0x29, 0x05, 0x05, 0x1f, 0x7d, 0x91, 0x99, 0xe0, 0x23,
	0x17, 0xe3, 0x0e, 0x5d, 0x2c, 0xe0, 0x00, 0x35, 0x75, 0xa5, 0x4b, 0x28, 0x92, 0x60, 0xb3, 0xe3,
	0x75, 0x5b, 0xf4, 0xf2, 0x11, 0x7f, 0x44, 0xfe, 0x99, 0x8c, 0x4f, 0x20, 0x2f, 0x33, 0x66, 0x20,
	0xd8, 0x72, 0xa9, 0x3d, 0x5a, 0x3d, 0x6c, 0x6f, 0x22, 0x76, 0x9b, 0x9c, 0xa6, 0xf1, 0x2b, 0xd4,
	0xd6, 0x15, 0xe7, 0xa0, 0xf5, 0x69, 0x95, 0xf5, 0x64, 0xac, 0x3f, 0x08, 0x6d, 0xa4, 0x1a, 0xf5,
	0x45, 0x2e, 0x4c, 0xb0, 0xdd, 0xf1, 0xba, 0x5b, 0xf4, 0x1f, 0x0a, 0x7c, 0x88, 0xee, 0x9c, 0x32,
	0x91, 0x41, 0xb2, 0xc0, 0x36, 0x1d, 0xbb, 0xa2, 0x1a, 0xfe, 0xf2, 0xd0, 0x8d, 0x99, 0x13, 0x82,
	0x5f, 0xa2, 0x6d, 0xc6, 0x8d, 0x18, 0xda, 0xc8, 0xed, 0xfa, 0x76, 0xa7, 0x27, 0xb2, 0xff, 0xf0,
	0xe4, 0x3c, 0x53, 0x38, 0x05, 0x9b, 0x15, 0xd0, 0x1a, 0xc1, 0x5f, 0xd1, 0xad, 0x8c, 0x69, 0x73,
	0x5c, 0x6f, 0xe9, 0x44, 0xe4, 0xe0, 0x62, 0xf3, 0x0f, 0x1e, 0x5f, 0xed, 0x38, 0x59, 0x82, 0x2e,
	0xf4, 0x08, 0x7f, 0x7a, 0xc8, 0x9f, 0xca, 0xef, 0x9a, 0x7f, 0xc4, 0xf7, 0xa8, 0x65, 0x2e, 0xd7,
	0xd8, 0xf8, 0xdf, 0x35, 0x8e, 0xd1, 0xf0, 0xbb, 0x87, 0x6e, 0xce, 0x55, 0xaf, 0xd9, 0xe8, 0xde,
	0xcc, 0x8d, 0x71, 0x6f, 0x99, 0x49, 0x32, 0x73, 0x51, 0xbc, 0xd9, 0xf9, 0xd6, 0xac, 0x8d, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xef, 0x32, 0x97, 0x52, 0x19, 0x06, 0x00, 0x00,
}

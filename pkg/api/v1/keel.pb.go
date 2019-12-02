// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keel.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AnnotationFilterOp int32

const (
	AnnotationFilterOp_OP_EQUALS      AnnotationFilterOp = 0
	AnnotationFilterOp_OP_STARTS_WITH AnnotationFilterOp = 1
	AnnotationFilterOp_OP_ENDS_WITH   AnnotationFilterOp = 2
	AnnotationFilterOp_OP_CONTAINS    AnnotationFilterOp = 3
	AnnotationFilterOp_OP_HAS_KEY     AnnotationFilterOp = 4
)

var AnnotationFilterOp_name = map[int32]string{
	0: "OP_EQUALS",
	1: "OP_STARTS_WITH",
	2: "OP_ENDS_WITH",
	3: "OP_CONTAINS",
	4: "OP_HAS_KEY",
}

var AnnotationFilterOp_value = map[string]int32{
	"OP_EQUALS":      0,
	"OP_STARTS_WITH": 1,
	"OP_ENDS_WITH":   2,
	"OP_CONTAINS":    3,
	"OP_HAS_KEY":     4,
}

func (x AnnotationFilterOp) String() string {
	return proto.EnumName(AnnotationFilterOp_name, int32(x))
}

func (AnnotationFilterOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{0}
}

type JobPhase int32

const (
	// Unknown means we don't know what state the job is in
	JobPhase_PHASE_UNKNOWN JobPhase = 0
	// Preparing means the job hasn't started yet and isn't consuming resources in the system
	JobPhase_PHASE_PREPARING JobPhase = 1
	// Starting means the job has been scheduled and is waiting to run. Things that might prevent it
	// from running already are pod scheduling, image pull or container startup.
	JobPhase_PHASE_STARTING JobPhase = 2
	// Running means the job is actually running and doing work.
	JobPhase_PHASE_RUNNING JobPhase = 3
	// Done means the job has run and is finished
	JobPhase_PHASE_DONE JobPhase = 4
	// Cleaning means the job is in post-run cleanup
	JobPhase_PHASE_CLEANUP JobPhase = 5
)

var JobPhase_name = map[int32]string{
	0: "PHASE_UNKNOWN",
	1: "PHASE_PREPARING",
	2: "PHASE_STARTING",
	3: "PHASE_RUNNING",
	4: "PHASE_DONE",
	5: "PHASE_CLEANUP",
}

var JobPhase_value = map[string]int32{
	"PHASE_UNKNOWN":   0,
	"PHASE_PREPARING": 1,
	"PHASE_STARTING":  2,
	"PHASE_RUNNING":   3,
	"PHASE_DONE":      4,
	"PHASE_CLEANUP":   5,
}

func (x JobPhase) String() string {
	return proto.EnumName(JobPhase_name, int32(x))
}

func (JobPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{1}
}

type LogSlicePhase int32

const (
	LogSlicePhase_SLICE_ABANDONED  LogSlicePhase = 0
	LogSlicePhase_SLICE_CHECKPOINT LogSlicePhase = 1
	LogSlicePhase_SLICE_START      LogSlicePhase = 2
	LogSlicePhase_SLICE_CONTENT    LogSlicePhase = 3
	LogSlicePhase_SLICE_END        LogSlicePhase = 4
)

var LogSlicePhase_name = map[int32]string{
	0: "SLICE_ABANDONED",
	1: "SLICE_CHECKPOINT",
	2: "SLICE_START",
	3: "SLICE_CONTENT",
	4: "SLICE_END",
}

var LogSlicePhase_value = map[string]int32{
	"SLICE_ABANDONED":  0,
	"SLICE_CHECKPOINT": 1,
	"SLICE_START":      2,
	"SLICE_CONTENT":    3,
	"SLICE_END":        4,
}

func (x LogSlicePhase) String() string {
	return proto.EnumName(LogSlicePhase_name, int32(x))
}

func (LogSlicePhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{2}
}

type ListJobsRequest struct {
	Filter               []*AnnotationFilter `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	RunningOnly          bool                `protobuf:"varint,2,opt,name=running_only,json=runningOnly,proto3" json:"running_only,omitempty"`
	Start                int32               `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	Limit                int32               `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{0}
}

func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (m *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(m, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

func (m *ListJobsRequest) GetFilter() []*AnnotationFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListJobsRequest) GetRunningOnly() bool {
	if m != nil {
		return m.RunningOnly
	}
	return false
}

func (m *ListJobsRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListJobsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type AnnotationFilter struct {
	Terms                []*AnnotationFilterTerm `protobuf:"bytes,1,rep,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AnnotationFilter) Reset()         { *m = AnnotationFilter{} }
func (m *AnnotationFilter) String() string { return proto.CompactTextString(m) }
func (*AnnotationFilter) ProtoMessage()    {}
func (*AnnotationFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{1}
}

func (m *AnnotationFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationFilter.Unmarshal(m, b)
}
func (m *AnnotationFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationFilter.Marshal(b, m, deterministic)
}
func (m *AnnotationFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationFilter.Merge(m, src)
}
func (m *AnnotationFilter) XXX_Size() int {
	return xxx_messageInfo_AnnotationFilter.Size(m)
}
func (m *AnnotationFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationFilter.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationFilter proto.InternalMessageInfo

func (m *AnnotationFilter) GetTerms() []*AnnotationFilterTerm {
	if m != nil {
		return m.Terms
	}
	return nil
}

type AnnotationFilterTerm struct {
	Annotation           string             `protobuf:"bytes,1,opt,name=annotation,proto3" json:"annotation,omitempty"`
	Value                string             `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Operation            AnnotationFilterOp `protobuf:"varint,3,opt,name=operation,proto3,enum=v1.AnnotationFilterOp" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AnnotationFilterTerm) Reset()         { *m = AnnotationFilterTerm{} }
func (m *AnnotationFilterTerm) String() string { return proto.CompactTextString(m) }
func (*AnnotationFilterTerm) ProtoMessage()    {}
func (*AnnotationFilterTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{2}
}

func (m *AnnotationFilterTerm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationFilterTerm.Unmarshal(m, b)
}
func (m *AnnotationFilterTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationFilterTerm.Marshal(b, m, deterministic)
}
func (m *AnnotationFilterTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationFilterTerm.Merge(m, src)
}
func (m *AnnotationFilterTerm) XXX_Size() int {
	return xxx_messageInfo_AnnotationFilterTerm.Size(m)
}
func (m *AnnotationFilterTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationFilterTerm.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationFilterTerm proto.InternalMessageInfo

func (m *AnnotationFilterTerm) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

func (m *AnnotationFilterTerm) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AnnotationFilterTerm) GetOperation() AnnotationFilterOp {
	if m != nil {
		return m.Operation
	}
	return AnnotationFilterOp_OP_EQUALS
}

type ListJobsResponse struct {
	Total                int32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Result               []*JobStatus `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{3}
}

func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (m *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(m, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

func (m *ListJobsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListJobsResponse) GetResult() []*JobStatus {
	if m != nil {
		return m.Result
	}
	return nil
}

type ListenRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Updates              bool     `protobuf:"varint,2,opt,name=updates,proto3" json:"updates,omitempty"`
	Logs                 bool     `protobuf:"varint,3,opt,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenRequest) Reset()         { *m = ListenRequest{} }
func (m *ListenRequest) String() string { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()    {}
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{4}
}

func (m *ListenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenRequest.Unmarshal(m, b)
}
func (m *ListenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenRequest.Marshal(b, m, deterministic)
}
func (m *ListenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenRequest.Merge(m, src)
}
func (m *ListenRequest) XXX_Size() int {
	return xxx_messageInfo_ListenRequest.Size(m)
}
func (m *ListenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenRequest proto.InternalMessageInfo

func (m *ListenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenRequest) GetUpdates() bool {
	if m != nil {
		return m.Updates
	}
	return false
}

func (m *ListenRequest) GetLogs() bool {
	if m != nil {
		return m.Logs
	}
	return false
}

type ListenResponse struct {
	// Types that are valid to be assigned to Content:
	//	*ListenResponse_Update
	//	*ListenResponse_Slice
	Content              isListenResponse_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListenResponse) Reset()         { *m = ListenResponse{} }
func (m *ListenResponse) String() string { return proto.CompactTextString(m) }
func (*ListenResponse) ProtoMessage()    {}
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{5}
}

func (m *ListenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenResponse.Unmarshal(m, b)
}
func (m *ListenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenResponse.Marshal(b, m, deterministic)
}
func (m *ListenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenResponse.Merge(m, src)
}
func (m *ListenResponse) XXX_Size() int {
	return xxx_messageInfo_ListenResponse.Size(m)
}
func (m *ListenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenResponse proto.InternalMessageInfo

type isListenResponse_Content interface {
	isListenResponse_Content()
}

type ListenResponse_Update struct {
	Update *JobStatus `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type ListenResponse_Slice struct {
	Slice *LogSliceEvent `protobuf:"bytes,2,opt,name=slice,proto3,oneof"`
}

func (*ListenResponse_Update) isListenResponse_Content() {}

func (*ListenResponse_Slice) isListenResponse_Content() {}

func (m *ListenResponse) GetContent() isListenResponse_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ListenResponse) GetUpdate() *JobStatus {
	if x, ok := m.GetContent().(*ListenResponse_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ListenResponse) GetSlice() *LogSliceEvent {
	if x, ok := m.GetContent().(*ListenResponse_Slice); ok {
		return x.Slice
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenResponse_Update)(nil),
		(*ListenResponse_Slice)(nil),
	}
}

type JobStatus struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Metadata             *JobMetadata   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Phase                JobPhase       `protobuf:"varint,3,opt,name=phase,proto3,enum=v1.JobPhase" json:"phase,omitempty"`
	Conditions           *JobConditions `protobuf:"bytes,4,opt,name=conditions,proto3" json:"conditions,omitempty"`
	Details              string         `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}
func (*JobStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{6}
}

func (m *JobStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatus.Unmarshal(m, b)
}
func (m *JobStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatus.Marshal(b, m, deterministic)
}
func (m *JobStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatus.Merge(m, src)
}
func (m *JobStatus) XXX_Size() int {
	return xxx_messageInfo_JobStatus.Size(m)
}
func (m *JobStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatus.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatus proto.InternalMessageInfo

func (m *JobStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobStatus) GetMetadata() *JobMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *JobStatus) GetPhase() JobPhase {
	if m != nil {
		return m.Phase
	}
	return JobPhase_PHASE_UNKNOWN
}

func (m *JobStatus) GetConditions() *JobConditions {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *JobStatus) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type JobMetadata struct {
	Annotations          []*Annotation `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JobMetadata) Reset()         { *m = JobMetadata{} }
func (m *JobMetadata) String() string { return proto.CompactTextString(m) }
func (*JobMetadata) ProtoMessage()    {}
func (*JobMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{7}
}

func (m *JobMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobMetadata.Unmarshal(m, b)
}
func (m *JobMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobMetadata.Marshal(b, m, deterministic)
}
func (m *JobMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobMetadata.Merge(m, src)
}
func (m *JobMetadata) XXX_Size() int {
	return xxx_messageInfo_JobMetadata.Size(m)
}
func (m *JobMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_JobMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_JobMetadata proto.InternalMessageInfo

func (m *JobMetadata) GetAnnotations() []*Annotation {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Annotation struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Annotation) Reset()         { *m = Annotation{} }
func (m *Annotation) String() string { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()    {}
func (*Annotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{8}
}

func (m *Annotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Annotation.Unmarshal(m, b)
}
func (m *Annotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Annotation.Marshal(b, m, deterministic)
}
func (m *Annotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Annotation.Merge(m, src)
}
func (m *Annotation) XXX_Size() int {
	return xxx_messageInfo_Annotation.Size(m)
}
func (m *Annotation) XXX_DiscardUnknown() {
	xxx_messageInfo_Annotation.DiscardUnknown(m)
}

var xxx_messageInfo_Annotation proto.InternalMessageInfo

func (m *Annotation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Annotation) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type JobConditions struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FailureCount         int32    `protobuf:"varint,2,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobConditions) Reset()         { *m = JobConditions{} }
func (m *JobConditions) String() string { return proto.CompactTextString(m) }
func (*JobConditions) ProtoMessage()    {}
func (*JobConditions) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{9}
}

func (m *JobConditions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobConditions.Unmarshal(m, b)
}
func (m *JobConditions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobConditions.Marshal(b, m, deterministic)
}
func (m *JobConditions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobConditions.Merge(m, src)
}
func (m *JobConditions) XXX_Size() int {
	return xxx_messageInfo_JobConditions.Size(m)
}
func (m *JobConditions) XXX_DiscardUnknown() {
	xxx_messageInfo_JobConditions.DiscardUnknown(m)
}

var xxx_messageInfo_JobConditions proto.InternalMessageInfo

func (m *JobConditions) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *JobConditions) GetFailureCount() int32 {
	if m != nil {
		return m.FailureCount
	}
	return 0
}

type LogSliceEvent struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phase                LogSlicePhase `protobuf:"varint,2,opt,name=phase,proto3,enum=v1.LogSlicePhase" json:"phase,omitempty"`
	Payload              string        `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogSliceEvent) Reset()         { *m = LogSliceEvent{} }
func (m *LogSliceEvent) String() string { return proto.CompactTextString(m) }
func (*LogSliceEvent) ProtoMessage()    {}
func (*LogSliceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc1f325416a5753, []int{10}
}

func (m *LogSliceEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogSliceEvent.Unmarshal(m, b)
}
func (m *LogSliceEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogSliceEvent.Marshal(b, m, deterministic)
}
func (m *LogSliceEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogSliceEvent.Merge(m, src)
}
func (m *LogSliceEvent) XXX_Size() int {
	return xxx_messageInfo_LogSliceEvent.Size(m)
}
func (m *LogSliceEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogSliceEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogSliceEvent proto.InternalMessageInfo

func (m *LogSliceEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogSliceEvent) GetPhase() LogSlicePhase {
	if m != nil {
		return m.Phase
	}
	return LogSlicePhase_SLICE_ABANDONED
}

func (m *LogSliceEvent) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterEnum("v1.AnnotationFilterOp", AnnotationFilterOp_name, AnnotationFilterOp_value)
	proto.RegisterEnum("v1.JobPhase", JobPhase_name, JobPhase_value)
	proto.RegisterEnum("v1.LogSlicePhase", LogSlicePhase_name, LogSlicePhase_value)
	proto.RegisterType((*ListJobsRequest)(nil), "v1.ListJobsRequest")
	proto.RegisterType((*AnnotationFilter)(nil), "v1.AnnotationFilter")
	proto.RegisterType((*AnnotationFilterTerm)(nil), "v1.AnnotationFilterTerm")
	proto.RegisterType((*ListJobsResponse)(nil), "v1.ListJobsResponse")
	proto.RegisterType((*ListenRequest)(nil), "v1.ListenRequest")
	proto.RegisterType((*ListenResponse)(nil), "v1.ListenResponse")
	proto.RegisterType((*JobStatus)(nil), "v1.JobStatus")
	proto.RegisterType((*JobMetadata)(nil), "v1.JobMetadata")
	proto.RegisterType((*Annotation)(nil), "v1.Annotation")
	proto.RegisterType((*JobConditions)(nil), "v1.JobConditions")
	proto.RegisterType((*LogSliceEvent)(nil), "v1.LogSliceEvent")
}

func init() { proto.RegisterFile("keel.proto", fileDescriptor_cfc1f325416a5753) }

var fileDescriptor_cfc1f325416a5753 = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x93, 0x38, 0x4d, 0x4e, 0x9a, 0xc4, 0x9d, 0x46, 0xc8, 0xda, 0x0b, 0x54, 0x8c, 0xd0,
	0x96, 0x82, 0xaa, 0x6d, 0x77, 0x25, 0x2e, 0x91, 0x9b, 0x1a, 0x92, 0x36, 0xd8, 0xde, 0x71, 0xaa,
	0x15, 0x57, 0x96, 0x9b, 0x4c, 0x8b, 0x59, 0xc7, 0x13, 0x3c, 0xe3, 0x48, 0x15, 0x5c, 0xf1, 0x00,
	0xbc, 0x11, 0xef, 0x86, 0xe6, 0xc7, 0xf9, 0x29, 0xb9, 0x9b, 0xf3, 0x9d, 0xef, 0x9c, 0xf9, 0xce,
	0x8f, 0xc7, 0x00, 0x9f, 0x09, 0xc9, 0x2e, 0x57, 0x05, 0xe5, 0x14, 0xd5, 0xd7, 0x57, 0xce, 0x3f,
	0x06, 0x0c, 0xa6, 0x29, 0xe3, 0x77, 0xf4, 0x91, 0x61, 0xf2, 0x47, 0x49, 0x18, 0x47, 0xdf, 0x43,
	0xeb, 0x29, 0xcd, 0x38, 0x29, 0x6c, 0xe3, 0xac, 0x71, 0xde, 0xbd, 0x1e, 0x5e, 0xae, 0xaf, 0x2e,
	0xdd, 0x3c, 0xa7, 0x3c, 0xe1, 0x29, 0xcd, 0x7f, 0x92, 0x3e, 0xac, 0x39, 0xe8, 0x2b, 0x38, 0x2e,
	0xca, 0x3c, 0x4f, 0xf3, 0xe7, 0x98, 0xe6, 0xd9, 0x8b, 0x5d, 0x3f, 0x33, 0xce, 0xdb, 0xb8, 0xab,
	0xb1, 0x20, 0xcf, 0x5e, 0xd0, 0x10, 0x4c, 0xc6, 0x93, 0x82, 0xdb, 0x8d, 0x33, 0xe3, 0xdc, 0xc4,
	0xca, 0x10, 0x68, 0x96, 0x2e, 0x53, 0x6e, 0x37, 0x15, 0x2a, 0x0d, 0xe7, 0x06, 0xac, 0xd7, 0x57,
	0xa1, 0x4b, 0x30, 0x39, 0x29, 0x96, 0x4c, 0xeb, 0xb1, 0x0f, 0xe9, 0x99, 0x91, 0x62, 0x89, 0x15,
	0xcd, 0xf9, 0xdb, 0x80, 0xe1, 0x21, 0x3f, 0xfa, 0x12, 0x20, 0xd9, 0xe0, 0xb6, 0x71, 0x66, 0x9c,
	0x77, 0xf0, 0x0e, 0x22, 0x24, 0xad, 0x93, 0xac, 0x24, 0xb2, 0x88, 0x0e, 0x56, 0x06, 0xfa, 0x00,
	0x1d, 0xba, 0x22, 0x85, 0x0a, 0x12, 0x25, 0xf4, 0xaf, 0xbf, 0x38, 0x24, 0x21, 0x58, 0xe1, 0x2d,
	0xd1, 0x09, 0xc0, 0xda, 0x36, 0x96, 0xad, 0x68, 0xce, 0x88, 0xc8, 0xcf, 0x29, 0x4f, 0x32, 0x79,
	0xb5, 0x89, 0x95, 0x81, 0xbe, 0x81, 0x56, 0x41, 0x58, 0x99, 0x71, 0xbb, 0x2e, 0xeb, 0xeb, 0x89,
	0xe4, 0x77, 0xf4, 0x31, 0xe2, 0x09, 0x2f, 0x19, 0xd6, 0x4e, 0xe7, 0x23, 0xf4, 0x44, 0x42, 0x92,
	0x57, 0x73, 0x42, 0xd0, 0xcc, 0x93, 0x25, 0xd1, 0x75, 0xc8, 0x33, 0xb2, 0xe1, 0xa8, 0x5c, 0x2d,
	0x12, 0x4e, 0x98, 0x1e, 0x44, 0x65, 0x0a, 0x76, 0x46, 0x9f, 0x99, 0x2c, 0xa0, 0x8d, 0xe5, 0xd9,
	0xa1, 0xd0, 0xaf, 0x52, 0x6a, 0x85, 0x6f, 0xa1, 0xa5, 0x02, 0x64, 0xd6, 0xd7, 0x5a, 0xc6, 0x35,
	0xac, 0xdd, 0xe8, 0x5b, 0x30, 0x59, 0x96, 0xce, 0x55, 0xab, 0xba, 0xd7, 0x27, 0x82, 0x37, 0xa5,
	0xcf, 0x91, 0xc0, 0xbc, 0x35, 0xc9, 0xf9, 0xb8, 0x86, 0x15, 0xe3, 0xa6, 0x03, 0x47, 0x73, 0x9a,
	0x73, 0x92, 0x73, 0xe7, 0x5f, 0x03, 0x3a, 0x9b, 0x6c, 0x07, 0x0b, 0xf8, 0x0e, 0xda, 0x4b, 0xc2,
	0x93, 0x45, 0xc2, 0x13, 0x9d, 0x7a, 0xa0, 0x25, 0xfc, 0xa2, 0x61, 0xbc, 0x21, 0x20, 0x07, 0xcc,
	0xd5, 0x6f, 0x09, 0x23, 0x7a, 0x2a, 0xc7, 0x9a, 0x19, 0x0a, 0x0c, 0x2b, 0x17, 0xba, 0x02, 0x98,
	0xd3, 0x7c, 0x91, 0x8a, 0xa1, 0x30, 0xb9, 0x6b, 0x5a, 0xed, 0x1d, 0x7d, 0x1c, 0x6d, 0x1c, 0x78,
	0x87, 0x24, 0x9a, 0xb8, 0x20, 0x3c, 0x49, 0x33, 0x66, 0x9b, 0x52, 0x5a, 0x65, 0x3a, 0x3f, 0x42,
	0x77, 0x47, 0x09, 0x7a, 0x07, 0xdd, 0xed, 0xf6, 0x54, 0xeb, 0xd9, 0xdf, 0xdf, 0x0d, 0xbc, 0x4b,
	0x71, 0x3e, 0x00, 0x6c, 0x5d, 0xc8, 0x82, 0xc6, 0x67, 0xf2, 0xa2, 0xeb, 0x17, 0xc7, 0xc3, 0x1b,
	0xe8, 0xf8, 0xd0, 0xdb, 0x53, 0x2b, 0x14, 0xb2, 0x72, 0x3e, 0x27, 0x8c, 0xc9, 0xe0, 0x36, 0xae,
	0x4c, 0xf4, 0x35, 0xf4, 0x9e, 0x92, 0x34, 0x2b, 0x0b, 0x12, 0xcf, 0x69, 0x99, 0x73, 0x99, 0xc8,
	0xc4, 0xc7, 0x1a, 0x1c, 0x09, 0xcc, 0x79, 0x82, 0xde, 0xde, 0xac, 0x0e, 0x4e, 0xe2, 0x6d, 0xd5,
	0xdc, 0xba, 0x6c, 0xee, 0xde, 0x84, 0xf7, 0x3a, 0x6c, 0xc3, 0xd1, 0x2a, 0x79, 0xc9, 0x68, 0xb2,
	0x90, 0x73, 0xe8, 0xe0, 0xca, 0xbc, 0xf8, 0x1d, 0xd0, 0xff, 0x3f, 0x12, 0xd4, 0x83, 0x4e, 0x10,
	0xc6, 0xde, 0xc7, 0x07, 0x77, 0x1a, 0x59, 0x35, 0x84, 0xa0, 0x1f, 0x84, 0x71, 0x34, 0x73, 0xf1,
	0x2c, 0x8a, 0x3f, 0x4d, 0x66, 0x63, 0xcb, 0x40, 0x16, 0x1c, 0x0b, 0x8a, 0x7f, 0xab, 0x91, 0x3a,
	0x1a, 0x40, 0x37, 0x08, 0xe3, 0x51, 0xe0, 0xcf, 0xdc, 0x89, 0x1f, 0x59, 0x0d, 0xd4, 0x07, 0x08,
	0xc2, 0x78, 0xec, 0x46, 0xf1, 0xbd, 0xf7, 0xab, 0xd5, 0xbc, 0xf8, 0x0b, 0xda, 0xd5, 0xe8, 0xd1,
	0x09, 0xf4, 0xc2, 0xb1, 0x1b, 0x79, 0xf1, 0x83, 0x7f, 0xef, 0x07, 0x9f, 0x7c, 0xab, 0x86, 0x4e,
	0x61, 0xa0, 0xa0, 0x10, 0x7b, 0xa1, 0x8b, 0x27, 0xfe, 0xcf, 0x96, 0x21, 0xae, 0x56, 0xa0, 0xbc,
	0x5d, 0x60, 0xf5, 0x6d, 0x2c, 0x7e, 0xf0, 0x7d, 0x01, 0xc9, 0xab, 0x14, 0x74, 0x1b, 0xf8, 0x9e,
	0xd5, 0xdc, 0x52, 0x46, 0x53, 0xcf, 0xf5, 0x1f, 0x42, 0xcb, 0xbc, 0x58, 0x6e, 0x3b, 0xaa, 0x24,
	0x9c, 0xc2, 0x20, 0x9a, 0x4e, 0x46, 0x5e, 0xec, 0xde, 0xb8, 0xbe, 0x88, 0xbb, 0xb5, 0x6a, 0x68,
	0x08, 0x96, 0x02, 0x47, 0x63, 0x6f, 0x74, 0x1f, 0x06, 0x13, 0x7f, 0x66, 0x19, 0xa2, 0x34, 0x85,
	0x4a, 0x15, 0x4a, 0x82, 0xa6, 0x05, 0xfe, 0xcc, 0xf3, 0x67, 0x56, 0x43, 0xf4, 0x4c, 0x41, 0x9e,
	0x7f, 0x6b, 0x35, 0xaf, 0xff, 0x84, 0xee, 0x3d, 0x21, 0x59, 0x44, 0x8a, 0x75, 0x3a, 0x27, 0xe8,
	0x07, 0x68, 0x57, 0x6f, 0x0d, 0x3a, 0x95, 0x73, 0xda, 0x7f, 0xd2, 0xdf, 0x0c, 0xf7, 0x41, 0xf5,
	0xb1, 0x3b, 0x35, 0xf4, 0x1e, 0x5a, 0xea, 0x01, 0x40, 0x27, 0x15, 0x63, 0xf3, 0xbe, 0xbc, 0x41,
	0xbb, 0x50, 0x15, 0xf2, 0xce, 0x78, 0x6c, 0xc9, 0xdf, 0xc7, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x29, 0xf3, 0x2b, 0xde, 0x4c, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeelServiceClient is the client API for KeelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeelServiceClient interface {
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (KeelService_ListenClient, error)
}

type keelServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeelServiceClient(cc *grpc.ClientConn) KeelServiceClient {
	return &keelServiceClient{cc}
}

func (c *keelServiceClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, "/v1.KeelService/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keelServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (KeelService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeelService_serviceDesc.Streams[0], "/v1.KeelService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &keelServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeelService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type keelServiceListenClient struct {
	grpc.ClientStream
}

func (x *keelServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeelServiceServer is the server API for KeelService service.
type KeelServiceServer interface {
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	Listen(*ListenRequest, KeelService_ListenServer) error
}

// UnimplementedKeelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeelServiceServer struct {
}

func (*UnimplementedKeelServiceServer) ListJobs(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (*UnimplementedKeelServiceServer) Listen(req *ListenRequest, srv KeelService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}

func RegisterKeelServiceServer(s *grpc.Server, srv KeelServiceServer) {
	s.RegisterService(&_KeelService_serviceDesc, srv)
}

func _KeelService_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeelServiceServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.KeelService/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeelServiceServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeelService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeelServiceServer).Listen(m, &keelServiceListenServer{stream})
}

type KeelService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type keelServiceListenServer struct {
	grpc.ServerStream
}

func (x *keelServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _KeelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.KeelService",
	HandlerType: (*KeelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListJobs",
			Handler:    _KeelService_ListJobs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _KeelService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "keel.proto",
}

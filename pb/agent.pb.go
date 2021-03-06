// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/agent.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// find or create group
type FCGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int64  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FCGroupReq) Reset() {
	*x = FCGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGroupReq) ProtoMessage() {}

func (x *FCGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCGroupReq.ProtoReflect.Descriptor instead.
func (*FCGroupReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{0}
}

func (x *FCGroupReq) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *FCGroupReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FCGroupRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FCGroupRes) Reset() {
	*x = FCGroupRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGroupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGroupRes) ProtoMessage() {}

func (x *FCGroupRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCGroupRes.ProtoReflect.Descriptor instead.
func (*FCGroupRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{1}
}

func (x *FCGroupRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// find or create graph
type FCGraphReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   int64  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Unit    string `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	GroupID int64  `protobuf:"varint,4,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *FCGraphReq) Reset() {
	*x = FCGraphReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGraphReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGraphReq) ProtoMessage() {}

func (x *FCGraphReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCGraphReq.ProtoReflect.Descriptor instead.
func (*FCGraphReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{2}
}

func (x *FCGraphReq) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *FCGraphReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FCGraphReq) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *FCGraphReq) GetGroupID() int64 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

type FCGraphRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FCGraphRes) Reset() {
	*x = FCGraphRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGraphRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGraphRes) ProtoMessage() {}

func (x *FCGraphRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCGraphRes.ProtoReflect.Descriptor instead.
func (*FCGraphRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{3}
}

func (x *FCGraphRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// find or create metric
type FCMetricReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   int64  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	GraphID int64  `protobuf:"varint,4,opt,name=graphID,proto3" json:"graphID,omitempty"`
}

func (x *FCMetricReq) Reset() {
	*x = FCMetricReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCMetricReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCMetricReq) ProtoMessage() {}

func (x *FCMetricReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCMetricReq.ProtoReflect.Descriptor instead.
func (*FCMetricReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{4}
}

func (x *FCMetricReq) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *FCMetricReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FCMetricReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FCMetricReq) GetGraphID() int64 {
	if x != nil {
		return x.GraphID
	}
	return 0
}

type FCMetricRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FCMetricRes) Reset() {
	*x = FCMetricRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCMetricRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCMetricRes) ProtoMessage() {}

func (x *FCMetricRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FCMetricRes.ProtoReflect.Descriptor instead.
func (*FCMetricRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{5}
}

func (x *FCMetricRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// histogram, counter, gauge
type BasedReqMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int64  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"` // app ID
	EID   string `protobuf:"bytes,2,opt,name=eID,proto3" json:"eID,omitempty"`      // executor ID
	MID   int64  `protobuf:"varint,3,opt,name=mID,proto3" json:"mID,omitempty"`     // metric ID
	Time  int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *BasedReqMetric) Reset() {
	*x = BasedReqMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasedReqMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasedReqMetric) ProtoMessage() {}

func (x *BasedReqMetric) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasedReqMetric.ProtoReflect.Descriptor instead.
func (*BasedReqMetric) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{6}
}

func (x *BasedReqMetric) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *BasedReqMetric) GetEID() string {
	if x != nil {
		return x.EID
	}
	return ""
}

func (x *BasedReqMetric) GetMID() int64 {
	if x != nil {
		return x.MID
	}
	return 0
}

func (x *BasedReqMetric) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type HistogramValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Min    int64   `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max    int64   `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Mean   float64 `protobuf:"fixed64,4,opt,name=mean,proto3" json:"mean,omitempty"`
	Stddev float64 `protobuf:"fixed64,5,opt,name=stddev,proto3" json:"stddev,omitempty"`
	Median float64 `protobuf:"fixed64,6,opt,name=median,proto3" json:"median,omitempty"`
	P75    float64 `protobuf:"fixed64,7,opt,name=p75,proto3" json:"p75,omitempty"`
	P95    float64 `protobuf:"fixed64,8,opt,name=p95,proto3" json:"p95,omitempty"`
	P99    float64 `protobuf:"fixed64,9,opt,name=p99,proto3" json:"p99,omitempty"`
	P999   float64 `protobuf:"fixed64,10,opt,name=p999,proto3" json:"p999,omitempty"`
}

func (x *HistogramValues) Reset() {
	*x = HistogramValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramValues) ProtoMessage() {}

func (x *HistogramValues) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramValues.ProtoReflect.Descriptor instead.
func (*HistogramValues) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{7}
}

func (x *HistogramValues) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *HistogramValues) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *HistogramValues) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *HistogramValues) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *HistogramValues) GetStddev() float64 {
	if x != nil {
		return x.Stddev
	}
	return 0
}

func (x *HistogramValues) GetMedian() float64 {
	if x != nil {
		return x.Median
	}
	return 0
}

func (x *HistogramValues) GetP75() float64 {
	if x != nil {
		return x.P75
	}
	return 0
}

func (x *HistogramValues) GetP95() float64 {
	if x != nil {
		return x.P95
	}
	return 0
}

func (x *HistogramValues) GetP99() float64 {
	if x != nil {
		return x.P99
	}
	return 0
}

func (x *HistogramValues) GetP999() float64 {
	if x != nil {
		return x.P999
	}
	return 0
}

type HistogramReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *BasedReqMetric  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Histogram *HistogramValues `protobuf:"bytes,2,opt,name=histogram,proto3" json:"histogram,omitempty"`
}

func (x *HistogramReq) Reset() {
	*x = HistogramReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramReq) ProtoMessage() {}

func (x *HistogramReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramReq.ProtoReflect.Descriptor instead.
func (*HistogramReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{8}
}

func (x *HistogramReq) GetBase() *BasedReqMetric {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *HistogramReq) GetHistogram() *HistogramValues {
	if x != nil {
		return x.Histogram
	}
	return nil
}

type HistogramRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HistogramRes) Reset() {
	*x = HistogramRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramRes) ProtoMessage() {}

func (x *HistogramRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramRes.ProtoReflect.Descriptor instead.
func (*HistogramRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{9}
}

type CounterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *BasedReqMetric `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Count int64           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CounterReq) Reset() {
	*x = CounterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterReq) ProtoMessage() {}

func (x *CounterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterReq.ProtoReflect.Descriptor instead.
func (*CounterReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{10}
}

func (x *CounterReq) GetBase() *BasedReqMetric {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CounterReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CounterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CounterRes) Reset() {
	*x = CounterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterRes) ProtoMessage() {}

func (x *CounterRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterRes.ProtoReflect.Descriptor instead.
func (*CounterRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{11}
}

type GaugeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *BasedReqMetric `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Gauge int64           `protobuf:"varint,2,opt,name=gauge,proto3" json:"gauge,omitempty"`
}

func (x *GaugeReq) Reset() {
	*x = GaugeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeReq) ProtoMessage() {}

func (x *GaugeReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeReq.ProtoReflect.Descriptor instead.
func (*GaugeReq) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{12}
}

func (x *GaugeReq) GetBase() *BasedReqMetric {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *GaugeReq) GetGauge() int64 {
	if x != nil {
		return x.Gauge
	}
	return 0
}

type GaugeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GaugeRes) Reset() {
	*x = GaugeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_agent_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeRes) ProtoMessage() {}

func (x *GaugeRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_agent_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeRes.ProtoReflect.Descriptor instead.
func (*GaugeRes) Descriptor() ([]byte, []int) {
	return file_pb_agent_proto_rawDescGZIP(), []int{13}
}

var File_pb_agent_proto protoreflect.FileDescriptor

var file_pb_agent_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x36, 0x0a, 0x0a, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x0a,
	0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x0a, 0x46, 0x43,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x22, 0x1c, 0x0a, 0x0a, 0x46, 0x43, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x67, 0x0a, 0x0b, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x43, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x64, 0x64, 0x65, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x74, 0x64, 0x64,
	0x65, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x37,
	0x35, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x37, 0x35, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x39, 0x35, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x39, 0x35, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x39, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x39, 0x39,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x39, 0x39, 0x39, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x70, 0x39, 0x39, 0x39, 0x22, 0x69, 0x0a, 0x0c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22,
	0x0e, 0x0a, 0x0c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22,
	0x4a, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x08, 0x47, 0x61, 0x75,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x32,
	0xa4, 0x02, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0f, 0x46, 0x69, 0x6e,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_agent_proto_rawDescOnce sync.Once
	file_pb_agent_proto_rawDescData = file_pb_agent_proto_rawDesc
)

func file_pb_agent_proto_rawDescGZIP() []byte {
	file_pb_agent_proto_rawDescOnce.Do(func() {
		file_pb_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_agent_proto_rawDescData)
	})
	return file_pb_agent_proto_rawDescData
}

var file_pb_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pb_agent_proto_goTypes = []interface{}{
	(*FCGroupReq)(nil),      // 0: pb.FCGroupReq
	(*FCGroupRes)(nil),      // 1: pb.FCGroupRes
	(*FCGraphReq)(nil),      // 2: pb.FCGraphReq
	(*FCGraphRes)(nil),      // 3: pb.FCGraphRes
	(*FCMetricReq)(nil),     // 4: pb.FCMetricReq
	(*FCMetricRes)(nil),     // 5: pb.FCMetricRes
	(*BasedReqMetric)(nil),  // 6: pb.BasedReqMetric
	(*HistogramValues)(nil), // 7: pb.HistogramValues
	(*HistogramReq)(nil),    // 8: pb.HistogramReq
	(*HistogramRes)(nil),    // 9: pb.HistogramRes
	(*CounterReq)(nil),      // 10: pb.CounterReq
	(*CounterRes)(nil),      // 11: pb.CounterRes
	(*GaugeReq)(nil),        // 12: pb.GaugeReq
	(*GaugeRes)(nil),        // 13: pb.GaugeRes
}
var file_pb_agent_proto_depIdxs = []int32{
	6,  // 0: pb.HistogramReq.base:type_name -> pb.BasedReqMetric
	7,  // 1: pb.HistogramReq.histogram:type_name -> pb.HistogramValues
	6,  // 2: pb.CounterReq.base:type_name -> pb.BasedReqMetric
	6,  // 3: pb.GaugeReq.base:type_name -> pb.BasedReqMetric
	0,  // 4: pb.Agent.FindCreateGroup:input_type -> pb.FCGroupReq
	2,  // 5: pb.Agent.FindCreateGraph:input_type -> pb.FCGraphReq
	4,  // 6: pb.Agent.FindCreateMetric:input_type -> pb.FCMetricReq
	8,  // 7: pb.Agent.Histogram:input_type -> pb.HistogramReq
	10, // 8: pb.Agent.Counter:input_type -> pb.CounterReq
	12, // 9: pb.Agent.Gauge:input_type -> pb.GaugeReq
	1,  // 10: pb.Agent.FindCreateGroup:output_type -> pb.FCGroupRes
	3,  // 11: pb.Agent.FindCreateGraph:output_type -> pb.FCGraphRes
	5,  // 12: pb.Agent.FindCreateMetric:output_type -> pb.FCMetricRes
	9,  // 13: pb.Agent.Histogram:output_type -> pb.HistogramRes
	11, // 14: pb.Agent.Counter:output_type -> pb.CounterRes
	13, // 15: pb.Agent.Gauge:output_type -> pb.GaugeRes
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pb_agent_proto_init() }
func file_pb_agent_proto_init() {
	if File_pb_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCGroupReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCGroupRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCGraphReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCGraphRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCMetricReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FCMetricRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasedReqMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramValues); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_agent_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_agent_proto_goTypes,
		DependencyIndexes: file_pb_agent_proto_depIdxs,
		MessageInfos:      file_pb_agent_proto_msgTypes,
	}.Build()
	File_pb_agent_proto = out.File
	file_pb_agent_proto_rawDesc = nil
	file_pb_agent_proto_goTypes = nil
	file_pb_agent_proto_depIdxs = nil
}

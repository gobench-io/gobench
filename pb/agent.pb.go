// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.8.0
// source: agent.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGroupReq) ProtoMessage() {}

func (x *FCGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
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
	return file_agent_proto_rawDescGZIP(), []int{0}
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
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGroupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGroupRes) ProtoMessage() {}

func (x *FCGroupRes) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
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
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *FCGroupRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

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
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGraphReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGraphReq) ProtoMessage() {}

func (x *FCGraphReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
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
	return file_agent_proto_rawDescGZIP(), []int{2}
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
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCGraphRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCGraphRes) ProtoMessage() {}

func (x *FCGraphRes) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
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
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *FCGraphRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

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
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCMetricReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCMetricReq) ProtoMessage() {}

func (x *FCMetricReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
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
	return file_agent_proto_rawDescGZIP(), []int{4}
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
		mi := &file_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FCMetricRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FCMetricRes) ProtoMessage() {}

func (x *FCMetricRes) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[5]
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
	return file_agent_proto_rawDescGZIP(), []int{5}
}

func (x *FCMetricRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x36, 0x0a, 0x0a, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x46, 0x43, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x0a, 0x46, 0x43, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22,
	0x1c, 0x0a, 0x0a, 0x46, 0x43, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a,
	0x0b, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa9, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x33, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x10, 0x46, 0x69, 0x6e,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x43, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_agent_proto_goTypes = []interface{}{
	(*FCGroupReq)(nil),  // 0: pb.FCGroupReq
	(*FCGroupRes)(nil),  // 1: pb.FCGroupRes
	(*FCGraphReq)(nil),  // 2: pb.FCGraphReq
	(*FCGraphRes)(nil),  // 3: pb.FCGraphRes
	(*FCMetricReq)(nil), // 4: pb.FCMetricReq
	(*FCMetricRes)(nil), // 5: pb.FCMetricRes
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: pb.Agent.FindCreateGroup:input_type -> pb.FCGroupReq
	2, // 1: pb.Agent.FindCreateGraph:input_type -> pb.FCGraphReq
	4, // 2: pb.Agent.FindCreateMetric:input_type -> pb.FCMetricReq
	1, // 3: pb.Agent.FindCreateGroup:output_type -> pb.FCGroupRes
	3, // 4: pb.Agent.FindCreateGraph:output_type -> pb.FCGraphRes
	5, // 5: pb.Agent.FindCreateMetric:output_type -> pb.FCMetricRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	FindCreateGroup(ctx context.Context, in *FCGroupReq, opts ...grpc.CallOption) (*FCGroupRes, error)
	FindCreateGraph(ctx context.Context, in *FCGraphReq, opts ...grpc.CallOption) (*FCGraphRes, error)
	FindCreateMetric(ctx context.Context, in *FCMetricReq, opts ...grpc.CallOption) (*FCMetricRes, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) FindCreateGroup(ctx context.Context, in *FCGroupReq, opts ...grpc.CallOption) (*FCGroupRes, error) {
	out := new(FCGroupRes)
	err := c.cc.Invoke(ctx, "/pb.Agent/FindCreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) FindCreateGraph(ctx context.Context, in *FCGraphReq, opts ...grpc.CallOption) (*FCGraphRes, error) {
	out := new(FCGraphRes)
	err := c.cc.Invoke(ctx, "/pb.Agent/FindCreateGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) FindCreateMetric(ctx context.Context, in *FCMetricReq, opts ...grpc.CallOption) (*FCMetricRes, error) {
	out := new(FCMetricRes)
	err := c.cc.Invoke(ctx, "/pb.Agent/FindCreateMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	FindCreateGroup(context.Context, *FCGroupReq) (*FCGroupRes, error)
	FindCreateGraph(context.Context, *FCGraphReq) (*FCGraphRes, error)
	FindCreateMetric(context.Context, *FCMetricReq) (*FCMetricRes, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) FindCreateGroup(context.Context, *FCGroupReq) (*FCGroupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCreateGroup not implemented")
}
func (*UnimplementedAgentServer) FindCreateGraph(context.Context, *FCGraphReq) (*FCGraphRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCreateGraph not implemented")
}
func (*UnimplementedAgentServer) FindCreateMetric(context.Context, *FCMetricReq) (*FCMetricRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCreateMetric not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_FindCreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FCGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).FindCreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/FindCreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).FindCreateGroup(ctx, req.(*FCGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_FindCreateGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FCGraphReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).FindCreateGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/FindCreateGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).FindCreateGraph(ctx, req.(*FCGraphReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_FindCreateMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FCMetricReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).FindCreateMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/FindCreateMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).FindCreateMetric(ctx, req.(*FCMetricReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCreateGroup",
			Handler:    _Agent_FindCreateGroup_Handler,
		},
		{
			MethodName: "FindCreateGraph",
			Handler:    _Agent_FindCreateGraph_Handler,
		},
		{
			MethodName: "FindCreateMetric",
			Handler:    _Agent_FindCreateMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

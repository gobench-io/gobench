// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pb/executor.proto

package pb

import (
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

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_pb_executor_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

type StartResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   int64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StartResult) Reset() {
	*x = StartResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResult) ProtoMessage() {}

func (x *StartResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResult.ProtoReflect.Descriptor instead.
func (*StartResult) Descriptor() ([]byte, []int) {
	return file_pb_executor_proto_rawDescGZIP(), []int{1}
}

func (x *StartResult) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *StartResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Code  int64 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *TermRequest) Reset() {
	*x = TermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermRequest) ProtoMessage() {}

func (x *TermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermRequest.ProtoReflect.Descriptor instead.
func (*TermRequest) Descriptor() ([]byte, []int) {
	return file_pb_executor_proto_rawDescGZIP(), []int{2}
}

func (x *TermRequest) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *TermRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type TermResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   int64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TermResult) Reset() {
	*x = TermResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermResult) ProtoMessage() {}

func (x *TermResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermResult.ProtoReflect.Descriptor instead.
func (*TermResult) Descriptor() ([]byte, []int) {
	return file_pb_executor_proto_rawDescGZIP(), []int{3}
}

func (x *TermResult) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *TermResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pb_executor_proto protoreflect.FileDescriptor

var file_pb_executor_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x3d, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0b,
	0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3c, 0x0a, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x64, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12,
	0x2a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_executor_proto_rawDescOnce sync.Once
	file_pb_executor_proto_rawDescData = file_pb_executor_proto_rawDesc
)

func file_pb_executor_proto_rawDescGZIP() []byte {
	file_pb_executor_proto_rawDescOnce.Do(func() {
		file_pb_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_executor_proto_rawDescData)
	})
	return file_pb_executor_proto_rawDescData
}

var file_pb_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_executor_proto_goTypes = []interface{}{
	(*StartRequest)(nil), // 0: pb.StartRequest
	(*StartResult)(nil),  // 1: pb.StartResult
	(*TermRequest)(nil),  // 2: pb.TermRequest
	(*TermResult)(nil),   // 3: pb.TermResult
}
var file_pb_executor_proto_depIdxs = []int32{
	0, // 0: pb.Executor.Start:input_type -> pb.StartRequest
	2, // 1: pb.Executor.Terminate:input_type -> pb.TermRequest
	1, // 2: pb.Executor.Start:output_type -> pb.StartResult
	3, // 3: pb.Executor.Terminate:output_type -> pb.TermResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_executor_proto_init() }
func file_pb_executor_proto_init() {
	if File_pb_executor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_pb_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResult); i {
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
		file_pb_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermRequest); i {
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
		file_pb_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermResult); i {
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
			RawDescriptor: file_pb_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_executor_proto_goTypes,
		DependencyIndexes: file_pb_executor_proto_depIdxs,
		MessageInfos:      file_pb_executor_proto_msgTypes,
	}.Build()
	File_pb_executor_proto = out.File
	file_pb_executor_proto_rawDesc = nil
	file_pb_executor_proto_goTypes = nil
	file_pb_executor_proto_depIdxs = nil
}
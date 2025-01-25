// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: lcache.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Group         string                 `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_lcache_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_lcache_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_lcache_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseForGet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseForGet) Reset() {
	*x = ResponseForGet{}
	mi := &file_lcache_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseForGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseForGet) ProtoMessage() {}

func (x *ResponseForGet) ProtoReflect() protoreflect.Message {
	mi := &file_lcache_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseForGet.ProtoReflect.Descriptor instead.
func (*ResponseForGet) Descriptor() ([]byte, []int) {
	return file_lcache_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseForGet) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ResponseForDelete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         bool                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseForDelete) Reset() {
	*x = ResponseForDelete{}
	mi := &file_lcache_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseForDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseForDelete) ProtoMessage() {}

func (x *ResponseForDelete) ProtoReflect() protoreflect.Message {
	mi := &file_lcache_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseForDelete.ProtoReflect.Descriptor instead.
func (*ResponseForDelete) Descriptor() ([]byte, []int) {
	return file_lcache_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseForDelete) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_lcache_proto protoreflect.FileDescriptor

var file_lcache_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x6c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x31, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x46, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x62, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_lcache_proto_rawDescOnce sync.Once
	file_lcache_proto_rawDescData []byte
)

func file_lcache_proto_rawDescGZIP() []byte {
	file_lcache_proto_rawDescOnce.Do(func() {
		file_lcache_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_lcache_proto_rawDesc), len(file_lcache_proto_rawDesc)))
	})
	return file_lcache_proto_rawDescData
}

var file_lcache_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lcache_proto_goTypes = []any{
	(*Request)(nil),           // 0: pb.Request
	(*ResponseForGet)(nil),    // 1: pb.ResponseForGet
	(*ResponseForDelete)(nil), // 2: pb.ResponseForDelete
}
var file_lcache_proto_depIdxs = []int32{
	0, // 0: pb.GroupCache.Get:input_type -> pb.Request
	0, // 1: pb.GroupCache.Delete:input_type -> pb.Request
	1, // 2: pb.GroupCache.Get:output_type -> pb.ResponseForGet
	2, // 3: pb.GroupCache.Delete:output_type -> pb.ResponseForDelete
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lcache_proto_init() }
func file_lcache_proto_init() {
	if File_lcache_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_lcache_proto_rawDesc), len(file_lcache_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lcache_proto_goTypes,
		DependencyIndexes: file_lcache_proto_depIdxs,
		MessageInfos:      file_lcache_proto_msgTypes,
	}.Build()
	File_lcache_proto = out.File
	file_lcache_proto_goTypes = nil
	file_lcache_proto_depIdxs = nil
}

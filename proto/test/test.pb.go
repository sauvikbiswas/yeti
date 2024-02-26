// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: proto/test/test.proto

package test

import (
	_ "github.com/sauvikbiswas/yeti/proto/options"
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

type TestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *TestProto) Reset() {
	*x = TestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProto) ProtoMessage() {}

func (x *TestProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProto.ProtoReflect.Descriptor instead.
func (*TestProto) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestProto) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type TestProtoWithCompositeKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AgeAsString string `protobuf:"bytes,2,opt,name=age_as_string,json=ageAsString,proto3" json:"age_as_string,omitempty"`
}

func (x *TestProtoWithCompositeKey) Reset() {
	*x = TestProtoWithCompositeKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProtoWithCompositeKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProtoWithCompositeKey) ProtoMessage() {}

func (x *TestProtoWithCompositeKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProtoWithCompositeKey.ProtoReflect.Descriptor instead.
func (*TestProtoWithCompositeKey) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestProtoWithCompositeKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestProtoWithCompositeKey) GetAgeAsString() string {
	if x != nil {
		return x.AgeAsString
	}
	return ""
}

type TestProtoWithNonStringPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *TestProtoWithNonStringPrimaryKey) Reset() {
	*x = TestProtoWithNonStringPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProtoWithNonStringPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProtoWithNonStringPrimaryKey) ProtoMessage() {}

func (x *TestProtoWithNonStringPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProtoWithNonStringPrimaryKey.ProtoReflect.Descriptor instead.
func (*TestProtoWithNonStringPrimaryKey) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestProtoWithNonStringPrimaryKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestProtoWithNonStringPrimaryKey) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type TestProtoWithNoString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age int32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *TestProtoWithNoString) Reset() {
	*x = TestProtoWithNoString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestProtoWithNoString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestProtoWithNoString) ProtoMessage() {}

func (x *TestProtoWithNoString) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestProtoWithNoString.ProtoReflect.Descriptor instead.
func (*TestProtoWithNoString) Descriptor() ([]byte, []int) {
	return file_proto_test_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestProtoWithNoString) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_proto_test_test_proto protoreflect.FileDescriptor

var file_proto_test_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x79, 0x65, 0x74, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xea, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x63, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xea, 0xb9, 0x19,
	0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0d, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xea, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x41, 0x73, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x58, 0x0a, 0x20, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xea, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x06, 0xea, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22,
	0x31, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68,
	0x4e, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xea, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x61, 0x75, 0x76, 0x69, 0x6b, 0x62, 0x69, 0x73, 0x77, 0x61, 0x73, 0x2f, 0x79, 0x65,
	0x74, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_test_test_proto_rawDescOnce sync.Once
	file_proto_test_test_proto_rawDescData = file_proto_test_test_proto_rawDesc
)

func file_proto_test_test_proto_rawDescGZIP() []byte {
	file_proto_test_test_proto_rawDescOnce.Do(func() {
		file_proto_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_test_proto_rawDescData)
	})
	return file_proto_test_test_proto_rawDescData
}

var file_proto_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_test_test_proto_goTypes = []interface{}{
	(*TestProto)(nil),                        // 0: yeti.proto.test.TestProto
	(*TestProtoWithCompositeKey)(nil),        // 1: yeti.proto.test.TestProtoWithCompositeKey
	(*TestProtoWithNonStringPrimaryKey)(nil), // 2: yeti.proto.test.TestProtoWithNonStringPrimaryKey
	(*TestProtoWithNoString)(nil),            // 3: yeti.proto.test.TestProtoWithNoString
}
var file_proto_test_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_test_test_proto_init() }
func file_proto_test_test_proto_init() {
	if File_proto_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProto); i {
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
		file_proto_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProtoWithCompositeKey); i {
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
		file_proto_test_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProtoWithNonStringPrimaryKey); i {
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
		file_proto_test_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestProtoWithNoString); i {
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
			RawDescriptor: file_proto_test_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_test_test_proto_goTypes,
		DependencyIndexes: file_proto_test_test_proto_depIdxs,
		MessageInfos:      file_proto_test_test_proto_msgTypes,
	}.Build()
	File_proto_test_test_proto = out.File
	file_proto_test_test_proto_rawDesc = nil
	file_proto_test_test_proto_goTypes = nil
	file_proto_test_test_proto_depIdxs = nil
}

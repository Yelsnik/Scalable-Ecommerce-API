// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: product-service/file.proto

package product

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

type File struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Buffer        []byte                 `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`
	Originalname  string                 `protobuf:"bytes,2,opt,name=originalname,proto3" json:"originalname,omitempty"`
	Fieldname     string                 `protobuf:"bytes,3,opt,name=fieldname,proto3" json:"fieldname,omitempty"`
	Size          int64                  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Mimetype      string                 `protobuf:"bytes,5,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Encoding      string                 `protobuf:"bytes,6,opt,name=encoding,proto3" json:"encoding,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_product_service_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_product_service_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *File) GetOriginalname() string {
	if x != nil {
		return x.Originalname
	}
	return ""
}

func (x *File) GetFieldname() string {
	if x != nil {
		return x.Fieldname
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

func (x *File) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

var File_product_service_file_proto protoreflect.FileDescriptor

var file_product_service_file_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x24, 0x5a, 0x22, 0x45, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_product_service_file_proto_rawDescOnce sync.Once
	file_product_service_file_proto_rawDescData = file_product_service_file_proto_rawDesc
)

func file_product_service_file_proto_rawDescGZIP() []byte {
	file_product_service_file_proto_rawDescOnce.Do(func() {
		file_product_service_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_service_file_proto_rawDescData)
	})
	return file_product_service_file_proto_rawDescData
}

var file_product_service_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_product_service_file_proto_goTypes = []any{
	(*File)(nil), // 0: product.File
}
var file_product_service_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_service_file_proto_init() }
func file_product_service_file_proto_init() {
	if File_product_service_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_service_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_service_file_proto_goTypes,
		DependencyIndexes: file_product_service_file_proto_depIdxs,
		MessageInfos:      file_product_service_file_proto_msgTypes,
	}.Build()
	File_product_service_file_proto = out.File
	file_product_service_file_proto_rawDesc = nil
	file_product_service_file_proto_goTypes = nil
	file_product_service_file_proto_depIdxs = nil
}

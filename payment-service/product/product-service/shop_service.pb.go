// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: product-service/shop_service.proto

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

type CreateShopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ShopOwner     string                 `protobuf:"bytes,3,opt,name=shopOwner,proto3" json:"shopOwner,omitempty"`
	Image         *File                  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShopRequest) Reset() {
	*x = CreateShopRequest{}
	mi := &file_product_service_shop_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShopRequest) ProtoMessage() {}

func (x *CreateShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShopRequest.ProtoReflect.Descriptor instead.
func (*CreateShopRequest) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateShopRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateShopRequest) GetShopOwner() string {
	if x != nil {
		return x.ShopOwner
	}
	return ""
}

func (x *CreateShopRequest) GetImage() *File {
	if x != nil {
		return x.Image
	}
	return nil
}

type ShopResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shop          *Shop                  `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShopResponse) Reset() {
	*x = ShopResponse{}
	mi := &file_product_service_shop_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopResponse) ProtoMessage() {}

func (x *ShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopResponse.ProtoReflect.Descriptor instead.
func (*ShopResponse) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShopResponse) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

type GetShopByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShopByIdRequest) Reset() {
	*x = GetShopByIdRequest{}
	mi := &file_product_service_shop_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShopByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopByIdRequest) ProtoMessage() {}

func (x *GetShopByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopByIdRequest.ProtoReflect.Descriptor instead.
func (*GetShopByIdRequest) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetShopByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetShopsByOwnerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShopsByOwnerRequest) Reset() {
	*x = GetShopsByOwnerRequest{}
	mi := &file_product_service_shop_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShopsByOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopsByOwnerRequest) ProtoMessage() {}

func (x *GetShopsByOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopsByOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetShopsByOwnerRequest) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetShopsByOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetShopsByOwnerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shops         []*Shop                `protobuf:"bytes,1,rep,name=shops,proto3" json:"shops,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShopsByOwnerResponse) Reset() {
	*x = GetShopsByOwnerResponse{}
	mi := &file_product_service_shop_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShopsByOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopsByOwnerResponse) ProtoMessage() {}

func (x *GetShopsByOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopsByOwnerResponse.ProtoReflect.Descriptor instead.
func (*GetShopsByOwnerResponse) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetShopsByOwnerResponse) GetShops() []*Shop {
	if x != nil {
		return x.Shops
	}
	return nil
}

type UpdateShopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShopRequest) Reset() {
	*x = UpdateShopRequest{}
	mi := &file_product_service_shop_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShopRequest) ProtoMessage() {}

func (x *UpdateShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShopRequest.ProtoReflect.Descriptor instead.
func (*UpdateShopRequest) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateShopRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateShopRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateShopRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type DeleteShopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShopRequest) Reset() {
	*x = DeleteShopRequest{}
	mi := &file_product_service_shop_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShopRequest) ProtoMessage() {}

func (x *DeleteShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShopRequest.ProtoReflect.Descriptor instead.
func (*DeleteShopRequest) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteShopRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_product_service_shop_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{7}
}

type Shop struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageName     string                 `protobuf:"bytes,4,opt,name=imageName,proto3" json:"imageName,omitempty"`
	ShopOwner     string                 `protobuf:"bytes,5,opt,name=shopOwner,proto3" json:"shopOwner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Shop) Reset() {
	*x = Shop{}
	mi := &file_product_service_shop_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_shop_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_product_service_shop_service_proto_rawDescGZIP(), []int{8}
}

func (x *Shop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Shop) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Shop) GetShopOwner() string {
	if x != nil {
		return x.ShopOwner
	}
	return ""
}

var File_product_service_shop_service_proto protoreflect.FileDescriptor

var file_product_service_shop_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x24, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x42, 0x79, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x22, 0x7c, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x32, 0xec, 0x02, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73,
	0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70,
	0x73, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x24, 0x5a, 0x22, 0x45, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_service_shop_service_proto_rawDescOnce sync.Once
	file_product_service_shop_service_proto_rawDescData = file_product_service_shop_service_proto_rawDesc
)

func file_product_service_shop_service_proto_rawDescGZIP() []byte {
	file_product_service_shop_service_proto_rawDescOnce.Do(func() {
		file_product_service_shop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_service_shop_service_proto_rawDescData)
	})
	return file_product_service_shop_service_proto_rawDescData
}

var file_product_service_shop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_service_shop_service_proto_goTypes = []any{
	(*CreateShopRequest)(nil),       // 0: product.CreateShopRequest
	(*ShopResponse)(nil),            // 1: product.ShopResponse
	(*GetShopByIdRequest)(nil),      // 2: product.GetShopByIdRequest
	(*GetShopsByOwnerRequest)(nil),  // 3: product.GetShopsByOwnerRequest
	(*GetShopsByOwnerResponse)(nil), // 4: product.GetShopsByOwnerResponse
	(*UpdateShopRequest)(nil),       // 5: product.UpdateShopRequest
	(*DeleteShopRequest)(nil),       // 6: product.DeleteShopRequest
	(*Empty)(nil),                   // 7: product.Empty
	(*Shop)(nil),                    // 8: product.Shop
	(*File)(nil),                    // 9: product.File
}
var file_product_service_shop_service_proto_depIdxs = []int32{
	9, // 0: product.CreateShopRequest.image:type_name -> product.File
	8, // 1: product.ShopResponse.shop:type_name -> product.Shop
	8, // 2: product.GetShopsByOwnerResponse.shops:type_name -> product.Shop
	0, // 3: product.ShopService.CreateShop:input_type -> product.CreateShopRequest
	2, // 4: product.ShopService.GetShopByID:input_type -> product.GetShopByIdRequest
	3, // 5: product.ShopService.GetShopsByOwner:input_type -> product.GetShopsByOwnerRequest
	5, // 6: product.ShopService.UpdateShop:input_type -> product.UpdateShopRequest
	6, // 7: product.ShopService.DeleteShop:input_type -> product.DeleteShopRequest
	1, // 8: product.ShopService.CreateShop:output_type -> product.ShopResponse
	1, // 9: product.ShopService.GetShopByID:output_type -> product.ShopResponse
	4, // 10: product.ShopService.GetShopsByOwner:output_type -> product.GetShopsByOwnerResponse
	1, // 11: product.ShopService.UpdateShop:output_type -> product.ShopResponse
	7, // 12: product.ShopService.DeleteShop:output_type -> product.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_service_shop_service_proto_init() }
func file_product_service_shop_service_proto_init() {
	if File_product_service_shop_service_proto != nil {
		return
	}
	file_product_service_file_proto_init()
	file_product_service_shop_service_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_service_shop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_service_shop_service_proto_goTypes,
		DependencyIndexes: file_product_service_shop_service_proto_depIdxs,
		MessageInfos:      file_product_service_shop_service_proto_msgTypes,
	}.Build()
	File_product_service_shop_service_proto = out.File
	file_product_service_shop_service_proto_rawDesc = nil
	file_product_service_shop_service_proto_goTypes = nil
	file_product_service_shop_service_proto_depIdxs = nil
}

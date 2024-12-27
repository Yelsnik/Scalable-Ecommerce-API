// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: product/product_service.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *anypb.Any             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category      string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	ProductName   string                 `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Brand         string                 `protobuf:"bytes,5,opt,name=brand,proto3" json:"brand,omitempty"`
	ImageName     string                 `protobuf:"bytes,6,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	CountInStock  int64                  `protobuf:"varint,7,opt,name=count_in_stock,json=countInStock,proto3" json:"count_in_stock,omitempty"`
	Price         float32                `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	Currency      string                 `protobuf:"bytes,9,opt,name=currency,proto3" json:"currency,omitempty"`
	Shop          string                 `protobuf:"bytes,10,opt,name=shop,proto3" json:"shop,omitempty"`
	Rating        int64                  `protobuf:"varint,11,opt,name=rating,proto3" json:"rating,omitempty"`
	IsFeatured    bool                   `protobuf:"varint,12,opt,name=is_featured,json=isFeatured,proto3" json:"is_featured,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_product_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() *anypb.Any {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Product) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Product) GetCountInStock() int64 {
	if x != nil {
		return x.CountInStock
	}
	return 0
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Product) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *Product) GetRating() int64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Product) GetIsFeatured() bool {
	if x != nil {
		return x.IsFeatured
	}
	return false
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      string                 `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Brand         string                 `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Image         *File                  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	CountInStock  int64                  `protobuf:"varint,6,opt,name=count_in_stock,json=countInStock,proto3" json:"count_in_stock,omitempty"`
	Price         float32                `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	Currency      string                 `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	Shop          string                 `protobuf:"bytes,9,opt,name=shop,proto3" json:"shop,omitempty"`
	Rating        int64                  `protobuf:"varint,10,opt,name=rating,proto3" json:"rating,omitempty"`
	IsFeatured    bool                   `protobuf:"varint,11,opt,name=is_featured,json=isFeatured,proto3" json:"is_featured,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_product_product_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateProductRequest) GetImage() *File {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *CreateProductRequest) GetCountInStock() int64 {
	if x != nil {
		return x.CountInStock
	}
	return 0
}

func (x *CreateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateProductRequest) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *CreateProductRequest) GetRating() int64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateProductRequest) GetIsFeatured() bool {
	if x != nil {
		return x.IsFeatured
	}
	return false
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_product_product_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetProductByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductByIdRequest) Reset() {
	*x = GetProductByIdRequest{}
	mi := &file_product_product_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRequest) ProtoMessage() {}

func (x *GetProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductsByShopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsByShopRequest) Reset() {
	*x = GetProductsByShopRequest{}
	mi := &file_product_product_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsByShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsByShopRequest) ProtoMessage() {}

func (x *GetProductsByShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsByShopRequest.ProtoReflect.Descriptor instead.
func (*GetProductsByShopRequest) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductsByShopRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductsByShopResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       []*Product             `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsByShopResponse) Reset() {
	*x = GetProductsByShopResponse{}
	mi := &file_product_product_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsByShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsByShopResponse) ProtoMessage() {}

func (x *GetProductsByShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsByShopResponse.ProtoReflect.Descriptor instead.
func (*GetProductsByShopResponse) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductsByShopResponse) GetProduct() []*Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category      *string                `protobuf:"bytes,2,opt,name=category,proto3,oneof" json:"category,omitempty"`
	ProductName   *string                `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3,oneof" json:"product_name,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Brand         *string                `protobuf:"bytes,5,opt,name=brand,proto3,oneof" json:"brand,omitempty"`
	Image         *File                  `protobuf:"bytes,6,opt,name=image,proto3,oneof" json:"image,omitempty"`
	CountInStock  *int64                 `protobuf:"varint,7,opt,name=count_in_stock,json=countInStock,proto3,oneof" json:"count_in_stock,omitempty"`
	Price         *float32               `protobuf:"fixed32,8,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Currency      *string                `protobuf:"bytes,9,opt,name=currency,proto3,oneof" json:"currency,omitempty"`
	Rating        *int64                 `protobuf:"varint,10,opt,name=rating,proto3,oneof" json:"rating,omitempty"`
	IsFeatured    *bool                  `protobuf:"varint,11,opt,name=is_featured,json=isFeatured,proto3,oneof" json:"is_featured,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_product_product_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *UpdateProductRequest) GetProductName() string {
	if x != nil && x.ProductName != nil {
		return *x.ProductName
	}
	return ""
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *UpdateProductRequest) GetImage() *File {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UpdateProductRequest) GetCountInStock() int64 {
	if x != nil && x.CountInStock != nil {
		return *x.CountInStock
	}
	return 0
}

func (x *UpdateProductRequest) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetCurrency() string {
	if x != nil && x.Currency != nil {
		return *x.Currency
	}
	return ""
}

func (x *UpdateProductRequest) GetRating() int64 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

func (x *UpdateProductRequest) GetIsFeatured() bool {
	if x != nil && x.IsFeatured != nil {
		return *x.IsFeatured
	}
	return false
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_product_product_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRes) Reset() {
	*x = EmptyRes{}
	mi := &file_product_product_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRes) ProtoMessage() {}

func (x *EmptyRes) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRes.ProtoReflect.Descriptor instead.
func (*EmptyRes) Descriptor() ([]byte, []int) {
	return file_product_product_service_proto_rawDescGZIP(), []int{8}
}

var File_product_product_service_proto protoreflect.FileDescriptor

var file_product_product_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd7, 0x02, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x8c, 0x04, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x04, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x48, 0x06, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x48, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x32, 0x96, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53,
	0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x63, 0x61, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_product_service_proto_rawDescOnce sync.Once
	file_product_product_service_proto_rawDescData = file_product_product_service_proto_rawDesc
)

func file_product_product_service_proto_rawDescGZIP() []byte {
	file_product_product_service_proto_rawDescOnce.Do(func() {
		file_product_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_product_service_proto_rawDescData)
	})
	return file_product_product_service_proto_rawDescData
}

var file_product_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_product_product_service_proto_goTypes = []any{
	(*Product)(nil),                   // 0: product.Product
	(*CreateProductRequest)(nil),      // 1: product.CreateProductRequest
	(*ProductResponse)(nil),           // 2: product.ProductResponse
	(*GetProductByIdRequest)(nil),     // 3: product.GetProductByIdRequest
	(*GetProductsByShopRequest)(nil),  // 4: product.GetProductsByShopRequest
	(*GetProductsByShopResponse)(nil), // 5: product.GetProductsByShopResponse
	(*UpdateProductRequest)(nil),      // 6: product.UpdateProductRequest
	(*DeleteProductRequest)(nil),      // 7: product.DeleteProductRequest
	(*EmptyRes)(nil),                  // 8: product.EmptyRes
	(*anypb.Any)(nil),                 // 9: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
	(*File)(nil),                      // 11: product.File
}
var file_product_product_service_proto_depIdxs = []int32{
	9,  // 0: product.Product.id:type_name -> google.protobuf.Any
	10, // 1: product.Product.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: product.Product.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: product.CreateProductRequest.image:type_name -> product.File
	0,  // 4: product.ProductResponse.product:type_name -> product.Product
	0,  // 5: product.GetProductsByShopResponse.product:type_name -> product.Product
	11, // 6: product.UpdateProductRequest.image:type_name -> product.File
	1,  // 7: product.ProductService.AddProduct:input_type -> product.CreateProductRequest
	3,  // 8: product.ProductService.GetProductByID:input_type -> product.GetProductByIdRequest
	4,  // 9: product.ProductService.GetProductsByShop:input_type -> product.GetProductsByShopRequest
	6,  // 10: product.ProductService.UpdateProduct:input_type -> product.UpdateProductRequest
	7,  // 11: product.ProductService.DeleteProduct:input_type -> product.DeleteProductRequest
	2,  // 12: product.ProductService.AddProduct:output_type -> product.ProductResponse
	2,  // 13: product.ProductService.GetProductByID:output_type -> product.ProductResponse
	5,  // 14: product.ProductService.GetProductsByShop:output_type -> product.GetProductsByShopResponse
	2,  // 15: product.ProductService.UpdateProduct:output_type -> product.ProductResponse
	8,  // 16: product.ProductService.DeleteProduct:output_type -> product.EmptyRes
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_product_product_service_proto_init() }
func file_product_product_service_proto_init() {
	if File_product_product_service_proto != nil {
		return
	}
	file_product_file_proto_init()
	file_product_product_service_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_service_proto_goTypes,
		DependencyIndexes: file_product_product_service_proto_depIdxs,
		MessageInfos:      file_product_product_service_proto_msgTypes,
	}.Build()
	File_product_product_service_proto = out.File
	file_product_product_service_proto_rawDesc = nil
	file_product_product_service_proto_goTypes = nil
	file_product_product_service_proto_depIdxs = nil
}

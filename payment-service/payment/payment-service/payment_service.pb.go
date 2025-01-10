// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: payment-service/payment_service.proto

package payment

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

type StripeCustomerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentId     string                 `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StripeCustomerRequest) Reset() {
	*x = StripeCustomerRequest{}
	mi := &file_payment_service_payment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StripeCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripeCustomerRequest) ProtoMessage() {}

func (x *StripeCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripeCustomerRequest.ProtoReflect.Descriptor instead.
func (*StripeCustomerRequest) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{0}
}

func (x *StripeCustomerRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StripeCustomerRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *StripeCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type StripeCustomerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	CustomerId    string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StripeCustomerResponse) Reset() {
	*x = StripeCustomerResponse{}
	mi := &file_payment_service_payment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StripeCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripeCustomerResponse) ProtoMessage() {}

func (x *StripeCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripeCustomerResponse.ProtoReflect.Descriptor instead.
func (*StripeCustomerResponse) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{1}
}

func (x *StripeCustomerResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *StripeCustomerResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type CreatePaymentRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentId       string                 `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Email           string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Amount          float32                `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	DeliveryAddress string                 `protobuf:"bytes,5,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	Country         string                 `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	CartItemId      string                 `protobuf:"bytes,7,opt,name=cart_item_id,json=cartItemId,proto3" json:"cart_item_id,omitempty"`
	Currency        string                 `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	SaveCard        bool                   `protobuf:"varint,9,opt,name=save_card,json=saveCard,proto3" json:"save_card,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	mi := &file_payment_service_payment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePaymentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePaymentRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *CreatePaymentRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreatePaymentRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreatePaymentRequest) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *CreatePaymentRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreatePaymentRequest) GetCartItemId() string {
	if x != nil {
		return x.CartItemId
	}
	return ""
}

func (x *CreatePaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreatePaymentRequest) GetSaveCard() bool {
	if x != nil {
		return x.SaveCard
	}
	return false
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientSecret  string                 `protobuf:"bytes,1,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	BuyerUserId   string                 `protobuf:"bytes,2,opt,name=buyer_user_id,json=buyerUserId,proto3" json:"buyer_user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	mi := &file_payment_service_payment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePaymentResponse) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *CreatePaymentResponse) GetBuyerUserId() string {
	if x != nil {
		return x.BuyerUserId
	}
	return ""
}

type WebhookRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Payload         string                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	StripeSignature string                 `protobuf:"bytes,2,opt,name=stripe_signature,json=stripeSignature,proto3" json:"stripe_signature,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *WebhookRequest) Reset() {
	*x = WebhookRequest{}
	mi := &file_payment_service_payment_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookRequest) ProtoMessage() {}

func (x *WebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookRequest.ProtoReflect.Descriptor instead.
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{4}
}

func (x *WebhookRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *WebhookRequest) GetStripeSignature() string {
	if x != nil {
		return x.StripeSignature
	}
	return ""
}

type WebhookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payment       *Payment               `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	Order         *Order                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebhookResponse) Reset() {
	*x = WebhookResponse{}
	mi := &file_payment_service_payment_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookResponse) ProtoMessage() {}

func (x *WebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_payment_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookResponse.ProtoReflect.Descriptor instead.
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return file_payment_service_payment_service_proto_rawDescGZIP(), []int{5}
}

func (x *WebhookResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *WebhookResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_payment_service_payment_service_proto protoreflect.FileDescriptor

var file_payment_service_payment_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x1d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x15,
	0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x4f, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x9c, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x22, 0x60, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x79, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x63, 0x0a, 0x0f,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x32, 0xf1, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_service_payment_service_proto_rawDescOnce sync.Once
	file_payment_service_payment_service_proto_rawDescData = file_payment_service_payment_service_proto_rawDesc
)

func file_payment_service_payment_service_proto_rawDescGZIP() []byte {
	file_payment_service_payment_service_proto_rawDescOnce.Do(func() {
		file_payment_service_payment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_service_payment_service_proto_rawDescData)
	})
	return file_payment_service_payment_service_proto_rawDescData
}

var file_payment_service_payment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payment_service_payment_service_proto_goTypes = []any{
	(*StripeCustomerRequest)(nil),  // 0: payment.StripeCustomerRequest
	(*StripeCustomerResponse)(nil), // 1: payment.StripeCustomerResponse
	(*CreatePaymentRequest)(nil),   // 2: payment.CreatePaymentRequest
	(*CreatePaymentResponse)(nil),  // 3: payment.CreatePaymentResponse
	(*WebhookRequest)(nil),         // 4: payment.WebhookRequest
	(*WebhookResponse)(nil),        // 5: payment.WebhookResponse
	(*Payment)(nil),                // 6: payment.Payment
	(*Order)(nil),                  // 7: payment.Order
}
var file_payment_service_payment_service_proto_depIdxs = []int32{
	6, // 0: payment.WebhookResponse.payment:type_name -> payment.Payment
	7, // 1: payment.WebhookResponse.order:type_name -> payment.Order
	0, // 2: payment.PaymentService.StripeCustomer:input_type -> payment.StripeCustomerRequest
	2, // 3: payment.PaymentService.CreatePayment:input_type -> payment.CreatePaymentRequest
	4, // 4: payment.PaymentService.Webhook:input_type -> payment.WebhookRequest
	1, // 5: payment.PaymentService.StripeCustomer:output_type -> payment.StripeCustomerResponse
	3, // 6: payment.PaymentService.CreatePayment:output_type -> payment.CreatePaymentResponse
	5, // 7: payment.PaymentService.Webhook:output_type -> payment.WebhookResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_service_payment_service_proto_init() }
func file_payment_service_payment_service_proto_init() {
	if File_payment_service_payment_service_proto != nil {
		return
	}
	file_payment_service_payment_proto_init()
	file_payment_service_order_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_service_payment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_service_payment_service_proto_goTypes,
		DependencyIndexes: file_payment_service_payment_service_proto_depIdxs,
		MessageInfos:      file_payment_service_payment_service_proto_msgTypes,
	}.Build()
	File_payment_service_payment_service_proto = out.File
	file_payment_service_payment_service_proto_rawDesc = nil
	file_payment_service_payment_service_proto_goTypes = nil
	file_payment_service_payment_service_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/order.proto

package __

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId       string  `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Symbol        string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Quantity      int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PricePerStock float64 `protobuf:"fixed64,4,opt,name=pricePerStock,proto3" json:"pricePerStock,omitempty"`
	TotalPrice    float64 `protobuf:"fixed64,5,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	OrderType     string  `protobuf:"bytes,6,opt,name=orderType,proto3" json:"orderType,omitempty"`
	OrderStatus   string  `protobuf:"bytes,7,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Order) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order) GetPricePerStock() float64 {
	if x != nil {
		return x.PricePerStock
	}
	return 0
}

func (x *Order) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *Order) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	OrderType string `protobuf:"bytes,3,opt,name=orderType,proto3" json:"orderType,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OrderRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderRequest) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order    *Order    `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *OrderResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *CancelOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order    *Order    `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *CancelOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *CancelOrderResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type OrderHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderHistoryRequest) Reset() {
	*x = OrderHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderHistoryRequest) ProtoMessage() {}

func (x *OrderHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderHistoryRequest.ProtoReflect.Descriptor instead.
func (*OrderHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{5}
}

type OrderHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders   []*Order  `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OrderHistoryResponse) Reset() {
	*x = OrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderHistoryResponse) ProtoMessage() {}

func (x *OrderHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*OrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderHistoryResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *OrderHistoryResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetCurrentPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *GetCurrentPriceRequest) Reset() {
	*x = GetCurrentPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentPriceRequest) ProtoMessage() {}

func (x *GetCurrentPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentPriceRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentPriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{7}
}

func (x *GetCurrentPriceRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type GetCurrentPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price    float64   `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetCurrentPriceResponse) Reset() {
	*x = GetCurrentPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentPriceResponse) ProtoMessage() {}

func (x *GetCurrentPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentPriceResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentPriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{8}
}

func (x *GetCurrentPriceResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetCurrentPriceResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{9}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xdb, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x60, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x60, 0x0a, 0x0d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x12,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x13,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x14, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xab, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_order_proto_goTypes = []any{
	(*Order)(nil),                   // 0: order.Order
	(*OrderRequest)(nil),            // 1: order.OrderRequest
	(*OrderResponse)(nil),           // 2: order.OrderResponse
	(*CancelOrderRequest)(nil),      // 3: order.CancelOrderRequest
	(*CancelOrderResponse)(nil),     // 4: order.CancelOrderResponse
	(*OrderHistoryRequest)(nil),     // 5: order.OrderHistoryRequest
	(*OrderHistoryResponse)(nil),    // 6: order.OrderHistoryResponse
	(*GetCurrentPriceRequest)(nil),  // 7: order.GetCurrentPriceRequest
	(*GetCurrentPriceResponse)(nil), // 8: order.GetCurrentPriceResponse
	(*Response)(nil),                // 9: order.Response
}
var file_proto_order_proto_depIdxs = []int32{
	0,  // 0: order.OrderResponse.order:type_name -> order.Order
	9,  // 1: order.OrderResponse.response:type_name -> order.Response
	0,  // 2: order.CancelOrderResponse.order:type_name -> order.Order
	9,  // 3: order.CancelOrderResponse.response:type_name -> order.Response
	0,  // 4: order.OrderHistoryResponse.orders:type_name -> order.Order
	9,  // 5: order.OrderHistoryResponse.response:type_name -> order.Response
	9,  // 6: order.GetCurrentPriceResponse.response:type_name -> order.Response
	1,  // 7: order.OrderService.PlaceOrder:input_type -> order.OrderRequest
	3,  // 8: order.OrderService.CancelOrder:input_type -> order.CancelOrderRequest
	5,  // 9: order.OrderService.GetOrderHistory:input_type -> order.OrderHistoryRequest
	7,  // 10: order.OrderService.GetCurrentPrice:input_type -> order.GetCurrentPriceRequest
	2,  // 11: order.OrderService.PlaceOrder:output_type -> order.OrderResponse
	4,  // 12: order.OrderService.CancelOrder:output_type -> order.CancelOrderResponse
	6,  // 13: order.OrderService.GetOrderHistory:output_type -> order.OrderHistoryResponse
	8,  // 14: order.OrderService.GetCurrentPrice:output_type -> order.GetCurrentPriceResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_proto_order_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OrderRequest); i {
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
		file_proto_order_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OrderResponse); i {
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
		file_proto_order_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CancelOrderRequest); i {
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
		file_proto_order_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CancelOrderResponse); i {
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
		file_proto_order_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OrderHistoryRequest); i {
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
		file_proto_order_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*OrderHistoryResponse); i {
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
		file_proto_order_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetCurrentPriceRequest); i {
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
		file_proto_order_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetCurrentPriceResponse); i {
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
		file_proto_order_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: courierServer.proto

package courierProto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type OrderCourierServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID           int64                  `protobuf:"varint,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	CourierServiceID  int64                  `protobuf:"varint,2,opt,name=CourierServiceID,proto3" json:"CourierServiceID,omitempty"`
	RestaurantAddress string                 `protobuf:"bytes,3,opt,name=RestaurantAddress,proto3" json:"RestaurantAddress,omitempty"`
	RestaurantName    string                 `protobuf:"bytes,4,opt,name=RestaurantName,proto3" json:"RestaurantName,omitempty"`
	ClientAddress     string                 `protobuf:"bytes,5,opt,name=ClientAddress,proto3" json:"ClientAddress,omitempty"`
	ClientFullName    string                 `protobuf:"bytes,6,opt,name=ClientFullName,proto3" json:"ClientFullName,omitempty"`
	ClientPhoneNumber string                 `protobuf:"bytes,7,opt,name=ClientPhoneNumber,proto3" json:"ClientPhoneNumber,omitempty"`
	DeliveryTime      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=DeliveryTime,proto3" json:"DeliveryTime,omitempty"`
	PaymentType       int64                  `protobuf:"varint,9,opt,name=PaymentType,proto3" json:"PaymentType,omitempty"`
}

func (x *OrderCourierServer) Reset() {
	*x = OrderCourierServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCourierServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCourierServer) ProtoMessage() {}

func (x *OrderCourierServer) ProtoReflect() protoreflect.Message {
	mi := &file_courierServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCourierServer.ProtoReflect.Descriptor instead.
func (*OrderCourierServer) Descriptor() ([]byte, []int) {
	return file_courierServer_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCourierServer) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *OrderCourierServer) GetCourierServiceID() int64 {
	if x != nil {
		return x.CourierServiceID
	}
	return 0
}

func (x *OrderCourierServer) GetRestaurantAddress() string {
	if x != nil {
		return x.RestaurantAddress
	}
	return ""
}

func (x *OrderCourierServer) GetRestaurantName() string {
	if x != nil {
		return x.RestaurantName
	}
	return ""
}

func (x *OrderCourierServer) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

func (x *OrderCourierServer) GetClientFullName() string {
	if x != nil {
		return x.ClientFullName
	}
	return ""
}

func (x *OrderCourierServer) GetClientPhoneNumber() string {
	if x != nil {
		return x.ClientPhoneNumber
	}
	return ""
}

func (x *OrderCourierServer) GetDeliveryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryTime
	}
	return nil
}

func (x *OrderCourierServer) GetPaymentType() int64 {
	if x != nil {
		return x.PaymentType
	}
	return 0
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*DeliveryService `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_courierServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_courierServer_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceResponse) GetServices() []*DeliveryService {
	if x != nil {
		return x.Services
	}
	return nil
}

type DeliveryService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId          int64  `protobuf:"varint,1,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
	ServiceName        string `protobuf:"bytes,2,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	ServiceEmail       string `protobuf:"bytes,3,opt,name=ServiceEmail,proto3" json:"ServiceEmail,omitempty"`
	ServicePhoto       string `protobuf:"bytes,4,opt,name=ServicePhoto,proto3" json:"ServicePhoto,omitempty"`
	ServiceDescription string `protobuf:"bytes,5,opt,name=ServiceDescription,proto3" json:"ServiceDescription,omitempty"`
	ServicePhone       string `protobuf:"bytes,6,opt,name=ServicePhone,proto3" json:"ServicePhone,omitempty"`
	ServiceManagerId   int64  `protobuf:"varint,7,opt,name=ServiceManagerId,proto3" json:"ServiceManagerId,omitempty"`
	ServiceStatus      string `protobuf:"bytes,8,opt,name=ServiceStatus,proto3" json:"ServiceStatus,omitempty"`
}

func (x *DeliveryService) Reset() {
	*x = DeliveryService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courierServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryService) ProtoMessage() {}

func (x *DeliveryService) ProtoReflect() protoreflect.Message {
	mi := &file_courierServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryService.ProtoReflect.Descriptor instead.
func (*DeliveryService) Descriptor() ([]byte, []int) {
	return file_courierServer_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryService) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *DeliveryService) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DeliveryService) GetServiceEmail() string {
	if x != nil {
		return x.ServiceEmail
	}
	return ""
}

func (x *DeliveryService) GetServicePhoto() string {
	if x != nil {
		return x.ServicePhoto
	}
	return ""
}

func (x *DeliveryService) GetServiceDescription() string {
	if x != nil {
		return x.ServiceDescription
	}
	return ""
}

func (x *DeliveryService) GetServicePhone() string {
	if x != nil {
		return x.ServicePhone
	}
	return ""
}

func (x *DeliveryService) GetServiceManagerId() int64 {
	if x != nil {
		return x.ServiceManagerId
	}
	return 0
}

func (x *DeliveryService) GetServiceStatus() string {
	if x != nil {
		return x.ServiceStatus
	}
	return ""
}

var File_courierServer_proto protoreflect.FileDescriptor

var file_courierServer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a,
	0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2a, 0x0a,
	0x10, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a,
	0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xbf, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9f, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x47, 0x52,
	0x50, 0x43, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_courierServer_proto_rawDescOnce sync.Once
	file_courierServer_proto_rawDescData = file_courierServer_proto_rawDesc
)

func file_courierServer_proto_rawDescGZIP() []byte {
	file_courierServer_proto_rawDescOnce.Do(func() {
		file_courierServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_courierServer_proto_rawDescData)
	})
	return file_courierServer_proto_rawDescData
}

var file_courierServer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_courierServer_proto_goTypes = []interface{}{
	(*OrderCourierServer)(nil),    // 0: courier.OrderCourierServer
	(*ServiceResponse)(nil),       // 1: courier.ServiceResponse
	(*DeliveryService)(nil),       // 2: courier.DeliveryService
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_courierServer_proto_depIdxs = []int32{
	3, // 0: courier.OrderCourierServer.DeliveryTime:type_name -> google.protobuf.Timestamp
	2, // 1: courier.ServiceResponse.services:type_name -> courier.DeliveryService
	0, // 2: courier.CourierServer.CreateOrder:input_type -> courier.OrderCourierServer
	4, // 3: courier.CourierServer.GetDeliveryService:input_type -> google.protobuf.Empty
	4, // 4: courier.CourierServer.CreateOrder:output_type -> google.protobuf.Empty
	1, // 5: courier.CourierServer.GetDeliveryService:output_type -> courier.ServiceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_courierServer_proto_init() }
func file_courierServer_proto_init() {
	if File_courierServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courierServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCourierServer); i {
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
		file_courierServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_courierServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryService); i {
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
			RawDescriptor: file_courierServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courierServer_proto_goTypes,
		DependencyIndexes: file_courierServer_proto_depIdxs,
		MessageInfos:      file_courierServer_proto_msgTypes,
	}.Build()
	File_courierServer_proto = out.File
	file_courierServer_proto_rawDesc = nil
	file_courierServer_proto_goTypes = nil
	file_courierServer_proto_depIdxs = nil
}

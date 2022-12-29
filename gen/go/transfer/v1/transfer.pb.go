// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: transfer/v1/transfer.proto

package transferv1

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

type TransferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAddress   string `protobuf:"bytes,1,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	SenderPassword  string `protobuf:"bytes,2,opt,name=sender_password,json=senderPassword,proto3" json:"sender_password,omitempty"`
	ReceiverAddress string `protobuf:"bytes,3,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	Amount          int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferReq) Reset() {
	*x = TransferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferReq) ProtoMessage() {}

func (x *TransferReq) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferReq.ProtoReflect.Descriptor instead.
func (*TransferReq) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferReq) GetSenderAddress() string {
	if x != nil {
		return x.SenderAddress
	}
	return ""
}

func (x *TransferReq) GetSenderPassword() string {
	if x != nil {
		return x.SenderPassword
	}
	return ""
}

func (x *TransferReq) GetReceiverAddress() string {
	if x != nil {
		return x.ReceiverAddress
	}
	return ""
}

func (x *TransferReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TransferResp) Reset() {
	*x = TransferResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResp) ProtoMessage() {}

func (x *TransferResp) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResp.ProtoReflect.Descriptor instead.
func (*TransferResp) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ApproveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress    string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	OwnerPassword   string `protobuf:"bytes,2,opt,name=owner_password,json=ownerPassword,proto3" json:"owner_password,omitempty"`
	DelegateAddress string `protobuf:"bytes,3,opt,name=delegate_address,json=delegateAddress,proto3" json:"delegate_address,omitempty"`
	Amount          int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ApproveReq) Reset() {
	*x = ApproveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveReq) ProtoMessage() {}

func (x *ApproveReq) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveReq.ProtoReflect.Descriptor instead.
func (*ApproveReq) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *ApproveReq) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *ApproveReq) GetOwnerPassword() string {
	if x != nil {
		return x.OwnerPassword
	}
	return ""
}

func (x *ApproveReq) GetDelegateAddress() string {
	if x != nil {
		return x.DelegateAddress
	}
	return ""
}

func (x *ApproveReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ApproveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ApproveResp) Reset() {
	*x = ApproveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveResp) ProtoMessage() {}

func (x *ApproveResp) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveResp.ProtoReflect.Descriptor instead.
func (*ApproveResp) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AllowanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress    string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	DelegateAddress string `protobuf:"bytes,2,opt,name=delegate_address,json=delegateAddress,proto3" json:"delegate_address,omitempty"`
}

func (x *AllowanceReq) Reset() {
	*x = AllowanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceReq) ProtoMessage() {}

func (x *AllowanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceReq.ProtoReflect.Descriptor instead.
func (*AllowanceReq) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *AllowanceReq) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *AllowanceReq) GetDelegateAddress() string {
	if x != nil {
		return x.DelegateAddress
	}
	return ""
}

type AllowanceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AllowanceResp) Reset() {
	*x = AllowanceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceResp) ProtoMessage() {}

func (x *AllowanceResp) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceResp.ProtoReflect.Descriptor instead.
func (*AllowanceResp) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *AllowanceResp) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferFromReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegateAddress  string `protobuf:"bytes,1,opt,name=delegate_address,json=delegateAddress,proto3" json:"delegate_address,omitempty"`
	DelegatePassword string `protobuf:"bytes,2,opt,name=delegate_password,json=delegatePassword,proto3" json:"delegate_password,omitempty"`
	OwnerAddress     string `protobuf:"bytes,3,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	BuyerAddress     string `protobuf:"bytes,4,opt,name=buyer_address,json=buyerAddress,proto3" json:"buyer_address,omitempty"`
	Amount           int64  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferFromReq) Reset() {
	*x = TransferFromReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFromReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFromReq) ProtoMessage() {}

func (x *TransferFromReq) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFromReq.ProtoReflect.Descriptor instead.
func (*TransferFromReq) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{6}
}

func (x *TransferFromReq) GetDelegateAddress() string {
	if x != nil {
		return x.DelegateAddress
	}
	return ""
}

func (x *TransferFromReq) GetDelegatePassword() string {
	if x != nil {
		return x.DelegatePassword
	}
	return ""
}

func (x *TransferFromReq) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *TransferFromReq) GetBuyerAddress() string {
	if x != nil {
		return x.BuyerAddress
	}
	return ""
}

func (x *TransferFromReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferFromResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TransferFromResp) Reset() {
	*x = TransferFromResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_v1_transfer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFromResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFromResp) ProtoMessage() {}

func (x *TransferFromResp) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_v1_transfer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFromResp.ProtoReflect.Descriptor instead.
func (*TransferFromResp) Descriptor() ([]byte, []int) {
	return file_transfer_v1_transfer_proto_rawDescGZIP(), []int{7}
}

func (x *TransferFromResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_transfer_v1_transfer_proto protoreflect.FileDescriptor

var file_transfer_v1_transfer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a,
	0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a,
	0x0d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x32, 0xa1, 0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x42, 0xac, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x73, 0x61, 0x70, 0x6f, 0x6c,
	0x6c, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2d, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58,
	0xaa, 0x02, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_v1_transfer_proto_rawDescOnce sync.Once
	file_transfer_v1_transfer_proto_rawDescData = file_transfer_v1_transfer_proto_rawDesc
)

func file_transfer_v1_transfer_proto_rawDescGZIP() []byte {
	file_transfer_v1_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_v1_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_v1_transfer_proto_rawDescData)
	})
	return file_transfer_v1_transfer_proto_rawDescData
}

var file_transfer_v1_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_transfer_v1_transfer_proto_goTypes = []interface{}{
	(*TransferReq)(nil),      // 0: transfer.v1.TransferReq
	(*TransferResp)(nil),     // 1: transfer.v1.TransferResp
	(*ApproveReq)(nil),       // 2: transfer.v1.ApproveReq
	(*ApproveResp)(nil),      // 3: transfer.v1.ApproveResp
	(*AllowanceReq)(nil),     // 4: transfer.v1.AllowanceReq
	(*AllowanceResp)(nil),    // 5: transfer.v1.AllowanceResp
	(*TransferFromReq)(nil),  // 6: transfer.v1.TransferFromReq
	(*TransferFromResp)(nil), // 7: transfer.v1.TransferFromResp
}
var file_transfer_v1_transfer_proto_depIdxs = []int32{
	0, // 0: transfer.v1.TransferService.Transfer:input_type -> transfer.v1.TransferReq
	2, // 1: transfer.v1.TransferService.Approve:input_type -> transfer.v1.ApproveReq
	4, // 2: transfer.v1.TransferService.Allowance:input_type -> transfer.v1.AllowanceReq
	6, // 3: transfer.v1.TransferService.TransferFrom:input_type -> transfer.v1.TransferFromReq
	1, // 4: transfer.v1.TransferService.Transfer:output_type -> transfer.v1.TransferResp
	3, // 5: transfer.v1.TransferService.Approve:output_type -> transfer.v1.ApproveResp
	5, // 6: transfer.v1.TransferService.Allowance:output_type -> transfer.v1.AllowanceResp
	7, // 7: transfer.v1.TransferService.TransferFrom:output_type -> transfer.v1.TransferFromResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_v1_transfer_proto_init() }
func file_transfer_v1_transfer_proto_init() {
	if File_transfer_v1_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_v1_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferReq); i {
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
		file_transfer_v1_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResp); i {
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
		file_transfer_v1_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveReq); i {
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
		file_transfer_v1_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveResp); i {
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
		file_transfer_v1_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceReq); i {
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
		file_transfer_v1_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceResp); i {
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
		file_transfer_v1_transfer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFromReq); i {
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
		file_transfer_v1_transfer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFromResp); i {
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
			RawDescriptor: file_transfer_v1_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_v1_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_v1_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_v1_transfer_proto_msgTypes,
	}.Build()
	File_transfer_v1_transfer_proto = out.File
	file_transfer_v1_transfer_proto_rawDesc = nil
	file_transfer_v1_transfer_proto_goTypes = nil
	file_transfer_v1_transfer_proto_depIdxs = nil
}

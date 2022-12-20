// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: balance/v1/balance.proto

package balancev1

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

type GetBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetBalanceReq) Reset() {
	*x = GetBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_balance_v1_balance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceReq) ProtoMessage() {}

func (x *GetBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_balance_v1_balance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceReq.ProtoReflect.Descriptor instead.
func (*GetBalanceReq) Descriptor() ([]byte, []int) {
	return file_balance_v1_balance_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetBalanceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetBalanceResp) Reset() {
	*x = GetBalanceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_balance_v1_balance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResp) ProtoMessage() {}

func (x *GetBalanceResp) ProtoReflect() protoreflect.Message {
	mi := &file_balance_v1_balance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResp.ProtoReflect.Descriptor instead.
func (*GetBalanceResp) Descriptor() ([]byte, []int) {
	return file_balance_v1_balance_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceResp) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_balance_v1_balance_proto protoreflect.FileDescriptor

var file_balance_v1_balance_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x55, 0x0a, 0x0e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x42, 0xa4, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x76, 0x61, 0x73, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2f, 0x65, 0x74, 0x68,
	0x2d, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x16, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_balance_v1_balance_proto_rawDescOnce sync.Once
	file_balance_v1_balance_proto_rawDescData = file_balance_v1_balance_proto_rawDesc
)

func file_balance_v1_balance_proto_rawDescGZIP() []byte {
	file_balance_v1_balance_proto_rawDescOnce.Do(func() {
		file_balance_v1_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_balance_v1_balance_proto_rawDescData)
	})
	return file_balance_v1_balance_proto_rawDescData
}

var file_balance_v1_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_balance_v1_balance_proto_goTypes = []interface{}{
	(*GetBalanceReq)(nil),  // 0: balance.v1.GetBalanceReq
	(*GetBalanceResp)(nil), // 1: balance.v1.GetBalanceResp
}
var file_balance_v1_balance_proto_depIdxs = []int32{
	0, // 0: balance.v1.BalanceService.GetBalance:input_type -> balance.v1.GetBalanceReq
	1, // 1: balance.v1.BalanceService.GetBalance:output_type -> balance.v1.GetBalanceResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_balance_v1_balance_proto_init() }
func file_balance_v1_balance_proto_init() {
	if File_balance_v1_balance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_balance_v1_balance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceReq); i {
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
		file_balance_v1_balance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResp); i {
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
			RawDescriptor: file_balance_v1_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_balance_v1_balance_proto_goTypes,
		DependencyIndexes: file_balance_v1_balance_proto_depIdxs,
		MessageInfos:      file_balance_v1_balance_proto_msgTypes,
	}.Build()
	File_balance_v1_balance_proto = out.File
	file_balance_v1_balance_proto_rawDesc = nil
	file_balance_v1_balance_proto_goTypes = nil
	file_balance_v1_balance_proto_depIdxs = nil
}

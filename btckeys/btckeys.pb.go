// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: btckeys.proto

package btckeys

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

type MultiSigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkeys []string `protobuf:"bytes,1,rep,name=pubkeys,proto3" json:"pubkeys,omitempty"`
	M       uint64   `protobuf:"varint,2,opt,name=m,proto3" json:"m,omitempty"`
	N       uint64   `protobuf:"varint,3,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *MultiSigRequest) Reset() {
	*x = MultiSigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btckeys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiSigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiSigRequest) ProtoMessage() {}

func (x *MultiSigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_btckeys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiSigRequest.ProtoReflect.Descriptor instead.
func (*MultiSigRequest) Descriptor() ([]byte, []int) {
	return file_btckeys_proto_rawDescGZIP(), []int{0}
}

func (x *MultiSigRequest) GetPubkeys() []string {
	if x != nil {
		return x.Pubkeys
	}
	return nil
}

func (x *MultiSigRequest) GetM() uint64 {
	if x != nil {
		return x.M
	}
	return 0
}

func (x *MultiSigRequest) GetN() uint64 {
	if x != nil {
		return x.N
	}
	return 0
}

type MultiSigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Redeemscript string `protobuf:"bytes,2,opt,name=redeemscript,proto3" json:"redeemscript,omitempty"`
}

func (x *MultiSigResponse) Reset() {
	*x = MultiSigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btckeys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiSigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiSigResponse) ProtoMessage() {}

func (x *MultiSigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_btckeys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiSigResponse.ProtoReflect.Descriptor instead.
func (*MultiSigResponse) Descriptor() ([]byte, []int) {
	return file_btckeys_proto_rawDescGZIP(), []int{1}
}

func (x *MultiSigResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MultiSigResponse) GetRedeemscript() string {
	if x != nil {
		return x.Redeemscript
	}
	return ""
}

type DerivationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xpub string `protobuf:"bytes,1,opt,name=xpub,proto3" json:"xpub,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DerivationRequest) Reset() {
	*x = DerivationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btckeys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivationRequest) ProtoMessage() {}

func (x *DerivationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_btckeys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivationRequest.ProtoReflect.Descriptor instead.
func (*DerivationRequest) Descriptor() ([]byte, []int) {
	return file_btckeys_proto_rawDescGZIP(), []int{2}
}

func (x *DerivationRequest) GetXpub() string {
	if x != nil {
		return x.Xpub
	}
	return ""
}

func (x *DerivationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btckeys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_btckeys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_btckeys_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_btckeys_proto protoreflect.FileDescriptor

var file_btckeys_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x47, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x01, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01,
	0x6e, 0x22, 0x50, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x70, 0x75, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x78, 0x70, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x23, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x07, 0x42, 0x74, 0x63, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x4d, 0x0a, 0x1b, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x42, 0x65, 0x63, 0x68, 0x33,
	0x32, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x58, 0x70, 0x75, 0x62,
	0x12, 0x1a, 0x2e, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x44, 0x65, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62,
	0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x69, 0x67, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x53, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x75, 0x67, 0x67,
	0x6c, 0x65, 0x2d, 0x44, 0x75, 0x2f, 0x62, 0x74, 0x63, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x62, 0x74,
	0x63, 0x6b, 0x65, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_btckeys_proto_rawDescOnce sync.Once
	file_btckeys_proto_rawDescData = file_btckeys_proto_rawDesc
)

func file_btckeys_proto_rawDescGZIP() []byte {
	file_btckeys_proto_rawDescOnce.Do(func() {
		file_btckeys_proto_rawDescData = protoimpl.X.CompressGZIP(file_btckeys_proto_rawDescData)
	})
	return file_btckeys_proto_rawDescData
}

var file_btckeys_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_btckeys_proto_goTypes = []interface{}{
	(*MultiSigRequest)(nil),   // 0: btckeys.MultiSigRequest
	(*MultiSigResponse)(nil),  // 1: btckeys.MultiSigResponse
	(*DerivationRequest)(nil), // 2: btckeys.DerivationRequest
	(*Address)(nil),           // 3: btckeys.Address
}
var file_btckeys_proto_depIdxs = []int32{
	2, // 0: btckeys.BtcKeys.DeriveBech32AddressFromXpub:input_type -> btckeys.DerivationRequest
	0, // 1: btckeys.BtcKeys.GetMultiSigAddress:input_type -> btckeys.MultiSigRequest
	3, // 2: btckeys.BtcKeys.DeriveBech32AddressFromXpub:output_type -> btckeys.Address
	1, // 3: btckeys.BtcKeys.GetMultiSigAddress:output_type -> btckeys.MultiSigResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_btckeys_proto_init() }
func file_btckeys_proto_init() {
	if File_btckeys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_btckeys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiSigRequest); i {
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
		file_btckeys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiSigResponse); i {
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
		file_btckeys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivationRequest); i {
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
		file_btckeys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_btckeys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_btckeys_proto_goTypes,
		DependencyIndexes: file_btckeys_proto_depIdxs,
		MessageInfos:      file_btckeys_proto_msgTypes,
	}.Build()
	File_btckeys_proto = out.File
	file_btckeys_proto_rawDesc = nil
	file_btckeys_proto_goTypes = nil
	file_btckeys_proto_depIdxs = nil
}

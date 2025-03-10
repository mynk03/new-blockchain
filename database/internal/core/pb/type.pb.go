// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: type.proto

package _type

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Index         uint64                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Transactions  []*Transaction         `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Validator     []byte                 `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
	PrevHash      []byte                 `protobuf:"bytes,5,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	Hash          []byte                 `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	StateRoot     []byte                 `protobuf:"bytes,7,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Block) GetValidator() []byte {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *Block) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *Block) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Block) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       []byte                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance       uint64                 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Stake         uint64                 `protobuf:"varint,3,opt,name=stake,proto3" json:"stake,omitempty"`
	Nonce         uint64                 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Account) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Account) GetStake() uint64 {
	if x != nil {
		return x.Stake
	}
	return 0
}

func (x *Account) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          []byte                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            []byte                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount        uint64                 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee           uint64                 `protobuf:"varint,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Nonce         uint64                 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasLimit      uint64                 `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice      uint64                 `protobuf:"varint,7,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Signature     []byte                 `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	PublicKey     []byte                 `protobuf:"bytes,9,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // Serialized ECDSA public key
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Transaction) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Transaction) GetGasPrice() uint64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *Transaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Transaction) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

var File_type_proto protoreflect.FileDescriptor

var file_type_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x22, 0x69, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0xe8, 0x01,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_type_proto_rawDescOnce sync.Once
	file_type_proto_rawDescData []byte
)

func file_type_proto_rawDescGZIP() []byte {
	file_type_proto_rawDescOnce.Do(func() {
		file_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_type_proto_rawDesc), len(file_type_proto_rawDesc)))
	})
	return file_type_proto_rawDescData
}

var file_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_type_proto_goTypes = []any{
	(*Block)(nil),                 // 0: core.Block
	(*Account)(nil),               // 1: core.Account
	(*Transaction)(nil),           // 2: core.Transaction
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_type_proto_depIdxs = []int32{
	3, // 0: core.Block.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: core.Block.transactions:type_name -> core.Transaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_type_proto_init() }
func file_type_proto_init() {
	if File_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_type_proto_rawDesc), len(file_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_proto_goTypes,
		DependencyIndexes: file_type_proto_depIdxs,
		MessageInfos:      file_type_proto_msgTypes,
	}.Build()
	File_type_proto = out.File
	file_type_proto_goTypes = nil
	file_type_proto_depIdxs = nil
}

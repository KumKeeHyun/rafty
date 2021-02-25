// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: appendEntries.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EntryType int32

const (
	EntryType_Command      EntryType = 0 // 일반적인 쓰기 작업
	EntryType_ConfigChange EntryType = 1 // 클러스터 구성 변경
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "Command",
		1: "ConfigChange",
	}
	EntryType_value = map[string]int32{
		"Command":      0,
		"ConfigChange": 1,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_appendEntries_proto_enumTypes[0].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_appendEntries_proto_enumTypes[0]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_appendEntries_proto_rawDescGZIP(), []int{0}
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term      uint64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Index     uint64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	RequestID uint64    `protobuf:"varint,3,opt,name=requestID,proto3" json:"requestID,omitempty"` // 클라이언트의 요청 ID
	Type      EntryType `protobuf:"varint,4,opt,name=type,proto3,enum=protobuf.EntryType" json:"type,omitempty"`
	Data      []byte    `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appendEntries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_appendEntries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_appendEntries_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Entry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Entry) GetRequestID() uint64 {
	if x != nil {
		return x.RequestID
	}
	return 0
}

func (x *Entry) GetType() EntryType {
	if x != nil {
		return x.Type
	}
	return EntryType_Command
}

func (x *Entry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AppendEntriesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Leader       uint64   `protobuf:"varint,2,opt,name=leader,proto3" json:"leader,omitempty"`
	PrevLogIndex uint64   `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  uint64   `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	LeaderCommit uint64   `protobuf:"varint,5,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
	Entries      []*Entry `protobuf:"bytes,6,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *AppendEntriesReq) Reset() {
	*x = AppendEntriesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appendEntries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesReq) ProtoMessage() {}

func (x *AppendEntriesReq) ProtoReflect() protoreflect.Message {
	mi := &file_appendEntries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesReq.ProtoReflect.Descriptor instead.
func (*AppendEntriesReq) Descriptor() ([]byte, []int) {
	return file_appendEntries_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntriesReq) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesReq) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *AppendEntriesReq) GetPrevLogIndex() uint64 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesReq) GetPrevLogTerm() uint64 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesReq) GetLeaderCommit() uint64 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

func (x *AppendEntriesReq) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AppendEntriesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Lag     uint64 `protobuf:"varint,3,opt,name=lag,proto3" json:"lag,omitempty"` // 요청의 entries의 Index가 로컬 로그와 얼마나 차이나는지
}

func (x *AppendEntriesResp) Reset() {
	*x = AppendEntriesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appendEntries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResp) ProtoMessage() {}

func (x *AppendEntriesResp) ProtoReflect() protoreflect.Message {
	mi := &file_appendEntries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResp.ProtoReflect.Descriptor instead.
func (*AppendEntriesResp) Descriptor() ([]byte, []int) {
	return file_appendEntries_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntriesResp) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AppendEntriesResp) GetLag() uint64 {
	if x != nil {
		return x.Lag
	}
	return 0
}

var File_appendEntries_proto protoreflect.FileDescriptor

var file_appendEntries_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x8c, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd3,
	0x01, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65,
	0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f,
	0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x61, 0x67, 0x2a, 0x2a, 0x0a, 0x09, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appendEntries_proto_rawDescOnce sync.Once
	file_appendEntries_proto_rawDescData = file_appendEntries_proto_rawDesc
)

func file_appendEntries_proto_rawDescGZIP() []byte {
	file_appendEntries_proto_rawDescOnce.Do(func() {
		file_appendEntries_proto_rawDescData = protoimpl.X.CompressGZIP(file_appendEntries_proto_rawDescData)
	})
	return file_appendEntries_proto_rawDescData
}

var file_appendEntries_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_appendEntries_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appendEntries_proto_goTypes = []interface{}{
	(EntryType)(0),            // 0: protobuf.EntryType
	(*Entry)(nil),             // 1: protobuf.Entry
	(*AppendEntriesReq)(nil),  // 2: protobuf.AppendEntriesReq
	(*AppendEntriesResp)(nil), // 3: protobuf.AppendEntriesResp
}
var file_appendEntries_proto_depIdxs = []int32{
	0, // 0: protobuf.Entry.type:type_name -> protobuf.EntryType
	1, // 1: protobuf.AppendEntriesReq.entries:type_name -> protobuf.Entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_appendEntries_proto_init() }
func file_appendEntries_proto_init() {
	if File_appendEntries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appendEntries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_appendEntries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesReq); i {
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
		file_appendEntries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesResp); i {
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
			RawDescriptor: file_appendEntries_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_appendEntries_proto_goTypes,
		DependencyIndexes: file_appendEntries_proto_depIdxs,
		EnumInfos:         file_appendEntries_proto_enumTypes,
		MessageInfos:      file_appendEntries_proto_msgTypes,
	}.Build()
	File_appendEntries_proto = out.File
	file_appendEntries_proto_rawDesc = nil
	file_appendEntries_proto_goTypes = nil
	file_appendEntries_proto_depIdxs = nil
}
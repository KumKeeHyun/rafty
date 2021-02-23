// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: installSnapshot.proto

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

type InstallSnapshotReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderID     int64  `protobuf:"varint,2,opt,name=leaderID,proto3" json:"leaderID,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
	Done         bool   `protobuf:"varint,5,opt,name=done,proto3" json:"done,omitempty"`
	Offset       uint64 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	DataChunk    []byte `protobuf:"bytes,7,opt,name=dataChunk,proto3" json:"dataChunk,omitempty"`
}

func (x *InstallSnapshotReq) Reset() {
	*x = InstallSnapshotReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installSnapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotReq) ProtoMessage() {}

func (x *InstallSnapshotReq) ProtoReflect() protoreflect.Message {
	mi := &file_installSnapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotReq.ProtoReflect.Descriptor instead.
func (*InstallSnapshotReq) Descriptor() ([]byte, []int) {
	return file_installSnapshot_proto_rawDescGZIP(), []int{0}
}

func (x *InstallSnapshotReq) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *InstallSnapshotReq) GetLeaderID() int64 {
	if x != nil {
		return x.LeaderID
	}
	return 0
}

func (x *InstallSnapshotReq) GetLastLogIndex() uint64 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *InstallSnapshotReq) GetLastLogTerm() uint64 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

func (x *InstallSnapshotReq) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *InstallSnapshotReq) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *InstallSnapshotReq) GetDataChunk() []byte {
	if x != nil {
		return x.DataChunk
	}
	return nil
}

type InstallSnapshotResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *InstallSnapshotResp) Reset() {
	*x = InstallSnapshotResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installSnapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallSnapshotResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallSnapshotResp) ProtoMessage() {}

func (x *InstallSnapshotResp) ProtoReflect() protoreflect.Message {
	mi := &file_installSnapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallSnapshotResp.ProtoReflect.Descriptor instead.
func (*InstallSnapshotResp) Descriptor() ([]byte, []int) {
	return file_installSnapshot_proto_rawDescGZIP(), []int{1}
}

func (x *InstallSnapshotResp) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_installSnapshot_proto protoreflect.FileDescriptor

var file_installSnapshot_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0xd4, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x29, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_installSnapshot_proto_rawDescOnce sync.Once
	file_installSnapshot_proto_rawDescData = file_installSnapshot_proto_rawDesc
)

func file_installSnapshot_proto_rawDescGZIP() []byte {
	file_installSnapshot_proto_rawDescOnce.Do(func() {
		file_installSnapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_installSnapshot_proto_rawDescData)
	})
	return file_installSnapshot_proto_rawDescData
}

var file_installSnapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_installSnapshot_proto_goTypes = []interface{}{
	(*InstallSnapshotReq)(nil),  // 0: protobuf.InstallSnapshotReq
	(*InstallSnapshotResp)(nil), // 1: protobuf.InstallSnapshotResp
}
var file_installSnapshot_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_installSnapshot_proto_init() }
func file_installSnapshot_proto_init() {
	if File_installSnapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_installSnapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotReq); i {
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
		file_installSnapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallSnapshotResp); i {
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
			RawDescriptor: file_installSnapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_installSnapshot_proto_goTypes,
		DependencyIndexes: file_installSnapshot_proto_depIdxs,
		MessageInfos:      file_installSnapshot_proto_msgTypes,
	}.Build()
	File_installSnapshot_proto = out.File
	file_installSnapshot_proto_rawDesc = nil
	file_installSnapshot_proto_goTypes = nil
	file_installSnapshot_proto_depIdxs = nil
}

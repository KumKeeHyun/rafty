// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RaftyClient is the client API for Rafty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftyClient interface {
	RequestVote(ctx context.Context, in *RequestVoteReq, opts ...grpc.CallOption) (*RequestVoteResp, error)
	AppendEntries(ctx context.Context, in *AppendEntriesReq, opts ...grpc.CallOption) (*AppendEntriesResp, error)
	InstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (Rafty_InstallSnapshotClient, error)
	Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (*WriteResp, error)
	Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error)
	FindLeader(ctx context.Context, in *FindLeaderReq, opts ...grpc.CallOption) (*FindLeaderResp, error)
}

type raftyClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftyClient(cc grpc.ClientConnInterface) RaftyClient {
	return &raftyClient{cc}
}

func (c *raftyClient) RequestVote(ctx context.Context, in *RequestVoteReq, opts ...grpc.CallOption) (*RequestVoteResp, error) {
	out := new(RequestVoteResp)
	err := c.cc.Invoke(ctx, "/protobuf.Rafty/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftyClient) AppendEntries(ctx context.Context, in *AppendEntriesReq, opts ...grpc.CallOption) (*AppendEntriesResp, error) {
	out := new(AppendEntriesResp)
	err := c.cc.Invoke(ctx, "/protobuf.Rafty/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftyClient) InstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (Rafty_InstallSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rafty_ServiceDesc.Streams[0], "/protobuf.Rafty/InstallSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftyInstallSnapshotClient{stream}
	return x, nil
}

type Rafty_InstallSnapshotClient interface {
	Send(*InstallSnapshotResp) error
	CloseAndRecv() (*InstallSnapshotResp, error)
	grpc.ClientStream
}

type raftyInstallSnapshotClient struct {
	grpc.ClientStream
}

func (x *raftyInstallSnapshotClient) Send(m *InstallSnapshotResp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftyInstallSnapshotClient) CloseAndRecv() (*InstallSnapshotResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InstallSnapshotResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *raftyClient) Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (*WriteResp, error) {
	out := new(WriteResp)
	err := c.cc.Invoke(ctx, "/protobuf.Rafty/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftyClient) Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error) {
	out := new(ReadResp)
	err := c.cc.Invoke(ctx, "/protobuf.Rafty/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftyClient) FindLeader(ctx context.Context, in *FindLeaderReq, opts ...grpc.CallOption) (*FindLeaderResp, error) {
	out := new(FindLeaderResp)
	err := c.cc.Invoke(ctx, "/protobuf.Rafty/FindLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftyServer is the server API for Rafty service.
// All implementations must embed UnimplementedRaftyServer
// for forward compatibility
type RaftyServer interface {
	RequestVote(context.Context, *RequestVoteReq) (*RequestVoteResp, error)
	AppendEntries(context.Context, *AppendEntriesReq) (*AppendEntriesResp, error)
	InstallSnapshot(Rafty_InstallSnapshotServer) error
	Write(context.Context, *WriteReq) (*WriteResp, error)
	Read(context.Context, *ReadReq) (*ReadResp, error)
	FindLeader(context.Context, *FindLeaderReq) (*FindLeaderResp, error)
	mustEmbedUnimplementedRaftyServer()
}

// UnimplementedRaftyServer must be embedded to have forward compatible implementations.
type UnimplementedRaftyServer struct {
}

func (UnimplementedRaftyServer) RequestVote(context.Context, *RequestVoteReq) (*RequestVoteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedRaftyServer) AppendEntries(context.Context, *AppendEntriesReq) (*AppendEntriesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedRaftyServer) InstallSnapshot(Rafty_InstallSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}
func (UnimplementedRaftyServer) Write(context.Context, *WriteReq) (*WriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedRaftyServer) Read(context.Context, *ReadReq) (*ReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRaftyServer) FindLeader(context.Context, *FindLeaderReq) (*FindLeaderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLeader not implemented")
}
func (UnimplementedRaftyServer) mustEmbedUnimplementedRaftyServer() {}

// UnsafeRaftyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftyServer will
// result in compilation errors.
type UnsafeRaftyServer interface {
	mustEmbedUnimplementedRaftyServer()
}

func RegisterRaftyServer(s grpc.ServiceRegistrar, srv RaftyServer) {
	s.RegisterService(&Rafty_ServiceDesc, srv)
}

func _Rafty_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftyServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Rafty/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftyServer).RequestVote(ctx, req.(*RequestVoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rafty_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftyServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Rafty/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftyServer).AppendEntries(ctx, req.(*AppendEntriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rafty_InstallSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftyServer).InstallSnapshot(&raftyInstallSnapshotServer{stream})
}

type Rafty_InstallSnapshotServer interface {
	SendAndClose(*InstallSnapshotResp) error
	Recv() (*InstallSnapshotResp, error)
	grpc.ServerStream
}

type raftyInstallSnapshotServer struct {
	grpc.ServerStream
}

func (x *raftyInstallSnapshotServer) SendAndClose(m *InstallSnapshotResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftyInstallSnapshotServer) Recv() (*InstallSnapshotResp, error) {
	m := new(InstallSnapshotResp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Rafty_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftyServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Rafty/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftyServer).Write(ctx, req.(*WriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rafty_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftyServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Rafty/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftyServer).Read(ctx, req.(*ReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rafty_FindLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLeaderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftyServer).FindLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Rafty/FindLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftyServer).FindLeader(ctx, req.(*FindLeaderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rafty_ServiceDesc is the grpc.ServiceDesc for Rafty service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rafty_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Rafty",
	HandlerType: (*RaftyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _Rafty_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _Rafty_AppendEntries_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Rafty_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Rafty_Read_Handler,
		},
		{
			MethodName: "FindLeader",
			Handler:    _Rafty_FindLeader_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InstallSnapshot",
			Handler:       _Rafty_InstallSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "raft.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	MakeRequests(ctx context.Context, opts ...grpc.CallOption) (ChatService_MakeRequestsClient, error)
	ConsumeMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_ConsumeMessagesClient, error)
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ServerResponse, error)
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ChatMessage, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) MakeRequests(ctx context.Context, opts ...grpc.CallOption) (ChatService_MakeRequestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/grpc.ChatService/MakeRequests", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceMakeRequestsClient{stream}
	return x, nil
}

type ChatService_MakeRequestsClient interface {
	Send(*ClientRequest) error
	Recv() (*ServerResponse, error)
	grpc.ClientStream
}

type chatServiceMakeRequestsClient struct {
	grpc.ClientStream
}

func (x *chatServiceMakeRequestsClient) Send(m *ClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceMakeRequestsClient) Recv() (*ServerResponse, error) {
	m := new(ServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ConsumeMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_ConsumeMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/grpc.ChatService/ConsumeMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceConsumeMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ConsumeMessagesClient interface {
	Recv() (*ServerResponse, error)
	grpc.ClientStream
}

type chatServiceConsumeMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceConsumeMessagesClient) Recv() (*ServerResponse, error) {
	m := new(ServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/grpc.ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessages(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ChatMessage, error) {
	out := new(ChatMessage)
	err := c.cc.Invoke(ctx, "/grpc.ChatService/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	MakeRequests(ChatService_MakeRequestsServer) error
	ConsumeMessages(*empty.Empty, ChatService_ConsumeMessagesServer) error
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	SendMessage(context.Context, *ChatMessage) (*ServerResponse, error)
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetMessages(context.Context, *empty.Empty) (*ChatMessage, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) MakeRequests(ChatService_MakeRequestsServer) error {
	return status.Errorf(codes.Unimplemented, "method MakeRequests not implemented")
}
func (UnimplementedChatServiceServer) ConsumeMessages(*empty.Empty, ChatService_ConsumeMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeMessages not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *ChatMessage) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetMessages(context.Context, *empty.Empty) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_MakeRequests_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).MakeRequests(&chatServiceMakeRequestsServer{stream})
}

type ChatService_MakeRequestsServer interface {
	Send(*ServerResponse) error
	Recv() (*ClientRequest, error)
	grpc.ServerStream
}

type chatServiceMakeRequestsServer struct {
	grpc.ServerStream
}

func (x *chatServiceMakeRequestsServer) Send(m *ServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceMakeRequestsServer) Recv() (*ClientRequest, error) {
	m := new(ClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_ConsumeMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ConsumeMessages(m, &chatServiceConsumeMessagesServer{stream})
}

type ChatService_ConsumeMessagesServer interface {
	Send(*ServerResponse) error
	grpc.ServerStream
}

type chatServiceConsumeMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceConsumeMessagesServer) Send(m *ServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ChatService/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetMessages(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _ChatService_GetMessages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeRequests",
			Handler:       _ChatService_MakeRequests_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ConsumeMessages",
			Handler:       _ChatService_ConsumeMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/test.proto",
}

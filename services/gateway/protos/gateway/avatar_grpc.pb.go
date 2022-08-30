// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// AvatarClient is the client API for Avatar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvatarClient interface {
	StreamOperatorSideVideo(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamOperatorSideVideoClient, error)
	StreamOperatorSideVoice(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamOperatorSideVoiceClient, error)
	StreamPOSSideVideo(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamPOSSideVideoClient, error)
	StreamPOSSideVoice(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamPOSSideVoiceClient, error)
	SpeechToText(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_SpeechToTextClient, error)
	ListenEventPOSSide(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenEventPOSSideClient, error)
	ListenEventOperatorSide(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenEventOperatorSideClient, error)
	ListenNotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenNotesClient, error)
	ListenListPos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenListPosClient, error)
}

type avatarClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarClient(cc grpc.ClientConnInterface) AvatarClient {
	return &avatarClient{cc}
}

func (c *avatarClient) StreamOperatorSideVideo(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamOperatorSideVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[0], "/avatar.Avatar/StreamOperatorSideVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarStreamOperatorSideVideoClient{stream}
	return x, nil
}

type Avatar_StreamOperatorSideVideoClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type avatarStreamOperatorSideVideoClient struct {
	grpc.ClientStream
}

func (x *avatarStreamOperatorSideVideoClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avatarStreamOperatorSideVideoClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) StreamOperatorSideVoice(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamOperatorSideVoiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[1], "/avatar.Avatar/StreamOperatorSideVoice", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarStreamOperatorSideVoiceClient{stream}
	return x, nil
}

type Avatar_StreamOperatorSideVoiceClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type avatarStreamOperatorSideVoiceClient struct {
	grpc.ClientStream
}

func (x *avatarStreamOperatorSideVoiceClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avatarStreamOperatorSideVoiceClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) StreamPOSSideVideo(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamPOSSideVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[2], "/avatar.Avatar/StreamPOSSideVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarStreamPOSSideVideoClient{stream}
	return x, nil
}

type Avatar_StreamPOSSideVideoClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type avatarStreamPOSSideVideoClient struct {
	grpc.ClientStream
}

func (x *avatarStreamPOSSideVideoClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avatarStreamPOSSideVideoClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) StreamPOSSideVoice(ctx context.Context, opts ...grpc.CallOption) (Avatar_StreamPOSSideVoiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[3], "/avatar.Avatar/StreamPOSSideVoice", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarStreamPOSSideVoiceClient{stream}
	return x, nil
}

type Avatar_StreamPOSSideVoiceClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type avatarStreamPOSSideVoiceClient struct {
	grpc.ClientStream
}

func (x *avatarStreamPOSSideVoiceClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *avatarStreamPOSSideVoiceClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) SpeechToText(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_SpeechToTextClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[4], "/avatar.Avatar/SpeechToText", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarSpeechToTextClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Avatar_SpeechToTextClient interface {
	Recv() (*SpeechToTextData, error)
	grpc.ClientStream
}

type avatarSpeechToTextClient struct {
	grpc.ClientStream
}

func (x *avatarSpeechToTextClient) Recv() (*SpeechToTextData, error) {
	m := new(SpeechToTextData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) ListenEventPOSSide(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenEventPOSSideClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[5], "/avatar.Avatar/ListenEventPOSSide", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarListenEventPOSSideClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Avatar_ListenEventPOSSideClient interface {
	Recv() (*ListenEventPOSSideResponse, error)
	grpc.ClientStream
}

type avatarListenEventPOSSideClient struct {
	grpc.ClientStream
}

func (x *avatarListenEventPOSSideClient) Recv() (*ListenEventPOSSideResponse, error) {
	m := new(ListenEventPOSSideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) ListenEventOperatorSide(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenEventOperatorSideClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[6], "/avatar.Avatar/ListenEventOperatorSide", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarListenEventOperatorSideClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Avatar_ListenEventOperatorSideClient interface {
	Recv() (*ListenEventOperatorSideResponse, error)
	grpc.ClientStream
}

type avatarListenEventOperatorSideClient struct {
	grpc.ClientStream
}

func (x *avatarListenEventOperatorSideClient) Recv() (*ListenEventOperatorSideResponse, error) {
	m := new(ListenEventOperatorSideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) ListenNotes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[7], "/avatar.Avatar/ListenNotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarListenNotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Avatar_ListenNotesClient interface {
	Recv() (*ListenEventResponse, error)
	grpc.ClientStream
}

type avatarListenNotesClient struct {
	grpc.ClientStream
}

func (x *avatarListenNotesClient) Recv() (*ListenEventResponse, error) {
	m := new(ListenEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *avatarClient) ListenListPos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Avatar_ListenListPosClient, error) {
	stream, err := c.cc.NewStream(ctx, &Avatar_ServiceDesc.Streams[8], "/avatar.Avatar/ListenListPos", opts...)
	if err != nil {
		return nil, err
	}
	x := &avatarListenListPosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Avatar_ListenListPosClient interface {
	Recv() (*ListenListPosResponse, error)
	grpc.ClientStream
}

type avatarListenListPosClient struct {
	grpc.ClientStream
}

func (x *avatarListenListPosClient) Recv() (*ListenListPosResponse, error) {
	m := new(ListenListPosResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvatarServer is the server API for Avatar service.
// All implementations must embed UnimplementedAvatarServer
// for forward compatibility
type AvatarServer interface {
	StreamOperatorSideVideo(Avatar_StreamOperatorSideVideoServer) error
	StreamOperatorSideVoice(Avatar_StreamOperatorSideVoiceServer) error
	StreamPOSSideVideo(Avatar_StreamPOSSideVideoServer) error
	StreamPOSSideVoice(Avatar_StreamPOSSideVoiceServer) error
	SpeechToText(*Empty, Avatar_SpeechToTextServer) error
	ListenEventPOSSide(*Empty, Avatar_ListenEventPOSSideServer) error
	ListenEventOperatorSide(*Empty, Avatar_ListenEventOperatorSideServer) error
	ListenNotes(*Empty, Avatar_ListenNotesServer) error
	ListenListPos(*Empty, Avatar_ListenListPosServer) error
	mustEmbedUnimplementedAvatarServer()
}

// UnimplementedAvatarServer must be embedded to have forward compatible implementations.
type UnimplementedAvatarServer struct {
}

func (UnimplementedAvatarServer) StreamOperatorSideVideo(Avatar_StreamOperatorSideVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOperatorSideVideo not implemented")
}
func (UnimplementedAvatarServer) StreamOperatorSideVoice(Avatar_StreamOperatorSideVoiceServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOperatorSideVoice not implemented")
}
func (UnimplementedAvatarServer) StreamPOSSideVideo(Avatar_StreamPOSSideVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPOSSideVideo not implemented")
}
func (UnimplementedAvatarServer) StreamPOSSideVoice(Avatar_StreamPOSSideVoiceServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPOSSideVoice not implemented")
}
func (UnimplementedAvatarServer) SpeechToText(*Empty, Avatar_SpeechToTextServer) error {
	return status.Errorf(codes.Unimplemented, "method SpeechToText not implemented")
}
func (UnimplementedAvatarServer) ListenEventPOSSide(*Empty, Avatar_ListenEventPOSSideServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenEventPOSSide not implemented")
}
func (UnimplementedAvatarServer) ListenEventOperatorSide(*Empty, Avatar_ListenEventOperatorSideServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenEventOperatorSide not implemented")
}
func (UnimplementedAvatarServer) ListenNotes(*Empty, Avatar_ListenNotesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenNotes not implemented")
}
func (UnimplementedAvatarServer) ListenListPos(*Empty, Avatar_ListenListPosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenListPos not implemented")
}
func (UnimplementedAvatarServer) mustEmbedUnimplementedAvatarServer() {}

// UnsafeAvatarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvatarServer will
// result in compilation errors.
type UnsafeAvatarServer interface {
	mustEmbedUnimplementedAvatarServer()
}

func RegisterAvatarServer(s grpc.ServiceRegistrar, srv AvatarServer) {
	s.RegisterService(&Avatar_ServiceDesc, srv)
}

func _Avatar_StreamOperatorSideVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvatarServer).StreamOperatorSideVideo(&avatarStreamOperatorSideVideoServer{stream})
}

type Avatar_StreamOperatorSideVideoServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type avatarStreamOperatorSideVideoServer struct {
	grpc.ServerStream
}

func (x *avatarStreamOperatorSideVideoServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avatarStreamOperatorSideVideoServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Avatar_StreamOperatorSideVoice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvatarServer).StreamOperatorSideVoice(&avatarStreamOperatorSideVoiceServer{stream})
}

type Avatar_StreamOperatorSideVoiceServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type avatarStreamOperatorSideVoiceServer struct {
	grpc.ServerStream
}

func (x *avatarStreamOperatorSideVoiceServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avatarStreamOperatorSideVoiceServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Avatar_StreamPOSSideVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvatarServer).StreamPOSSideVideo(&avatarStreamPOSSideVideoServer{stream})
}

type Avatar_StreamPOSSideVideoServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type avatarStreamPOSSideVideoServer struct {
	grpc.ServerStream
}

func (x *avatarStreamPOSSideVideoServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avatarStreamPOSSideVideoServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Avatar_StreamPOSSideVoice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AvatarServer).StreamPOSSideVoice(&avatarStreamPOSSideVoiceServer{stream})
}

type Avatar_StreamPOSSideVoiceServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type avatarStreamPOSSideVoiceServer struct {
	grpc.ServerStream
}

func (x *avatarStreamPOSSideVoiceServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *avatarStreamPOSSideVoiceServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Avatar_SpeechToText_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvatarServer).SpeechToText(m, &avatarSpeechToTextServer{stream})
}

type Avatar_SpeechToTextServer interface {
	Send(*SpeechToTextData) error
	grpc.ServerStream
}

type avatarSpeechToTextServer struct {
	grpc.ServerStream
}

func (x *avatarSpeechToTextServer) Send(m *SpeechToTextData) error {
	return x.ServerStream.SendMsg(m)
}

func _Avatar_ListenEventPOSSide_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvatarServer).ListenEventPOSSide(m, &avatarListenEventPOSSideServer{stream})
}

type Avatar_ListenEventPOSSideServer interface {
	Send(*ListenEventPOSSideResponse) error
	grpc.ServerStream
}

type avatarListenEventPOSSideServer struct {
	grpc.ServerStream
}

func (x *avatarListenEventPOSSideServer) Send(m *ListenEventPOSSideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Avatar_ListenEventOperatorSide_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvatarServer).ListenEventOperatorSide(m, &avatarListenEventOperatorSideServer{stream})
}

type Avatar_ListenEventOperatorSideServer interface {
	Send(*ListenEventOperatorSideResponse) error
	grpc.ServerStream
}

type avatarListenEventOperatorSideServer struct {
	grpc.ServerStream
}

func (x *avatarListenEventOperatorSideServer) Send(m *ListenEventOperatorSideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Avatar_ListenNotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvatarServer).ListenNotes(m, &avatarListenNotesServer{stream})
}

type Avatar_ListenNotesServer interface {
	Send(*ListenEventResponse) error
	grpc.ServerStream
}

type avatarListenNotesServer struct {
	grpc.ServerStream
}

func (x *avatarListenNotesServer) Send(m *ListenEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Avatar_ListenListPos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvatarServer).ListenListPos(m, &avatarListenListPosServer{stream})
}

type Avatar_ListenListPosServer interface {
	Send(*ListenListPosResponse) error
	grpc.ServerStream
}

type avatarListenListPosServer struct {
	grpc.ServerStream
}

func (x *avatarListenListPosServer) Send(m *ListenListPosResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Avatar_ServiceDesc is the grpc.ServiceDesc for Avatar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Avatar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "avatar.Avatar",
	HandlerType: (*AvatarServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamOperatorSideVideo",
			Handler:       _Avatar_StreamOperatorSideVideo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamOperatorSideVoice",
			Handler:       _Avatar_StreamOperatorSideVoice_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamPOSSideVideo",
			Handler:       _Avatar_StreamPOSSideVideo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamPOSSideVoice",
			Handler:       _Avatar_StreamPOSSideVoice_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SpeechToText",
			Handler:       _Avatar_SpeechToText_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenEventPOSSide",
			Handler:       _Avatar_ListenEventPOSSide_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenEventOperatorSide",
			Handler:       _Avatar_ListenEventOperatorSide_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenNotes",
			Handler:       _Avatar_ListenNotes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenListPos",
			Handler:       _Avatar_ListenListPos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "avatar.proto",
}

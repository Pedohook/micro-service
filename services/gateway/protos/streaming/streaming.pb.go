// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: streaming.proto

package protos

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SpeechToTextData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speaker     string `protobuf:"bytes,1,opt,name=speaker,proto3" json:"speaker,omitempty"`
	Content     string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	SendingTime string `protobuf:"bytes,3,opt,name=sendingTime,proto3" json:"sendingTime,omitempty"`
}

func (x *SpeechToTextData) Reset() {
	*x = SpeechToTextData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechToTextData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechToTextData) ProtoMessage() {}

func (x *SpeechToTextData) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechToTextData.ProtoReflect.Descriptor instead.
func (*SpeechToTextData) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *SpeechToTextData) GetSpeaker() string {
	if x != nil {
		return x.Speaker
	}
	return ""
}

func (x *SpeechToTextData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SpeechToTextData) GetSendingTime() string {
	if x != nil {
		return x.SendingTime
	}
	return ""
}

type ListenEventOperatorSideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ListenEventOperatorSideResponse) Reset() {
	*x = ListenEventOperatorSideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenEventOperatorSideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenEventOperatorSideResponse) ProtoMessage() {}

func (x *ListenEventOperatorSideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenEventOperatorSideResponse.ProtoReflect.Descriptor instead.
func (*ListenEventOperatorSideResponse) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{2}
}

func (x *ListenEventOperatorSideResponse) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *ListenEventOperatorSideResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ListenEventOperatorSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterID string `protobuf:"bytes,1,opt,name=centerID,proto3" json:"centerID,omitempty"`
}

func (x *ListenEventOperatorSideRequest) Reset() {
	*x = ListenEventOperatorSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenEventOperatorSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenEventOperatorSideRequest) ProtoMessage() {}

func (x *ListenEventOperatorSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenEventOperatorSideRequest.ProtoReflect.Descriptor instead.
func (*ListenEventOperatorSideRequest) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{3}
}

func (x *ListenEventOperatorSideRequest) GetCenterID() string {
	if x != nil {
		return x.CenterID
	}
	return ""
}

type ListenEventPOSSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosId string `protobuf:"bytes,1,opt,name=posId,proto3" json:"posId,omitempty"`
}

func (x *ListenEventPOSSideRequest) Reset() {
	*x = ListenEventPOSSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenEventPOSSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenEventPOSSideRequest) ProtoMessage() {}

func (x *ListenEventPOSSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenEventPOSSideRequest.ProtoReflect.Descriptor instead.
func (*ListenEventPOSSideRequest) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{4}
}

func (x *ListenEventPOSSideRequest) GetPosId() string {
	if x != nil {
		return x.PosId
	}
	return ""
}

type ListenEventPOSSideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ListenEventPOSSideResponse) Reset() {
	*x = ListenEventPOSSideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenEventPOSSideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenEventPOSSideResponse) ProtoMessage() {}

func (x *ListenEventPOSSideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenEventPOSSideResponse.ProtoReflect.Descriptor instead.
func (*ListenEventPOSSideResponse) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{5}
}

func (x *ListenEventPOSSideResponse) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *ListenEventPOSSideResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ListenEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   int64  `protobuf:"varint,1,opt,name=event,proto3" json:"event,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ListenEventResponse) Reset() {
	*x = ListenEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenEventResponse) ProtoMessage() {}

func (x *ListenEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenEventResponse.ProtoReflect.Descriptor instead.
func (*ListenEventResponse) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{6}
}

func (x *ListenEventResponse) GetEvent() int64 {
	if x != nil {
		return x.Event
	}
	return 0
}

func (x *ListenEventResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ListenListPosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *ListenListPosRequest) Reset() {
	*x = ListenListPosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenListPosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenListPosRequest) ProtoMessage() {}

func (x *ListenListPosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenListPosRequest.ProtoReflect.Descriptor instead.
func (*ListenListPosRequest) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{7}
}

func (x *ListenListPosRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ListPosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosId            int64   `protobuf:"varint,1,opt,name=posId,proto3" json:"posId,omitempty"`
	TalkSessionId    int64   `protobuf:"varint,2,opt,name=talkSessionId,proto3" json:"talkSessionId,omitempty"`
	Status           int32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	MainPos          bool    `protobuf:"varint,4,opt,name=mainPos,proto3" json:"mainPos,omitempty"`
	CameraId         []int64 `protobuf:"varint,5,rep,packed,name=cameraId,proto3" json:"cameraId,omitempty"`
	DefaultCameraId  int64   `protobuf:"varint,6,opt,name=defaultCameraId,proto3" json:"defaultCameraId,omitempty"`
	ServerUri        string  `protobuf:"bytes,7,opt,name=serverUri,proto3" json:"serverUri,omitempty"`
	Name             string  `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Address          string  `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	StartTimeDeteted string  `protobuf:"bytes,10,opt,name=startTimeDeteted,proto3" json:"startTimeDeteted,omitempty"`
}

func (x *ListPosResponse) Reset() {
	*x = ListPosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPosResponse) ProtoMessage() {}

func (x *ListPosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPosResponse.ProtoReflect.Descriptor instead.
func (*ListPosResponse) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{8}
}

func (x *ListPosResponse) GetPosId() int64 {
	if x != nil {
		return x.PosId
	}
	return 0
}

func (x *ListPosResponse) GetTalkSessionId() int64 {
	if x != nil {
		return x.TalkSessionId
	}
	return 0
}

func (x *ListPosResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListPosResponse) GetMainPos() bool {
	if x != nil {
		return x.MainPos
	}
	return false
}

func (x *ListPosResponse) GetCameraId() []int64 {
	if x != nil {
		return x.CameraId
	}
	return nil
}

func (x *ListPosResponse) GetDefaultCameraId() int64 {
	if x != nil {
		return x.DefaultCameraId
	}
	return 0
}

func (x *ListPosResponse) GetServerUri() string {
	if x != nil {
		return x.ServerUri
	}
	return ""
}

func (x *ListPosResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListPosResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListPosResponse) GetStartTimeDeteted() string {
	if x != nil {
		return x.StartTimeDeteted
	}
	return ""
}

type ListenListPosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListPosResponse []*ListPosResponse `protobuf:"bytes,1,rep,name=ListPosResponse,proto3" json:"ListPosResponse,omitempty"`
}

func (x *ListenListPosResponse) Reset() {
	*x = ListenListPosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streaming_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenListPosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenListPosResponse) ProtoMessage() {}

func (x *ListenListPosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streaming_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenListPosResponse.ProtoReflect.Descriptor instead.
func (*ListenListPosResponse) Descriptor() ([]byte, []int) {
	return file_streaming_proto_rawDescGZIP(), []int{9}
}

func (x *ListenListPosResponse) GetListPosResponse() []*ListPosResponse {
	if x != nil {
		return x.ListPosResponse
	}
	return nil
}

var File_streaming_proto protoreflect.FileDescriptor

var file_streaming_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54,
	0x6f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x51, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x3c, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x31, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6f, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f,
	0x73, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xbd, 0x02, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x6f, 0x73, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x61, 0x6c, 0x6b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x61, 0x6c,
	0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72, 0x69, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x44, 0x65, 0x74, 0x65, 0x74, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd6, 0x05, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x41, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x64, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x17, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x64, 0x65,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a,
	0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x12, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0c, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x65, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x4f, 0x53, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x74, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x69, 0x64, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x49, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streaming_proto_rawDescOnce sync.Once
	file_streaming_proto_rawDescData = file_streaming_proto_rawDesc
)

func file_streaming_proto_rawDescGZIP() []byte {
	file_streaming_proto_rawDescOnce.Do(func() {
		file_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_streaming_proto_rawDescData)
	})
	return file_streaming_proto_rawDescData
}

var file_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_streaming_proto_goTypes = []interface{}{
	(*Data)(nil),                            // 0: streaming.Data
	(*SpeechToTextData)(nil),                // 1: streaming.SpeechToTextData
	(*ListenEventOperatorSideResponse)(nil), // 2: streaming.ListenEventOperatorSideResponse
	(*ListenEventOperatorSideRequest)(nil),  // 3: streaming.ListenEventOperatorSideRequest
	(*ListenEventPOSSideRequest)(nil),       // 4: streaming.ListenEventPOSSideRequest
	(*ListenEventPOSSideResponse)(nil),      // 5: streaming.ListenEventPOSSideResponse
	(*ListenEventResponse)(nil),             // 6: streaming.ListenEventResponse
	(*ListenListPosRequest)(nil),            // 7: streaming.ListenListPosRequest
	(*ListPosResponse)(nil),                 // 8: streaming.ListPosResponse
	(*ListenListPosResponse)(nil),           // 9: streaming.ListenListPosResponse
	(*empty.Empty)(nil),                     // 10: google.protobuf.Empty
}
var file_streaming_proto_depIdxs = []int32{
	8,  // 0: streaming.ListenListPosResponse.ListPosResponse:type_name -> streaming.ListPosResponse
	0,  // 1: streaming.Streaming.StreamOperatorSideVideo:input_type -> streaming.Data
	0,  // 2: streaming.Streaming.StreamOperatorSideVoice:input_type -> streaming.Data
	0,  // 3: streaming.Streaming.StreamPOSSideVideo:input_type -> streaming.Data
	0,  // 4: streaming.Streaming.StreamPOSSideVoice:input_type -> streaming.Data
	10, // 5: streaming.Streaming.SpeechToText:input_type -> google.protobuf.Empty
	4,  // 6: streaming.Streaming.ListenEventPOSSide:input_type -> streaming.ListenEventPOSSideRequest
	3,  // 7: streaming.Streaming.ListenEventOperatorSide:input_type -> streaming.ListenEventOperatorSideRequest
	10, // 8: streaming.Streaming.ListenNotes:input_type -> google.protobuf.Empty
	7,  // 9: streaming.Streaming.ListenListPos:input_type -> streaming.ListenListPosRequest
	0,  // 10: streaming.Streaming.StreamOperatorSideVideo:output_type -> streaming.Data
	0,  // 11: streaming.Streaming.StreamOperatorSideVoice:output_type -> streaming.Data
	0,  // 12: streaming.Streaming.StreamPOSSideVideo:output_type -> streaming.Data
	0,  // 13: streaming.Streaming.StreamPOSSideVoice:output_type -> streaming.Data
	1,  // 14: streaming.Streaming.SpeechToText:output_type -> streaming.SpeechToTextData
	5,  // 15: streaming.Streaming.ListenEventPOSSide:output_type -> streaming.ListenEventPOSSideResponse
	2,  // 16: streaming.Streaming.ListenEventOperatorSide:output_type -> streaming.ListenEventOperatorSideResponse
	6,  // 17: streaming.Streaming.ListenNotes:output_type -> streaming.ListenEventResponse
	9,  // 18: streaming.Streaming.ListenListPos:output_type -> streaming.ListenListPosResponse
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_streaming_proto_init() }
func file_streaming_proto_init() {
	if File_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_streaming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeechToTextData); i {
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
		file_streaming_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenEventOperatorSideResponse); i {
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
		file_streaming_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenEventOperatorSideRequest); i {
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
		file_streaming_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenEventPOSSideRequest); i {
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
		file_streaming_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenEventPOSSideResponse); i {
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
		file_streaming_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenEventResponse); i {
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
		file_streaming_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenListPosRequest); i {
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
		file_streaming_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPosResponse); i {
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
		file_streaming_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenListPosResponse); i {
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
			RawDescriptor: file_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streaming_proto_goTypes,
		DependencyIndexes: file_streaming_proto_depIdxs,
		MessageInfos:      file_streaming_proto_msgTypes,
	}.Build()
	File_streaming_proto = out.File
	file_streaming_proto_rawDesc = nil
	file_streaming_proto_goTypes = nil
	file_streaming_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: proto/hb.proto

package heartbeat_pb

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

type HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bpm      int32  `protobuf:"varint,1,opt,name=bpm,proto3" json:"bpm,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *HeartBeat) Reset() {
	*x = HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeat) ProtoMessage() {}

func (x *HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeat.ProtoReflect.Descriptor instead.
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{0}
}

func (x *HeartBeat) GetBpm() int32 {
	if x != nil {
		return x.Bpm
	}
	return 0
}

func (x *HeartBeat) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type NormalAbnormalHeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bpm int32 `protobuf:"varint,1,opt,name=bpm,proto3" json:"bpm,omitempty"`
}

func (x *NormalAbnormalHeartBeat) Reset() {
	*x = NormalAbnormalHeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalAbnormalHeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalAbnormalHeartBeat) ProtoMessage() {}

func (x *NormalAbnormalHeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalAbnormalHeartBeat.ProtoReflect.Descriptor instead.
func (*NormalAbnormalHeartBeat) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{1}
}

func (x *NormalAbnormalHeartBeat) GetBpm() int32 {
	if x != nil {
		return x.Bpm
	}
	return 0
}

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *HeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{2}
}

func (x *HeartBeatRequest) GetHeartbeat() *HeartBeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type HeartBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HeartBeatResponse) Reset() {
	*x = HeartBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatResponse) ProtoMessage() {}

func (x *HeartBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatResponse.ProtoReflect.Descriptor instead.
func (*HeartBeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{3}
}

func (x *HeartBeatResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type LiveHeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *HeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *LiveHeartBeatRequest) Reset() {
	*x = LiveHeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveHeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveHeartBeatRequest) ProtoMessage() {}

func (x *LiveHeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveHeartBeatRequest.ProtoReflect.Descriptor instead.
func (*LiveHeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{4}
}

func (x *LiveHeartBeatRequest) GetHeartbeat() *HeartBeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type LiveHeartBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LiveHeartBeatResponse) Reset() {
	*x = LiveHeartBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveHeartBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveHeartBeatResponse) ProtoMessage() {}

func (x *LiveHeartBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveHeartBeatResponse.ProtoReflect.Descriptor instead.
func (*LiveHeartBeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{5}
}

func (x *LiveHeartBeatResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type HeartBeatHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *HeartBeatHistoryRequest) Reset() {
	*x = HeartBeatHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatHistoryRequest) ProtoMessage() {}

func (x *HeartBeatHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatHistoryRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{6}
}

func (x *HeartBeatHistoryRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type HeartBeatHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *HeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *HeartBeatHistoryResponse) Reset() {
	*x = HeartBeatHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatHistoryResponse) ProtoMessage() {}

func (x *HeartBeatHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatHistoryResponse.ProtoReflect.Descriptor instead.
func (*HeartBeatHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{7}
}

func (x *HeartBeatHistoryResponse) GetHeartbeat() *HeartBeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type NormalAbnormalHeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heartbeat *NormalAbnormalHeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *NormalAbnormalHeartBeatRequest) Reset() {
	*x = NormalAbnormalHeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalAbnormalHeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalAbnormalHeartBeatRequest) ProtoMessage() {}

func (x *NormalAbnormalHeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalAbnormalHeartBeatRequest.ProtoReflect.Descriptor instead.
func (*NormalAbnormalHeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{8}
}

func (x *NormalAbnormalHeartBeatRequest) GetHeartbeat() *NormalAbnormalHeartBeat {
	if x != nil {
		return x.Heartbeat
	}
	return nil
}

type NormalAbnormalHeatBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NormalAbnormalHeatBeatResponse) Reset() {
	*x = NormalAbnormalHeatBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalAbnormalHeatBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalAbnormalHeatBeatResponse) ProtoMessage() {}

func (x *NormalAbnormalHeatBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalAbnormalHeatBeatResponse.ProtoReflect.Descriptor instead.
func (*NormalAbnormalHeatBeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_hb_proto_rawDescGZIP(), []int{9}
}

func (x *NormalAbnormalHeatBeatResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_proto_hb_proto protoreflect.FileDescriptor

var file_proto_hb_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x70, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x17, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41,
	0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x70, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62,
	0x70, 0x6d, 0x22, 0x45, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x09,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x14, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x22, 0x2f, 0x0a, 0x15, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x35, 0x0a, 0x17, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x18, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x09, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x61, 0x0a, 0x1e, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41,
	0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x38, 0x0a, 0x1e, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x88, 0x03, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x68, 0x62,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x5d, 0x0a, 0x10,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x21, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x73, 0x0a, 0x17, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x28, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x68, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x41, 0x62, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hb_proto_rawDescOnce sync.Once
	file_proto_hb_proto_rawDescData = file_proto_hb_proto_rawDesc
)

func file_proto_hb_proto_rawDescGZIP() []byte {
	file_proto_hb_proto_rawDescOnce.Do(func() {
		file_proto_hb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hb_proto_rawDescData)
	})
	return file_proto_hb_proto_rawDescData
}

var file_proto_hb_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_hb_proto_goTypes = []interface{}{
	(*HeartBeat)(nil),                      // 0: hb_proto.HeartBeat
	(*NormalAbnormalHeartBeat)(nil),        // 1: hb_proto.NormalAbnormalHeartBeat
	(*HeartBeatRequest)(nil),               // 2: hb_proto.HeartBeatRequest
	(*HeartBeatResponse)(nil),              // 3: hb_proto.HeartBeatResponse
	(*LiveHeartBeatRequest)(nil),           // 4: hb_proto.LiveHeartBeatRequest
	(*LiveHeartBeatResponse)(nil),          // 5: hb_proto.LiveHeartBeatResponse
	(*HeartBeatHistoryRequest)(nil),        // 6: hb_proto.HeartBeatHistoryRequest
	(*HeartBeatHistoryResponse)(nil),       // 7: hb_proto.HeartBeatHistoryResponse
	(*NormalAbnormalHeartBeatRequest)(nil), // 8: hb_proto.NormalAbnormalHeartBeatRequest
	(*NormalAbnormalHeatBeatResponse)(nil), // 9: hb_proto.NormalAbnormalHeatBeatResponse
}
var file_proto_hb_proto_depIdxs = []int32{
	0, // 0: hb_proto.HeartBeatRequest.heartbeat:type_name -> hb_proto.HeartBeat
	0, // 1: hb_proto.LiveHeartBeatRequest.heartbeat:type_name -> hb_proto.HeartBeat
	0, // 2: hb_proto.HeartBeatHistoryResponse.heartbeat:type_name -> hb_proto.HeartBeat
	1, // 3: hb_proto.NormalAbnormalHeartBeatRequest.heartbeat:type_name -> hb_proto.NormalAbnormalHeartBeat
	2, // 4: hb_proto.HeartBeatService.UserHeartBeat:input_type -> hb_proto.HeartBeatRequest
	4, // 5: hb_proto.HeartBeatService.LiveHeartBeat:input_type -> hb_proto.LiveHeartBeatRequest
	6, // 6: hb_proto.HeartBeatService.HeartBeatHistory:input_type -> hb_proto.HeartBeatHistoryRequest
	8, // 7: hb_proto.HeartBeatService.NormalAbnormalHeartBeat:input_type -> hb_proto.NormalAbnormalHeartBeatRequest
	3, // 8: hb_proto.HeartBeatService.UserHeartBeat:output_type -> hb_proto.HeartBeatResponse
	5, // 9: hb_proto.HeartBeatService.LiveHeartBeat:output_type -> hb_proto.LiveHeartBeatResponse
	7, // 10: hb_proto.HeartBeatService.HeartBeatHistory:output_type -> hb_proto.HeartBeatHistoryResponse
	9, // 11: hb_proto.HeartBeatService.NormalAbnormalHeartBeat:output_type -> hb_proto.NormalAbnormalHeatBeatResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_hb_proto_init() }
func file_proto_hb_proto_init() {
	if File_proto_hb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeat); i {
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
		file_proto_hb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalAbnormalHeartBeat); i {
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
		file_proto_hb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatRequest); i {
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
		file_proto_hb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatResponse); i {
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
		file_proto_hb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveHeartBeatRequest); i {
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
		file_proto_hb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveHeartBeatResponse); i {
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
		file_proto_hb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatHistoryRequest); i {
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
		file_proto_hb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatHistoryResponse); i {
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
		file_proto_hb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalAbnormalHeartBeatRequest); i {
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
		file_proto_hb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalAbnormalHeatBeatResponse); i {
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
			RawDescriptor: file_proto_hb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hb_proto_goTypes,
		DependencyIndexes: file_proto_hb_proto_depIdxs,
		MessageInfos:      file_proto_hb_proto_msgTypes,
	}.Build()
	File_proto_hb_proto = out.File
	file_proto_hb_proto_rawDesc = nil
	file_proto_hb_proto_goTypes = nil
	file_proto_hb_proto_depIdxs = nil
}

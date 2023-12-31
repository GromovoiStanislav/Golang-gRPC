// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pokemon/pokemon.proto

package pokemonpc

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

type Pokemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid         string `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Power       string `protobuf:"bytes,4,opt,name=power,proto3" json:"power,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{0}
}

func (x *Pokemon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pokemon) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

func (x *Pokemon) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreatePokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *CreatePokemonRequest) Reset() {
	*x = CreatePokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePokemonRequest) ProtoMessage() {}

func (x *CreatePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePokemonRequest.ProtoReflect.Descriptor instead.
func (*CreatePokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePokemonRequest) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type CreatePokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"` // будет иметь идентификатор покемона
}

func (x *CreatePokemonResponse) Reset() {
	*x = CreatePokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePokemonResponse) ProtoMessage() {}

func (x *CreatePokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePokemonResponse.ProtoReflect.Descriptor instead.
func (*CreatePokemonResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePokemonResponse) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type ReadPokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ReadPokemonRequest) Reset() {
	*x = ReadPokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPokemonRequest) ProtoMessage() {}

func (x *ReadPokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPokemonRequest.ProtoReflect.Descriptor instead.
func (*ReadPokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPokemonRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type ReadPokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *ReadPokemonResponse) Reset() {
	*x = ReadPokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPokemonResponse) ProtoMessage() {}

func (x *ReadPokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPokemonResponse.ProtoReflect.Descriptor instead.
func (*ReadPokemonResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPokemonResponse) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type UpdatePokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *UpdatePokemonRequest) Reset() {
	*x = UpdatePokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePokemonRequest) ProtoMessage() {}

func (x *UpdatePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePokemonRequest.ProtoReflect.Descriptor instead.
func (*UpdatePokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePokemonRequest) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type UpdatePokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *UpdatePokemonResponse) Reset() {
	*x = UpdatePokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePokemonResponse) ProtoMessage() {}

func (x *UpdatePokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePokemonResponse.ProtoReflect.Descriptor instead.
func (*UpdatePokemonResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePokemonResponse) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type DeletePokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *DeletePokemonRequest) Reset() {
	*x = DeletePokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePokemonRequest) ProtoMessage() {}

func (x *DeletePokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePokemonRequest.ProtoReflect.Descriptor instead.
func (*DeletePokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePokemonRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type DeletePokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *DeletePokemonResponse) Reset() {
	*x = DeletePokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePokemonResponse) ProtoMessage() {}

func (x *DeletePokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePokemonResponse.ProtoReflect.Descriptor instead.
func (*DeletePokemonResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePokemonResponse) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type ListPokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPokemonRequest) Reset() {
	*x = ListPokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPokemonRequest) ProtoMessage() {}

func (x *ListPokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPokemonRequest.ProtoReflect.Descriptor instead.
func (*ListPokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{9}
}

type ListPokemonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon *Pokemon `protobuf:"bytes,1,opt,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *ListPokemonResponse) Reset() {
	*x = ListPokemonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_pokemon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPokemonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPokemonResponse) ProtoMessage() {}

func (x *ListPokemonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_pokemon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPokemonResponse.ProtoReflect.Descriptor instead.
func (*ListPokemonResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_pokemon_proto_rawDescGZIP(), []int{10}
}

func (x *ListPokemonResponse) GetPokemon() *Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

var File_pokemon_pokemon_proto protoreflect.FileDescriptor

var file_pokemon_pokemon_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x22, 0x77, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x43, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x13, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x42, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x41, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x32, 0x96, 0x03, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x19, 0x5a,
	0x17, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x3b, 0x70,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokemon_pokemon_proto_rawDescOnce sync.Once
	file_pokemon_pokemon_proto_rawDescData = file_pokemon_pokemon_proto_rawDesc
)

func file_pokemon_pokemon_proto_rawDescGZIP() []byte {
	file_pokemon_pokemon_proto_rawDescOnce.Do(func() {
		file_pokemon_pokemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokemon_pokemon_proto_rawDescData)
	})
	return file_pokemon_pokemon_proto_rawDescData
}

var file_pokemon_pokemon_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pokemon_pokemon_proto_goTypes = []interface{}{
	(*Pokemon)(nil),               // 0: pokemon.Pokemon
	(*CreatePokemonRequest)(nil),  // 1: pokemon.CreatePokemonRequest
	(*CreatePokemonResponse)(nil), // 2: pokemon.CreatePokemonResponse
	(*ReadPokemonRequest)(nil),    // 3: pokemon.ReadPokemonRequest
	(*ReadPokemonResponse)(nil),   // 4: pokemon.ReadPokemonResponse
	(*UpdatePokemonRequest)(nil),  // 5: pokemon.UpdatePokemonRequest
	(*UpdatePokemonResponse)(nil), // 6: pokemon.UpdatePokemonResponse
	(*DeletePokemonRequest)(nil),  // 7: pokemon.DeletePokemonRequest
	(*DeletePokemonResponse)(nil), // 8: pokemon.DeletePokemonResponse
	(*ListPokemonRequest)(nil),    // 9: pokemon.ListPokemonRequest
	(*ListPokemonResponse)(nil),   // 10: pokemon.ListPokemonResponse
}
var file_pokemon_pokemon_proto_depIdxs = []int32{
	0,  // 0: pokemon.CreatePokemonRequest.pokemon:type_name -> pokemon.Pokemon
	0,  // 1: pokemon.CreatePokemonResponse.pokemon:type_name -> pokemon.Pokemon
	0,  // 2: pokemon.ReadPokemonResponse.pokemon:type_name -> pokemon.Pokemon
	0,  // 3: pokemon.UpdatePokemonRequest.pokemon:type_name -> pokemon.Pokemon
	0,  // 4: pokemon.UpdatePokemonResponse.pokemon:type_name -> pokemon.Pokemon
	0,  // 5: pokemon.ListPokemonResponse.pokemon:type_name -> pokemon.Pokemon
	1,  // 6: pokemon.PokemonService.CreatePokemon:input_type -> pokemon.CreatePokemonRequest
	3,  // 7: pokemon.PokemonService.ReadPokemon:input_type -> pokemon.ReadPokemonRequest
	5,  // 8: pokemon.PokemonService.UpdatePokemon:input_type -> pokemon.UpdatePokemonRequest
	7,  // 9: pokemon.PokemonService.DeletePokemon:input_type -> pokemon.DeletePokemonRequest
	9,  // 10: pokemon.PokemonService.ListPokemon:input_type -> pokemon.ListPokemonRequest
	2,  // 11: pokemon.PokemonService.CreatePokemon:output_type -> pokemon.CreatePokemonResponse
	4,  // 12: pokemon.PokemonService.ReadPokemon:output_type -> pokemon.ReadPokemonResponse
	6,  // 13: pokemon.PokemonService.UpdatePokemon:output_type -> pokemon.UpdatePokemonResponse
	8,  // 14: pokemon.PokemonService.DeletePokemon:output_type -> pokemon.DeletePokemonResponse
	10, // 15: pokemon.PokemonService.ListPokemon:output_type -> pokemon.ListPokemonResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pokemon_pokemon_proto_init() }
func file_pokemon_pokemon_proto_init() {
	if File_pokemon_pokemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokemon_pokemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pokemon); i {
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
		file_pokemon_pokemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePokemonRequest); i {
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
		file_pokemon_pokemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePokemonResponse); i {
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
		file_pokemon_pokemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPokemonRequest); i {
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
		file_pokemon_pokemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPokemonResponse); i {
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
		file_pokemon_pokemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePokemonRequest); i {
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
		file_pokemon_pokemon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePokemonResponse); i {
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
		file_pokemon_pokemon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePokemonRequest); i {
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
		file_pokemon_pokemon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePokemonResponse); i {
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
		file_pokemon_pokemon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPokemonRequest); i {
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
		file_pokemon_pokemon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPokemonResponse); i {
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
			RawDescriptor: file_pokemon_pokemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokemon_pokemon_proto_goTypes,
		DependencyIndexes: file_pokemon_pokemon_proto_depIdxs,
		MessageInfos:      file_pokemon_pokemon_proto_msgTypes,
	}.Build()
	File_pokemon_pokemon_proto = out.File
	file_pokemon_pokemon_proto_rawDesc = nil
	file_pokemon_pokemon_proto_goTypes = nil
	file_pokemon_pokemon_proto_depIdxs = nil
}

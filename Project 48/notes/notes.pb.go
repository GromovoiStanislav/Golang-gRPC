// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: notes/notes.proto

package notes

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

// The request message containing the note title
type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type NoteSaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Saved bool `protobuf:"varint,1,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *NoteSaveReply) Reset() {
	*x = NoteSaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteSaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteSaveReply) ProtoMessage() {}

func (x *NoteSaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteSaveReply.ProtoReflect.Descriptor instead.
func (*NoteSaveReply) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{1}
}

func (x *NoteSaveReply) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

// The request message containing the note title
type NoteSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *NoteSearch) Reset() {
	*x = NoteSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteSearch) ProtoMessage() {}

func (x *NoteSearch) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteSearch.ProtoReflect.Descriptor instead.
func (*NoteSearch) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{2}
}

func (x *NoteSearch) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

var File_notes_notes_proto protoreflect.FileDescriptor

var file_notes_notes_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x25, 0x0a, 0x0d,
	0x4e, 0x6f, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x61,
	0x76, 0x65, 0x64, 0x22, 0x26, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x96, 0x01, 0x0a, 0x05,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0b, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x14, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x11, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x0b, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d,
	0x53, 0x61, 0x76, 0x65, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0b, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x14, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_notes_notes_proto_rawDescOnce sync.Once
	file_notes_notes_proto_rawDescData = file_notes_notes_proto_rawDesc
)

func file_notes_notes_proto_rawDescGZIP() []byte {
	file_notes_notes_proto_rawDescOnce.Do(func() {
		file_notes_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_notes_proto_rawDescData)
	})
	return file_notes_notes_proto_rawDescData
}

var file_notes_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notes_notes_proto_goTypes = []interface{}{
	(*Note)(nil),          // 0: notes.Note
	(*NoteSaveReply)(nil), // 1: notes.NoteSaveReply
	(*NoteSearch)(nil),    // 2: notes.NoteSearch
}
var file_notes_notes_proto_depIdxs = []int32{
	0, // 0: notes.Notes.Save:input_type -> notes.Note
	2, // 1: notes.Notes.Load:input_type -> notes.NoteSearch
	0, // 2: notes.Notes.SaveLargeNote:input_type -> notes.Note
	1, // 3: notes.Notes.Save:output_type -> notes.NoteSaveReply
	0, // 4: notes.Notes.Load:output_type -> notes.Note
	1, // 5: notes.Notes.SaveLargeNote:output_type -> notes.NoteSaveReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notes_notes_proto_init() }
func file_notes_notes_proto_init() {
	if File_notes_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_notes_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteSaveReply); i {
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
		file_notes_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteSearch); i {
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
			RawDescriptor: file_notes_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_notes_proto_goTypes,
		DependencyIndexes: file_notes_notes_proto_depIdxs,
		MessageInfos:      file_notes_notes_proto_msgTypes,
	}.Build()
	File_notes_notes_proto = out.File
	file_notes_notes_proto_rawDesc = nil
	file_notes_notes_proto_goTypes = nil
	file_notes_notes_proto_depIdxs = nil
}

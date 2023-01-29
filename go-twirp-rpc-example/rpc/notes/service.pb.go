// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc/notes/service.proto

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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_notes_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_notes_service_proto_msgTypes[0]
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
	return file_rpc_notes_service_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Note) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type CreateNoteParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateNoteParams) Reset() {
	*x = CreateNoteParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_notes_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteParams) ProtoMessage() {}

func (x *CreateNoteParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_notes_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteParams.ProtoReflect.Descriptor instead.
func (*CreateNoteParams) Descriptor() ([]byte, []int) {
	return file_rpc_notes_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNoteParams) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetAllNotesParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllNotesParams) Reset() {
	*x = GetAllNotesParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_notes_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesParams) ProtoMessage() {}

func (x *GetAllNotesParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_notes_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesParams.ProtoReflect.Descriptor instead.
func (*GetAllNotesParams) Descriptor() ([]byte, []int) {
	return file_rpc_notes_service_proto_rawDescGZIP(), []int{2}
}

type GetAllNotesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GetAllNotesResult) Reset() {
	*x = GetAllNotesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_notes_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesResult) ProtoMessage() {}

func (x *GetAllNotesResult) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_notes_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesResult.ProtoReflect.Descriptor instead.
func (*GetAllNotesResult) Descriptor() ([]byte, []int) {
	return file_rpc_notes_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllNotesResult) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

var File_rpc_notes_service_proto protoreflect.FileDescriptor

var file_rpc_notes_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x74, 0x77, 0x69,
	0x72, 0x70, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x26,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x49, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x34, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xd1, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x67, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x2b, 0x2e,
	0x67, 0x6f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x72, 0x70,
	0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_notes_service_proto_rawDescOnce sync.Once
	file_rpc_notes_service_proto_rawDescData = file_rpc_notes_service_proto_rawDesc
)

func file_rpc_notes_service_proto_rawDescGZIP() []byte {
	file_rpc_notes_service_proto_rawDescOnce.Do(func() {
		file_rpc_notes_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_notes_service_proto_rawDescData)
	})
	return file_rpc_notes_service_proto_rawDescData
}

var file_rpc_notes_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_notes_service_proto_goTypes = []interface{}{
	(*Note)(nil),              // 0: gotwirpexample.rpc.notes.Note
	(*CreateNoteParams)(nil),  // 1: gotwirpexample.rpc.notes.CreateNoteParams
	(*GetAllNotesParams)(nil), // 2: gotwirpexample.rpc.notes.GetAllNotesParams
	(*GetAllNotesResult)(nil), // 3: gotwirpexample.rpc.notes.GetAllNotesResult
}
var file_rpc_notes_service_proto_depIdxs = []int32{
	0, // 0: gotwirpexample.rpc.notes.GetAllNotesResult.notes:type_name -> gotwirpexample.rpc.notes.Note
	1, // 1: gotwirpexample.rpc.notes.NotesService.CreateNote:input_type -> gotwirpexample.rpc.notes.CreateNoteParams
	2, // 2: gotwirpexample.rpc.notes.NotesService.GetAllNotes:input_type -> gotwirpexample.rpc.notes.GetAllNotesParams
	0, // 3: gotwirpexample.rpc.notes.NotesService.CreateNote:output_type -> gotwirpexample.rpc.notes.Note
	3, // 4: gotwirpexample.rpc.notes.NotesService.GetAllNotes:output_type -> gotwirpexample.rpc.notes.GetAllNotesResult
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_notes_service_proto_init() }
func file_rpc_notes_service_proto_init() {
	if File_rpc_notes_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_notes_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_notes_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteParams); i {
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
		file_rpc_notes_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotesParams); i {
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
		file_rpc_notes_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotesResult); i {
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
			RawDescriptor: file_rpc_notes_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_notes_service_proto_goTypes,
		DependencyIndexes: file_rpc_notes_service_proto_depIdxs,
		MessageInfos:      file_rpc_notes_service_proto_msgTypes,
	}.Build()
	File_rpc_notes_service_proto = out.File
	file_rpc_notes_service_proto_rawDesc = nil
	file_rpc_notes_service_proto_goTypes = nil
	file_rpc_notes_service_proto_depIdxs = nil
}

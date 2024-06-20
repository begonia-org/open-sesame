// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: constant.proto

package v1

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

type ChunkType int32

const (
	ChunkType_CHUNK_TYPE_UNKNOWN ChunkType = 0
	ChunkType_CHUNK_TYPE_NAIVE   ChunkType = 1
	ChunkType_CHUNK_TYPE_PAPER   ChunkType = 2
)

// Enum value maps for ChunkType.
var (
	ChunkType_name = map[int32]string{
		0: "CHUNK_TYPE_UNKNOWN",
		1: "CHUNK_TYPE_NAIVE",
		2: "CHUNK_TYPE_PAPER",
	}
	ChunkType_value = map[string]int32{
		"CHUNK_TYPE_UNKNOWN": 0,
		"CHUNK_TYPE_NAIVE":   1,
		"CHUNK_TYPE_PAPER":   2,
	}
)

func (x ChunkType) Enum() *ChunkType {
	p := new(ChunkType)
	*p = x
	return p
}

func (x ChunkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChunkType) Descriptor() protoreflect.EnumDescriptor {
	return file_constant_proto_enumTypes[0].Descriptor()
}

func (ChunkType) Type() protoreflect.EnumType {
	return &file_constant_proto_enumTypes[0]
}

func (x ChunkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChunkType.Descriptor instead.
func (ChunkType) EnumDescriptor() ([]byte, []int) {
	return file_constant_proto_rawDescGZIP(), []int{0}
}

// class FileType(StrEnum):
//
//	PDF = 'pdf'
//	DOC = 'doc'
//	VISUAL = 'visual'
//	AURAL = 'aural'
//	VIRTUAL = 'virtual'
//	FOLDER = 'folder'
//	OTHER = "other"
type FileType int32

const (
	FileType_FILE_TYPE_UNKNOWN FileType = 0
	FileType_FILE_TYPE_PDF     FileType = 1
	FileType_FILE_TYPE_DOC     FileType = 2
	FileType_FILE_TYPE_VISUAL  FileType = 3
	FileType_FILE_TYPE_AURAL   FileType = 4
	FileType_FILE_TYPE_VIRTUAL FileType = 5
	FileType_FILE_TYPE_FOLDER  FileType = 6
	FileType_FILE_TYPE_OTHER   FileType = 7
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "FILE_TYPE_UNKNOWN",
		1: "FILE_TYPE_PDF",
		2: "FILE_TYPE_DOC",
		3: "FILE_TYPE_VISUAL",
		4: "FILE_TYPE_AURAL",
		5: "FILE_TYPE_VIRTUAL",
		6: "FILE_TYPE_FOLDER",
		7: "FILE_TYPE_OTHER",
	}
	FileType_value = map[string]int32{
		"FILE_TYPE_UNKNOWN": 0,
		"FILE_TYPE_PDF":     1,
		"FILE_TYPE_DOC":     2,
		"FILE_TYPE_VISUAL":  3,
		"FILE_TYPE_AURAL":   4,
		"FILE_TYPE_VIRTUAL": 5,
		"FILE_TYPE_FOLDER":  6,
		"FILE_TYPE_OTHER":   7,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_constant_proto_enumTypes[1].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_constant_proto_enumTypes[1]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_constant_proto_rawDescGZIP(), []int{1}
}

type LLMType int32

const (
	LLMType_LLM_TYPE_UNKNOWN     LLMType = 0
	LLMType_LLM_TYPE_CHAT        LLMType = 1
	LLMType_LLM_TYPE_EMBEDDING   LLMType = 2
	LLMType_LLM_TYPE_SPEECH2TEXT LLMType = 3
	LLMType_LLM_TYPE_IMAGE2TEXT  LLMType = 4
	LLMType_LLM_TYPE_RERANK      LLMType = 5
)

// Enum value maps for LLMType.
var (
	LLMType_name = map[int32]string{
		0: "LLM_TYPE_UNKNOWN",
		1: "LLM_TYPE_CHAT",
		2: "LLM_TYPE_EMBEDDING",
		3: "LLM_TYPE_SPEECH2TEXT",
		4: "LLM_TYPE_IMAGE2TEXT",
		5: "LLM_TYPE_RERANK",
	}
	LLMType_value = map[string]int32{
		"LLM_TYPE_UNKNOWN":     0,
		"LLM_TYPE_CHAT":        1,
		"LLM_TYPE_EMBEDDING":   2,
		"LLM_TYPE_SPEECH2TEXT": 3,
		"LLM_TYPE_IMAGE2TEXT":  4,
		"LLM_TYPE_RERANK":      5,
	}
)

func (x LLMType) Enum() *LLMType {
	p := new(LLMType)
	*p = x
	return p
}

func (x LLMType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LLMType) Descriptor() protoreflect.EnumDescriptor {
	return file_constant_proto_enumTypes[2].Descriptor()
}

func (LLMType) Type() protoreflect.EnumType {
	return &file_constant_proto_enumTypes[2]
}

func (x LLMType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LLMType.Descriptor instead.
func (LLMType) EnumDescriptor() ([]byte, []int) {
	return file_constant_proto_rawDescGZIP(), []int{2}
}

var File_constant_proto protoreflect.FileDescriptor

var file_constant_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2a, 0x4f, 0x0a, 0x09, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x41, 0x50, 0x45, 0x52, 0x10, 0x02, 0x2a, 0xb4, 0x01, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x44, 0x46, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x4f, 0x43, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x56, 0x49, 0x53, 0x55, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x52, 0x41, 0x4c, 0x10, 0x04, 0x12,
	0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x52,
	0x54, 0x55, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10,
	0x07, 0x2a, 0x92, 0x01, 0x0a, 0x07, 0x4c, 0x4c, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x4d, 0x42, 0x45, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x45, 0x43,
	0x48, 0x32, 0x54, 0x45, 0x58, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x4c, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x32, 0x54, 0x45, 0x58, 0x54, 0x10,
	0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4c, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45,
	0x52, 0x41, 0x4e, 0x4b, 0x10, 0x05, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2d, 0x6f, 0x72, 0x67,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_constant_proto_rawDescOnce sync.Once
	file_constant_proto_rawDescData = file_constant_proto_rawDesc
)

func file_constant_proto_rawDescGZIP() []byte {
	file_constant_proto_rawDescOnce.Do(func() {
		file_constant_proto_rawDescData = protoimpl.X.CompressGZIP(file_constant_proto_rawDescData)
	})
	return file_constant_proto_rawDescData
}

var file_constant_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_constant_proto_goTypes = []interface{}{
	(ChunkType)(0), // 0: openrag.ChunkType
	(FileType)(0),  // 1: openrag.FileType
	(LLMType)(0),   // 2: openrag.LLMType
}
var file_constant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_constant_proto_init() }
func file_constant_proto_init() {
	if File_constant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_constant_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_constant_proto_goTypes,
		DependencyIndexes: file_constant_proto_depIdxs,
		EnumInfos:         file_constant_proto_enumTypes,
	}.Build()
	File_constant_proto = out.File
	file_constant_proto_rawDesc = nil
	file_constant_proto_goTypes = nil
	file_constant_proto_depIdxs = nil
}
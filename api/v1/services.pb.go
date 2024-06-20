// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: services.proto

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

type DocumentEmbedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId       string        `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Lang         string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	ChunkType    ChunkType     `protobuf:"varint,3,opt,name=chunk_type,proto3,enum=openrag.ChunkType" json:"chunk_type,omitempty"`
	LlmType      LLMType       `protobuf:"varint,4,opt,name=llm_type,proto3,enum=openrag.LLMType" json:"llm_type,omitempty"`
	ParserConfig *ParserConfig `protobuf:"bytes,5,opt,name=parser_config,proto3" json:"parser_config,omitempty"`
	Llm          *LLM          `protobuf:"bytes,6,opt,name=llm,proto3" json:"llm,omitempty"`
}

func (x *DocumentEmbedRequest) Reset() {
	*x = DocumentEmbedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentEmbedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentEmbedRequest) ProtoMessage() {}

func (x *DocumentEmbedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentEmbedRequest.ProtoReflect.Descriptor instead.
func (*DocumentEmbedRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentEmbedRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *DocumentEmbedRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DocumentEmbedRequest) GetChunkType() ChunkType {
	if x != nil {
		return x.ChunkType
	}
	return ChunkType_CHUNK_TYPE_UNKNOWN
}

func (x *DocumentEmbedRequest) GetLlmType() LLMType {
	if x != nil {
		return x.LlmType
	}
	return LLMType_LLM_TYPE_UNKNOWN
}

func (x *DocumentEmbedRequest) GetParserConfig() *ParserConfig {
	if x != nil {
		return x.ParserConfig
	}
	return nil
}

func (x *DocumentEmbedRequest) GetLlm() *LLM {
	if x != nil {
		return x.Llm
	}
	return nil
}

type DocumentEmbedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document    *Doc  `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	TotalTokens int64 `protobuf:"varint,2,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
}

func (x *DocumentEmbedResponse) Reset() {
	*x = DocumentEmbedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentEmbedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentEmbedResponse) ProtoMessage() {}

func (x *DocumentEmbedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentEmbedResponse.ProtoReflect.Descriptor instead.
func (*DocumentEmbedResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentEmbedResponse) GetDocument() *Doc {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *DocumentEmbedResponse) GetTotalTokens() int64 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

type CallbaclRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,proto3" json:"file_id,omitempty"`
}

func (x *CallbaclRequest) Reset() {
	*x = CallbaclRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbaclRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbaclRequest) ProtoMessage() {}

func (x *CallbaclRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbaclRequest.ProtoReflect.Descriptor instead.
func (*CallbaclRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *CallbaclRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type CallbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallbackResponse) Reset() {
	*x = CallbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackResponse) ProtoMessage() {}

func (x *CallbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackResponse.ProtoReflect.Descriptor instead.
func (*CallbackResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x1a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x64, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6c, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x82, 0x02, 0x0a, 0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x72, 0x61, 0x67, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x6c, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x4c, 0x4c, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c,
	0x6c, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x03, 0x6c, 0x6c, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x4c, 0x4c, 0x4d, 0x52,
	0x03, 0x6c, 0x6c, 0x6d, 0x22, 0x64, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x01, 0x0a, 0x14,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x1d, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x6d, 0x62, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x41,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x72, 0x61, 0x67, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x32, 0x54, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x72, 0x61, 0x67, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2d, 0x6f, 0x72,
	0x67, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_proto_goTypes = []interface{}{
	(*DocumentEmbedRequest)(nil),  // 0: openrag.DocumentEmbedRequest
	(*DocumentEmbedResponse)(nil), // 1: openrag.DocumentEmbedResponse
	(*CallbaclRequest)(nil),       // 2: openrag.CallbaclRequest
	(*CallbackResponse)(nil),      // 3: openrag.CallbackResponse
	(ChunkType)(0),                // 4: openrag.ChunkType
	(LLMType)(0),                  // 5: openrag.LLMType
	(*ParserConfig)(nil),          // 6: openrag.ParserConfig
	(*LLM)(nil),                   // 7: openrag.LLM
	(*Doc)(nil),                   // 8: openrag.Doc
}
var file_services_proto_depIdxs = []int32{
	4, // 0: openrag.DocumentEmbedRequest.chunk_type:type_name -> openrag.ChunkType
	5, // 1: openrag.DocumentEmbedRequest.llm_type:type_name -> openrag.LLMType
	6, // 2: openrag.DocumentEmbedRequest.parser_config:type_name -> openrag.ParserConfig
	7, // 3: openrag.DocumentEmbedRequest.llm:type_name -> openrag.LLM
	8, // 4: openrag.DocumentEmbedResponse.document:type_name -> openrag.Doc
	0, // 5: openrag.DocumentEmbedService.embed:input_type -> openrag.DocumentEmbedRequest
	2, // 6: openrag.DocumentEmbedService.progress:input_type -> openrag.CallbaclRequest
	2, // 7: openrag.CallbackService.progress:input_type -> openrag.CallbaclRequest
	1, // 8: openrag.DocumentEmbedService.embed:output_type -> openrag.DocumentEmbedResponse
	3, // 9: openrag.DocumentEmbedService.progress:output_type -> openrag.CallbackResponse
	3, // 10: openrag.CallbackService.progress:output_type -> openrag.CallbackResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	file_constant_proto_init()
	file_doc_proto_init()
	file_llm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentEmbedRequest); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentEmbedResponse); i {
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
		file_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbaclRequest); i {
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
		file_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackResponse); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}

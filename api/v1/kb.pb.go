// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: kb.proto

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

// class Meta:
//
//	db_table = "openkb_knowledge_base"
type KnowledgeBaseStatus int32

const (
	KnowledgeBaseStatus_WASTED   KnowledgeBaseStatus = 0
	KnowledgeBaseStatus_VALIDATE KnowledgeBaseStatus = 1
)

// Enum value maps for KnowledgeBaseStatus.
var (
	KnowledgeBaseStatus_name = map[int32]string{
		0: "WASTED",
		1: "VALIDATE",
	}
	KnowledgeBaseStatus_value = map[string]int32{
		"WASTED":   0,
		"VALIDATE": 1,
	}
)

func (x KnowledgeBaseStatus) Enum() *KnowledgeBaseStatus {
	p := new(KnowledgeBaseStatus)
	*p = x
	return p
}

func (x KnowledgeBaseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KnowledgeBaseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kb_proto_enumTypes[0].Descriptor()
}

func (KnowledgeBaseStatus) Type() protoreflect.EnumType {
	return &file_kb_proto_enumTypes[0]
}

func (x KnowledgeBaseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KnowledgeBaseStatus.Descriptor instead.
func (KnowledgeBaseStatus) EnumDescriptor() ([]byte, []int) {
	return file_kb_proto_rawDescGZIP(), []int{0}
}

type KnowledgeBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Avatar                 string              `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	TenantId               string              `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name                   string              `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Language               string              `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	Description            string              `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	EmbdId                 string              `protobuf:"bytes,7,opt,name=embd_id,json=embdId,proto3" json:"embd_id,omitempty"`
	Permission             string              `protobuf:"bytes,8,opt,name=permission,proto3" json:"permission,omitempty"`
	CreatedBy              string              `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	DocNum                 int64               `protobuf:"varint,10,opt,name=doc_num,json=docNum,proto3" json:"doc_num,omitempty"`
	TokenNum               int64               `protobuf:"varint,11,opt,name=token_num,json=tokenNum,proto3" json:"token_num,omitempty"`
	ChunkNum               int64               `protobuf:"varint,12,opt,name=chunk_num,json=chunkNum,proto3" json:"chunk_num,omitempty"`
	SimilarityThreshold    float32             `protobuf:"fixed32,13,opt,name=similarity_threshold,json=similarityThreshold,proto3" json:"similarity_threshold,omitempty"`
	VectorSimilarityWeight float32             `protobuf:"fixed32,14,opt,name=vector_similarity_weight,json=vectorSimilarityWeight,proto3" json:"vector_similarity_weight,omitempty"`
	ParserId               string              `protobuf:"bytes,15,opt,name=parser_id,json=parserId,proto3" json:"parser_id,omitempty"`
	ParserConfig           map[string]string   `protobuf:"bytes,16,rep,name=parser_config,json=parserConfig,proto3" json:"parser_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status                 KnowledgeBaseStatus `protobuf:"varint,17,opt,name=status,proto3,enum=openrag.KnowledgeBaseStatus" json:"status,omitempty"`
}

func (x *KnowledgeBase) Reset() {
	*x = KnowledgeBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowledgeBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowledgeBase) ProtoMessage() {}

func (x *KnowledgeBase) ProtoReflect() protoreflect.Message {
	mi := &file_kb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowledgeBase.ProtoReflect.Descriptor instead.
func (*KnowledgeBase) Descriptor() ([]byte, []int) {
	return file_kb_proto_rawDescGZIP(), []int{0}
}

func (x *KnowledgeBase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KnowledgeBase) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *KnowledgeBase) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *KnowledgeBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KnowledgeBase) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *KnowledgeBase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *KnowledgeBase) GetEmbdId() string {
	if x != nil {
		return x.EmbdId
	}
	return ""
}

func (x *KnowledgeBase) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *KnowledgeBase) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *KnowledgeBase) GetDocNum() int64 {
	if x != nil {
		return x.DocNum
	}
	return 0
}

func (x *KnowledgeBase) GetTokenNum() int64 {
	if x != nil {
		return x.TokenNum
	}
	return 0
}

func (x *KnowledgeBase) GetChunkNum() int64 {
	if x != nil {
		return x.ChunkNum
	}
	return 0
}

func (x *KnowledgeBase) GetSimilarityThreshold() float32 {
	if x != nil {
		return x.SimilarityThreshold
	}
	return 0
}

func (x *KnowledgeBase) GetVectorSimilarityWeight() float32 {
	if x != nil {
		return x.VectorSimilarityWeight
	}
	return 0
}

func (x *KnowledgeBase) GetParserId() string {
	if x != nil {
		return x.ParserId
	}
	return ""
}

func (x *KnowledgeBase) GetParserConfig() map[string]string {
	if x != nil {
		return x.ParserConfig
	}
	return nil
}

func (x *KnowledgeBase) GetStatus() KnowledgeBaseStatus {
	if x != nil {
		return x.Status
	}
	return KnowledgeBaseStatus_WASTED
}

var File_kb_proto protoreflect.FileDescriptor

var file_kb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e,
	0x72, 0x61, 0x67, 0x22, 0xa1, 0x05, 0x0a, 0x0d, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x65, 0x6d, 0x62, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x6d, 0x62, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x69, 0x6d, 0x69, 0x6c,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x16, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x4d, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72,
	0x61, 0x67, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x72, 0x61, 0x67, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2f, 0x0a, 0x13, 0x4b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a,
	0x0a, 0x06, 0x57, 0x41, 0x53, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x6e, 0x69, 0x61, 0x2d, 0x6f,
	0x72, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x73, 0x61, 0x6d, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kb_proto_rawDescOnce sync.Once
	file_kb_proto_rawDescData = file_kb_proto_rawDesc
)

func file_kb_proto_rawDescGZIP() []byte {
	file_kb_proto_rawDescOnce.Do(func() {
		file_kb_proto_rawDescData = protoimpl.X.CompressGZIP(file_kb_proto_rawDescData)
	})
	return file_kb_proto_rawDescData
}

var file_kb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kb_proto_goTypes = []interface{}{
	(KnowledgeBaseStatus)(0), // 0: openrag.KnowledgeBaseStatus
	(*KnowledgeBase)(nil),    // 1: openrag.KnowledgeBase
	nil,                      // 2: openrag.KnowledgeBase.ParserConfigEntry
}
var file_kb_proto_depIdxs = []int32{
	2, // 0: openrag.KnowledgeBase.parser_config:type_name -> openrag.KnowledgeBase.ParserConfigEntry
	0, // 1: openrag.KnowledgeBase.status:type_name -> openrag.KnowledgeBaseStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kb_proto_init() }
func file_kb_proto_init() {
	if File_kb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowledgeBase); i {
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
			RawDescriptor: file_kb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kb_proto_goTypes,
		DependencyIndexes: file_kb_proto_depIdxs,
		EnumInfos:         file_kb_proto_enumTypes,
		MessageInfos:      file_kb_proto_msgTypes,
	}.Build()
	File_kb_proto = out.File
	file_kb_proto_rawDesc = nil
	file_kb_proto_goTypes = nil
	file_kb_proto_depIdxs = nil
}

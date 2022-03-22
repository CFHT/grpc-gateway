// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: examples/internal/proto/examplepb/visibility_rule_echo_service.proto

// Visibility Rule Echo Service
// Similar to echo_service.proto but with annotations to change visibility
// of services, methods, fields and enum values.
//
// `google.api.VisibilityRule` annotations are added to customize where they are generated.
// Combined with the option `visibility_restriction_selectors` overlapping rules will appear in the OpenAPI output.
// Elements without `google.api.VisibilityRule` annotations will appear as usual in the generated output.
//
// These restrictions and selectors are completely arbitrary and you can define whatever values or hierarchies you want.
// In this example `INTERNAL`, `PREVIEW` are used, but `INTERNAL`, `ALPHA`, `BETA`, `RELEASED`, or anything else could be used if you wish.

package examplepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

type VisibilityRuleSimpleMessage_VisibilityEnum int32

const (
	VisibilityRuleSimpleMessage_VISIBILITY_ENUM_UNSPECIFIED VisibilityRuleSimpleMessage_VisibilityEnum = 0
	VisibilityRuleSimpleMessage_VISIBILITY_ENUM_VISIBLE     VisibilityRuleSimpleMessage_VisibilityEnum = 1
	VisibilityRuleSimpleMessage_VISIBILITY_ENUM_INTERNAL    VisibilityRuleSimpleMessage_VisibilityEnum = 2
	VisibilityRuleSimpleMessage_VISIBILITY_ENUM_PREVIEW     VisibilityRuleSimpleMessage_VisibilityEnum = 3
)

// Enum value maps for VisibilityRuleSimpleMessage_VisibilityEnum.
var (
	VisibilityRuleSimpleMessage_VisibilityEnum_name = map[int32]string{
		0: "VISIBILITY_ENUM_UNSPECIFIED",
		1: "VISIBILITY_ENUM_VISIBLE",
		2: "VISIBILITY_ENUM_INTERNAL",
		3: "VISIBILITY_ENUM_PREVIEW",
	}
	VisibilityRuleSimpleMessage_VisibilityEnum_value = map[string]int32{
		"VISIBILITY_ENUM_UNSPECIFIED": 0,
		"VISIBILITY_ENUM_VISIBLE":     1,
		"VISIBILITY_ENUM_INTERNAL":    2,
		"VISIBILITY_ENUM_PREVIEW":     3,
	}
)

func (x VisibilityRuleSimpleMessage_VisibilityEnum) Enum() *VisibilityRuleSimpleMessage_VisibilityEnum {
	p := new(VisibilityRuleSimpleMessage_VisibilityEnum)
	*p = x
	return p
}

func (x VisibilityRuleSimpleMessage_VisibilityEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisibilityRuleSimpleMessage_VisibilityEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_enumTypes[0].Descriptor()
}

func (VisibilityRuleSimpleMessage_VisibilityEnum) Type() protoreflect.EnumType {
	return &file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_enumTypes[0]
}

func (x VisibilityRuleSimpleMessage_VisibilityEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisibilityRuleSimpleMessage_VisibilityEnum.Descriptor instead.
func (VisibilityRuleSimpleMessage_VisibilityEnum) EnumDescriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescGZIP(), []int{1, 0}
}

// Embedded represents a message embedded in SimpleMessage.
type VisibilityRuleEmbedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mark:
	//	*VisibilityRuleEmbedded_Progress
	//	*VisibilityRuleEmbedded_Note
	//	*VisibilityRuleEmbedded_InternalField
	//	*VisibilityRuleEmbedded_PreviewField
	Mark isVisibilityRuleEmbedded_Mark `protobuf_oneof:"mark"`
}

func (x *VisibilityRuleEmbedded) Reset() {
	*x = VisibilityRuleEmbedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityRuleEmbedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityRuleEmbedded) ProtoMessage() {}

func (x *VisibilityRuleEmbedded) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityRuleEmbedded.ProtoReflect.Descriptor instead.
func (*VisibilityRuleEmbedded) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescGZIP(), []int{0}
}

func (m *VisibilityRuleEmbedded) GetMark() isVisibilityRuleEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (x *VisibilityRuleEmbedded) GetProgress() int64 {
	if x, ok := x.GetMark().(*VisibilityRuleEmbedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (x *VisibilityRuleEmbedded) GetNote() string {
	if x, ok := x.GetMark().(*VisibilityRuleEmbedded_Note); ok {
		return x.Note
	}
	return ""
}

func (x *VisibilityRuleEmbedded) GetInternalField() string {
	if x, ok := x.GetMark().(*VisibilityRuleEmbedded_InternalField); ok {
		return x.InternalField
	}
	return ""
}

func (x *VisibilityRuleEmbedded) GetPreviewField() string {
	if x, ok := x.GetMark().(*VisibilityRuleEmbedded_PreviewField); ok {
		return x.PreviewField
	}
	return ""
}

type isVisibilityRuleEmbedded_Mark interface {
	isVisibilityRuleEmbedded_Mark()
}

type VisibilityRuleEmbedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type VisibilityRuleEmbedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

type VisibilityRuleEmbedded_InternalField struct {
	InternalField string `protobuf:"bytes,3,opt,name=internal_field,json=internalField,proto3,oneof"`
}

type VisibilityRuleEmbedded_PreviewField struct {
	PreviewField string `protobuf:"bytes,4,opt,name=preview_field,json=previewField,proto3,oneof"`
}

func (*VisibilityRuleEmbedded_Progress) isVisibilityRuleEmbedded_Mark() {}

func (*VisibilityRuleEmbedded_Note) isVisibilityRuleEmbedded_Mark() {}

func (*VisibilityRuleEmbedded_InternalField) isVisibilityRuleEmbedded_Mark() {}

func (*VisibilityRuleEmbedded_PreviewField) isVisibilityRuleEmbedded_Mark() {}

// SimpleMessage represents a simple message sent to the Echo service.
type VisibilityRuleSimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// Types that are assignable to Code:
	//	*VisibilityRuleSimpleMessage_LineNum
	//	*VisibilityRuleSimpleMessage_Lang
	Code   isVisibilityRuleSimpleMessage_Code `protobuf_oneof:"code"`
	Status *VisibilityRuleEmbedded            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to Ext:
	//	*VisibilityRuleSimpleMessage_En
	//	*VisibilityRuleSimpleMessage_No
	Ext           isVisibilityRuleSimpleMessage_Ext          `protobuf_oneof:"ext"`
	InternalField string                                     `protobuf:"bytes,8,opt,name=internal_field,json=internalField,proto3" json:"internal_field,omitempty"`
	PreviewField  string                                     `protobuf:"bytes,9,opt,name=preview_field,json=previewField,proto3" json:"preview_field,omitempty"`
	AnEnum        VisibilityRuleSimpleMessage_VisibilityEnum `protobuf:"varint,10,opt,name=an_enum,json=anEnum,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage_VisibilityEnum" json:"an_enum,omitempty"`
}

func (x *VisibilityRuleSimpleMessage) Reset() {
	*x = VisibilityRuleSimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisibilityRuleSimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisibilityRuleSimpleMessage) ProtoMessage() {}

func (x *VisibilityRuleSimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisibilityRuleSimpleMessage.ProtoReflect.Descriptor instead.
func (*VisibilityRuleSimpleMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescGZIP(), []int{1}
}

func (x *VisibilityRuleSimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VisibilityRuleSimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (m *VisibilityRuleSimpleMessage) GetCode() isVisibilityRuleSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (x *VisibilityRuleSimpleMessage) GetLineNum() int64 {
	if x, ok := x.GetCode().(*VisibilityRuleSimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (x *VisibilityRuleSimpleMessage) GetLang() string {
	if x, ok := x.GetCode().(*VisibilityRuleSimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (x *VisibilityRuleSimpleMessage) GetStatus() *VisibilityRuleEmbedded {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *VisibilityRuleSimpleMessage) GetExt() isVisibilityRuleSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (x *VisibilityRuleSimpleMessage) GetEn() int64 {
	if x, ok := x.GetExt().(*VisibilityRuleSimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (x *VisibilityRuleSimpleMessage) GetNo() *VisibilityRuleEmbedded {
	if x, ok := x.GetExt().(*VisibilityRuleSimpleMessage_No); ok {
		return x.No
	}
	return nil
}

func (x *VisibilityRuleSimpleMessage) GetInternalField() string {
	if x != nil {
		return x.InternalField
	}
	return ""
}

func (x *VisibilityRuleSimpleMessage) GetPreviewField() string {
	if x != nil {
		return x.PreviewField
	}
	return ""
}

func (x *VisibilityRuleSimpleMessage) GetAnEnum() VisibilityRuleSimpleMessage_VisibilityEnum {
	if x != nil {
		return x.AnEnum
	}
	return VisibilityRuleSimpleMessage_VISIBILITY_ENUM_UNSPECIFIED
}

type isVisibilityRuleSimpleMessage_Code interface {
	isVisibilityRuleSimpleMessage_Code()
}

type VisibilityRuleSimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type VisibilityRuleSimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,proto3,oneof"`
}

func (*VisibilityRuleSimpleMessage_LineNum) isVisibilityRuleSimpleMessage_Code() {}

func (*VisibilityRuleSimpleMessage_Lang) isVisibilityRuleSimpleMessage_Code() {}

type isVisibilityRuleSimpleMessage_Ext interface {
	isVisibilityRuleSimpleMessage_Ext()
}

type VisibilityRuleSimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,proto3,oneof"`
}

type VisibilityRuleSimpleMessage_No struct {
	No *VisibilityRuleEmbedded `protobuf:"bytes,7,opt,name=no,proto3,oneof"`
}

func (*VisibilityRuleSimpleMessage_En) isVisibilityRuleSimpleMessage_Ext() {}

func (*VisibilityRuleSimpleMessage_No) isVisibilityRuleSimpleMessage_Ext() {}

var File_examples_internal_proto_examplepb_visibility_rule_echo_service_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDesc = []byte{
	0x0a, 0x44, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x16, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x75, 0x6c, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x12, 0x39, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x2c, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x48, 0x00, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x06, 0x0a, 0x04,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0xf2, 0x05, 0x0a, 0x1b, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x64, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x02, 0x65, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x02, 0x65, 0x6e, 0x12, 0x58, 0x0a, 0x02, 0x6e,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x01, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x37, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3d,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x2c, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x52,
	0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x73, 0x0a,
	0x07, 0x61, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6e, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xb5, 0x01, 0x0a, 0x0e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x2e, 0x0a, 0x18, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10,
	0x02, 0x1a, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x12, 0x35, 0x0a, 0x17, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x03,
	0x1a, 0x18, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x2c, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x42, 0x06, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x32, 0x92, 0x07, 0x0a, 0x19, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x63, 0x68, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f,
	0x12, 0x4b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x4b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xdb, 0x01, 0x0a, 0x0c, 0x45, 0x63,
	0x68, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x4b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x4b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x12, 0xd8, 0x01, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x4b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x4b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x50, 0x52, 0x45, 0x56, 0x49,
	0x45, 0x57, 0x12, 0xf9, 0x01, 0x0a, 0x16, 0x45, 0x63, 0x68, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x4b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x4b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12,
	0x25, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x2c, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x32, 0x80,
	0x02, 0x0a, 0x21, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xc8, 0x01, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x4b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x4b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a,
	0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62,
	0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescData = file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDesc
)

func file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDescData
}

var file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_goTypes = []interface{}{
	(VisibilityRuleSimpleMessage_VisibilityEnum)(0), // 0: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage.VisibilityEnum
	(*VisibilityRuleEmbedded)(nil),                  // 1: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEmbedded
	(*VisibilityRuleSimpleMessage)(nil),             // 2: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
}
var file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_depIdxs = []int32{
	1, // 0: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage.status:type_name -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEmbedded
	1, // 1: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage.no:type_name -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEmbedded
	0, // 2: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage.an_enum:type_name -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage.VisibilityEnum
	2, // 3: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 4: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoInternal:input_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 5: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoPreview:input_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 6: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoInternalAndPreview:input_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 7: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleInternalEchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 8: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 9: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoInternal:output_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 10: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoPreview:output_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 11: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService.EchoInternalAndPreview:output_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	2, // 12: grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleInternalEchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleSimpleMessage
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_init() }
func file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_init() {
	if File_examples_internal_proto_examplepb_visibility_rule_echo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisibilityRuleEmbedded); i {
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
		file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisibilityRuleSimpleMessage); i {
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
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*VisibilityRuleEmbedded_Progress)(nil),
		(*VisibilityRuleEmbedded_Note)(nil),
		(*VisibilityRuleEmbedded_InternalField)(nil),
		(*VisibilityRuleEmbedded_PreviewField)(nil),
	}
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*VisibilityRuleSimpleMessage_LineNum)(nil),
		(*VisibilityRuleSimpleMessage_Lang)(nil),
		(*VisibilityRuleSimpleMessage_En)(nil),
		(*VisibilityRuleSimpleMessage_No)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_depIdxs,
		EnumInfos:         file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_enumTypes,
		MessageInfos:      file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_visibility_rule_echo_service_proto = out.File
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_goTypes = nil
	file_examples_internal_proto_examplepb_visibility_rule_echo_service_proto_depIdxs = nil
}
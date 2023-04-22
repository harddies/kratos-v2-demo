// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.11
// source: admin/v1/feature_flag.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFeatureFlagRequest) Reset() {
	*x = CreateFeatureFlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeatureFlagRequest) ProtoMessage() {}

func (x *CreateFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*CreateFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{0}
}

type CreateFeatureFlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFeatureFlagReply) Reset() {
	*x = CreateFeatureFlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeatureFlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeatureFlagReply) ProtoMessage() {}

func (x *CreateFeatureFlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeatureFlagReply.ProtoReflect.Descriptor instead.
func (*CreateFeatureFlagReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{1}
}

type UpdateFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFeatureFlagRequest) Reset() {
	*x = UpdateFeatureFlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeatureFlagRequest) ProtoMessage() {}

func (x *UpdateFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{2}
}

type UpdateFeatureFlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFeatureFlagReply) Reset() {
	*x = UpdateFeatureFlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeatureFlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeatureFlagReply) ProtoMessage() {}

func (x *UpdateFeatureFlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeatureFlagReply.ProtoReflect.Descriptor instead.
func (*UpdateFeatureFlagReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{3}
}

type DeleteFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFeatureFlagRequest) Reset() {
	*x = DeleteFeatureFlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFeatureFlagRequest) ProtoMessage() {}

func (x *DeleteFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*DeleteFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{4}
}

type DeleteFeatureFlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFeatureFlagReply) Reset() {
	*x = DeleteFeatureFlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFeatureFlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFeatureFlagReply) ProtoMessage() {}

func (x *DeleteFeatureFlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFeatureFlagReply.ProtoReflect.Descriptor instead.
func (*DeleteFeatureFlagReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{5}
}

type GetFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureFlagId string `protobuf:"bytes,1,opt,name=feature_flag_id,json=featureFlagId,proto3" json:"feature_flag_id,omitempty"`
}

func (x *GetFeatureFlagRequest) Reset() {
	*x = GetFeatureFlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureFlagRequest) ProtoMessage() {}

func (x *GetFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{6}
}

func (x *GetFeatureFlagRequest) GetFeatureFlagId() string {
	if x != nil {
		return x.FeatureFlagId
	}
	return ""
}

type GetFeatureFlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureFlagId string `protobuf:"bytes,1,opt,name=feature_flag_id,json=featureFlagId,proto3" json:"feature_flag_id,omitempty"`
	FeatureKey    string `protobuf:"bytes,2,opt,name=feature_key,json=featureKey,proto3" json:"feature_key,omitempty"`
}

func (x *GetFeatureFlagReply) Reset() {
	*x = GetFeatureFlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureFlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureFlagReply) ProtoMessage() {}

func (x *GetFeatureFlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureFlagReply.ProtoReflect.Descriptor instead.
func (*GetFeatureFlagReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{7}
}

func (x *GetFeatureFlagReply) GetFeatureFlagId() string {
	if x != nil {
		return x.FeatureFlagId
	}
	return ""
}

func (x *GetFeatureFlagReply) GetFeatureKey() string {
	if x != nil {
		return x.FeatureKey
	}
	return ""
}

type ListFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFeatureFlagRequest) Reset() {
	*x = ListFeatureFlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureFlagRequest) ProtoMessage() {}

func (x *ListFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*ListFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{8}
}

type ListFeatureFlagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFeatureFlagReply) Reset() {
	*x = ListFeatureFlagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_v1_feature_flag_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeatureFlagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeatureFlagReply) ProtoMessage() {}

func (x *ListFeatureFlagReply) ProtoReflect() protoreflect.Message {
	mi := &file_admin_v1_feature_flag_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeatureFlagReply.ProtoReflect.Descriptor instead.
func (*ListFeatureFlagReply) Descriptor() ([]byte, []int) {
	return file_admin_v1_feature_flag_proto_rawDescGZIP(), []int{9}
}

var File_admin_v1_feature_flag_proto protoreflect.FileDescriptor

var file_admin_v1_feature_flag_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3f, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x66, 0x6c, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x18, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xbb,
	0x05, 0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x8c,
	0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x78, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6c, 0x61,
	0x67, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x8c, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x78, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6c, 0x61, 0x67,
	0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x8c, 0x01, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x78, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6c, 0x61, 0x67, 0x2f,
	0x66, 0x6c, 0x61, 0x67, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7c, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f,
	0x78, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d,
	0x66, 0x6c, 0x61, 0x67, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x78, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2d, 0x66, 0x6c, 0x61, 0x67, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x1e, 0x5a, 0x1c,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6c, 0x61, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_v1_feature_flag_proto_rawDescOnce sync.Once
	file_admin_v1_feature_flag_proto_rawDescData = file_admin_v1_feature_flag_proto_rawDesc
)

func file_admin_v1_feature_flag_proto_rawDescGZIP() []byte {
	file_admin_v1_feature_flag_proto_rawDescOnce.Do(func() {
		file_admin_v1_feature_flag_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_v1_feature_flag_proto_rawDescData)
	})
	return file_admin_v1_feature_flag_proto_rawDescData
}

var file_admin_v1_feature_flag_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_admin_v1_feature_flag_proto_goTypes = []interface{}{
	(*CreateFeatureFlagRequest)(nil), // 0: api.admin.v1.CreateFeatureFlagRequest
	(*CreateFeatureFlagReply)(nil),   // 1: api.admin.v1.CreateFeatureFlagReply
	(*UpdateFeatureFlagRequest)(nil), // 2: api.admin.v1.UpdateFeatureFlagRequest
	(*UpdateFeatureFlagReply)(nil),   // 3: api.admin.v1.UpdateFeatureFlagReply
	(*DeleteFeatureFlagRequest)(nil), // 4: api.admin.v1.DeleteFeatureFlagRequest
	(*DeleteFeatureFlagReply)(nil),   // 5: api.admin.v1.DeleteFeatureFlagReply
	(*GetFeatureFlagRequest)(nil),    // 6: api.admin.v1.GetFeatureFlagRequest
	(*GetFeatureFlagReply)(nil),      // 7: api.admin.v1.GetFeatureFlagReply
	(*ListFeatureFlagRequest)(nil),   // 8: api.admin.v1.ListFeatureFlagRequest
	(*ListFeatureFlagReply)(nil),     // 9: api.admin.v1.ListFeatureFlagReply
}
var file_admin_v1_feature_flag_proto_depIdxs = []int32{
	0, // 0: api.admin.v1.FeatureFlag.CreateFeatureFlag:input_type -> api.admin.v1.CreateFeatureFlagRequest
	2, // 1: api.admin.v1.FeatureFlag.UpdateFeatureFlag:input_type -> api.admin.v1.UpdateFeatureFlagRequest
	4, // 2: api.admin.v1.FeatureFlag.DeleteFeatureFlag:input_type -> api.admin.v1.DeleteFeatureFlagRequest
	6, // 3: api.admin.v1.FeatureFlag.GetFeatureFlag:input_type -> api.admin.v1.GetFeatureFlagRequest
	8, // 4: api.admin.v1.FeatureFlag.ListFeatureFlag:input_type -> api.admin.v1.ListFeatureFlagRequest
	1, // 5: api.admin.v1.FeatureFlag.CreateFeatureFlag:output_type -> api.admin.v1.CreateFeatureFlagReply
	3, // 6: api.admin.v1.FeatureFlag.UpdateFeatureFlag:output_type -> api.admin.v1.UpdateFeatureFlagReply
	5, // 7: api.admin.v1.FeatureFlag.DeleteFeatureFlag:output_type -> api.admin.v1.DeleteFeatureFlagReply
	7, // 8: api.admin.v1.FeatureFlag.GetFeatureFlag:output_type -> api.admin.v1.GetFeatureFlagReply
	9, // 9: api.admin.v1.FeatureFlag.ListFeatureFlag:output_type -> api.admin.v1.ListFeatureFlagReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_v1_feature_flag_proto_init() }
func file_admin_v1_feature_flag_proto_init() {
	if File_admin_v1_feature_flag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_v1_feature_flag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeatureFlagRequest); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeatureFlagReply); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeatureFlagRequest); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeatureFlagReply); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFeatureFlagRequest); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFeatureFlagReply); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureFlagRequest); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureFlagReply); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeatureFlagRequest); i {
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
		file_admin_v1_feature_flag_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeatureFlagReply); i {
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
			RawDescriptor: file_admin_v1_feature_flag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_v1_feature_flag_proto_goTypes,
		DependencyIndexes: file_admin_v1_feature_flag_proto_depIdxs,
		MessageInfos:      file_admin_v1_feature_flag_proto_msgTypes,
	}.Build()
	File_admin_v1_feature_flag_proto = out.File
	file_admin_v1_feature_flag_proto_rawDesc = nil
	file_admin_v1_feature_flag_proto_goTypes = nil
	file_admin_v1_feature_flag_proto_depIdxs = nil
}

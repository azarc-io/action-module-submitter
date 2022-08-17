// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/domain/module/api/v1/model.proto

package module_v1

import (
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Spark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Readme      []byte `protobuf:"bytes,2,opt,name=readme,proto3" json:"readme,omitempty" bson:"readme"`
	InputSchema []byte `protobuf:"bytes,3,opt,name=input_schema,json=inputSchema,proto3" json:"input_schema,omitempty" bson:"input_schema"`
	Label       string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty" yaml:"label" bson:"label"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" yaml:"description" bson:"description"`
	Icon        string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty" yaml:"icon" bson:"icon"`
}

func (x *Spark) Reset() {
	*x = Spark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spark) ProtoMessage() {}

func (x *Spark) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spark.ProtoReflect.Descriptor instead.
func (*Spark) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *Spark) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spark) GetReadme() []byte {
	if x != nil {
		return x.Readme
	}
	return nil
}

func (x *Spark) GetInputSchema() []byte {
	if x != nil {
		return x.InputSchema
	}
	return nil
}

func (x *Spark) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Spark) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Spark) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type DetailEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Repo             string                 `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty" bson:"repo"`
	Version          string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" bson:"version"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at"`
	AggregateVersion uint32                 `protobuf:"varint,6,opt,name=aggregate_version,json=aggregateVersion,proto3" json:"aggregate_version,omitempty" bson:"aggregate_version"`
	Readme           []byte                 `protobuf:"bytes,7,opt,name=readme,proto3" json:"readme,omitempty" bson:"readme"`
	Label            string                 `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty" bson:"label"`
	Description      string                 `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	Icon             []byte                 `protobuf:"bytes,10,opt,name=icon,proto3" json:"icon,omitempty" bson:"icon"`
	Tags             []string               `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty" bson:"tags"`
	Sparks           []*Spark               `protobuf:"bytes,12,rep,name=sparks,proto3" json:"sparks,omitempty" bson:"sparks"`
}

func (x *DetailEntity) Reset() {
	*x = DetailEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailEntity) ProtoMessage() {}

func (x *DetailEntity) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailEntity.ProtoReflect.Descriptor instead.
func (*DetailEntity) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *DetailEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DetailEntity) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *DetailEntity) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DetailEntity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DetailEntity) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DetailEntity) GetAggregateVersion() uint32 {
	if x != nil {
		return x.AggregateVersion
	}
	return 0
}

func (x *DetailEntity) GetReadme() []byte {
	if x != nil {
		return x.Readme
	}
	return nil
}

func (x *DetailEntity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DetailEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DetailEntity) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *DetailEntity) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DetailEntity) GetSparks() []*Spark {
	if x != nil {
		return x.Sparks
	}
	return nil
}

type MasterEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Repo             string                 `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty" bson:"repo"`
	Version          string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" bson:"version"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bson:"updated_at"`
	AggregateVersion uint32                 `protobuf:"varint,6,opt,name=aggregate_version,json=aggregateVersion,proto3" json:"aggregate_version,omitempty" bson:"aggregate_version"`
	Label            string                 `protobuf:"bytes,7,opt,name=label,proto3" json:"label,omitempty" bson:"label"`
	Description      string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	Icon             []byte                 `protobuf:"bytes,9,opt,name=icon,proto3" json:"icon,omitempty" bson:"icon"`
	Tags             []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty" bson:"tags"`
	Sparks           uint32                 `protobuf:"varint,11,opt,name=sparks,proto3" json:"sparks,omitempty" bson:"sparks"`
}

func (x *MasterEntity) Reset() {
	*x = MasterEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterEntity) ProtoMessage() {}

func (x *MasterEntity) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterEntity.ProtoReflect.Descriptor instead.
func (*MasterEntity) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_model_proto_rawDescGZIP(), []int{2}
}

func (x *MasterEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MasterEntity) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *MasterEntity) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MasterEntity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MasterEntity) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MasterEntity) GetAggregateVersion() uint32 {
	if x != nil {
		return x.AggregateVersion
	}
	return 0
}

func (x *MasterEntity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MasterEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MasterEntity) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *MasterEntity) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *MasterEntity) GetSparks() uint32 {
	if x != nil {
		return x.Sparks
	}
	return 0
}

var File_internal_domain_module_api_v1_model_proto protoreflect.FileDescriptor

var file_internal_domain_module_api_v1_model_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x73, 0x6f, 0x6b, 0x6f, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61,
	0x67, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x12,
	0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a,
	0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x22, 0x52, 0x06, 0x72, 0x65, 0x61, 0x64, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x22, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3b,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x9a, 0x84, 0x9e, 0x03, 0x19, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x31, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x9a, 0x84, 0x9e, 0x03, 0x25, 0x79, 0x61,
	0x6d, 0x6c, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x20, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c,
	0x9a, 0x84, 0x9e, 0x03, 0x17, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22,
	0x20, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x22, 0x90, 0x05, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72,
	0x65, 0x70, 0x6f, 0x22, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0x9a, 0x84, 0x9e,
	0x03, 0x0e, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x22, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x51, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84,
	0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x4a, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x1d, 0x9a, 0x84, 0x9e, 0x03,
	0x18, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x64, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x12, 0x9a, 0x84, 0x9e,
	0x03, 0x0d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x22, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x52, 0x06, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x73, 0x22, 0xd2, 0x04, 0x0a, 0x0c, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f,
	0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x2d, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84,
	0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x51, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x4a, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x1d, 0x9a,
	0x84, 0x9e, 0x03, 0x18, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x10, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x9a,
	0x84, 0x9e, 0x03, 0x0c, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84,
	0x9e, 0x03, 0x12, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x63, 0x6f,
	0x6e, 0x22, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x73, 0x22, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x12,
	0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x73, 0x22, 0x52, 0x06, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x7a, 0x61, 0x72, 0x63, 0x2d, 0x69,
	0x6f, 0x2f, 0x76, 0x65, 0x72, 0x61, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_domain_module_api_v1_model_proto_rawDescOnce sync.Once
	file_internal_domain_module_api_v1_model_proto_rawDescData = file_internal_domain_module_api_v1_model_proto_rawDesc
)

func file_internal_domain_module_api_v1_model_proto_rawDescGZIP() []byte {
	file_internal_domain_module_api_v1_model_proto_rawDescOnce.Do(func() {
		file_internal_domain_module_api_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_domain_module_api_v1_model_proto_rawDescData)
	})
	return file_internal_domain_module_api_v1_model_proto_rawDescData
}

var file_internal_domain_module_api_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_domain_module_api_v1_model_proto_goTypes = []interface{}{
	(*Spark)(nil),                 // 0: module_v1.Spark
	(*DetailEntity)(nil),          // 1: module_v1.DetailEntity
	(*MasterEntity)(nil),          // 2: module_v1.MasterEntity
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_domain_module_api_v1_model_proto_depIdxs = []int32{
	3, // 0: module_v1.DetailEntity.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: module_v1.DetailEntity.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: module_v1.DetailEntity.sparks:type_name -> module_v1.Spark
	3, // 3: module_v1.MasterEntity.created_at:type_name -> google.protobuf.Timestamp
	3, // 4: module_v1.MasterEntity.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_domain_module_api_v1_model_proto_init() }
func file_internal_domain_module_api_v1_model_proto_init() {
	if File_internal_domain_module_api_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_domain_module_api_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spark); i {
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
		file_internal_domain_module_api_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailEntity); i {
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
		file_internal_domain_module_api_v1_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterEntity); i {
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
			RawDescriptor: file_internal_domain_module_api_v1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_domain_module_api_v1_model_proto_goTypes,
		DependencyIndexes: file_internal_domain_module_api_v1_model_proto_depIdxs,
		MessageInfos:      file_internal_domain_module_api_v1_model_proto_msgTypes,
	}.Build()
	File_internal_domain_module_api_v1_model_proto = out.File
	file_internal_domain_module_api_v1_model_proto_rawDesc = nil
	file_internal_domain_module_api_v1_model_proto_goTypes = nil
	file_internal_domain_module_api_v1_model_proto_depIdxs = nil
}

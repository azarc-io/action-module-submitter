// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/domain/module/api/v1/command.proto

package module_v1

import (
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Readme      []byte   `protobuf:"bytes,1,opt,name=readme,proto3" json:"readme,omitempty"`
	Label       string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty" yaml:"label"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Icon        []byte   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty" yaml:"icon"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" yaml:"tags"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_command_proto_rawDescGZIP(), []int{0}
}

func (x *Module) GetReadme() []byte {
	if x != nil {
		return x.Readme
	}
	return nil
}

func (x *Module) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Module) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Module) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *Module) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreateAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo    string   `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Version string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Module  *Module  `protobuf:"bytes,3,opt,name=module,proto3" json:"module,omitempty"`
	Sparks  []*Spark `protobuf:"bytes,4,rep,name=sparks,proto3" json:"sparks,omitempty"`
}

func (x *CreateAction) Reset() {
	*x = CreateAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAction) ProtoMessage() {}

func (x *CreateAction) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAction.ProtoReflect.Descriptor instead.
func (*CreateAction) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_command_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAction) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *CreateAction) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateAction) GetModule() *Module {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *CreateAction) GetSparks() []*Spark {
	if x != nil {
		return x.Sparks
	}
	return nil
}

type CreateGlobalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo    string  `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Version string  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Module  *Module `protobuf:"bytes,3,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *CreateGlobalCommand) Reset() {
	*x = CreateGlobalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGlobalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGlobalCommand) ProtoMessage() {}

func (x *CreateGlobalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGlobalCommand.ProtoReflect.Descriptor instead.
func (*CreateGlobalCommand) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_command_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGlobalCommand) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *CreateGlobalCommand) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateGlobalCommand) GetModule() *Module {
	if x != nil {
		return x.Module
	}
	return nil
}

type AddSparkGlobalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo    string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Spark   *Spark `protobuf:"bytes,3,opt,name=spark,proto3" json:"spark,omitempty"`
}

func (x *AddSparkGlobalCommand) Reset() {
	*x = AddSparkGlobalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSparkGlobalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSparkGlobalCommand) ProtoMessage() {}

func (x *AddSparkGlobalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSparkGlobalCommand.ProtoReflect.Descriptor instead.
func (*AddSparkGlobalCommand) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_command_proto_rawDescGZIP(), []int{3}
}

func (x *AddSparkGlobalCommand) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *AddSparkGlobalCommand) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AddSparkGlobalCommand) GetSpark() *Spark {
	if x != nil {
		return x.Spark
	}
	return nil
}

type MarkAsOutdatedCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkAsOutdatedCommand) Reset() {
	*x = MarkAsOutdatedCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkAsOutdatedCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkAsOutdatedCommand) ProtoMessage() {}

func (x *MarkAsOutdatedCommand) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkAsOutdatedCommand.ProtoReflect.Descriptor instead.
func (*MarkAsOutdatedCommand) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_command_proto_rawDescGZIP(), []int{4}
}

var File_internal_domain_module_api_v1_command_proto protoreflect.FileDescriptor

var file_internal_domain_module_api_v1_command_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x73, 0x6f, 0x6b, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x64, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x9a, 0x84, 0x9e, 0x03,
	0x0c, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x72, 0x17,
	0x32, 0x15, 0x76, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x2e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b,
	0x2e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x73, 0x22,
	0x6e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22,
	0x6d, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x52, 0x05, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x22, 0x17,
	0x0a, 0x15, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x4f, 0x75, 0x74, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x7a, 0x61, 0x72, 0x63, 0x2d, 0x69, 0x6f, 0x2f, 0x76,
	0x65, 0x72, 0x61, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_domain_module_api_v1_command_proto_rawDescOnce sync.Once
	file_internal_domain_module_api_v1_command_proto_rawDescData = file_internal_domain_module_api_v1_command_proto_rawDesc
)

func file_internal_domain_module_api_v1_command_proto_rawDescGZIP() []byte {
	file_internal_domain_module_api_v1_command_proto_rawDescOnce.Do(func() {
		file_internal_domain_module_api_v1_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_domain_module_api_v1_command_proto_rawDescData)
	})
	return file_internal_domain_module_api_v1_command_proto_rawDescData
}

var file_internal_domain_module_api_v1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_domain_module_api_v1_command_proto_goTypes = []interface{}{
	(*Module)(nil),                // 0: module_v1.Module
	(*CreateAction)(nil),          // 1: module_v1.CreateAction
	(*CreateGlobalCommand)(nil),   // 2: module_v1.CreateGlobalCommand
	(*AddSparkGlobalCommand)(nil), // 3: module_v1.AddSparkGlobalCommand
	(*MarkAsOutdatedCommand)(nil), // 4: module_v1.MarkAsOutdatedCommand
	(*Spark)(nil),                 // 5: module_v1.Spark
}
var file_internal_domain_module_api_v1_command_proto_depIdxs = []int32{
	0, // 0: module_v1.CreateAction.module:type_name -> module_v1.Module
	5, // 1: module_v1.CreateAction.sparks:type_name -> module_v1.Spark
	0, // 2: module_v1.CreateGlobalCommand.module:type_name -> module_v1.Module
	5, // 3: module_v1.AddSparkGlobalCommand.spark:type_name -> module_v1.Spark
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_domain_module_api_v1_command_proto_init() }
func file_internal_domain_module_api_v1_command_proto_init() {
	if File_internal_domain_module_api_v1_command_proto != nil {
		return
	}
	file_internal_domain_module_api_v1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_domain_module_api_v1_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_internal_domain_module_api_v1_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAction); i {
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
		file_internal_domain_module_api_v1_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGlobalCommand); i {
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
		file_internal_domain_module_api_v1_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSparkGlobalCommand); i {
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
		file_internal_domain_module_api_v1_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkAsOutdatedCommand); i {
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
			RawDescriptor: file_internal_domain_module_api_v1_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_domain_module_api_v1_command_proto_goTypes,
		DependencyIndexes: file_internal_domain_module_api_v1_command_proto_depIdxs,
		MessageInfos:      file_internal_domain_module_api_v1_command_proto_msgTypes,
	}.Build()
	File_internal_domain_module_api_v1_command_proto = out.File
	file_internal_domain_module_api_v1_command_proto_rawDesc = nil
	file_internal_domain_module_api_v1_command_proto_goTypes = nil
	file_internal_domain_module_api_v1_command_proto_depIdxs = nil
}

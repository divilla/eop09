// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: crud.proto

package crudproto

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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Entity) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PkEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldKey string `protobuf:"bytes,1,opt,name=oldKey,proto3" json:"oldKey,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PkEntity) Reset() {
	*x = PkEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkEntity) ProtoMessage() {}

func (x *PkEntity) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkEntity.ProtoReflect.Descriptor instead.
func (*PkEntity) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{1}
}

func (x *PkEntity) GetOldKey() string {
	if x != nil {
		return x.OldKey
	}
	return ""
}

func (x *PkEntity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PkEntity) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type IndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *IndexRequest) Reset() {
	*x = IndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRequest) ProtoMessage() {}

func (x *IndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRequest.ProtoReflect.Descriptor instead.
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{2}
}

func (x *IndexRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IndexRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PkRequest) Reset() {
	*x = PkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkRequest) ProtoMessage() {}

func (x *PkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkRequest.ProtoReflect.Descriptor instead.
func (*PkRequest) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{3}
}

func (x *PkRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type IndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         int64     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int64     `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Results      []*Entity `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	TotalResults int64     `protobuf:"varint,4,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	TotalPages   int64     `protobuf:"varint,5,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
}

func (x *IndexResponse) Reset() {
	*x = IndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexResponse) ProtoMessage() {}

func (x *IndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexResponse.ProtoReflect.Descriptor instead.
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{4}
}

func (x *IndexResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IndexResponse) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *IndexResponse) GetResults() []*Entity {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *IndexResponse) GetTotalResults() int64 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

func (x *IndexResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64 `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{5}
}

func (x *CommandResponse) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64  `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
	Errors       []byte `protobuf:"bytes,2,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{6}
}

func (x *ImportResponse) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *ImportResponse) GetErrors() []byte {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_crud_proto protoreflect.FileDescriptor

var file_crud_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x72,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x08, 0x50, 0x6b, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6c, 0x64, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x50, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x4c,
	0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xa0, 0x03, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x50,
	0x75, 0x74, 0x12, 0x13, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x11, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x19, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x76, 0x69, 0x6c, 0x6c, 0x61, 0x2f, 0x65, 0x6f, 0x70, 0x30, 0x39, 0x2f, 0x63, 0x72, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_proto_rawDescOnce sync.Once
	file_crud_proto_rawDescData = file_crud_proto_rawDesc
)

func file_crud_proto_rawDescGZIP() []byte {
	file_crud_proto_rawDescOnce.Do(func() {
		file_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_proto_rawDescData)
	})
	return file_crud_proto_rawDescData
}

var file_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_crud_proto_goTypes = []interface{}{
	(*Entity)(nil),          // 0: crudproto.Entity
	(*PkEntity)(nil),        // 1: crudproto.PkEntity
	(*IndexRequest)(nil),    // 2: crudproto.IndexRequest
	(*PkRequest)(nil),       // 3: crudproto.PkRequest
	(*IndexResponse)(nil),   // 4: crudproto.IndexResponse
	(*CommandResponse)(nil), // 5: crudproto.CommandResponse
	(*ImportResponse)(nil),  // 6: crudproto.ImportResponse
}
var file_crud_proto_depIdxs = []int32{
	0, // 0: crudproto.IndexResponse.results:type_name -> crudproto.Entity
	2, // 1: crudproto.RPC.Index:input_type -> crudproto.IndexRequest
	3, // 2: crudproto.RPC.Get:input_type -> crudproto.PkRequest
	0, // 3: crudproto.RPC.Create:input_type -> crudproto.Entity
	1, // 4: crudproto.RPC.Patch:input_type -> crudproto.PkEntity
	1, // 5: crudproto.RPC.Put:input_type -> crudproto.PkEntity
	3, // 6: crudproto.RPC.Delete:input_type -> crudproto.PkRequest
	0, // 7: crudproto.RPC.Import:input_type -> crudproto.Entity
	4, // 8: crudproto.RPC.Index:output_type -> crudproto.IndexResponse
	0, // 9: crudproto.RPC.Get:output_type -> crudproto.Entity
	5, // 10: crudproto.RPC.Create:output_type -> crudproto.CommandResponse
	5, // 11: crudproto.RPC.Patch:output_type -> crudproto.CommandResponse
	5, // 12: crudproto.RPC.Put:output_type -> crudproto.CommandResponse
	5, // 13: crudproto.RPC.Delete:output_type -> crudproto.CommandResponse
	6, // 14: crudproto.RPC.Import:output_type -> crudproto.ImportResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_crud_proto_init() }
func file_crud_proto_init() {
	if File_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkEntity); i {
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
		file_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRequest); i {
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
		file_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkRequest); i {
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
		file_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexResponse); i {
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
		file_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
			RawDescriptor: file_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_proto_goTypes,
		DependencyIndexes: file_crud_proto_depIdxs,
		MessageInfos:      file_crud_proto_msgTypes,
	}.Build()
	File_crud_proto = out.File
	file_crud_proto_rawDesc = nil
	file_crud_proto_goTypes = nil
	file_crud_proto_depIdxs = nil
}

// Copyright 2025, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pulumi/resource_status.proto

package pulumirpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the kind of operation performed by a step.
type ViewStep_Op int32

const (
	// An unspecified operation.
	ViewStep_UNSPECIFIED ViewStep_Op = 0
	// Nothing to do.
	ViewStep_SAME ViewStep_Op = 1
	// Creating a new resource.
	ViewStep_CREATE ViewStep_Op = 2
	// Updating an existing resource.
	ViewStep_UPDATE ViewStep_Op = 3
	// Deleting an existing resource.
	ViewStep_DELETE ViewStep_Op = 4
	// Replacing a resource with a new one.
	ViewStep_REPLACE ViewStep_Op = 5
	// Creating a new resource for a replacement.
	ViewStep_CREATE_REPLACEMENT ViewStep_Op = 6
	// Deleting an existing resource after replacement.
	ViewStep_DELETE_REPLACED ViewStep_Op = 7
	// Reading an existing resource.
	ViewStep_READ ViewStep_Op = 8
	// Reading an existing resource for a replacement.
	ViewStep_READ_REPLACEMENT ViewStep_Op = 9
	// Refreshing an existing resource.
	ViewStep_REFRESH ViewStep_Op = 10
	// Removing a resource that was read.
	ViewStep_READ_DISCARD ViewStep_Op = 11
	// Discarding a read resource that was replaced.
	ViewStep_DISCARD_REPLACED ViewStep_Op = 12
	// Removing a pending replace resource.
	ViewStep_REMOVE_PENDING_REPLACE ViewStep_Op = 13
	// Import an existing resource.
	ViewStep_IMPORT ViewStep_Op = 14
	// Replace an existing resource.
	ViewStep_IMPORT_REPLACEMENT ViewStep_Op = 15
)

// Enum value maps for ViewStep_Op.
var (
	ViewStep_Op_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "SAME",
		2:  "CREATE",
		3:  "UPDATE",
		4:  "DELETE",
		5:  "REPLACE",
		6:  "CREATE_REPLACEMENT",
		7:  "DELETE_REPLACED",
		8:  "READ",
		9:  "READ_REPLACEMENT",
		10: "REFRESH",
		11: "READ_DISCARD",
		12: "DISCARD_REPLACED",
		13: "REMOVE_PENDING_REPLACE",
		14: "IMPORT",
		15: "IMPORT_REPLACEMENT",
	}
	ViewStep_Op_value = map[string]int32{
		"UNSPECIFIED":            0,
		"SAME":                   1,
		"CREATE":                 2,
		"UPDATE":                 3,
		"DELETE":                 4,
		"REPLACE":                5,
		"CREATE_REPLACEMENT":     6,
		"DELETE_REPLACED":        7,
		"READ":                   8,
		"READ_REPLACEMENT":       9,
		"REFRESH":                10,
		"READ_DISCARD":           11,
		"DISCARD_REPLACED":       12,
		"REMOVE_PENDING_REPLACE": 13,
		"IMPORT":                 14,
		"IMPORT_REPLACEMENT":     15,
	}
)

func (x ViewStep_Op) Enum() *ViewStep_Op {
	p := new(ViewStep_Op)
	*p = x
	return p
}

func (x ViewStep_Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ViewStep_Op) Descriptor() protoreflect.EnumDescriptor {
	return file_pulumi_resource_status_proto_enumTypes[0].Descriptor()
}

func (ViewStep_Op) Type() protoreflect.EnumType {
	return &file_pulumi_resource_status_proto_enumTypes[0]
}

func (x ViewStep_Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ViewStep_Op.Descriptor instead.
func (ViewStep_Op) EnumDescriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{2, 0}
}

// Status is returned when an error has occurred during a resource provider operation.
// It indicates whether the operation could be rolled back cleanly (OK). If not, it
// means the resource was left in an indeterminate state.
type ViewStep_Status int32

const (
	ViewStep_OK              ViewStep_Status = 0
	ViewStep_PARTIAL_FAILURE ViewStep_Status = 1
	ViewStep_UNKNOWN         ViewStep_Status = 2
)

// Enum value maps for ViewStep_Status.
var (
	ViewStep_Status_name = map[int32]string{
		0: "OK",
		1: "PARTIAL_FAILURE",
		2: "UNKNOWN",
	}
	ViewStep_Status_value = map[string]int32{
		"OK":              0,
		"PARTIAL_FAILURE": 1,
		"UNKNOWN":         2,
	}
)

func (x ViewStep_Status) Enum() *ViewStep_Status {
	p := new(ViewStep_Status)
	*p = x
	return p
}

func (x ViewStep_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ViewStep_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pulumi_resource_status_proto_enumTypes[1].Descriptor()
}

func (ViewStep_Status) Type() protoreflect.EnumType {
	return &file_pulumi_resource_status_proto_enumTypes[1]
}

func (x ViewStep_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ViewStep_Status.Descriptor instead.
func (ViewStep_Status) EnumDescriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{2, 1}
}

// `PublishViewStepsRequest` is the type of requests sent as part of a
// [](pulumirpc.ResourceStatus.PublishViewSteps) call.
type PublishViewStepsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service context token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The steps to publish.
	Steps []*ViewStep `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *PublishViewStepsRequest) Reset() {
	*x = PublishViewStepsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_resource_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishViewStepsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishViewStepsRequest) ProtoMessage() {}

func (x *PublishViewStepsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_resource_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishViewStepsRequest.ProtoReflect.Descriptor instead.
func (*PublishViewStepsRequest) Descriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{0}
}

func (x *PublishViewStepsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PublishViewStepsRequest) GetSteps() []*ViewStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

// `PublishViewStepsResponse` is the type of responses sent as part of a
// [](pulumirpc.ResourceStatus.PublishViewSteps) call.
type PublishViewStepsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishViewStepsResponse) Reset() {
	*x = PublishViewStepsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_resource_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishViewStepsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishViewStepsResponse) ProtoMessage() {}

func (x *PublishViewStepsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_resource_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishViewStepsResponse.ProtoReflect.Descriptor instead.
func (*PublishViewStepsResponse) Descriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{1}
}

// `ViewStep` represents a deployment operation step for a view resource.
type ViewStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the operation.
	Status ViewStep_Status `protobuf:"varint,1,opt,name=status,proto3,enum=pulumirpc.ViewStep_Status" json:"status,omitempty"`
	// An optional error message indicating the operation failed.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The operation being performed.
	Op ViewStep_Op `protobuf:"varint,3,opt,name=op,proto3,enum=pulumirpc.ViewStep_Op" json:"op,omitempty"`
	// The type of the view resource.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// The name of the view resource.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// The state of the view resource before performing the step.
	Old *ViewStepState `protobuf:"bytes,6,opt,name=old,proto3" json:"old,omitempty"`
	// The state of the view resource after performing the step.
	New *ViewStepState `protobuf:"bytes,7,opt,name=new,proto3" json:"new,omitempty"`
	// The keys causing a replacement (only applicable for "create" and "replace" ops).
	Keys []string `protobuf:"bytes,8,rep,name=keys,proto3" json:"keys,omitempty"`
	// The keys that changed with this step.
	Diffs []string `protobuf:"bytes,9,rep,name=diffs,proto3" json:"diffs,omitempty"`
	// A detailed diff is a map from [property paths](property-paths) to [](pulumirpc.PropertyDiff)s,
	// which describe the kind of change that occurred to the property located at that path.
	DetailedDiff map[string]*PropertyDiff `protobuf:"bytes,10,rep,name=detailed_diff,json=detailedDiff,proto3" json:"detailed_diff,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Whether the detailed diff is present.
	HasDetailedDiff bool `protobuf:"varint,11,opt,name=has_detailed_diff,json=hasDetailedDiff,proto3" json:"has_detailed_diff,omitempty"`
}

func (x *ViewStep) Reset() {
	*x = ViewStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_resource_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewStep) ProtoMessage() {}

func (x *ViewStep) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_resource_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewStep.ProtoReflect.Descriptor instead.
func (*ViewStep) Descriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{2}
}

func (x *ViewStep) GetStatus() ViewStep_Status {
	if x != nil {
		return x.Status
	}
	return ViewStep_OK
}

func (x *ViewStep) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ViewStep) GetOp() ViewStep_Op {
	if x != nil {
		return x.Op
	}
	return ViewStep_UNSPECIFIED
}

func (x *ViewStep) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ViewStep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewStep) GetOld() *ViewStepState {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *ViewStep) GetNew() *ViewStepState {
	if x != nil {
		return x.New
	}
	return nil
}

func (x *ViewStep) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *ViewStep) GetDiffs() []string {
	if x != nil {
		return x.Diffs
	}
	return nil
}

func (x *ViewStep) GetDetailedDiff() map[string]*PropertyDiff {
	if x != nil {
		return x.DetailedDiff
	}
	return nil
}

func (x *ViewStep) GetHasDetailedDiff() bool {
	if x != nil {
		return x.HasDetailedDiff
	}
	return false
}

// `ViewStepState` represents the state of a view resource.
type ViewStepState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the view resource.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The name of the view resource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// An optional type of the parent view resource. By default, a view resource will
	// be a child of the owning resource, but views can be nested within other views,
	// as long as they're all under the same owner. Both parent_type and parent_name
	// must be set together.
	ParentType string `protobuf:"bytes,3,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty"`
	// An optional name of the parent view resource. By default, a view resource will
	// be a child of the owning resource, but views can be nested within other views,
	// as long as they're all under the same owner. Both parent_type and parent_name
	// must be set together.
	ParentName string `protobuf:"bytes,4,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	// The view resource's inputs.
	Inputs *structpb.Struct `protobuf:"bytes,5,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// The view resource's outputs.
	Outputs *structpb.Struct `protobuf:"bytes,6,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ViewStepState) Reset() {
	*x = ViewStepState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pulumi_resource_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewStepState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewStepState) ProtoMessage() {}

func (x *ViewStepState) ProtoReflect() protoreflect.Message {
	mi := &file_pulumi_resource_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewStepState.ProtoReflect.Descriptor instead.
func (*ViewStepState) Descriptor() ([]byte, []int) {
	return file_pulumi_resource_status_proto_rawDescGZIP(), []int{3}
}

func (x *ViewStepState) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ViewStepState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ViewStepState) GetParentType() string {
	if x != nil {
		return x.ParentType
	}
	return ""
}

func (x *ViewStepState) GetParentName() string {
	if x != nil {
		return x.ParentName
	}
	return ""
}

func (x *ViewStepState) GetInputs() *structpb.Struct {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ViewStepState) GetOutputs() *structpb.Struct {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_pulumi_resource_status_proto protoreflect.FileDescriptor

var file_pulumi_resource_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a,
	0x0a, 0x17, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x29, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc1, 0x06, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a,
	0x02, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x4f,
	0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x03, 0x6f, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x03, 0x6e, 0x65, 0x77,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72,
	0x70, 0x63, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x03, 0x6e, 0x65, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x66,
	0x66, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x69, 0x66, 0x66, 0x73, 0x12,
	0x4a, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72,
	0x70, 0x63, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x44, 0x69, 0x66, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x69, 0x66, 0x66, 0x12, 0x2a, 0x0a, 0x11, 0x68,
	0x61, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x44, 0x69, 0x66, 0x66, 0x1a, 0x58, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x44, 0x69, 0x66, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x44, 0x69, 0x66, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x92, 0x02, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x41, 0x4d,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x44, 0x10,
	0x07, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x52,
	0x45, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x09, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x10, 0x0a, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x10, 0x0b,
	0x12, 0x14, 0x0a, 0x10, 0x44, 0x49, 0x53, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x50, 0x4c,
	0x41, 0x43, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x0e, 0x12, 0x16,
	0x0a, 0x12, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0f, 0x22, 0x32, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54,
	0x49, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x56,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x32, 0x6f, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x10,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x22, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x65, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pulumi_resource_status_proto_rawDescOnce sync.Once
	file_pulumi_resource_status_proto_rawDescData = file_pulumi_resource_status_proto_rawDesc
)

func file_pulumi_resource_status_proto_rawDescGZIP() []byte {
	file_pulumi_resource_status_proto_rawDescOnce.Do(func() {
		file_pulumi_resource_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_pulumi_resource_status_proto_rawDescData)
	})
	return file_pulumi_resource_status_proto_rawDescData
}

var file_pulumi_resource_status_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pulumi_resource_status_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pulumi_resource_status_proto_goTypes = []interface{}{
	(ViewStep_Op)(0),                 // 0: pulumirpc.ViewStep.Op
	(ViewStep_Status)(0),             // 1: pulumirpc.ViewStep.Status
	(*PublishViewStepsRequest)(nil),  // 2: pulumirpc.PublishViewStepsRequest
	(*PublishViewStepsResponse)(nil), // 3: pulumirpc.PublishViewStepsResponse
	(*ViewStep)(nil),                 // 4: pulumirpc.ViewStep
	(*ViewStepState)(nil),            // 5: pulumirpc.ViewStepState
	nil,                              // 6: pulumirpc.ViewStep.DetailedDiffEntry
	(*structpb.Struct)(nil),          // 7: google.protobuf.Struct
	(*PropertyDiff)(nil),             // 8: pulumirpc.PropertyDiff
}
var file_pulumi_resource_status_proto_depIdxs = []int32{
	4,  // 0: pulumirpc.PublishViewStepsRequest.steps:type_name -> pulumirpc.ViewStep
	1,  // 1: pulumirpc.ViewStep.status:type_name -> pulumirpc.ViewStep.Status
	0,  // 2: pulumirpc.ViewStep.op:type_name -> pulumirpc.ViewStep.Op
	5,  // 3: pulumirpc.ViewStep.old:type_name -> pulumirpc.ViewStepState
	5,  // 4: pulumirpc.ViewStep.new:type_name -> pulumirpc.ViewStepState
	6,  // 5: pulumirpc.ViewStep.detailed_diff:type_name -> pulumirpc.ViewStep.DetailedDiffEntry
	7,  // 6: pulumirpc.ViewStepState.inputs:type_name -> google.protobuf.Struct
	7,  // 7: pulumirpc.ViewStepState.outputs:type_name -> google.protobuf.Struct
	8,  // 8: pulumirpc.ViewStep.DetailedDiffEntry.value:type_name -> pulumirpc.PropertyDiff
	2,  // 9: pulumirpc.ResourceStatus.PublishViewSteps:input_type -> pulumirpc.PublishViewStepsRequest
	3,  // 10: pulumirpc.ResourceStatus.PublishViewSteps:output_type -> pulumirpc.PublishViewStepsResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pulumi_resource_status_proto_init() }
func file_pulumi_resource_status_proto_init() {
	if File_pulumi_resource_status_proto != nil {
		return
	}
	file_pulumi_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pulumi_resource_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishViewStepsRequest); i {
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
		file_pulumi_resource_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishViewStepsResponse); i {
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
		file_pulumi_resource_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewStep); i {
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
		file_pulumi_resource_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewStepState); i {
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
			RawDescriptor: file_pulumi_resource_status_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pulumi_resource_status_proto_goTypes,
		DependencyIndexes: file_pulumi_resource_status_proto_depIdxs,
		EnumInfos:         file_pulumi_resource_status_proto_enumTypes,
		MessageInfos:      file_pulumi_resource_status_proto_msgTypes,
	}.Build()
	File_pulumi_resource_status_proto = out.File
	file_pulumi_resource_status_proto_rawDesc = nil
	file_pulumi_resource_status_proto_goTypes = nil
	file_pulumi_resource_status_proto_depIdxs = nil
}

// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: cloud.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cloud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin           *CloudOrigin `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	AccountId        string       `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountName      string       `protobuf:"bytes,3,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	AvailabilityZone string       `protobuf:"bytes,4,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	InstanceId       string       `protobuf:"bytes,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstanceName     string       `protobuf:"bytes,6,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	MachineType      string       `protobuf:"bytes,7,opt,name=machine_type,json=machineType,proto3" json:"machine_type,omitempty"`
	ProjectId        string       `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName      string       `protobuf:"bytes,9,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Provider         string       `protobuf:"bytes,10,opt,name=provider,proto3" json:"provider,omitempty"`
	Region           string       `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	ServiceName      string       `protobuf:"bytes,12,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *Cloud) Reset() {
	*x = Cloud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cloud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cloud) ProtoMessage() {}

func (x *Cloud) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cloud.ProtoReflect.Descriptor instead.
func (*Cloud) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{0}
}

func (x *Cloud) GetOrigin() *CloudOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Cloud) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Cloud) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *Cloud) GetAvailabilityZone() string {
	if x != nil {
		return x.AvailabilityZone
	}
	return ""
}

func (x *Cloud) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Cloud) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *Cloud) GetMachineType() string {
	if x != nil {
		return x.MachineType
	}
	return ""
}

func (x *Cloud) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Cloud) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *Cloud) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Cloud) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Cloud) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type CloudOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Region      string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *CloudOrigin) Reset() {
	*x = CloudOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudOrigin) ProtoMessage() {}

func (x *CloudOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudOrigin.ProtoReflect.Descriptor instead.
func (*CloudOrigin) Descriptor() ([]byte, []int) {
	return file_cloud_proto_rawDescGZIP(), []int{1}
}

func (x *CloudOrigin) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CloudOrigin) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CloudOrigin) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CloudOrigin) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

var File_cloud_proto protoreflect.FileDescriptor

var file_cloud_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0xad, 0x03,
	0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x83, 0x01,
	0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloud_proto_rawDescOnce sync.Once
	file_cloud_proto_rawDescData = file_cloud_proto_rawDesc
)

func file_cloud_proto_rawDescGZIP() []byte {
	file_cloud_proto_rawDescOnce.Do(func() {
		file_cloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_proto_rawDescData)
	})
	return file_cloud_proto_rawDescData
}

var file_cloud_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloud_proto_goTypes = []interface{}{
	(*Cloud)(nil),       // 0: elastic.apm.v1.Cloud
	(*CloudOrigin)(nil), // 1: elastic.apm.v1.CloudOrigin
}
var file_cloud_proto_depIdxs = []int32{
	1, // 0: elastic.apm.v1.Cloud.origin:type_name -> elastic.apm.v1.CloudOrigin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloud_proto_init() }
func file_cloud_proto_init() {
	if File_cloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cloud); i {
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
		file_cloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudOrigin); i {
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
			RawDescriptor: file_cloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_proto_goTypes,
		DependencyIndexes: file_cloud_proto_depIdxs,
		MessageInfos:      file_cloud_proto_msgTypes,
	}.Build()
	File_cloud_proto = out.File
	file_cloud_proto_rawDesc = nil
	file_cloud_proto_goTypes = nil
	file_cloud_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: job_service.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_job_service_proto protoreflect.FileDescriptor

var file_job_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a,
	0x6f, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72,
	0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc7, 0x0a, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x4a, 0x6f, 0x62,
	0x12, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x42, 0x92, 0x41, 0x24, 0x12, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x1a, 0x18,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x70, 0x6f, 0x73, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01,
	0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x6a, 0x6f, 0x62, 0x12, 0x8c, 0x01, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x12,
	0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x45,
	0x64, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42,
	0x92, 0x41, 0x24, 0x12, 0x08, 0x45, 0x64, 0x69, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x1a, 0x18, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x65,
	0x64, 0x69, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x6a,
	0x6f, 0x62, 0x12, 0x92, 0x01, 0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x12,
	0x1e, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x45, 0x92, 0x41, 0x26, 0x12, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x20, 0x6a, 0x6f, 0x62,
	0x1a, 0x19, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74,
	0x6f, 0x20, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x20, 0x6a, 0x6f, 0x62, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x12, 0xd4, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x2c, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a,
	0x6f, 0x62, 0x42, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x6f, 0x62,
	0x42, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5d, 0x92, 0x41, 0x36, 0x12, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x6a, 0x6f, 0x62,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20,
	0x6a, 0x6f, 0x62, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x9e,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4b, 0x92, 0x41, 0x30, 0x12, 0x0d, 0x67, 0x65, 0x74, 0x20, 0x6a, 0x6f, 0x62,
	0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x1a, 0x1f, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6a, 0x6f,
	0x62, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x8e, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d,
	0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x92,
	0x41, 0x2a, 0x12, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x6a, 0x6f, 0x62, 0x73, 0x1a,
	0x1b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x6a, 0x6f, 0x62, 0x73, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73,
	0x12, 0xb1, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6a,
	0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41,
	0x36, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x73, 0x20, 0x62, 0x79, 0x20, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x73, 0x20, 0x62,
	0x79, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0xc8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x27, 0x2e,
	0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x44, 0x12, 0x18, 0x47, 0x65, 0x74,
	0x20, 0x6a, 0x6f, 0x62, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x62, 0x79, 0x20, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x1a, 0x28, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x6a, 0x6f, 0x62, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x62, 0x79, 0x20, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x42,
	0x98, 0x01, 0x92, 0x41, 0x64, 0x12, 0x62, 0x0a, 0x14, 0x4a, 0x6f, 0x62, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x67, 0x52, 0x50, 0x43, 0x20, 0x41, 0x50, 0x49, 0x22, 0x45, 0x0a,
	0x09, 0x4a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x20, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x45, 0x43, 0x2d, 0x4a, 0x6f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x1a, 0x16, 0x74, 0x68,
	0x61, 0x6e, 0x68, 0x71, 0x75, 0x79, 0x31, 0x31, 0x30, 0x35, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x45, 0x43, 0x2d, 0x4a, 0x6f, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x6a, 0x6f, 0x62, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_job_service_proto_goTypes = []interface{}{
	(*PostJobRequest)(nil),                 // 0: jobstreet.job.PostJobRequest
	(*EditJobRequest)(nil),                 // 1: jobstreet.job.EditJobRequest
	(*CloseJobRequest)(nil),                // 2: jobstreet.job.CloseJobRequest
	(*ChangeStatusJobByAdminRequest)(nil),  // 3: jobstreet.job.ChangeStatusJobByAdminRequest
	(*GetJobByIDRequest)(nil),              // 4: jobstreet.job.GetJobByIDRequest
	(*JobListRequest)(nil),                 // 5: jobstreet.job.JobListRequest
	(*JobListByAdminRequest)(nil),          // 6: jobstreet.job.JobListByAdminRequest
	(*JobListByEmployerRequest)(nil),       // 7: jobstreet.job.JobListByEmployerRequest
	(*PostJobResponse)(nil),                // 8: jobstreet.job.PostJobResponse
	(*EditJobResponse)(nil),                // 9: jobstreet.job.EditJobResponse
	(*CloseJobResponse)(nil),               // 10: jobstreet.job.CloseJobResponse
	(*ChangeStatusJobByAdminResponse)(nil), // 11: jobstreet.job.ChangeStatusJobByAdminResponse
	(*GetJobByIDResponse)(nil),             // 12: jobstreet.job.GetJobByIDResponse
	(*JobListResponse)(nil),                // 13: jobstreet.job.JobListResponse
}
var file_job_service_proto_depIdxs = []int32{
	0,  // 0: jobstreet.job.JobService.PostJob:input_type -> jobstreet.job.PostJobRequest
	1,  // 1: jobstreet.job.JobService.EditJob:input_type -> jobstreet.job.EditJobRequest
	2,  // 2: jobstreet.job.JobService.CloseJob:input_type -> jobstreet.job.CloseJobRequest
	3,  // 3: jobstreet.job.JobService.ChangeStatusJobByAdmin:input_type -> jobstreet.job.ChangeStatusJobByAdminRequest
	4,  // 4: jobstreet.job.JobService.GetJobByID:input_type -> jobstreet.job.GetJobByIDRequest
	5,  // 5: jobstreet.job.JobService.GetJobList:input_type -> jobstreet.job.JobListRequest
	6,  // 6: jobstreet.job.JobService.GetJobListByAdmin:input_type -> jobstreet.job.JobListByAdminRequest
	7,  // 7: jobstreet.job.JobService.GetJobListByEmployer:input_type -> jobstreet.job.JobListByEmployerRequest
	8,  // 8: jobstreet.job.JobService.PostJob:output_type -> jobstreet.job.PostJobResponse
	9,  // 9: jobstreet.job.JobService.EditJob:output_type -> jobstreet.job.EditJobResponse
	10, // 10: jobstreet.job.JobService.CloseJob:output_type -> jobstreet.job.CloseJobResponse
	11, // 11: jobstreet.job.JobService.ChangeStatusJobByAdmin:output_type -> jobstreet.job.ChangeStatusJobByAdminResponse
	12, // 12: jobstreet.job.JobService.GetJobByID:output_type -> jobstreet.job.GetJobByIDResponse
	13, // 13: jobstreet.job.JobService.GetJobList:output_type -> jobstreet.job.JobListResponse
	13, // 14: jobstreet.job.JobService.GetJobListByAdmin:output_type -> jobstreet.job.JobListResponse
	13, // 15: jobstreet.job.JobService.GetJobListByEmployer:output_type -> jobstreet.job.JobListResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_job_service_proto_init() }
func file_job_service_proto_init() {
	if File_job_service_proto != nil {
		return
	}
	file_rpc_get_job_by_id_proto_init()
	file_rpc_get_job_list_by_employer_proto_init()
	file_rpc_get_job_list_proto_init()
	file_rpc_post_job_proto_init()
	file_rpc_edit_job_proto_init()
	file_rpc_close_job_proto_init()
	file_rpc_change_status_by_admin_proto_init()
	file_rpc_get_job_list_by_admin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_service_proto_goTypes,
		DependencyIndexes: file_job_service_proto_depIdxs,
	}.Build()
	File_job_service_proto = out.File
	file_job_service_proto_rawDesc = nil
	file_job_service_proto_goTypes = nil
	file_job_service_proto_depIdxs = nil
}
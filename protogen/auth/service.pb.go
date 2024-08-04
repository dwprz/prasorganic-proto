// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/auth/service.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_auth_service_proto protoreflect.FileDescriptor

var file_proto_auth_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64,
	0x4f, 0x74, 0x70, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x4f, 0x74, 0x70, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x77, 0x70, 0x72, 0x7a, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_auth_service_proto_goTypes = []any{
	(*SendOtpRequest)(nil),    // 0: auth.type.SendOtpRequest
	(*VerifyOtpRequest)(nil),  // 1: auth.type.VerifyOtpRequest
	(*emptypb.Empty)(nil),     // 2: google.protobuf.Empty
	(*VerifyOtpResponse)(nil), // 3: auth.type.VerifyOtpResponse
}
var file_proto_auth_service_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.SendOtp:input_type -> auth.type.SendOtpRequest
	1, // 1: auth.AuthService.VerifyOtp:input_type -> auth.type.VerifyOtpRequest
	2, // 2: auth.AuthService.SendOtp:output_type -> google.protobuf.Empty
	3, // 3: auth.AuthService.VerifyOtp:output_type -> auth.type.VerifyOtpResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_service_proto_init() }
func file_proto_auth_service_proto_init() {
	if File_proto_auth_service_proto != nil {
		return
	}
	file_proto_auth_type_otp_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_service_proto_goTypes,
		DependencyIndexes: file_proto_auth_service_proto_depIdxs,
	}.Build()
	File_proto_auth_service_proto = out.File
	file_proto_auth_service_proto_rawDesc = nil
	file_proto_auth_service_proto_goTypes = nil
	file_proto_auth_service_proto_depIdxs = nil
}

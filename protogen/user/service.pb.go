// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/user/service.proto

package user

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

var File_proto_user_service_proto protoreflect.FileDescriptor

var file_proto_user_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12,
	0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x77, 0x70, 0x72, 0x7a, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_user_service_proto_goTypes = []any{
	(*Email)(nil),                  // 0: user.type.Email
	(*RegisterRequest)(nil),        // 1: user.type.RegisterRequest
	(*LoginWithGoogleRequest)(nil), // 2: user.type.LoginWithGoogleRequest
	(*FindUserResponse)(nil),       // 3: user.type.FindUserResponse
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_proto_user_service_proto_depIdxs = []int32{
	0, // 0: user.UserService.FindByEmail:input_type -> user.type.Email
	1, // 1: user.UserService.Create:input_type -> user.type.RegisterRequest
	2, // 2: user.UserService.Upsert:input_type -> user.type.LoginWithGoogleRequest
	3, // 3: user.UserService.FindByEmail:output_type -> user.type.FindUserResponse
	4, // 4: user.UserService.Create:output_type -> google.protobuf.Empty
	3, // 5: user.UserService.Upsert:output_type -> user.type.FindUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_service_proto_init() }
func file_proto_user_service_proto_init() {
	if File_proto_user_service_proto != nil {
		return
	}
	file_proto_user_type_login_google_proto_init()
	file_proto_user_type_register_request_proto_init()
	file_proto_user_type_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_service_proto_goTypes,
		DependencyIndexes: file_proto_user_service_proto_depIdxs,
	}.Build()
	File_proto_user_service_proto = out.File
	file_proto_user_service_proto_rawDesc = nil
	file_proto_user_service_proto_goTypes = nil
	file_proto_user_service_proto_depIdxs = nil
}

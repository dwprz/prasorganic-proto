// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/user/type/user.proto

package user

import (
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string                 `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Email        string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName     string                 `protobuf:"bytes,3,opt,name=full_name,proto3" json:"full_name,omitempty"`
	Whatsapp     string                 `protobuf:"bytes,4,opt,name=whatsapp,proto3" json:"whatsapp,omitempty"`
	PhotoProfile string                 `protobuf:"bytes,5,opt,name=photo_profile,proto3" json:"photo_profile,omitempty"`
	Password     string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	RefreshToken string                 `protobuf:"bytes,7,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	Role         string                 `protobuf:"bytes,10,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_type_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_type_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_user_type_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetWhatsapp() string {
	if x != nil {
		return x.Whatsapp
	}
	return ""
}

func (x *User) GetPhotoProfile() string {
	if x != nil {
		return x.PhotoProfile
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type FindUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *User `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_type_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_type_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_type_user_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserResponse) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_user_type_user_proto protoreflect.FileDescriptor

var file_proto_user_type_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x37, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x70,
	0x72, 0x7a, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_type_user_proto_rawDescOnce sync.Once
	file_proto_user_type_user_proto_rawDescData = file_proto_user_type_user_proto_rawDesc
)

func file_proto_user_type_user_proto_rawDescGZIP() []byte {
	file_proto_user_type_user_proto_rawDescOnce.Do(func() {
		file_proto_user_type_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_type_user_proto_rawDescData)
	})
	return file_proto_user_type_user_proto_rawDescData
}

var file_proto_user_type_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_type_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: user.type.User
	(*FindUserResponse)(nil),      // 1: user.type.FindUserResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_user_type_user_proto_depIdxs = []int32{
	2, // 0: user.type.User.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: user.type.User.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: user.type.FindUserResponse.data:type_name -> user.type.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_user_type_user_proto_init() }
func file_proto_user_type_user_proto_init() {
	if File_proto_user_type_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_type_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_proto_user_type_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FindUserResponse); i {
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
			RawDescriptor: file_proto_user_type_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_type_user_proto_goTypes,
		DependencyIndexes: file_proto_user_type_user_proto_depIdxs,
		MessageInfos:      file_proto_user_type_user_proto_msgTypes,
	}.Build()
	File_proto_user_type_user_proto = out.File
	file_proto_user_type_user_proto_rawDesc = nil
	file_proto_user_type_user_proto_goTypes = nil
	file_proto_user_type_user_proto_depIdxs = nil
}

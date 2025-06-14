// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: auth/auth.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	Role_ADMIN            Role = 1
	Role_USER             Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ADMIN",
		2: "USER",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ADMIN":            1,
		"USER":             2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_auth_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_auth_auth_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            string                 `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_auth_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Role          Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=auth.Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_auth_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LoginResponse) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

type ConfirmRegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Code          string                 `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmRegisterRequest) Reset() {
	*x = ConfirmRegisterRequest{}
	mi := &file_auth_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRegisterRequest) ProtoMessage() {}

func (x *ConfirmRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRegisterRequest.ProtoReflect.Descriptor instead.
func (*ConfirmRegisterRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmRegisterRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ConfirmRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ConfirmRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConfirmRegisterRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ConfirmRegisterRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ConfirmRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Role          Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=auth.Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmRegisterResponse) Reset() {
	*x = ConfirmRegisterResponse{}
	mi := &file_auth_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRegisterResponse) ProtoMessage() {}

func (x *ConfirmRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRegisterResponse.ProtoReflect.Descriptor instead.
func (*ConfirmRegisterResponse) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmRegisterResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ConfirmRegisterResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ConfirmRegisterResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ConfirmRegisterResponse) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

var File_auth_auth_proto protoreflect.FileDescriptor

const file_auth_auth_proto_rawDesc = "" +
	"\n" +
	"\x0fauth/auth.proto\x12\x04auth\"o\n" +
	"\x0fRegisterRequest\x12\x14\n" +
	"\x05phone\x18\x01 \x01(\tR\x05phone\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\"\"\n" +
	"\x10RegisterResponse\x12\x0e\n" +
	"\x02ok\x18\x01 \x01(\tR\x02ok\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"q\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x02 \x01(\tR\x05phone\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x1e\n" +
	"\x04role\x18\x04 \x01(\x0e2\n" +
	".auth.RoleR\x04role\"\x8a\x01\n" +
	"\x16ConfirmRegisterRequest\x12\x14\n" +
	"\x05phone\x18\x01 \x01(\tR\x05phone\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12\x12\n" +
	"\x04code\x18\x05 \x01(\tR\x04code\"{\n" +
	"\x17ConfirmRegisterResponse\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x02 \x01(\tR\x05phone\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x1e\n" +
	"\x04role\x18\x04 \x01(\x0e2\n" +
	".auth.RoleR\x04role*1\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05ADMIN\x10\x01\x12\b\n" +
	"\x04USER\x10\x022\xc3\x01\n" +
	"\x04Auth\x129\n" +
	"\bRegister\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x120\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x12N\n" +
	"\x0fConfirmRegister\x12\x1c.auth.ConfirmRegisterRequest\x1a\x1d.auth.ConfirmRegisterResponseB\x19Z\x17emgushov.auth.v1;authv1b\x06proto3"

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData []byte
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)))
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_auth_proto_goTypes = []any{
	(Role)(0),                       // 0: auth.Role
	(*RegisterRequest)(nil),         // 1: auth.RegisterRequest
	(*RegisterResponse)(nil),        // 2: auth.RegisterResponse
	(*LoginRequest)(nil),            // 3: auth.LoginRequest
	(*LoginResponse)(nil),           // 4: auth.LoginResponse
	(*ConfirmRegisterRequest)(nil),  // 5: auth.ConfirmRegisterRequest
	(*ConfirmRegisterResponse)(nil), // 6: auth.ConfirmRegisterResponse
}
var file_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.LoginResponse.role:type_name -> auth.Role
	0, // 1: auth.ConfirmRegisterResponse.role:type_name -> auth.Role
	1, // 2: auth.Auth.Register:input_type -> auth.RegisterRequest
	3, // 3: auth.Auth.Login:input_type -> auth.LoginRequest
	5, // 4: auth.Auth.ConfirmRegister:input_type -> auth.ConfirmRegisterRequest
	2, // 5: auth.Auth.Register:output_type -> auth.RegisterResponse
	4, // 6: auth.Auth.Login:output_type -> auth.LoginResponse
	6, // 7: auth.Auth.ConfirmRegister:output_type -> auth.ConfirmRegisterResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		EnumInfos:         file_auth_auth_proto_enumTypes,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}

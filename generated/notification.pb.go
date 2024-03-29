// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: notification.proto

package qreeket

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

type NotificationChannelType int32

const (
	NotificationChannelType_UNKNOWN            NotificationChannelType = 0
	NotificationChannelType_E2E_PERSONAL_CHAT  NotificationChannelType = 1
	NotificationChannelType_E2E_GROUP_CHAT     NotificationChannelType = 2
	NotificationChannelType_CHANNEL_INVITATION NotificationChannelType = 3
	NotificationChannelType_TOPIC              NotificationChannelType = 4
	NotificationChannelType_SUBSCRIPTION       NotificationChannelType = 5
	NotificationChannelType_ACCOUNT            NotificationChannelType = 6
	NotificationChannelType_UPDATE             NotificationChannelType = 7
	NotificationChannelType_BROADCAST          NotificationChannelType = 8
)

// Enum value maps for NotificationChannelType.
var (
	NotificationChannelType_name = map[int32]string{
		0: "UNKNOWN",
		1: "E2E_PERSONAL_CHAT",
		2: "E2E_GROUP_CHAT",
		3: "CHANNEL_INVITATION",
		4: "TOPIC",
		5: "SUBSCRIPTION",
		6: "ACCOUNT",
		7: "UPDATE",
		8: "BROADCAST",
	}
	NotificationChannelType_value = map[string]int32{
		"UNKNOWN":            0,
		"E2E_PERSONAL_CHAT":  1,
		"E2E_GROUP_CHAT":     2,
		"CHANNEL_INVITATION": 3,
		"TOPIC":              4,
		"SUBSCRIPTION":       5,
		"ACCOUNT":            6,
		"UPDATE":             7,
		"BROADCAST":          8,
	}
)

func (x NotificationChannelType) Enum() *NotificationChannelType {
	p := new(NotificationChannelType)
	*p = x
	return p
}

func (x NotificationChannelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationChannelType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationChannelType) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[0]
}

func (x NotificationChannelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationChannelType.Descriptor instead.
func (NotificationChannelType) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

type SendNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Types that are assignable to NotificationTarget:
	//
	//	*SendNotificationRequest_Token
	//	*SendNotificationRequest_Topic
	NotificationTarget isSendNotificationRequest_NotificationTarget `protobuf_oneof:"notification_target"`
	Data               map[string]string                            `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ChannelType        NotificationChannelType                      `protobuf:"varint,6,opt,name=channel_type,json=channelType,proto3,enum=qreeket.NotificationChannelType" json:"channel_type,omitempty"`
}

func (x *SendNotificationRequest) Reset() {
	*x = SendNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationRequest) ProtoMessage() {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *SendNotificationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotificationRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (m *SendNotificationRequest) GetNotificationTarget() isSendNotificationRequest_NotificationTarget {
	if m != nil {
		return m.NotificationTarget
	}
	return nil
}

func (x *SendNotificationRequest) GetToken() string {
	if x, ok := x.GetNotificationTarget().(*SendNotificationRequest_Token); ok {
		return x.Token
	}
	return ""
}

func (x *SendNotificationRequest) GetTopic() string {
	if x, ok := x.GetNotificationTarget().(*SendNotificationRequest_Topic); ok {
		return x.Topic
	}
	return ""
}

func (x *SendNotificationRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SendNotificationRequest) GetChannelType() NotificationChannelType {
	if x != nil {
		return x.ChannelType
	}
	return NotificationChannelType_UNKNOWN
}

type isSendNotificationRequest_NotificationTarget interface {
	isSendNotificationRequest_NotificationTarget()
}

type SendNotificationRequest_Token struct {
	Token string `protobuf:"bytes,3,opt,name=token,proto3,oneof"`
}

type SendNotificationRequest_Topic struct {
	Topic string `protobuf:"bytes,4,opt,name=topic,proto3,oneof"`
}

func (*SendNotificationRequest_Token) isSendNotificationRequest_NotificationTarget() {}

func (*SendNotificationRequest_Topic) isSendNotificationRequest_NotificationTarget() {}

type SendNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendNotificationResponse) Reset() {
	*x = SendNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationResponse) ProtoMessage() {}

func (x *SendNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationResponse.ProtoReflect.Descriptor instead.
func (*SendNotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

func (x *SendNotificationResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

type RegisterDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *RegisterDeviceRequest) Reset() {
	*x = RegisterDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceRequest) ProtoMessage() {}

func (x *RegisterDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterDeviceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterDeviceRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x22, 0xc8, 0x02,
	0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x71, 0x72, 0x65, 0x65,
	0x6b, 0x65, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x15, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x39, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2a, 0xae, 0x01, 0x0a, 0x17, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x32, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41,
	0x4c, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x32, 0x45, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x04, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52,
	0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x08, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x6c, 0x6c, 0x63, 0x2f, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notification_proto_goTypes = []interface{}{
	(NotificationChannelType)(0),     // 0: qreeket.NotificationChannelType
	(*SendNotificationRequest)(nil),  // 1: qreeket.SendNotificationRequest
	(*SendNotificationResponse)(nil), // 2: qreeket.SendNotificationResponse
	(*RegisterDeviceRequest)(nil),    // 3: qreeket.RegisterDeviceRequest
	nil,                              // 4: qreeket.SendNotificationRequest.DataEntry
}
var file_notification_proto_depIdxs = []int32{
	4, // 0: qreeket.SendNotificationRequest.data:type_name -> qreeket.SendNotificationRequest.DataEntry
	0, // 1: qreeket.SendNotificationRequest.channel_type:type_name -> qreeket.NotificationChannelType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationRequest); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationResponse); i {
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
		file_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeviceRequest); i {
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
	file_notification_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SendNotificationRequest_Token)(nil),
		(*SendNotificationRequest_Topic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		EnumInfos:         file_notification_proto_enumTypes,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}

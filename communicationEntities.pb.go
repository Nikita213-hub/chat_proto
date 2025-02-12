// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.15.8
// source: communicationEntities.proto

package chat_proto

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

type ChatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        int32                  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_communicationEntities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_communicationEntities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_communicationEntities_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ErrorMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	mi := &file_communicationEntities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_communicationEntities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_communicationEntities_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_communicationEntities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_communicationEntities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_communicationEntities_proto_rawDescGZIP(), []int{2}
}

func (x *Notification) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type WrapperMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Msg:
	//
	//	*WrapperMessage_Cm
	//	*WrapperMessage_Em
	//	*WrapperMessage_Nm
	Msg           isWrapperMessage_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WrapperMessage) Reset() {
	*x = WrapperMessage{}
	mi := &file_communicationEntities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WrapperMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperMessage) ProtoMessage() {}

func (x *WrapperMessage) ProtoReflect() protoreflect.Message {
	mi := &file_communicationEntities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperMessage.ProtoReflect.Descriptor instead.
func (*WrapperMessage) Descriptor() ([]byte, []int) {
	return file_communicationEntities_proto_rawDescGZIP(), []int{3}
}

func (x *WrapperMessage) GetMsg() isWrapperMessage_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *WrapperMessage) GetCm() *ChatMessage {
	if x != nil {
		if x, ok := x.Msg.(*WrapperMessage_Cm); ok {
			return x.Cm
		}
	}
	return nil
}

func (x *WrapperMessage) GetEm() *ErrorMessage {
	if x != nil {
		if x, ok := x.Msg.(*WrapperMessage_Em); ok {
			return x.Em
		}
	}
	return nil
}

func (x *WrapperMessage) GetNm() *Notification {
	if x != nil {
		if x, ok := x.Msg.(*WrapperMessage_Nm); ok {
			return x.Nm
		}
	}
	return nil
}

type isWrapperMessage_Msg interface {
	isWrapperMessage_Msg()
}

type WrapperMessage_Cm struct {
	Cm *ChatMessage `protobuf:"bytes,1,opt,name=cm,proto3,oneof"`
}

type WrapperMessage_Em struct {
	Em *ErrorMessage `protobuf:"bytes,2,opt,name=em,proto3,oneof"`
}

type WrapperMessage_Nm struct {
	Nm *Notification `protobuf:"bytes,3,opt,name=nm,proto3,oneof"`
}

func (*WrapperMessage_Cm) isWrapperMessage_Msg() {}

func (*WrapperMessage_Em) isWrapperMessage_Msg() {}

func (*WrapperMessage_Nm) isWrapperMessage_Msg() {}

var File_communicationEntities_proto protoreflect.FileDescriptor

var file_communicationEntities_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x28, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x0e, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x02,
	0x63, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x02,
	0x63, 0x6d, 0x12, 0x35, 0x0a, 0x02, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x02, 0x65, 0x6d, 0x12, 0x35, 0x0a, 0x02, 0x6e, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x6e, 0x6d,
	0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_communicationEntities_proto_rawDescOnce sync.Once
	file_communicationEntities_proto_rawDescData []byte
)

func file_communicationEntities_proto_rawDescGZIP() []byte {
	file_communicationEntities_proto_rawDescOnce.Do(func() {
		file_communicationEntities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_communicationEntities_proto_rawDesc), len(file_communicationEntities_proto_rawDesc)))
	})
	return file_communicationEntities_proto_rawDescData
}

var file_communicationEntities_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_communicationEntities_proto_goTypes = []any{
	(*ChatMessage)(nil),    // 0: communicationEntities.ChatMessage
	(*ErrorMessage)(nil),   // 1: communicationEntities.ErrorMessage
	(*Notification)(nil),   // 2: communicationEntities.Notification
	(*WrapperMessage)(nil), // 3: communicationEntities.WrapperMessage
}
var file_communicationEntities_proto_depIdxs = []int32{
	0, // 0: communicationEntities.WrapperMessage.cm:type_name -> communicationEntities.ChatMessage
	1, // 1: communicationEntities.WrapperMessage.em:type_name -> communicationEntities.ErrorMessage
	2, // 2: communicationEntities.WrapperMessage.nm:type_name -> communicationEntities.Notification
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_communicationEntities_proto_init() }
func file_communicationEntities_proto_init() {
	if File_communicationEntities_proto != nil {
		return
	}
	file_communicationEntities_proto_msgTypes[3].OneofWrappers = []any{
		(*WrapperMessage_Cm)(nil),
		(*WrapperMessage_Em)(nil),
		(*WrapperMessage_Nm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_communicationEntities_proto_rawDesc), len(file_communicationEntities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_communicationEntities_proto_goTypes,
		DependencyIndexes: file_communicationEntities_proto_depIdxs,
		MessageInfos:      file_communicationEntities_proto_msgTypes,
	}.Build()
	File_communicationEntities_proto = out.File
	file_communicationEntities_proto_goTypes = nil
	file_communicationEntities_proto_depIdxs = nil
}

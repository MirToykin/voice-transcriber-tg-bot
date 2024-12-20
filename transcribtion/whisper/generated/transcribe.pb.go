// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: transcribe.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request containing the path to the audio file
type TranscribeByPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string  `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Lang     *string `protobuf:"bytes,2,opt,name=lang,proto3,oneof" json:"lang,omitempty"`
}

func (x *TranscribeByPathRequest) Reset() {
	*x = TranscribeByPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscribeByPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscribeByPathRequest) ProtoMessage() {}

func (x *TranscribeByPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transcribe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscribeByPathRequest.ProtoReflect.Descriptor instead.
func (*TranscribeByPathRequest) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{0}
}

func (x *TranscribeByPathRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *TranscribeByPathRequest) GetLang() string {
	if x != nil && x.Lang != nil {
		return *x.Lang
	}
	return ""
}

// Request containing the binary data of the audio file
type TranscribeByBinaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioData []byte  `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
	Lang      *string `protobuf:"bytes,2,opt,name=lang,proto3,oneof" json:"lang,omitempty"`
}

func (x *TranscribeByBinaryRequest) Reset() {
	*x = TranscribeByBinaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscribeByBinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscribeByBinaryRequest) ProtoMessage() {}

func (x *TranscribeByBinaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transcribe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscribeByBinaryRequest.ProtoReflect.Descriptor instead.
func (*TranscribeByBinaryRequest) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{1}
}

func (x *TranscribeByBinaryRequest) GetAudioData() []byte {
	if x != nil {
		return x.AudioData
	}
	return nil
}

func (x *TranscribeByBinaryRequest) GetLang() string {
	if x != nil && x.Lang != nil {
		return *x.Lang
	}
	return ""
}

// Response containing the transcription text
type TranscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text             string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Status           bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ErrorCode        int32  `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorDescription string `protobuf:"bytes,4,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
}

func (x *TranscriptionResponse) Reset() {
	*x = TranscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptionResponse) ProtoMessage() {}

func (x *TranscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transcribe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptionResponse.ProtoReflect.Descriptor instead.
func (*TranscriptionResponse) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{2}
}

func (x *TranscriptionResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TranscriptionResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *TranscriptionResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *TranscriptionResponse) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

// Response containing the available languages list
type AvailableLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []string `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
}

func (x *AvailableLanguagesResponse) Reset() {
	*x = AvailableLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableLanguagesResponse) ProtoMessage() {}

func (x *AvailableLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transcribe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableLanguagesResponse.ProtoReflect.Descriptor instead.
func (*AvailableLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{3}
}

func (x *AvailableLanguagesResponse) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

var File_transcribe_proto protoreflect.FileDescriptor

var file_transcribe_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x58, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x5c, 0x0a, 0x19, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x1a, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x32, 0xf4, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1b, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x72, 0x54, 0x6f,
	0x79, 0x6b, 0x69, 0x6e, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2d, 0x74, 0x67, 0x2d, 0x62, 0x6f, 0x74, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transcribe_proto_rawDescOnce sync.Once
	file_transcribe_proto_rawDescData = file_transcribe_proto_rawDesc
)

func file_transcribe_proto_rawDescGZIP() []byte {
	file_transcribe_proto_rawDescOnce.Do(func() {
		file_transcribe_proto_rawDescData = protoimpl.X.CompressGZIP(file_transcribe_proto_rawDescData)
	})
	return file_transcribe_proto_rawDescData
}

var file_transcribe_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transcribe_proto_goTypes = []interface{}{
	(*TranscribeByPathRequest)(nil),    // 0: TranscribeByPathRequest
	(*TranscribeByBinaryRequest)(nil),  // 1: TranscribeByBinaryRequest
	(*TranscriptionResponse)(nil),      // 2: TranscriptionResponse
	(*AvailableLanguagesResponse)(nil), // 3: AvailableLanguagesResponse
	(*emptypb.Empty)(nil),              // 4: google.protobuf.Empty
}
var file_transcribe_proto_depIdxs = []int32{
	0, // 0: TranscriptionService.TranscribeByPath:input_type -> TranscribeByPathRequest
	1, // 1: TranscriptionService.TranscribeByBinary:input_type -> TranscribeByBinaryRequest
	4, // 2: TranscriptionService.GetAvailableLanguages:input_type -> google.protobuf.Empty
	2, // 3: TranscriptionService.TranscribeByPath:output_type -> TranscriptionResponse
	2, // 4: TranscriptionService.TranscribeByBinary:output_type -> TranscriptionResponse
	3, // 5: TranscriptionService.GetAvailableLanguages:output_type -> AvailableLanguagesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transcribe_proto_init() }
func file_transcribe_proto_init() {
	if File_transcribe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transcribe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscribeByPathRequest); i {
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
		file_transcribe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscribeByBinaryRequest); i {
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
		file_transcribe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscriptionResponse); i {
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
		file_transcribe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableLanguagesResponse); i {
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
	file_transcribe_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_transcribe_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transcribe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transcribe_proto_goTypes,
		DependencyIndexes: file_transcribe_proto_depIdxs,
		MessageInfos:      file_transcribe_proto_msgTypes,
	}.Build()
	File_transcribe_proto = out.File
	file_transcribe_proto_rawDesc = nil
	file_transcribe_proto_goTypes = nil
	file_transcribe_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: transcribe.proto

package generated

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

// Request containing the path to the audio file
type TranscribePathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *TranscribePathRequest) Reset() {
	*x = TranscribePathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscribePathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscribePathRequest) ProtoMessage() {}

func (x *TranscribePathRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TranscribePathRequest.ProtoReflect.Descriptor instead.
func (*TranscribePathRequest) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{0}
}

func (x *TranscribePathRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

// Request containing the binary data of the audio file
type TranscribeBinaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioData []byte `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
}

func (x *TranscribeBinaryRequest) Reset() {
	*x = TranscribeBinaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transcribe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscribeBinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscribeBinaryRequest) ProtoMessage() {}

func (x *TranscribeBinaryRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TranscribeBinaryRequest.ProtoReflect.Descriptor instead.
func (*TranscribeBinaryRequest) Descriptor() ([]byte, []int) {
	return file_transcribe_proto_rawDescGZIP(), []int{1}
}

func (x *TranscribeBinaryRequest) GetAudioData() []byte {
	if x != nil {
		return x.AudioData
	}
	return nil
}

// Response containing the transcription text
type TranscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
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

var File_transcribe_proto protoreflect.FileDescriptor

var file_transcribe_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x34,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x38, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2b,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x14,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_transcribe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transcribe_proto_goTypes = []interface{}{
	(*TranscribePathRequest)(nil),   // 0: transcribe.TranscribePathRequest
	(*TranscribeBinaryRequest)(nil), // 1: transcribe.TranscribeBinaryRequest
	(*TranscriptionResponse)(nil),   // 2: transcribe.TranscriptionResponse
}
var file_transcribe_proto_depIdxs = []int32{
	0, // 0: transcribe.TranscriptionService.TranscribeByPath:input_type -> transcribe.TranscribePathRequest
	1, // 1: transcribe.TranscriptionService.TranscribeByBinary:input_type -> transcribe.TranscribeBinaryRequest
	2, // 2: transcribe.TranscriptionService.TranscribeByPath:output_type -> transcribe.TranscriptionResponse
	2, // 3: transcribe.TranscriptionService.TranscribeByBinary:output_type -> transcribe.TranscriptionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
			switch v := v.(*TranscribePathRequest); i {
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
			switch v := v.(*TranscribeBinaryRequest); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transcribe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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

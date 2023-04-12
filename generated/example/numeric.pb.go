// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: validate/numeric.proto

package example

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Numerics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// x must equal 1.23 exactly
	X1 float32 `protobuf:"fixed32,1,opt,name=x1,proto3" json:"x1,omitempty"` //lt/lte/gt/gte: these inequalities (<, <=, >, >=, respectively) allow for deriving ranges in which the field must reside.
	// x must be less than 10
	X2 int32 `protobuf:"varint,2,opt,name=x2,proto3" json:"x2,omitempty"`
	// x must be greater than or equal to 20
	X3 uint64 `protobuf:"varint,3,opt,name=x3,proto3" json:"x3,omitempty"`
	// x must be in the range [30, 40)
	X4 uint32 `protobuf:"fixed32,4,opt,name=x4,proto3" json:"x4,omitempty"` //Inverting the values of lt(e) and gt(e) is valid and creates an exclusive range.
	// x must be outside the range [30, 40)
	X5 float64 `protobuf:"fixed64,5,opt,name=x5,proto3" json:"x5,omitempty"` //in/not_in: these two rules permit specifying allow/denylists for the values of a field.
	// x must be either 1, 2, or 3
	X6 uint32 `protobuf:"varint,6,opt,name=x6,proto3" json:"x6,omitempty"`
	// x cannot be 0 nor 0.99
	X7 float32 `protobuf:"fixed32,7,opt,name=x7,proto3" json:"x7,omitempty"` //ignore_empty: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.
	X8 uint32  `protobuf:"varint,8,opt,name=x8,proto3" json:"x8,omitempty"`  //unint32
}

func (x *Numerics) Reset() {
	*x = Numerics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validate_numeric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numerics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numerics) ProtoMessage() {}

func (x *Numerics) ProtoReflect() protoreflect.Message {
	mi := &file_validate_numeric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numerics.ProtoReflect.Descriptor instead.
func (*Numerics) Descriptor() ([]byte, []int) {
	return file_validate_numeric_proto_rawDescGZIP(), []int{0}
}

func (x *Numerics) GetX1() float32 {
	if x != nil {
		return x.X1
	}
	return 0
}

func (x *Numerics) GetX2() int32 {
	if x != nil {
		return x.X2
	}
	return 0
}

func (x *Numerics) GetX3() uint64 {
	if x != nil {
		return x.X3
	}
	return 0
}

func (x *Numerics) GetX4() uint32 {
	if x != nil {
		return x.X4
	}
	return 0
}

func (x *Numerics) GetX5() float64 {
	if x != nil {
		return x.X5
	}
	return 0
}

func (x *Numerics) GetX6() uint32 {
	if x != nil {
		return x.X6
	}
	return 0
}

func (x *Numerics) GetX7() float32 {
	if x != nil {
		return x.X7
	}
	return 0
}

func (x *Numerics) GetX8() uint32 {
	if x != nil {
		return x.X8
	}
	return 0
}

var File_validate_numeric_proto protoreflect.FileDescriptor

var file_validate_numeric_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a,
	0x08, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1a, 0x0a, 0x02, 0x78, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x0d, 0xa4, 0x70, 0x9d,
	0x3f, 0x52, 0x02, 0x78, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x78, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x10, 0x0a, 0x52, 0x02, 0x78, 0x32, 0x12, 0x17,
	0x0a, 0x02, 0x78, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x28, 0x14, 0x52, 0x02, 0x78, 0x33, 0x12, 0x1f, 0x0a, 0x02, 0x78, 0x34, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x07, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x4a, 0x0a, 0x15, 0x28, 0x00, 0x00, 0x00, 0x2d,
	0x1e, 0x00, 0x00, 0x00, 0x52, 0x02, 0x78, 0x34, 0x12, 0x27, 0x0a, 0x02, 0x78, 0x35, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x11, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x3e, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x40, 0x52, 0x02, 0x78,
	0x35, 0x12, 0x1b, 0x0a, 0x02, 0x78, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa,
	0x42, 0x08, 0x2a, 0x06, 0x30, 0x01, 0x30, 0x02, 0x30, 0x03, 0x52, 0x02, 0x78, 0x36, 0x12, 0x1f,
	0x0a, 0x02, 0x78, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x0a,
	0x0a, 0x3d, 0x00, 0x00, 0x00, 0x00, 0x3d, 0xa4, 0x70, 0x7d, 0x3f, 0x52, 0x02, 0x78, 0x37, 0x12,
	0x1a, 0x0a, 0x02, 0x78, 0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x2a, 0x05, 0x28, 0xc8, 0x01, 0x40, 0x01, 0x52, 0x02, 0x78, 0x38, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_validate_numeric_proto_rawDescOnce sync.Once
	file_validate_numeric_proto_rawDescData = file_validate_numeric_proto_rawDesc
)

func file_validate_numeric_proto_rawDescGZIP() []byte {
	file_validate_numeric_proto_rawDescOnce.Do(func() {
		file_validate_numeric_proto_rawDescData = protoimpl.X.CompressGZIP(file_validate_numeric_proto_rawDescData)
	})
	return file_validate_numeric_proto_rawDescData
}

var file_validate_numeric_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_validate_numeric_proto_goTypes = []interface{}{
	(*Numerics)(nil), // 0: examplepb.Numerics
}
var file_validate_numeric_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_validate_numeric_proto_init() }
func file_validate_numeric_proto_init() {
	if File_validate_numeric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validate_numeric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numerics); i {
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
			RawDescriptor: file_validate_numeric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_validate_numeric_proto_goTypes,
		DependencyIndexes: file_validate_numeric_proto_depIdxs,
		MessageInfos:      file_validate_numeric_proto_msgTypes,
	}.Build()
	File_validate_numeric_proto = out.File
	file_validate_numeric_proto_rawDesc = nil
	file_validate_numeric_proto_goTypes = nil
	file_validate_numeric_proto_depIdxs = nil
}
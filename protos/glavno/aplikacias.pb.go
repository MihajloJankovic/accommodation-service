// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: aplikacias.proto

package glavno

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

type DummyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy []*AccommodationResponse `protobuf:"bytes,1,rep,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *DummyList) Reset() {
	*x = DummyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aplikacias_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyList) ProtoMessage() {}

func (x *DummyList) ProtoReflect() protoreflect.Message {
	mi := &file_aplikacias_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyList.ProtoReflect.Descriptor instead.
func (*DummyList) Descriptor() ([]byte, []int) {
	return file_aplikacias_proto_rawDescGZIP(), []int{0}
}

func (x *DummyList) GetDummy() []*AccommodationResponse {
	if x != nil {
		return x.Dummy
	}
	return nil
}

type AccommodationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AccommodationRequest) Reset() {
	*x = AccommodationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aplikacias_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRequest) ProtoMessage() {}

func (x *AccommodationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aplikacias_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRequest.ProtoReflect.Descriptor instead.
func (*AccommodationRequest) Descriptor() ([]byte, []int) {
	return file_aplikacias_proto_rawDescGZIP(), []int{1}
}

func (x *AccommodationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AccommodationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price    string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Adress   string `protobuf:"bytes,4,opt,name=adress,proto3" json:"adress,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AccommodationResponse) Reset() {
	*x = AccommodationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aplikacias_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationResponse) ProtoMessage() {}

func (x *AccommodationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aplikacias_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationResponse.ProtoReflect.Descriptor instead.
func (*AccommodationResponse) Descriptor() ([]byte, []int) {
	return file_aplikacias_proto_rawDescGZIP(), []int{2}
}

func (x *AccommodationResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccommodationResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *AccommodationResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AccommodationResponse) GetAdress() string {
	if x != nil {
		return x.Adress
	}
	return ""
}

func (x *AccommodationResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Emptya struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptya) Reset() {
	*x = Emptya{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aplikacias_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptya) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptya) ProtoMessage() {}

func (x *Emptya) ProtoReflect() protoreflect.Message {
	mi := &file_aplikacias_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptya.ProtoReflect.Descriptor instead.
func (*Emptya) Descriptor() ([]byte, []int) {
	return file_aplikacias_proto_rawDescGZIP(), []int{3}
}

var File_aplikacias_proto protoreflect.FileDescriptor

var file_aplikacias_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x6c, 0x69, 0x6b, 0x61, 0x63, 0x69, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x09, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x2c, 0x0a,
	0x14, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x8b, 0x01, 0x0a, 0x15,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x08, 0x0a, 0x06, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x61, 0x32, 0xdf, 0x01, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x1a, 0x0a, 0x2e, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x07, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x61, 0x12, 0x36, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x07, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x61, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x67, 0x6c, 0x61, 0x76,
	0x6e, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aplikacias_proto_rawDescOnce sync.Once
	file_aplikacias_proto_rawDescData = file_aplikacias_proto_rawDesc
)

func file_aplikacias_proto_rawDescGZIP() []byte {
	file_aplikacias_proto_rawDescOnce.Do(func() {
		file_aplikacias_proto_rawDescData = protoimpl.X.CompressGZIP(file_aplikacias_proto_rawDescData)
	})
	return file_aplikacias_proto_rawDescData
}

var file_aplikacias_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aplikacias_proto_goTypes = []interface{}{
	(*DummyList)(nil),             // 0: DummyList
	(*AccommodationRequest)(nil),  // 1: AccommodationRequest
	(*AccommodationResponse)(nil), // 2: AccommodationResponse
	(*Emptya)(nil),                // 3: Emptya
}
var file_aplikacias_proto_depIdxs = []int32{
	2, // 0: DummyList.dummy:type_name -> AccommodationResponse
	1, // 1: accommodation.GetAccommodation:input_type -> AccommodationRequest
	3, // 2: accommodation.GetAllAccommodation:input_type -> Emptya
	2, // 3: accommodation.SetAccommodation:input_type -> AccommodationResponse
	2, // 4: accommodation.UpdateAccommodation:input_type -> AccommodationResponse
	0, // 5: accommodation.GetAccommodation:output_type -> DummyList
	0, // 6: accommodation.GetAllAccommodation:output_type -> DummyList
	3, // 7: accommodation.SetAccommodation:output_type -> Emptya
	3, // 8: accommodation.UpdateAccommodation:output_type -> Emptya
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aplikacias_proto_init() }
func file_aplikacias_proto_init() {
	if File_aplikacias_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aplikacias_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyList); i {
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
		file_aplikacias_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRequest); i {
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
		file_aplikacias_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationResponse); i {
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
		file_aplikacias_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptya); i {
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
			RawDescriptor: file_aplikacias_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aplikacias_proto_goTypes,
		DependencyIndexes: file_aplikacias_proto_depIdxs,
		MessageInfos:      file_aplikacias_proto_msgTypes,
	}.Build()
	File_aplikacias_proto = out.File
	file_aplikacias_proto_rawDesc = nil
	file_aplikacias_proto_goTypes = nil
	file_aplikacias_proto_depIdxs = nil
}

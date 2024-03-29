// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: dbadapter.proto

package dbadapter

import (
	message "github.com/autotest-plan/rpcdefine/go/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_dbadapter_proto protoreflect.FileDescriptor

var file_dbadapter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x5b, 0x0a, 0x09, 0x44, 0x42, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x06,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x06, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x06,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x74, 0x65, 0x73, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x62, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_dbadapter_proto_goTypes = []interface{}{
	(*message.Filter)(nil), // 0: Filter
	(*message.Tasks)(nil),  // 1: Tasks
	(*message.Task)(nil),   // 2: Task
}
var file_dbadapter_proto_depIdxs = []int32{
	0, // 0: DBAdapter.Load:input_type -> Filter
	0, // 1: DBAdapter.LoadSorted:input_type -> Filter
	1, // 2: DBAdapter.Store:input_type -> Tasks
	1, // 3: DBAdapter.Load:output_type -> Tasks
	1, // 4: DBAdapter.LoadSorted:output_type -> Tasks
	2, // 5: DBAdapter.Store:output_type -> Task
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dbadapter_proto_init() }
func file_dbadapter_proto_init() {
	if File_dbadapter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dbadapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbadapter_proto_goTypes,
		DependencyIndexes: file_dbadapter_proto_depIdxs,
	}.Build()
	File_dbadapter_proto = out.File
	file_dbadapter_proto_rawDesc = nil
	file_dbadapter_proto_goTypes = nil
	file_dbadapter_proto_depIdxs = nil
}

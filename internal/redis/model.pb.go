// Copyright 2020 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: internal/redis/model.proto

package redis

import (
	pb "github.com/daangn/eboolkiq/pb"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JobQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job     *pb.Job              `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Queue   *pb.Queue            `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	StartAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
}

func (x *JobQueue) Reset() {
	*x = JobQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_redis_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobQueue) ProtoMessage() {}

func (x *JobQueue) ProtoReflect() protoreflect.Message {
	mi := &file_internal_redis_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobQueue.ProtoReflect.Descriptor instead.
func (*JobQueue) Descriptor() ([]byte, []int) {
	return file_internal_redis_model_proto_rawDescGZIP(), []int{0}
}

func (x *JobQueue) GetJob() *pb.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *JobQueue) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *JobQueue) GetStartAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

var File_internal_redis_model_proto protoreflect.FileDescriptor

var file_internal_redis_model_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x4a,
	0x6f, 0x62, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x25, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x42, 0x56, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2f, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_redis_model_proto_rawDescOnce sync.Once
	file_internal_redis_model_proto_rawDescData = file_internal_redis_model_proto_rawDesc
)

func file_internal_redis_model_proto_rawDescGZIP() []byte {
	file_internal_redis_model_proto_rawDescOnce.Do(func() {
		file_internal_redis_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_redis_model_proto_rawDescData)
	})
	return file_internal_redis_model_proto_rawDescData
}

var file_internal_redis_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_redis_model_proto_goTypes = []interface{}{
	(*JobQueue)(nil),            // 0: eboolkiq.internal.redis.JobQueue
	(*pb.Job)(nil),              // 1: eboolkiq.Job
	(*pb.Queue)(nil),            // 2: eboolkiq.Queue
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_redis_model_proto_depIdxs = []int32{
	1, // 0: eboolkiq.internal.redis.JobQueue.job:type_name -> eboolkiq.Job
	2, // 1: eboolkiq.internal.redis.JobQueue.queue:type_name -> eboolkiq.Queue
	3, // 2: eboolkiq.internal.redis.JobQueue.start_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_redis_model_proto_init() }
func file_internal_redis_model_proto_init() {
	if File_internal_redis_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_redis_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobQueue); i {
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
			RawDescriptor: file_internal_redis_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_redis_model_proto_goTypes,
		DependencyIndexes: file_internal_redis_model_proto_depIdxs,
		MessageInfos:      file_internal_redis_model_proto_msgTypes,
	}.Build()
	File_internal_redis_model_proto = out.File
	file_internal_redis_model_proto_rawDesc = nil
	file_internal_redis_model_proto_goTypes = nil
	file_internal_redis_model_proto_depIdxs = nil
}

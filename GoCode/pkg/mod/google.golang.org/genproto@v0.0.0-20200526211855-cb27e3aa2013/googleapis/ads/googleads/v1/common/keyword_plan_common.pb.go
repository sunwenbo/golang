// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v1/common/keyword_plan_common.proto

package common

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Historical metrics specific to the targeting options selected.
// Targeting options include geographies, network, etc.
// Refer to https://support.google.com/google-ads/answer/3022575 for more
// details.
type KeywordPlanHistoricalMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Approximate number of monthly searches on this query averaged
	// for the past 12 months.
	AvgMonthlySearches *wrappers.Int64Value `protobuf:"bytes,1,opt,name=avg_monthly_searches,json=avgMonthlySearches,proto3" json:"avg_monthly_searches,omitempty"`
	// The competition level for the query.
	Competition enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel `protobuf:"varint,2,opt,name=competition,proto3,enum=google.ads.googleads.v1.enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel" json:"competition,omitempty"`
}

func (x *KeywordPlanHistoricalMetrics) Reset() {
	*x = KeywordPlanHistoricalMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_common_keyword_plan_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanHistoricalMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanHistoricalMetrics) ProtoMessage() {}

func (x *KeywordPlanHistoricalMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_common_keyword_plan_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanHistoricalMetrics.ProtoReflect.Descriptor instead.
func (*KeywordPlanHistoricalMetrics) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordPlanHistoricalMetrics) GetAvgMonthlySearches() *wrappers.Int64Value {
	if x != nil {
		return x.AvgMonthlySearches
	}
	return nil
}

func (x *KeywordPlanHistoricalMetrics) GetCompetition() enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel {
	if x != nil {
		return x.Competition
	}
	return enums.KeywordPlanCompetitionLevelEnum_UNSPECIFIED
}

var File_google_ads_googleads_v1_common_keyword_plan_common_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x42, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a,
	0x1c, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x4d, 0x0a,
	0x14, 0x61, 0x76, 0x67, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x61, 0x76, 0x67, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x73, 0x12, 0x7c, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xf1, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x42, 0x16, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescData = file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDesc
)

func file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDescData
}

var file_google_ads_googleads_v1_common_keyword_plan_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v1_common_keyword_plan_common_proto_goTypes = []interface{}{
	(*KeywordPlanHistoricalMetrics)(nil),                                   // 0: google.ads.googleads.v1.common.KeywordPlanHistoricalMetrics
	(*wrappers.Int64Value)(nil),                                            // 1: google.protobuf.Int64Value
	(enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel)(0), // 2: google.ads.googleads.v1.enums.KeywordPlanCompetitionLevelEnum.KeywordPlanCompetitionLevel
}
var file_google_ads_googleads_v1_common_keyword_plan_common_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v1.common.KeywordPlanHistoricalMetrics.avg_monthly_searches:type_name -> google.protobuf.Int64Value
	2, // 1: google.ads.googleads.v1.common.KeywordPlanHistoricalMetrics.competition:type_name -> google.ads.googleads.v1.enums.KeywordPlanCompetitionLevelEnum.KeywordPlanCompetitionLevel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_common_keyword_plan_common_proto_init() }
func file_google_ads_googleads_v1_common_keyword_plan_common_proto_init() {
	if File_google_ads_googleads_v1_common_keyword_plan_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_common_keyword_plan_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanHistoricalMetrics); i {
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
			RawDescriptor: file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v1_common_keyword_plan_common_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_common_keyword_plan_common_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v1_common_keyword_plan_common_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_common_keyword_plan_common_proto = out.File
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_rawDesc = nil
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_goTypes = nil
	file_google_ads_googleads_v1_common_keyword_plan_common_proto_depIdxs = nil
}
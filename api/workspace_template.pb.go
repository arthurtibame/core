// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workspace_template.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateWorkspaceTemplateRequest struct {
	Namespace            string             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkspaceTemplate    *WorkspaceTemplate `protobuf:"bytes,2,opt,name=workspaceTemplate,proto3" json:"workspaceTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateWorkspaceTemplateRequest) Reset()         { *m = CreateWorkspaceTemplateRequest{} }
func (m *CreateWorkspaceTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkspaceTemplateRequest) ProtoMessage()    {}
func (*CreateWorkspaceTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{0}
}

func (m *CreateWorkspaceTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkspaceTemplateRequest.Unmarshal(m, b)
}
func (m *CreateWorkspaceTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkspaceTemplateRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkspaceTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkspaceTemplateRequest.Merge(m, src)
}
func (m *CreateWorkspaceTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkspaceTemplateRequest.Size(m)
}
func (m *CreateWorkspaceTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkspaceTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkspaceTemplateRequest proto.InternalMessageInfo

func (m *CreateWorkspaceTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateWorkspaceTemplateRequest) GetWorkspaceTemplate() *WorkspaceTemplate {
	if m != nil {
		return m.WorkspaceTemplate
	}
	return nil
}

type ListWorkspaceTemplatesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkspaceTemplatesRequest) Reset()         { *m = ListWorkspaceTemplatesRequest{} }
func (m *ListWorkspaceTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkspaceTemplatesRequest) ProtoMessage()    {}
func (*ListWorkspaceTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{1}
}

func (m *ListWorkspaceTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkspaceTemplatesRequest.Unmarshal(m, b)
}
func (m *ListWorkspaceTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkspaceTemplatesRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkspaceTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkspaceTemplatesRequest.Merge(m, src)
}
func (m *ListWorkspaceTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkspaceTemplatesRequest.Size(m)
}
func (m *ListWorkspaceTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkspaceTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkspaceTemplatesRequest proto.InternalMessageInfo

func (m *ListWorkspaceTemplatesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListWorkspaceTemplatesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListWorkspaceTemplatesRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListWorkspaceTemplatesResponse struct {
	Count                int32                `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkspaceTemplates   []*WorkspaceTemplate `protobuf:"bytes,2,rep,name=workspaceTemplates,proto3" json:"workspaceTemplates,omitempty"`
	Page                 int32                `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Pages                int32                `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	TotalCount           int32                `protobuf:"varint,5,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListWorkspaceTemplatesResponse) Reset()         { *m = ListWorkspaceTemplatesResponse{} }
func (m *ListWorkspaceTemplatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkspaceTemplatesResponse) ProtoMessage()    {}
func (*ListWorkspaceTemplatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{2}
}

func (m *ListWorkspaceTemplatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkspaceTemplatesResponse.Unmarshal(m, b)
}
func (m *ListWorkspaceTemplatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkspaceTemplatesResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkspaceTemplatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkspaceTemplatesResponse.Merge(m, src)
}
func (m *ListWorkspaceTemplatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkspaceTemplatesResponse.Size(m)
}
func (m *ListWorkspaceTemplatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkspaceTemplatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkspaceTemplatesResponse proto.InternalMessageInfo

func (m *ListWorkspaceTemplatesResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkspaceTemplatesResponse) GetWorkspaceTemplates() []*WorkspaceTemplate {
	if m != nil {
		return m.WorkspaceTemplates
	}
	return nil
}

func (m *ListWorkspaceTemplatesResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListWorkspaceTemplatesResponse) GetPages() int32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ListWorkspaceTemplatesResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type ListWorkspaceTemplateVersionsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkspaceTemplateVersionsRequest) Reset()         { *m = ListWorkspaceTemplateVersionsRequest{} }
func (m *ListWorkspaceTemplateVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkspaceTemplateVersionsRequest) ProtoMessage()    {}
func (*ListWorkspaceTemplateVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{3}
}

func (m *ListWorkspaceTemplateVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkspaceTemplateVersionsRequest.Unmarshal(m, b)
}
func (m *ListWorkspaceTemplateVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkspaceTemplateVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkspaceTemplateVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkspaceTemplateVersionsRequest.Merge(m, src)
}
func (m *ListWorkspaceTemplateVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkspaceTemplateVersionsRequest.Size(m)
}
func (m *ListWorkspaceTemplateVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkspaceTemplateVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkspaceTemplateVersionsRequest proto.InternalMessageInfo

func (m *ListWorkspaceTemplateVersionsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListWorkspaceTemplateVersionsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WorkspaceTemplate struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              int64             `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Manifest             string            `protobuf:"bytes,4,opt,name=manifest,proto3" json:"manifest,omitempty"`
	IsLatest             bool              `protobuf:"varint,5,opt,name=isLatest,proto3" json:"isLatest,omitempty"`
	CreatedAt            string            `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,7,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkspaceTemplate) Reset()         { *m = WorkspaceTemplate{} }
func (m *WorkspaceTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkspaceTemplate) ProtoMessage()    {}
func (*WorkspaceTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{4}
}

func (m *WorkspaceTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceTemplate.Unmarshal(m, b)
}
func (m *WorkspaceTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceTemplate.Marshal(b, m, deterministic)
}
func (m *WorkspaceTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceTemplate.Merge(m, src)
}
func (m *WorkspaceTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkspaceTemplate.Size(m)
}
func (m *WorkspaceTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceTemplate proto.InternalMessageInfo

func (m *WorkspaceTemplate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WorkspaceTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkspaceTemplate) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkspaceTemplate) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *WorkspaceTemplate) GetIsLatest() bool {
	if m != nil {
		return m.IsLatest
	}
	return false
}

func (m *WorkspaceTemplate) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WorkspaceTemplate) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type GenerateWorkspaceTemplateWorkflowTemplateRequest struct {
	Namespace            string             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string             `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	WorkspaceTemplate    *WorkspaceTemplate `protobuf:"bytes,3,opt,name=workspaceTemplate,proto3" json:"workspaceTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) Reset() {
	*m = GenerateWorkspaceTemplateWorkflowTemplateRequest{}
}
func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) String() string {
	return proto.CompactTextString(m)
}
func (*GenerateWorkspaceTemplateWorkflowTemplateRequest) ProtoMessage() {}
func (*GenerateWorkspaceTemplateWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b00f4e2babeeb4c, []int{5}
}

func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest.Merge(m, src)
}
func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest.Size(m)
}
func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateWorkspaceTemplateWorkflowTemplateRequest proto.InternalMessageInfo

func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GenerateWorkspaceTemplateWorkflowTemplateRequest) GetWorkspaceTemplate() *WorkspaceTemplate {
	if m != nil {
		return m.WorkspaceTemplate
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateWorkspaceTemplateRequest)(nil), "api.CreateWorkspaceTemplateRequest")
	proto.RegisterType((*ListWorkspaceTemplatesRequest)(nil), "api.ListWorkspaceTemplatesRequest")
	proto.RegisterType((*ListWorkspaceTemplatesResponse)(nil), "api.ListWorkspaceTemplatesResponse")
	proto.RegisterType((*ListWorkspaceTemplateVersionsRequest)(nil), "api.ListWorkspaceTemplateVersionsRequest")
	proto.RegisterType((*WorkspaceTemplate)(nil), "api.WorkspaceTemplate")
	proto.RegisterType((*GenerateWorkspaceTemplateWorkflowTemplateRequest)(nil), "api.GenerateWorkspaceTemplateWorkflowTemplateRequest")
}

func init() { proto.RegisterFile("workspace_template.proto", fileDescriptor_1b00f4e2babeeb4c) }

var fileDescriptor_1b00f4e2babeeb4c = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xb1, 0x6e, 0x13, 0x4d,
	0x10, 0xd6, 0xfa, 0xe2, 0x24, 0x9e, 0xbf, 0x49, 0x46, 0x3f, 0xc9, 0xe9, 0x08, 0x91, 0xb5, 0xa1,
	0x48, 0x0a, 0x7c, 0x24, 0x28, 0x0d, 0x12, 0x42, 0x21, 0x04, 0x28, 0x52, 0x5d, 0x10, 0x20, 0x1a,
	0x58, 0xdb, 0x6b, 0x6b, 0x85, 0x7d, 0x7b, 0xdc, 0xae, 0x6d, 0x01, 0x4a, 0x83, 0x78, 0x03, 0x1a,
	0x4a, 0x4a, 0x1e, 0x82, 0x07, 0xa0, 0xa7, 0xa4, 0xa5, 0xa4, 0xa5, 0x47, 0xbb, 0x6b, 0x9f, 0x1d,
	0xdf, 0xd9, 0xb2, 0xab, 0xdb, 0xd9, 0x99, 0xd9, 0xf9, 0x66, 0xee, 0xfb, 0x76, 0xc1, 0x1f, 0xc8,
	0xf4, 0x8d, 0x4a, 0x58, 0x83, 0xbf, 0xd2, 0xbc, 0x9b, 0x74, 0x98, 0xe6, 0xb5, 0x24, 0x95, 0x5a,
	0xa2, 0xc7, 0x12, 0x11, 0xec, 0xb4, 0xa5, 0x6c, 0x77, 0x78, 0xc8, 0x12, 0x11, 0xb2, 0x38, 0x96,
	0x9a, 0x69, 0x21, 0x63, 0xe5, 0x42, 0x82, 0x6d, 0x93, 0xdc, 0xea, 0xc8, 0xc1, 0x54, 0x6e, 0x70,
	0x7d, 0x98, 0x66, 0xad, 0x7a, 0xaf, 0x15, 0xf2, 0x6e, 0xa2, 0xdf, 0x39, 0x27, 0xfd, 0x44, 0x60,
	0xf7, 0x34, 0xe5, 0x4c, 0xf3, 0xe7, 0xa3, 0xda, 0x4f, 0x87, 0xe9, 0x11, 0x7f, 0xdb, 0xe3, 0x4a,
	0xe3, 0x0e, 0x54, 0x62, 0xd6, 0xe5, 0xd6, 0xe7, 0x93, 0x2a, 0xd9, 0xaf, 0x44, 0xe3, 0x0d, 0x7c,
	0x08, 0x9b, 0x83, 0xe9, 0x4c, 0xbf, 0x54, 0x25, 0xfb, 0xff, 0x1d, 0x6d, 0xd5, 0x58, 0x22, 0x6a,
	0xf9, 0x73, 0xf3, 0x09, 0xb4, 0x0b, 0x37, 0xce, 0x85, 0xd2, 0xb9, 0x58, 0xb5, 0x18, 0x88, 0x00,
	0xd6, 0x13, 0xd6, 0xe6, 0x17, 0xe2, 0xbd, 0xab, 0x5d, 0x8e, 0x32, 0x1b, 0x11, 0x56, 0xcc, 0xda,
	0xf7, 0xec, 0xbe, 0x5d, 0xd3, 0x1f, 0x04, 0x76, 0x67, 0xd5, 0x53, 0x89, 0x8c, 0x15, 0xc7, 0xff,
	0xa1, 0xdc, 0x90, 0xbd, 0x58, 0xdb, 0x62, 0xe5, 0xc8, 0x19, 0xf8, 0x08, 0x30, 0x07, 0x5e, 0xf9,
	0xa5, 0xaa, 0x37, 0xa7, 0xdd, 0x82, 0x8c, 0x22, 0x50, 0xa6, 0xa2, 0xf9, 0x2a, 0x7f, 0xc5, 0x55,
	0xb4, 0x06, 0xee, 0x02, 0x68, 0xa9, 0x59, 0xe7, 0xd4, 0x82, 0x29, 0x5b, 0xd7, 0xc4, 0x0e, 0x7d,
	0x01, 0x37, 0x0b, 0x3b, 0x79, 0xc6, 0x53, 0x65, 0xd8, 0xb1, 0xd8, 0x00, 0x11, 0x56, 0x8c, 0x61,
	0x87, 0x57, 0x89, 0xec, 0x9a, 0xfe, 0x21, 0xb0, 0x99, 0x3b, 0x16, 0x37, 0xc0, 0xeb, 0x89, 0xe6,
	0xf0, 0x04, 0xb3, 0x2c, 0xca, 0x45, 0x1f, 0xd6, 0xfa, 0x0e, 0x80, 0x6d, 0xd1, 0x8b, 0x46, 0xa6,
	0xf9, 0x55, 0x5d, 0x16, 0x8b, 0x16, 0x57, 0xda, 0x36, 0x5a, 0x89, 0x32, 0xdb, 0xf8, 0x84, 0x3a,
	0x37, 0x03, 0x72, 0x9d, 0xae, 0x47, 0x99, 0x6d, 0xf0, 0x37, 0x2c, 0x4f, 0x9b, 0x27, 0xda, 0x5f,
	0x75, 0xf8, 0xb3, 0x0d, 0x3c, 0x81, 0x8d, 0x11, 0xfd, 0x33, 0x12, 0xae, 0x59, 0x12, 0x5e, 0xcb,
	0xfe, 0xca, 0xa4, 0x33, 0xca, 0x85, 0xd3, 0x6f, 0x04, 0x6e, 0x3f, 0xe6, 0x31, 0x4f, 0x8b, 0xb4,
	0x90, 0xcb, 0x5f, 0x68, 0xaa, 0xc3, 0x59, 0x95, 0xc6, 0xb3, 0x2a, 0x54, 0x8b, 0xb7, 0xa4, 0x5a,
	0x8e, 0xfe, 0x96, 0xc1, 0xcf, 0x05, 0x5e, 0xf0, 0xb4, 0x2f, 0x1a, 0x1c, 0xbf, 0x12, 0xd8, 0x9e,
	0xa1, 0x68, 0xdc, 0xb3, 0x35, 0xe6, 0xeb, 0x3d, 0x98, 0x01, 0x84, 0x3e, 0xf9, 0xf8, 0xf3, 0xf7,
	0xe7, 0xd2, 0x03, 0x7a, 0xcb, 0x5c, 0x40, 0x2a, 0xec, 0x1f, 0xd6, 0xb9, 0x66, 0x87, 0xe1, 0x87,
	0xac, 0xdf, 0xcb, 0x30, 0x7f, 0x7f, 0xa9, 0xbb, 0x79, 0xfc, 0xf8, 0x85, 0xc0, 0x56, 0xb1, 0xfc,
	0x90, 0xda, 0xe2, 0x73, 0xef, 0x82, 0x60, 0x6f, 0x6e, 0x8c, 0xd3, 0x2f, 0x3d, 0xb6, 0x68, 0x43,
	0x5c, 0x0e, 0x2d, 0x7e, 0x27, 0x33, 0x6e, 0xa2, 0x91, 0x9e, 0xf0, 0x60, 0x76, 0xf5, 0x29, 0xcd,
	0x05, 0x57, 0x43, 0x27, 0xf9, 0x33, 0x8e, 0x1c, 0xc2, 0x3d, 0xb3, 0x70, 0xef, 0xe3, 0xbd, 0xa5,
	0xe0, 0x3a, 0xff, 0x65, 0xd8, 0x1f, 0x81, 0xfb, 0x45, 0xe0, 0x60, 0x61, 0x12, 0xe3, 0xb1, 0xc5,
	0xb7, 0x2c, 0xe9, 0x83, 0x62, 0x49, 0xd1, 0xd7, 0xb6, 0x85, 0x97, 0xf4, 0x6c, 0xc9, 0x16, 0x7a,
	0xa2, 0xe9, 0x3c, 0x57, 0x1e, 0xaf, 0x02, 0xde, 0xd4, 0x57, 0xed, 0x9b, 0x75, 0xe7, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0xd3, 0xec, 0x31, 0x28, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkspaceTemplateServiceClient is the client API for WorkspaceTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkspaceTemplateServiceClient interface {
	// Creates a WorkspaceTemplate
	CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
	ListWorkspaceTemplates(ctx context.Context, in *ListWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesResponse, error)
	ListWorkspaceTemplateVersions(ctx context.Context, in *ListWorkspaceTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error)
	// Get the generated WorkflowTemplate for a WorkspaceTemplate
	GenerateWorkspaceTemplateWorkflowTemplate(ctx context.Context, in *GenerateWorkspaceTemplateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
}

type workspaceTemplateServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkspaceTemplateServiceClient(cc *grpc.ClientConn) WorkspaceTemplateServiceClient {
	return &workspaceTemplateServiceClient{cc}
}

func (c *workspaceTemplateServiceClient) CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/CreateWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ListWorkspaceTemplates(ctx context.Context, in *ListWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesResponse, error) {
	out := new(ListWorkspaceTemplatesResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ListWorkspaceTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ListWorkspaceTemplateVersions(ctx context.Context, in *ListWorkspaceTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error) {
	out := new(ListWorkflowTemplateVersionsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ListWorkspaceTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) GenerateWorkspaceTemplateWorkflowTemplate(ctx context.Context, in *GenerateWorkspaceTemplateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/GenerateWorkspaceTemplateWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceTemplateServiceServer is the server API for WorkspaceTemplateService service.
type WorkspaceTemplateServiceServer interface {
	// Creates a WorkspaceTemplate
	CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
	ListWorkspaceTemplates(context.Context, *ListWorkspaceTemplatesRequest) (*ListWorkspaceTemplatesResponse, error)
	ListWorkspaceTemplateVersions(context.Context, *ListWorkspaceTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error)
	// Get the generated WorkflowTemplate for a WorkspaceTemplate
	GenerateWorkspaceTemplateWorkflowTemplate(context.Context, *GenerateWorkspaceTemplateWorkflowTemplateRequest) (*WorkflowTemplate, error)
}

// UnimplementedWorkspaceTemplateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkspaceTemplateServiceServer struct {
}

func (*UnimplementedWorkspaceTemplateServiceServer) CreateWorkspaceTemplate(ctx context.Context, req *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspaceTemplate not implemented")
}
func (*UnimplementedWorkspaceTemplateServiceServer) ListWorkspaceTemplates(ctx context.Context, req *ListWorkspaceTemplatesRequest) (*ListWorkspaceTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceTemplates not implemented")
}
func (*UnimplementedWorkspaceTemplateServiceServer) ListWorkspaceTemplateVersions(ctx context.Context, req *ListWorkspaceTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceTemplateVersions not implemented")
}
func (*UnimplementedWorkspaceTemplateServiceServer) GenerateWorkspaceTemplateWorkflowTemplate(ctx context.Context, req *GenerateWorkspaceTemplateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWorkspaceTemplateWorkflowTemplate not implemented")
}

func RegisterWorkspaceTemplateServiceServer(s *grpc.Server, srv WorkspaceTemplateServiceServer) {
	s.RegisterService(&_WorkspaceTemplateService_serviceDesc, srv)
}

func _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/CreateWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, req.(*CreateWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ListWorkspaceTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ListWorkspaceTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplates(ctx, req.(*ListWorkspaceTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ListWorkspaceTemplateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceTemplateVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ListWorkspaceTemplateVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplateVersions(ctx, req.(*ListWorkspaceTemplateVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_GenerateWorkspaceTemplateWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWorkspaceTemplateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).GenerateWorkspaceTemplateWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/GenerateWorkspaceTemplateWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).GenerateWorkspaceTemplateWorkflowTemplate(ctx, req.(*GenerateWorkspaceTemplateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceTemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkspaceTemplateService",
	HandlerType: (*WorkspaceTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler,
		},
		{
			MethodName: "ListWorkspaceTemplates",
			Handler:    _WorkspaceTemplateService_ListWorkspaceTemplates_Handler,
		},
		{
			MethodName: "ListWorkspaceTemplateVersions",
			Handler:    _WorkspaceTemplateService_ListWorkspaceTemplateVersions_Handler,
		},
		{
			MethodName: "GenerateWorkspaceTemplateWorkflowTemplate",
			Handler:    _WorkspaceTemplateService_GenerateWorkspaceTemplateWorkflowTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace_template.proto",
}

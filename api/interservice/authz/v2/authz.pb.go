// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authz/v2/authz.proto

package v2 // import "github.com/chef/automate/api/interservice/authz/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TODO (sr): consider dropping this validation:
// a) it takes time
// b) bad input will with certainty lead to an "unauthorized" response
type IsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedReq) Reset()         { *m = IsAuthorizedReq{} }
func (m *IsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReq) ProtoMessage()    {}
func (*IsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{0}
}
func (m *IsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReq.Unmarshal(m, b)
}
func (m *IsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReq.Marshal(b, m, deterministic)
}
func (dst *IsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReq.Merge(dst, src)
}
func (m *IsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReq.Size(m)
}
func (m *IsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReq proto.InternalMessageInfo

func (m *IsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *IsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *IsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type IsAuthorizedResp struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty" toml:"authorized,omitempty" mapstructure:"authorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedResp) Reset()         { *m = IsAuthorizedResp{} }
func (m *IsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedResp) ProtoMessage()    {}
func (*IsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{1}
}
func (m *IsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedResp.Unmarshal(m, b)
}
func (m *IsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedResp.Marshal(b, m, deterministic)
}
func (dst *IsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedResp.Merge(dst, src)
}
func (m *IsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedResp.Size(m)
}
func (m *IsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedResp proto.InternalMessageInfo

func (m *IsAuthorizedResp) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

type ProjectsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	ProjectsFilter       []string `protobuf:"bytes,4,rep,name=projects_filter,json=projectsFilter,proto3" json:"projects_filter,omitempty" toml:"projects_filter,omitempty" mapstructure:"projects_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedReq) Reset()         { *m = ProjectsAuthorizedReq{} }
func (m *ProjectsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedReq) ProtoMessage()    {}
func (*ProjectsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{2}
}
func (m *ProjectsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedReq.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedReq.Marshal(b, m, deterministic)
}
func (dst *ProjectsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedReq.Merge(dst, src)
}
func (m *ProjectsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedReq.Size(m)
}
func (m *ProjectsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedReq proto.InternalMessageInfo

func (m *ProjectsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ProjectsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetProjectsFilter() []string {
	if m != nil {
		return m.ProjectsFilter
	}
	return nil
}

type ProjectsAuthorizedResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedResp) Reset()         { *m = ProjectsAuthorizedResp{} }
func (m *ProjectsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedResp) ProtoMessage()    {}
func (*ProjectsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{3}
}
func (m *ProjectsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedResp.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedResp.Marshal(b, m, deterministic)
}
func (dst *ProjectsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedResp.Merge(dst, src)
}
func (m *ProjectsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedResp.Size(m)
}
func (m *ProjectsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedResp proto.InternalMessageInfo

func (m *ProjectsAuthorizedResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

// FilterAuthorizedProjects uses this SAME request
type FilterAuthorizedPairsReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsReq) Reset()         { *m = FilterAuthorizedPairsReq{} }
func (m *FilterAuthorizedPairsReq) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsReq) ProtoMessage()    {}
func (*FilterAuthorizedPairsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{4}
}
func (m *FilterAuthorizedPairsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Marshal(b, m, deterministic)
}
func (dst *FilterAuthorizedPairsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsReq.Merge(dst, src)
}
func (m *FilterAuthorizedPairsReq) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Size(m)
}
func (m *FilterAuthorizedPairsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsReq proto.InternalMessageInfo

func (m *FilterAuthorizedPairsReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *FilterAuthorizedPairsReq) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type FilterAuthorizedPairsResp struct {
	Pairs                []*Pair  `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsResp) Reset()         { *m = FilterAuthorizedPairsResp{} }
func (m *FilterAuthorizedPairsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsResp) ProtoMessage()    {}
func (*FilterAuthorizedPairsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{5}
}
func (m *FilterAuthorizedPairsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Marshal(b, m, deterministic)
}
func (dst *FilterAuthorizedPairsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsResp.Merge(dst, src)
}
func (m *FilterAuthorizedPairsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Size(m)
}
func (m *FilterAuthorizedPairsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsResp proto.InternalMessageInfo

func (m *FilterAuthorizedPairsResp) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type FilterAuthorizedProjectsResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedProjectsResp) Reset()         { *m = FilterAuthorizedProjectsResp{} }
func (m *FilterAuthorizedProjectsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedProjectsResp) ProtoMessage()    {}
func (*FilterAuthorizedProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{6}
}
func (m *FilterAuthorizedProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Marshal(b, m, deterministic)
}
func (dst *FilterAuthorizedProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedProjectsResp.Merge(dst, src)
}
func (m *FilterAuthorizedProjectsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Size(m)
}
func (m *FilterAuthorizedProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedProjectsResp proto.InternalMessageInfo

func (m *FilterAuthorizedProjectsResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Pair struct {
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_authz_f1d9420b5d8a6335, []int{7}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Pair) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*IsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.IsAuthorizedReq")
	proto.RegisterType((*IsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.IsAuthorizedResp")
	proto.RegisterType((*ProjectsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedReq")
	proto.RegisterType((*ProjectsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedResp")
	proto.RegisterType((*FilterAuthorizedPairsReq)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsReq")
	proto.RegisterType((*FilterAuthorizedPairsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsResp")
	proto.RegisterType((*FilterAuthorizedProjectsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedProjectsResp")
	proto.RegisterType((*Pair)(nil), "chef.automate.domain.authz.v2.Pair")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error) {
	out := new(IsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error) {
	out := new(FilterAuthorizedPairsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error) {
	out := new(FilterAuthorizedProjectsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error) {
	out := new(ProjectsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	IsAuthorized(context.Context, *IsAuthorizedReq) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(context.Context, *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsAuthorized(ctx, req.(*IsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedPairsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, req.(*FilterAuthorizedPairsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedPairsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, req.(*FilterAuthorizedPairsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_ProjectsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, req.(*ProjectsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authz.v2.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Authorization_IsAuthorized_Handler,
		},
		{
			MethodName: "FilterAuthorizedPairs",
			Handler:    _Authorization_FilterAuthorizedPairs_Handler,
		},
		{
			MethodName: "FilterAuthorizedProjects",
			Handler:    _Authorization_FilterAuthorizedProjects_Handler,
		},
		{
			MethodName: "ProjectsAuthorized",
			Handler:    _Authorization_ProjectsAuthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authz/v2/authz.proto",
}

func init() {
	proto.RegisterFile("api/interservice/authz/v2/authz.proto", fileDescriptor_authz_f1d9420b5d8a6335)
}

var fileDescriptor_authz_f1d9420b5d8a6335 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xee, 0x24, 0xbd, 0x55, 0x3a, 0xf7, 0x5e, 0x0a, 0x53, 0x85, 0xba, 0xe9, 0x8f, 0xc2, 0x34,
	0x20, 0x27, 0x10, 0x1b, 0x86, 0xf2, 0x53, 0x23, 0xa8, 0xda, 0x05, 0x55, 0x05, 0x12, 0x95, 0x25,
	0xa8, 0x68, 0x95, 0xc0, 0xd4, 0x99, 0x12, 0x83, 0x13, 0xbb, 0x33, 0xe3, 0x2c, 0xd2, 0x74, 0x81,
	0xd8, 0x22, 0x21, 0xe5, 0x01, 0x78, 0x02, 0x9e, 0x80, 0x15, 0x6b, 0xde, 0x84, 0x2d, 0x4f, 0x80,
	0xec, 0xd8, 0x69, 0xda, 0x24, 0x44, 0xed, 0x8a, 0x05, 0xbb, 0xf1, 0xf1, 0x7c, 0xdf, 0x39, 0xe7,
	0xfb, 0xce, 0xcc, 0xc0, 0xab, 0xd4, 0xb3, 0x75, 0xbb, 0x2e, 0x19, 0x17, 0x8c, 0x37, 0x6c, 0x8b,
	0xe9, 0xd4, 0x97, 0xd5, 0xa6, 0xde, 0x20, 0x9d, 0x85, 0xe6, 0x71, 0x57, 0xba, 0x68, 0xc1, 0xaa,
	0xb2, 0x7d, 0x8d, 0xfa, 0xd2, 0xad, 0x51, 0xc9, 0xb4, 0x8a, 0x5b, 0xa3, 0x76, 0x5d, 0xeb, 0xec,
	0x68, 0x90, 0xcc, 0x4c, 0x83, 0x3a, 0x76, 0x85, 0x4a, 0xa6, 0xc7, 0x8b, 0x0e, 0x0e, 0x7f, 0x4e,
	0xc0, 0xa9, 0x4d, 0xb1, 0xe6, 0xcb, 0xaa, 0xcb, 0xed, 0x26, 0xab, 0x98, 0xec, 0x00, 0xbd, 0x07,
	0x30, 0x25, 0xfc, 0xbd, 0xb7, 0xcc, 0x92, 0x42, 0x01, 0xd9, 0xa4, 0x3a, 0xb9, 0xce, 0xbe, 0xfe,
	0xf8, 0x96, 0x7c, 0xdd, 0x06, 0xa5, 0x14, 0xc0, 0x2f, 0xf9, 0x36, 0x79, 0x5e, 0x56, 0x57, 0x0d,
	0xc9, 0x68, 0xad, 0xe5, 0x0b, 0xc6, 0xf3, 0x86, 0xba, 0x6a, 0x38, 0xae, 0x45, 0x9d, 0x96, 0x53,
	0xa1, 0x5e, 0x4b, 0xd0, 0x9a, 0x93, 0x37, 0x76, 0xcb, 0x46, 0xa1, 0x74, 0x3d, 0xd7, 0x2a, 0x4b,
	0xf7, 0x1d, 0xab, 0xf7, 0x7c, 0x3a, 0xc2, 0x88, 0x7a, 0x89, 0x82, 0xf1, 0x3f, 0xb3, 0x9b, 0x16,
	0x3d, 0x82, 0x29, 0xce, 0x84, 0xeb, 0x73, 0x8b, 0x29, 0x89, 0x2c, 0x50, 0x27, 0xd7, 0x71, 0x50,
	0xc2, 0x02, 0x9f, 0x23, 0xb3, 0xe5, 0x5d, 0x5a, 0x6c, 0x96, 0x42, 0x4c, 0x41, 0x5d, 0x35, 0x22,
	0x74, 0xbe, 0x90, 0x33, 0xbb, 0x18, 0xb4, 0x01, 0x27, 0xa8, 0x25, 0x6d, 0xb7, 0xae, 0x24, 0x43,
	0xb4, 0x1e, 0xa0, 0x0b, 0x5c, 0x25, 0xd7, 0x22, 0x34, 0x2d, 0x36, 0xd7, 0x8a, 0x3b, 0x11, 0xc1,
	0x89, 0x48, 0xfe, 0x90, 0x1c, 0xe5, 0xcc, 0x08, 0x8e, 0x09, 0xbc, 0x78, 0x52, 0x1f, 0xe1, 0xa1,
	0x45, 0x08, 0x69, 0x37, 0xa2, 0x80, 0x2c, 0x50, 0x53, 0x66, 0x4f, 0x04, 0xff, 0x4c, 0xc0, 0xf4,
	0x16, 0x77, 0xc3, 0x4e, 0xfe, 0x4a, 0x3b, 0x4c, 0x5a, 0xf4, 0x14, 0x4e, 0x79, 0x91, 0x4a, 0xaf,
	0xf6, 0x6d, 0x47, 0x32, 0xae, 0x8c, 0x87, 0x92, 0x2c, 0x05, 0x8c, 0x8b, 0x6d, 0x30, 0xa7, 0x00,
	0x3c, 0xc3, 0xd3, 0x64, 0x3a, 0x24, 0xbe, 0x59, 0x5c, 0x51, 0xf3, 0xc5, 0xd2, 0xe1, 0xad, 0x1b,
	0x77, 0x97, 0x8f, 0x72, 0xe6, 0x85, 0x18, 0xfb, 0x38, 0x84, 0xe2, 0x6d, 0x78, 0x79, 0x90, 0xe6,
	0xc2, 0x43, 0x0f, 0x61, 0x2a, 0xde, 0x1b, 0x69, 0x7e, 0x25, 0x48, 0x30, 0xdf, 0x06, 0xb3, 0x0a,
	0xc0, 0x69, 0x3e, 0x4d, 0x2e, 0xc5, 0x09, 0x8e, 0xe9, 0xbb, 0x10, 0xfc, 0x1d, 0x40, 0xa5, 0x93,
	0xe3, 0x98, 0x77, 0x8b, 0xda, 0x5c, 0x04, 0x86, 0x8a, 0x3e, 0x3f, 0xb7, 0x03, 0x6e, 0xb3, 0x0d,
	0x9e, 0xa5, 0x00, 0x7e, 0xc2, 0x37, 0xc9, 0x46, 0xe0, 0xe7, 0x48, 0x4b, 0x5b, 0xa1, 0x93, 0xad,
	0x7e, 0x03, 0xf3, 0x03, 0x1c, 0x5c, 0x81, 0xff, 0x78, 0x41, 0x01, 0x4a, 0x22, 0x9b, 0x54, 0xff,
	0x25, 0x4b, 0xda, 0x6f, 0x0f, 0xbf, 0x16, 0x14, 0x6b, 0x76, 0x10, 0xf8, 0x05, 0x9c, 0x1d, 0xd2,
	0x8b, 0xf0, 0x8e, 0x79, 0xc1, 0x99, 0x79, 0x0d, 0x38, 0xdf, 0xc7, 0x1b, 0x09, 0x18, 0x52, 0x67,
	0x4e, 0x7b, 0xd0, 0x23, 0xf0, 0x27, 0x00, 0xc7, 0x03, 0xae, 0x3f, 0x66, 0x32, 0xc9, 0x97, 0x71,
	0xf8, 0x7f, 0xdc, 0x08, 0x0d, 0x67, 0xf5, 0x00, 0xfe, 0xd7, 0x7b, 0x0d, 0x20, 0x6d, 0x84, 0x36,
	0xa7, 0xee, 0xd4, 0x8c, 0x7e, 0xa6, 0xfd, 0xc2, 0xc3, 0x63, 0xe8, 0x23, 0x80, 0xe9, 0x81, 0x5e,
	0xa1, 0x7b, 0x23, 0xc8, 0x86, 0x4d, 0x6b, 0xe6, 0xfe, 0xf9, 0x80, 0x61, 0x39, 0xed, 0x41, 0xc7,
	0x20, 0xb2, 0xf0, 0xfc, 0x15, 0x3d, 0x38, 0x2b, 0xb0, 0x67, 0xa8, 0xf0, 0x18, 0xfa, 0x00, 0x20,
	0xea, 0x3f, 0xf5, 0x68, 0x79, 0xd4, 0xe4, 0x0e, 0xba, 0x9c, 0x33, 0x77, 0xce, 0x81, 0x0a, 0xaa,
	0x58, 0x5f, 0xde, 0x21, 0x6f, 0x6c, 0x59, 0xf5, 0xf7, 0x34, 0xcb, 0xad, 0xe9, 0x01, 0x89, 0x1e,
	0x93, 0xe8, 0x43, 0x9f, 0xef, 0xbd, 0x89, 0xf0, 0x05, 0xbe, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x06, 0x4f, 0xc8, 0xe2, 0x07, 0x00, 0x00,
}

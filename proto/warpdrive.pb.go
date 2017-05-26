// Code generated by protoc-gen-go.
// source: proto/warpdrive.proto
// DO NOT EDIT!

/*
Package warpdrive is a generated protocol buffer package.

It is generated from these files:
	proto/warpdrive.proto

It has these top-level messages:
	Empty
	App
	Release
	Upgrade
	Chunck
*/
package warpdrive

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Platform int32

const (
	Platform_IOS     Platform = 0
	Platform_ANDROID Platform = 1
)

var Platform_name = map[int32]string{
	0: "IOS",
	1: "ANDROID",
}
var Platform_value = map[string]int32{
	"IOS":     0,
	"ANDROID": 1,
}

func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type App struct {
	// @inject_tag: storm:"id"
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty" storm:"id"`
	// @inject_tag: storm:"unique"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" storm:"unique"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *App) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Release struct {
	// @inject_tag: storm:"id"
	Id    uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty" storm:"id"`
	AppId uint64 `protobuf:"varint,2,opt,name=appId" json:"appId,omitempty"`
	// this is just for label. it's not unique
	// becuase you might want to rollback
	Version  string   `protobuf:"bytes,3,opt,name=version" json:"version,omitempty" `
	Notes    string   `protobuf:"bytes,4,opt,name=notes" json:"notes,omitempty"`
	Platform Platform `protobuf:"varint,5,opt,name=platform,enum=warpdrive.Platform" json:"platform,omitempty"`
	// this is list of releases that can safely upgrade to this
	// version
	UpgradableReleases []uint64 `protobuf:"varint,6,rep,packed,name=upgradableReleases" json:"upgradableReleases,omitempty" `
	// this is used as what kind of release is. As an example `beta`
	RolloutAt string `protobuf:"bytes,7,opt,name=rolloutAt" json:"rolloutAt,omitempty" `
	// this is the hash value of bundle package
	Bundle string `protobuf:"bytes,8,opt,name=bundle" json:"bundle,omitempty" `
	// if the lock value is true, it means that this release can not be ultered or modified.
	// this is used to make sure the production doesn't download the unlock one.
	Lock      bool   `protobuf:"varint,9,opt,name=lock" json:"lock,omitempty" `
	CreatedAt string `protobuf:"bytes,10,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,11,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Release) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Release) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Release) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Release) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_IOS
}

func (m *Release) GetUpgradableReleases() []uint64 {
	if m != nil {
		return m.UpgradableReleases
	}
	return nil
}

func (m *Release) GetRolloutAt() string {
	if m != nil {
		return m.RolloutAt
	}
	return ""
}

func (m *Release) GetBundle() string {
	if m != nil {
		return m.Bundle
	}
	return ""
}

func (m *Release) GetLock() bool {
	if m != nil {
		return m.Lock
	}
	return false
}

func (m *Release) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Release) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Upgrade struct {
	ReleaseId uint64 `protobuf:"varint,1,opt,name=releaseId" json:"releaseId,omitempty"`
	// so if shouldBeLock is `true`, then we only retern locked releases.
	// if it's `false`, we are returning locked and unlocked releases
	ShouldBeLock bool `protobuf:"varint,2,opt,name=shouldBeLock" json:"shouldBeLock,omitempty" `
}

func (m *Upgrade) Reset()                    { *m = Upgrade{} }
func (m *Upgrade) String() string            { return proto.CompactTextString(m) }
func (*Upgrade) ProtoMessage()               {}
func (*Upgrade) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Upgrade) GetReleaseId() uint64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

func (m *Upgrade) GetShouldBeLock() bool {
	if m != nil {
		return m.ShouldBeLock
	}
	return false
}

type Chunck struct {
	// Types that are valid to be assigned to Value:
	//	*Chunck_Header_
	//	*Chunck_Body_
	Value isChunck_Value `protobuf_oneof:"value" `
}

func (m *Chunck) Reset()                    { *m = Chunck{} }
func (m *Chunck) String() string            { return proto.CompactTextString(m) }
func (*Chunck) ProtoMessage()               {}
func (*Chunck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isChunck_Value interface {
	isChunck_Value()
}

type Chunck_Header_ struct {
	Header *Chunck_Header `protobuf:"bytes,1,opt,name=header,oneof"`
}
type Chunck_Body_ struct {
	Body *Chunck_Body `protobuf:"bytes,2,opt,name=body,oneof"`
}

func (*Chunck_Header_) isChunck_Value() {}
func (*Chunck_Body_) isChunck_Value()   {}

func (m *Chunck) GetValue() isChunck_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Chunck) GetHeader() *Chunck_Header {
	if x, ok := m.GetValue().(*Chunck_Header_); ok {
		return x.Header
	}
	return nil
}

func (m *Chunck) GetBody() *Chunck_Body {
	if x, ok := m.GetValue().(*Chunck_Body_); ok {
		return x.Body
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Chunck) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Chunck_OneofMarshaler, _Chunck_OneofUnmarshaler, _Chunck_OneofSizer, []interface{}{
		(*Chunck_Header_)(nil),
		(*Chunck_Body_)(nil),
	}
}

func _Chunck_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Chunck)
	// value
	switch x := m.Value.(type) {
	case *Chunck_Header_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case *Chunck_Body_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Body); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Chunck.Value has unexpected type %T", x)
	}
	return nil
}

func _Chunck_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Chunck)
	switch tag {
	case 1: // value.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Chunck_Header)
		err := b.DecodeMessage(msg)
		m.Value = &Chunck_Header_{msg}
		return true, err
	case 2: // value.body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Chunck_Body)
		err := b.DecodeMessage(msg)
		m.Value = &Chunck_Body_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Chunck_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Chunck)
	// value
	switch x := m.Value.(type) {
	case *Chunck_Header_:
		s := proto.Size(x.Header)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Chunck_Body_:
		s := proto.Size(x.Body)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Chunck_Header struct {
	ReleaseId uint64 `protobuf:"varint,1,opt,name=releaseId" json:"releaseId,omitempty"`
	Total     int64  `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *Chunck_Header) Reset()                    { *m = Chunck_Header{} }
func (m *Chunck_Header) String() string            { return proto.CompactTextString(m) }
func (*Chunck_Header) ProtoMessage()               {}
func (*Chunck_Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Chunck_Header) GetReleaseId() uint64 {
	if m != nil {
		return m.ReleaseId
	}
	return 0
}

func (m *Chunck_Header) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Chunck_Body struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Chunck_Body) Reset()                    { *m = Chunck_Body{} }
func (m *Chunck_Body) String() string            { return proto.CompactTextString(m) }
func (*Chunck_Body) ProtoMessage()               {}
func (*Chunck_Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *Chunck_Body) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "warpdrive.Empty")
	proto.RegisterType((*App)(nil), "warpdrive.App")
	proto.RegisterType((*Release)(nil), "warpdrive.Release")
	proto.RegisterType((*Upgrade)(nil), "warpdrive.Upgrade")
	proto.RegisterType((*Chunck)(nil), "warpdrive.Chunck")
	proto.RegisterType((*Chunck_Header)(nil), "warpdrive.Chunck.Header")
	proto.RegisterType((*Chunck_Body)(nil), "warpdrive.Chunck.Body")
	proto.RegisterEnum("warpdrive.Platform", Platform_name, Platform_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Command service

type CommandClient interface {
	CreateApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*App, error)
	GetAllApps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Command_GetAllAppsClient, error)
	RemoveApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Empty, error)
	CreateRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error)
	GetRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error)
	UpdateRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error)
	UploadRelease(ctx context.Context, opts ...grpc.CallOption) (Command_UploadReleaseClient, error)
}

type commandClient struct {
	cc *grpc.ClientConn
}

func NewCommandClient(cc *grpc.ClientConn) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) CreateApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := grpc.Invoke(ctx, "/warpdrive.Command/CreateApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) GetAllApps(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Command_GetAllAppsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Command_serviceDesc.Streams[0], c.cc, "/warpdrive.Command/GetAllApps", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandGetAllAppsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Command_GetAllAppsClient interface {
	Recv() (*App, error)
	grpc.ClientStream
}

type commandGetAllAppsClient struct {
	grpc.ClientStream
}

func (x *commandGetAllAppsClient) Recv() (*App, error) {
	m := new(App)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commandClient) RemoveApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/warpdrive.Command/RemoveApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) CreateRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := grpc.Invoke(ctx, "/warpdrive.Command/CreateRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) GetRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := grpc.Invoke(ctx, "/warpdrive.Command/GetRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) UpdateRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := grpc.Invoke(ctx, "/warpdrive.Command/UpdateRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) UploadRelease(ctx context.Context, opts ...grpc.CallOption) (Command_UploadReleaseClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Command_serviceDesc.Streams[1], c.cc, "/warpdrive.Command/UploadRelease", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandUploadReleaseClient{stream}
	return x, nil
}

type Command_UploadReleaseClient interface {
	Send(*Chunck) error
	CloseAndRecv() (*Release, error)
	grpc.ClientStream
}

type commandUploadReleaseClient struct {
	grpc.ClientStream
}

func (x *commandUploadReleaseClient) Send(m *Chunck) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandUploadReleaseClient) CloseAndRecv() (*Release, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Release)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Command service

type CommandServer interface {
	CreateApp(context.Context, *App) (*App, error)
	GetAllApps(*Empty, Command_GetAllAppsServer) error
	RemoveApp(context.Context, *App) (*Empty, error)
	CreateRelease(context.Context, *Release) (*Release, error)
	GetRelease(context.Context, *Release) (*Release, error)
	UpdateRelease(context.Context, *Release) (*Release, error)
	UploadRelease(Command_UploadReleaseServer) error
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Command/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).CreateApp(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_GetAllApps_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommandServer).GetAllApps(m, &commandGetAllAppsServer{stream})
}

type Command_GetAllAppsServer interface {
	Send(*App) error
	grpc.ServerStream
}

type commandGetAllAppsServer struct {
	grpc.ServerStream
}

func (x *commandGetAllAppsServer) Send(m *App) error {
	return x.ServerStream.SendMsg(m)
}

func _Command_RemoveApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).RemoveApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Command/RemoveApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).RemoveApp(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_CreateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).CreateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Command/CreateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).CreateRelease(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_GetRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).GetRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Command/GetRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).GetRelease(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_UpdateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).UpdateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Command/UpdateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).UpdateRelease(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_UploadRelease_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandServer).UploadRelease(&commandUploadReleaseServer{stream})
}

type Command_UploadReleaseServer interface {
	SendAndClose(*Release) error
	Recv() (*Chunck, error)
	grpc.ServerStream
}

type commandUploadReleaseServer struct {
	grpc.ServerStream
}

func (x *commandUploadReleaseServer) SendAndClose(m *Release) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandUploadReleaseServer) Recv() (*Chunck, error) {
	m := new(Chunck)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warpdrive.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _Command_CreateApp_Handler,
		},
		{
			MethodName: "RemoveApp",
			Handler:    _Command_RemoveApp_Handler,
		},
		{
			MethodName: "CreateRelease",
			Handler:    _Command_CreateRelease_Handler,
		},
		{
			MethodName: "GetRelease",
			Handler:    _Command_GetRelease_Handler,
		},
		{
			MethodName: "UpdateRelease",
			Handler:    _Command_UpdateRelease_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllApps",
			Handler:       _Command_GetAllApps_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadRelease",
			Handler:       _Command_UploadRelease_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/warpdrive.proto",
}

// Client API for Query service

type QueryClient interface {
	GetUpgrade(ctx context.Context, in *Upgrade, opts ...grpc.CallOption) (*Release, error)
	DownloadRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (Query_DownloadReleaseClient, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetUpgrade(ctx context.Context, in *Upgrade, opts ...grpc.CallOption) (*Release, error) {
	out := new(Release)
	err := grpc.Invoke(ctx, "/warpdrive.Query/GetUpgrade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DownloadRelease(ctx context.Context, in *Release, opts ...grpc.CallOption) (Query_DownloadReleaseClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Query_serviceDesc.Streams[0], c.cc, "/warpdrive.Query/DownloadRelease", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryDownloadReleaseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_DownloadReleaseClient interface {
	Recv() (*Chunck, error)
	grpc.ClientStream
}

type queryDownloadReleaseClient struct {
	grpc.ClientStream
}

func (x *queryDownloadReleaseClient) Recv() (*Chunck, error) {
	m := new(Chunck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Query service

type QueryServer interface {
	GetUpgrade(context.Context, *Upgrade) (*Release, error)
	DownloadRelease(*Release, Query_DownloadReleaseServer) error
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Upgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warpdrive.Query/GetUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetUpgrade(ctx, req.(*Upgrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DownloadRelease_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Release)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).DownloadRelease(m, &queryDownloadReleaseServer{stream})
}

type Query_DownloadReleaseServer interface {
	Send(*Chunck) error
	grpc.ServerStream
}

type queryDownloadReleaseServer struct {
	grpc.ServerStream
}

func (x *queryDownloadReleaseServer) Send(m *Chunck) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warpdrive.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpgrade",
			Handler:    _Query_GetUpgrade_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadRelease",
			Handler:       _Query_DownloadRelease_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/warpdrive.proto",
}

func init() { proto.RegisterFile("proto/warpdrive.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0xda, 0xb4, 0x69, 0x6f, 0xb7, 0xfd, 0xf6, 0x33, 0x63, 0xb2, 0x2a, 0x1e, 0xaa, 0x3c,
	0x85, 0x7f, 0xdd, 0x14, 0x90, 0x86, 0x10, 0x2f, 0xd9, 0x86, 0x58, 0x05, 0x62, 0x60, 0xb4, 0x0f,
	0xe0, 0xce, 0x86, 0x55, 0x73, 0x62, 0x2b, 0x71, 0x3a, 0xf5, 0x0b, 0xf2, 0xc6, 0x07, 0xe2, 0x0d,
	0xe5, 0x26, 0xe9, 0xc2, 0x1a, 0x09, 0xed, 0xcd, 0xf7, 0x9c, 0x73, 0x8f, 0x4f, 0xae, 0xed, 0xc0,
	0x63, 0x93, 0x6a, 0xab, 0x0f, 0x6f, 0x79, 0x6a, 0x44, 0xba, 0x58, 0xca, 0x29, 0xd6, 0x64, 0xb8,
	0x06, 0x7c, 0x0f, 0x7a, 0xef, 0x63, 0x63, 0x57, 0xfe, 0x53, 0xe8, 0x46, 0xc6, 0x90, 0x5d, 0xe8,
	0x2c, 0x04, 0x75, 0x26, 0x4e, 0xe0, 0xb2, 0xce, 0x42, 0x10, 0x02, 0x6e, 0xc2, 0x63, 0x49, 0x3b,
	0x13, 0x27, 0x18, 0x32, 0x5c, 0xfb, 0x3f, 0x3b, 0xe0, 0x31, 0xa9, 0x24, 0xcf, 0xe4, 0x86, 0x7e,
	0x1f, 0x7a, 0xdc, 0x98, 0x99, 0xc0, 0x06, 0x97, 0x95, 0x05, 0xa1, 0xe0, 0x2d, 0x65, 0x9a, 0x2d,
	0x74, 0x42, 0xbb, 0x68, 0x54, 0x97, 0x85, 0x3e, 0xd1, 0x56, 0x66, 0xd4, 0x45, 0xbc, 0x2c, 0xc8,
	0x21, 0x0c, 0x8c, 0xe2, 0xf6, 0xbb, 0x4e, 0x63, 0xda, 0x9b, 0x38, 0xc1, 0x6e, 0xf8, 0x68, 0x7a,
	0xf7, 0x11, 0x5f, 0x2a, 0x8a, 0xad, 0x45, 0x64, 0x0a, 0x24, 0x37, 0x3f, 0x52, 0x2e, 0xf8, 0x5c,
	0xc9, 0x2a, 0x5b, 0x46, 0xfb, 0x93, 0x6e, 0xe0, 0xb2, 0x16, 0x86, 0x3c, 0x81, 0x61, 0xaa, 0x95,
	0xd2, 0xb9, 0x8d, 0x2c, 0xf5, 0x70, 0xeb, 0x3b, 0x80, 0x1c, 0x40, 0x7f, 0x9e, 0x27, 0x42, 0x49,
	0x3a, 0x40, 0xaa, 0xaa, 0x8a, 0x61, 0x28, 0x7d, 0x75, 0x43, 0x87, 0x13, 0x27, 0x18, 0x30, 0x5c,
	0x17, 0x4e, 0x57, 0xa9, 0xe4, 0x56, 0x8a, 0xc8, 0x52, 0x28, 0x9d, 0xd6, 0x40, 0xc1, 0xe6, 0x46,
	0x54, 0xec, 0xa8, 0x64, 0xd7, 0x80, 0xff, 0x11, 0xbc, 0x4b, 0xcc, 0x26, 0x31, 0x50, 0x19, 0x6e,
	0x56, 0x8f, 0xf3, 0x0e, 0x20, 0x3e, 0x6c, 0x67, 0xd7, 0x3a, 0x57, 0xe2, 0x44, 0x7e, 0x2a, 0x02,
	0x74, 0x30, 0xc0, 0x5f, 0x98, 0xff, 0xcb, 0x81, 0xfe, 0xe9, 0x75, 0x9e, 0x5c, 0xdd, 0x90, 0x10,
	0xfa, 0xd7, 0x92, 0x0b, 0x99, 0xa2, 0xd3, 0x28, 0xa4, 0x8d, 0xe1, 0x95, 0x92, 0xe9, 0x39, 0xf2,
	0xe7, 0x5b, 0xac, 0x52, 0x92, 0x17, 0xe0, 0xce, 0xb5, 0x58, 0xa1, 0xf5, 0x28, 0x3c, 0xd8, 0xec,
	0x38, 0xd1, 0x62, 0x75, 0xbe, 0xc5, 0x50, 0x35, 0x7e, 0x07, 0xfd, 0xd2, 0xe1, 0x1f, 0xc1, 0xf7,
	0xa1, 0x67, 0xb5, 0xe5, 0x0a, 0x6d, 0xbb, 0xac, 0x2c, 0xc6, 0x63, 0x70, 0x0b, 0xb7, 0x62, 0x9e,
	0x82, 0x5b, 0x8e, 0x6d, 0xdb, 0x0c, 0xd7, 0x27, 0x1e, 0xf4, 0x96, 0x5c, 0xe5, 0xf2, 0xd9, 0x04,
	0x06, 0xf5, 0x41, 0x13, 0x0f, 0xba, 0xb3, 0x8b, 0x6f, 0x7b, 0x5b, 0x64, 0x04, 0x5e, 0xf4, 0xf9,
	0x8c, 0x5d, 0xcc, 0xce, 0xf6, 0x9c, 0xf0, 0x77, 0x07, 0xbc, 0x53, 0x1d, 0xc7, 0x3c, 0x11, 0xe4,
	0x39, 0x0c, 0x4f, 0x71, 0xea, 0x78, 0x89, 0x1b, 0xe9, 0x23, 0x63, 0xc6, 0xf7, 0x6a, 0x72, 0x04,
	0xf0, 0x41, 0xda, 0x48, 0xa9, 0xc8, 0x98, 0x8c, 0xec, 0x35, 0x58, 0x7c, 0x0b, 0xf7, 0xf5, 0x47,
	0x0e, 0x79, 0x09, 0x43, 0x26, 0x63, 0xbd, 0x6c, 0xb5, 0xdf, 0x30, 0x20, 0xc7, 0xb0, 0x53, 0xa6,
	0xa9, 0x9f, 0x09, 0x69, 0x48, 0x2a, 0x6c, 0xdc, 0x82, 0x91, 0xd7, 0x98, 0xec, 0xa1, 0x5d, 0xc7,
	0xb0, 0x73, 0x89, 0x97, 0xea, 0xa1, 0x8d, 0x6f, 0x8a, 0x46, 0xa5, 0xb9, 0xa8, 0x81, 0xff, 0x37,
	0xce, 0xbd, 0xad, 0x2f, 0x70, 0xc2, 0x15, 0xf4, 0xbe, 0xe6, 0x32, 0x5d, 0x55, 0x89, 0xeb, 0x6b,
	0xdc, 0x14, 0x57, 0x58, 0xeb, 0xc6, 0x6f, 0xe1, 0xbf, 0x33, 0x7d, 0x9b, 0x34, 0xb7, 0x6e, 0xcb,
	0xbc, 0x19, 0xe7, 0xc8, 0x99, 0xf7, 0xf1, 0x27, 0xf6, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x93, 0xce, 0x8d, 0x2a, 0xdd, 0x04, 0x00, 0x00,
}

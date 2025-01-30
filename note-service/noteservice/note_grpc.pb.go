// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: note.proto

package noteservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NoteService_GetAllNotes_FullMethodName   = "/noteservice.NoteService/GetAllNotes"
	NoteService_AddNote_FullMethodName       = "/noteservice.NoteService/AddNote"
	NoteService_GetNoteByText_FullMethodName = "/noteservice.NoteService/GetNoteByText"
	NoteService_GetNoteByDate_FullMethodName = "/noteservice.NoteService/GetNoteByDate"
	NoteService_GetNoteByTag_FullMethodName  = "/noteservice.NoteService/GetNoteByTag"
)

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис заметок
type NoteServiceClient interface {
	GetAllNotes(ctx context.Context, in *Request, opts ...grpc.CallOption) (*NoteList, error)
	AddNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error)
	GetNoteByText(ctx context.Context, in *NoteTextRequest, opts ...grpc.CallOption) (*Note, error)
	GetNoteByDate(ctx context.Context, in *NoteDateRequest, opts ...grpc.CallOption) (*Note, error)
	GetNoteByTag(ctx context.Context, in *NoteTagRequest, opts ...grpc.CallOption) (*NoteList, error)
}

type noteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteServiceClient(cc grpc.ClientConnInterface) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) GetAllNotes(ctx context.Context, in *Request, opts ...grpc.CallOption) (*NoteList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NoteList)
	err := c.cc.Invoke(ctx, NoteService_GetAllNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) AddNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*Note, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Note)
	err := c.cc.Invoke(ctx, NoteService_AddNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) GetNoteByText(ctx context.Context, in *NoteTextRequest, opts ...grpc.CallOption) (*Note, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Note)
	err := c.cc.Invoke(ctx, NoteService_GetNoteByText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) GetNoteByDate(ctx context.Context, in *NoteDateRequest, opts ...grpc.CallOption) (*Note, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Note)
	err := c.cc.Invoke(ctx, NoteService_GetNoteByDate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) GetNoteByTag(ctx context.Context, in *NoteTagRequest, opts ...grpc.CallOption) (*NoteList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NoteList)
	err := c.cc.Invoke(ctx, NoteService_GetNoteByTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServiceServer is the server API for NoteService service.
// All implementations must embed UnimplementedNoteServiceServer
// for forward compatibility.
//
// Сервис заметок
type NoteServiceServer interface {
	GetAllNotes(context.Context, *Request) (*NoteList, error)
	AddNote(context.Context, *Note) (*Note, error)
	GetNoteByText(context.Context, *NoteTextRequest) (*Note, error)
	GetNoteByDate(context.Context, *NoteDateRequest) (*Note, error)
	GetNoteByTag(context.Context, *NoteTagRequest) (*NoteList, error)
	mustEmbedUnimplementedNoteServiceServer()
}

// UnimplementedNoteServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNoteServiceServer struct{}

func (UnimplementedNoteServiceServer) GetAllNotes(context.Context, *Request) (*NoteList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotes not implemented")
}
func (UnimplementedNoteServiceServer) AddNote(context.Context, *Note) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNote not implemented")
}
func (UnimplementedNoteServiceServer) GetNoteByText(context.Context, *NoteTextRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByText not implemented")
}
func (UnimplementedNoteServiceServer) GetNoteByDate(context.Context, *NoteDateRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByDate not implemented")
}
func (UnimplementedNoteServiceServer) GetNoteByTag(context.Context, *NoteTagRequest) (*NoteList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByTag not implemented")
}
func (UnimplementedNoteServiceServer) mustEmbedUnimplementedNoteServiceServer() {}
func (UnimplementedNoteServiceServer) testEmbeddedByValue()                     {}

// UnsafeNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteServiceServer will
// result in compilation errors.
type UnsafeNoteServiceServer interface {
	mustEmbedUnimplementedNoteServiceServer()
}

func RegisterNoteServiceServer(s grpc.ServiceRegistrar, srv NoteServiceServer) {
	// If the following call pancis, it indicates UnimplementedNoteServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NoteService_ServiceDesc, srv)
}

func _NoteService_GetAllNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetAllNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetAllNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetAllNotes(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_AddNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Note)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).AddNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_AddNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).AddNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_GetNoteByText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNoteByText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetNoteByText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNoteByText(ctx, req.(*NoteTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_GetNoteByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNoteByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetNoteByDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNoteByDate(ctx, req.(*NoteDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_GetNoteByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNoteByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetNoteByTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNoteByTag(ctx, req.(*NoteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteService_ServiceDesc is the grpc.ServiceDesc for NoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noteservice.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllNotes",
			Handler:    _NoteService_GetAllNotes_Handler,
		},
		{
			MethodName: "AddNote",
			Handler:    _NoteService_AddNote_Handler,
		},
		{
			MethodName: "GetNoteByText",
			Handler:    _NoteService_GetNoteByText_Handler,
		},
		{
			MethodName: "GetNoteByDate",
			Handler:    _NoteService_GetNoteByDate_Handler,
		},
		{
			MethodName: "GetNoteByTag",
			Handler:    _NoteService_GetNoteByTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}

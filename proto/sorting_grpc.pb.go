// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sorting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SortingClient is the client API for Sorting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SortingClient interface {
	StringSort(ctx context.Context, in *StringInput, opts ...grpc.CallOption) (*StringResult, error)
	Int32Sort(ctx context.Context, in *Int32Input, opts ...grpc.CallOption) (*Int32Result, error)
	PersonSortByAge(ctx context.Context, in *PersonList, opts ...grpc.CallOption) (*PersonList, error)
	PersonSortByName(ctx context.Context, in *PersonList, opts ...grpc.CallOption) (*PersonList, error)
}

type sortingClient struct {
	cc grpc.ClientConnInterface
}

func NewSortingClient(cc grpc.ClientConnInterface) SortingClient {
	return &sortingClient{cc}
}

var sortingStringSortStreamDesc = &grpc.StreamDesc{
	StreamName: "StringSort",
}

func (c *sortingClient) StringSort(ctx context.Context, in *StringInput, opts ...grpc.CallOption) (*StringResult, error) {
	out := new(StringResult)
	err := c.cc.Invoke(ctx, "/sorting.Sorting/StringSort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sortingInt32SortStreamDesc = &grpc.StreamDesc{
	StreamName: "Int32Sort",
}

func (c *sortingClient) Int32Sort(ctx context.Context, in *Int32Input, opts ...grpc.CallOption) (*Int32Result, error) {
	out := new(Int32Result)
	err := c.cc.Invoke(ctx, "/sorting.Sorting/Int32Sort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sortingPersonSortByAgeStreamDesc = &grpc.StreamDesc{
	StreamName: "PersonSortByAge",
}

func (c *sortingClient) PersonSortByAge(ctx context.Context, in *PersonList, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, "/sorting.Sorting/PersonSortByAge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sortingPersonSortByNameStreamDesc = &grpc.StreamDesc{
	StreamName: "PersonSortByName",
}

func (c *sortingClient) PersonSortByName(ctx context.Context, in *PersonList, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, "/sorting.Sorting/PersonSortByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SortingService is the service API for Sorting service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSortingService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SortingService struct {
	StringSort       func(context.Context, *StringInput) (*StringResult, error)
	Int32Sort        func(context.Context, *Int32Input) (*Int32Result, error)
	PersonSortByAge  func(context.Context, *PersonList) (*PersonList, error)
	PersonSortByName func(context.Context, *PersonList) (*PersonList, error)
}

func (s *SortingService) stringSort(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.StringSort == nil {
		return nil, status.Errorf(codes.Unimplemented, "method StringSort not implemented")
	}
	in := new(StringInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.StringSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sorting.Sorting/StringSort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.StringSort(ctx, req.(*StringInput))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SortingService) int32Sort(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.Int32Sort == nil {
		return nil, status.Errorf(codes.Unimplemented, "method Int32Sort not implemented")
	}
	in := new(Int32Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Int32Sort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sorting.Sorting/Int32Sort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Int32Sort(ctx, req.(*Int32Input))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SortingService) personSortByAge(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.PersonSortByAge == nil {
		return nil, status.Errorf(codes.Unimplemented, "method PersonSortByAge not implemented")
	}
	in := new(PersonList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.PersonSortByAge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sorting.Sorting/PersonSortByAge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PersonSortByAge(ctx, req.(*PersonList))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SortingService) personSortByName(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.PersonSortByName == nil {
		return nil, status.Errorf(codes.Unimplemented, "method PersonSortByName not implemented")
	}
	in := new(PersonList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.PersonSortByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sorting.Sorting/PersonSortByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.PersonSortByName(ctx, req.(*PersonList))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSortingService registers a service implementation with a gRPC server.
func RegisterSortingService(s grpc.ServiceRegistrar, srv *SortingService) {
	sd := grpc.ServiceDesc{
		ServiceName: "sorting.Sorting",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "StringSort",
				Handler:    srv.stringSort,
			},
			{
				MethodName: "Int32Sort",
				Handler:    srv.int32Sort,
			},
			{
				MethodName: "PersonSortByAge",
				Handler:    srv.personSortByAge,
			},
			{
				MethodName: "PersonSortByName",
				Handler:    srv.personSortByName,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proto/sorting.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSortingService creates a new SortingService containing the
// implemented methods of the Sorting service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSortingService(s interface{}) *SortingService {
	ns := &SortingService{}
	if h, ok := s.(interface {
		StringSort(context.Context, *StringInput) (*StringResult, error)
	}); ok {
		ns.StringSort = h.StringSort
	}
	if h, ok := s.(interface {
		Int32Sort(context.Context, *Int32Input) (*Int32Result, error)
	}); ok {
		ns.Int32Sort = h.Int32Sort
	}
	if h, ok := s.(interface {
		PersonSortByAge(context.Context, *PersonList) (*PersonList, error)
	}); ok {
		ns.PersonSortByAge = h.PersonSortByAge
	}
	if h, ok := s.(interface {
		PersonSortByName(context.Context, *PersonList) (*PersonList, error)
	}); ok {
		ns.PersonSortByName = h.PersonSortByName
	}
	return ns
}

// UnstableSortingService is the service API for Sorting service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSortingService interface {
	StringSort(context.Context, *StringInput) (*StringResult, error)
	Int32Sort(context.Context, *Int32Input) (*Int32Result, error)
	PersonSortByAge(context.Context, *PersonList) (*PersonList, error)
	PersonSortByName(context.Context, *PersonList) (*PersonList, error)
}
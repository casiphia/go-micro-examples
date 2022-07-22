// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/test/test.proto

package test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Test service

func NewTestEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Test service

type TestService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Test_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Test_PingPongService, error)
}

type testService struct {
	c    client.Client
	name string
}

func NewTestService(name string, c client.Client) TestService {
	return &testService{
		c:    c,
		name: name,
	}
}

func (c *testService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Test.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Test_StreamService, error) {
	req := c.c.NewRequest(c.name, "Test.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &testServiceStream{stream}, nil
}

type Test_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type testServiceStream struct {
	stream client.Stream
}

func (x *testServiceStream) Close() error {
	return x.stream.Close()
}

func (x *testServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *testServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testService) PingPong(ctx context.Context, opts ...client.CallOption) (Test_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Test.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &testServicePingPong{stream}, nil
}

type Test_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type testServicePingPong struct {
	stream client.Stream
}

func (x *testServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *testServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *testServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *testServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *testServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Test service

type TestHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Test_StreamStream) error
	PingPong(context.Context, Test_PingPongStream) error
}

func RegisterTestHandler(s server.Server, hdlr TestHandler, opts ...server.HandlerOption) error {
	type test interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Test struct {
		test
	}
	h := &testHandler{hdlr}
	return s.Handle(s.NewHandler(&Test{h}, opts...))
}

type testHandler struct {
	TestHandler
}

func (h *testHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.TestHandler.Call(ctx, in, out)
}

func (h *testHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TestHandler.Stream(ctx, m, &testStreamStream{stream})
}

type Test_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type testStreamStream struct {
	stream server.Stream
}

func (x *testStreamStream) Close() error {
	return x.stream.Close()
}

func (x *testStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *testStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *testHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.TestHandler.PingPong(ctx, &testPingPongStream{stream})
}

type Test_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type testPingPongStream struct {
	stream server.Stream
}

func (x *testPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *testPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *testPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *testPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

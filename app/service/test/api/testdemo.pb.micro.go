// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: testdemo.proto

package api

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

// Api Endpoints for TestDemo service

func NewTestDemoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TestDemo service

type TestDemoService interface {
	TestHello(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*TestResponse, error)
}

type testDemoService struct {
	c    client.Client
	name string
}

func NewTestDemoService(name string, c client.Client) TestDemoService {
	return &testDemoService{
		c:    c,
		name: name,
	}
}

func (c *testDemoService) TestHello(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*TestResponse, error) {
	req := c.c.NewRequest(c.name, "TestDemo.TestHello", in)
	out := new(TestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestDemo service

type TestDemoHandler interface {
	TestHello(context.Context, *TestRequest, *TestResponse) error
}

func RegisterTestDemoHandler(s server.Server, hdlr TestDemoHandler, opts ...server.HandlerOption) error {
	type testDemo interface {
		TestHello(ctx context.Context, in *TestRequest, out *TestResponse) error
	}
	type TestDemo struct {
		testDemo
	}
	h := &testDemoHandler{hdlr}
	return s.Handle(s.NewHandler(&TestDemo{h}, opts...))
}

type testDemoHandler struct {
	TestDemoHandler
}

func (h *testDemoHandler) TestHello(ctx context.Context, in *TestRequest, out *TestResponse) error {
	return h.TestDemoHandler.TestHello(ctx, in, out)
}

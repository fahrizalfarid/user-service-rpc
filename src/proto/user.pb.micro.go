// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: src/proto/user.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...client.CallOption) (*UserResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...client.CallOption) (User_FindService, error)
	FindWithArray(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*UserFoundArray, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	UpdateById(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UserResponse, error)
	DeleteById(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*Error, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetById(ctx context.Context, in *GetByIdRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetById", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Find(ctx context.Context, in *FindRequest, opts ...client.CallOption) (User_FindService, error) {
	req := c.c.NewRequest(c.name, "User.Find", &FindRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceFind{stream}, nil
}

type User_FindService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*UserFound, error)
}

type userServiceFind struct {
	stream client.Stream
}

func (x *userServiceFind) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *userServiceFind) Close() error {
	return x.stream.Close()
}

func (x *userServiceFind) Context() context.Context {
	return x.stream.Context()
}

func (x *userServiceFind) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceFind) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceFind) Recv() (*UserFound, error) {
	m := new(UserFound)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) FindWithArray(ctx context.Context, in *FindRequest, opts ...client.CallOption) (*UserFoundArray, error) {
	req := c.c.NewRequest(c.name, "User.FindWithArray", in)
	out := new(UserFoundArray)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateById(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.UpdateById", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteById(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*Error, error) {
	req := c.c.NewRequest(c.name, "User.DeleteById", in)
	out := new(Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	GetById(context.Context, *GetByIdRequest, *UserResponse) error
	Find(context.Context, *FindRequest, User_FindStream) error
	FindWithArray(context.Context, *FindRequest, *UserFoundArray) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	UpdateById(context.Context, *UpdateRequest, *UserResponse) error
	DeleteById(context.Context, *DeleteRequest, *Error) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		GetById(ctx context.Context, in *GetByIdRequest, out *UserResponse) error
		Find(ctx context.Context, stream server.Stream) error
		FindWithArray(ctx context.Context, in *FindRequest, out *UserFoundArray) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		UpdateById(ctx context.Context, in *UpdateRequest, out *UserResponse) error
		DeleteById(ctx context.Context, in *DeleteRequest, out *Error) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) GetById(ctx context.Context, in *GetByIdRequest, out *UserResponse) error {
	return h.UserHandler.GetById(ctx, in, out)
}

func (h *userHandler) Find(ctx context.Context, stream server.Stream) error {
	m := new(FindRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserHandler.Find(ctx, m, &userFindStream{stream})
}

type User_FindStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UserFound) error
}

type userFindStream struct {
	stream server.Stream
}

func (x *userFindStream) Close() error {
	return x.stream.Close()
}

func (x *userFindStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userFindStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userFindStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userFindStream) Send(m *UserFound) error {
	return x.stream.Send(m)
}

func (h *userHandler) FindWithArray(ctx context.Context, in *FindRequest, out *UserFoundArray) error {
	return h.UserHandler.FindWithArray(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) UpdateById(ctx context.Context, in *UpdateRequest, out *UserResponse) error {
	return h.UserHandler.UpdateById(ctx, in, out)
}

func (h *userHandler) DeleteById(ctx context.Context, in *DeleteRequest, out *Error) error {
	return h.UserHandler.DeleteById(ctx, in, out)
}

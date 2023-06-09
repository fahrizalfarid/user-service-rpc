// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: src/proto/uservalidator.proto

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

// Api Endpoints for UserValidator service

func NewUserValidatorEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserValidator service

type UserValidatorService interface {
	IsUsernameExists(ctx context.Context, in *UsernameRequest, opts ...client.CallOption) (*Found, error)
	IsEmailExists(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*Found, error)
	IsUserExists(ctx context.Context, in *EmailOrUsernameRequest, opts ...client.CallOption) (*Found, error)
}

type userValidatorService struct {
	c    client.Client
	name string
}

func NewUserValidatorService(name string, c client.Client) UserValidatorService {
	return &userValidatorService{
		c:    c,
		name: name,
	}
}

func (c *userValidatorService) IsUsernameExists(ctx context.Context, in *UsernameRequest, opts ...client.CallOption) (*Found, error) {
	req := c.c.NewRequest(c.name, "UserValidator.IsUsernameExists", in)
	out := new(Found)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userValidatorService) IsEmailExists(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*Found, error) {
	req := c.c.NewRequest(c.name, "UserValidator.IsEmailExists", in)
	out := new(Found)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userValidatorService) IsUserExists(ctx context.Context, in *EmailOrUsernameRequest, opts ...client.CallOption) (*Found, error) {
	req := c.c.NewRequest(c.name, "UserValidator.IsUserExists", in)
	out := new(Found)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserValidator service

type UserValidatorHandler interface {
	IsUsernameExists(context.Context, *UsernameRequest, *Found) error
	IsEmailExists(context.Context, *EmailRequest, *Found) error
	IsUserExists(context.Context, *EmailOrUsernameRequest, *Found) error
}

func RegisterUserValidatorHandler(s server.Server, hdlr UserValidatorHandler, opts ...server.HandlerOption) error {
	type userValidator interface {
		IsUsernameExists(ctx context.Context, in *UsernameRequest, out *Found) error
		IsEmailExists(ctx context.Context, in *EmailRequest, out *Found) error
		IsUserExists(ctx context.Context, in *EmailOrUsernameRequest, out *Found) error
	}
	type UserValidator struct {
		userValidator
	}
	h := &userValidatorHandler{hdlr}
	return s.Handle(s.NewHandler(&UserValidator{h}, opts...))
}

type userValidatorHandler struct {
	UserValidatorHandler
}

func (h *userValidatorHandler) IsUsernameExists(ctx context.Context, in *UsernameRequest, out *Found) error {
	return h.UserValidatorHandler.IsUsernameExists(ctx, in, out)
}

func (h *userValidatorHandler) IsEmailExists(ctx context.Context, in *EmailRequest, out *Found) error {
	return h.UserValidatorHandler.IsEmailExists(ctx, in, out)
}

func (h *userValidatorHandler) IsUserExists(ctx context.Context, in *EmailOrUsernameRequest, out *Found) error {
	return h.UserValidatorHandler.IsUserExists(ctx, in, out)
}

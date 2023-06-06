package delivery

import (
	"context"

	"github.com/fahrizalfarid/user-service-rpc/src/model"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
)

type UserValidator struct {
	UserValidator model.UserValidatorSvcUsecase
}

func (u *UserValidator) IsUsernameExists(ctx context.Context, req *pb.UsernameRequest, res *pb.Found) error {
	exists := u.UserValidator.IsUsernameExists(ctx, req.Username)
	res.Found = exists
	return nil
}

func (u *UserValidator) IsEmailExists(ctx context.Context, req *pb.EmailRequest, res *pb.Found) error {
	exists := u.UserValidator.IsEmailExists(ctx, req.Email)
	res.Found = exists
	return nil
}

func (u *UserValidator) IsUserExists(ctx context.Context, req *pb.EmailOrUsernameRequest, res *pb.Found) error {
	exists := u.UserValidator.IsUserExists(ctx, req.EmailOrUsername)
	res.Found = exists
	return nil
}

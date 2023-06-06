package delivery

import (
	"context"
	"log"
	"sync"

	"github.com/fahrizalfarid/user-service-rpc/src/constant"
	"github.com/fahrizalfarid/user-service-rpc/src/model"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
	"github.com/fahrizalfarid/user-service-rpc/src/request"
	"github.com/fahrizalfarid/user-service-rpc/utils"
	"go-micro.dev/v4/metadata"
)

type User struct {
	UserUsecase    model.UserSvcUsecase
	Authentication utils.Authentication
	Mu             sync.Mutex
}

func (u *User) getToken(ctx context.Context) (*utils.Token, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, constant.ErrAuth
	}

	token, exist := md.Get("token")
	if !exist {
		return nil, constant.ErrAuth
	}
	claims, err := u.Authentication.ParsingToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (u *User) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	hashedPassword, err := u.Authentication.EncryptPassword(req.Password)
	if err != nil {
		res.Id = 0
		return err
	}

	id, err := u.UserUsecase.CreateUser(ctx, &request.UserRequestService{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address:   req.Address,
		Phone:     req.Phone,
		Username:  req.Username,
		Password:  hashedPassword,
	})
	if err != nil {
		res.Id = id
		return err
	}

	res.Id = id
	return nil
}

func (u *User) GetById(ctx context.Context, req *pb.GetByIdRequest, res *pb.UserResponse) error {
	_, err := u.getToken(ctx)
	if err != nil {
		return err
	}

	user, err := u.UserUsecase.GetUserById(ctx, req.Id)
	if err != nil {
		return err
	}

	res.Id = user.Id
	res.Firstname = user.Firstname
	res.Lastname = user.Lastname
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt
	res.Phone = user.Phone
	res.Address = user.Address
	res.Username = user.Username

	return nil
}

func (u *User) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	userCredential, err := u.UserUsecase.GetUserCredentials(ctx, req.UsernameOrEmail)
	if err != nil {
		return err
	}

	res.Id = userCredential.Id
	res.Username = userCredential.Username
	res.Password = userCredential.Password
	return nil
}

func (u *User) Find(ctx context.Context, req *pb.FindRequest, res pb.User_FindStream) error {
	_, err := u.getToken(ctx)
	if err != nil {
		return err
	}

	defer res.Close()

	user, err := u.UserUsecase.GetUserByEmailOrUsername(ctx, req.Word)
	if err != nil {
		return err
	}

	u.Mu.Lock()
	for _, v := range user {
		rsp := &pb.UserFound{
			Id:       v.Id,
			Username: v.Username,
			Fullname: v.Fullname,
			Email:    v.Email,
		}

		log.Println("delivery", rsp.Fullname, rsp.Email, rsp.Id)
		if err := res.Send(rsp); err != nil {
			return err
		}
	}
	u.Mu.Unlock()

	return nil
}

func (u *User) FindWithArray(ctx context.Context, req *pb.FindRequest, res *pb.UserFoundArray) error {
	_, err := u.getToken(ctx)
	if err != nil {
		return err
	}

	var data []*pb.UserFound

	user, err := u.UserUsecase.GetUserByEmailOrUsername(ctx, req.Word)
	if err != nil {
		return err
	}

	for _, v := range user {
		data = append(data, &pb.UserFound{
			Id:       v.Id,
			Username: v.Username,
			Fullname: v.Fullname,
			Email:    v.Email,
		})
	}

	res.Users = data

	return nil
}

func (u *User) UpdateById(ctx context.Context, req *pb.UpdateRequest, res *pb.UserResponse) error {
	claims, err := u.getToken(ctx)
	if err != nil {
		return err
	}

	user, err := u.UserUsecase.UpdateById(ctx, &request.UserUpdateService{
		Id:        claims.Id,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Username:  req.Address,
		Password:  req.Password,
	})
	if err != nil {
		return err
	}

	res.Id = user.Id
	res.Firstname = user.Firstname
	res.Lastname = user.Lastname
	res.Email = user.Email
	res.Phone = user.Phone
	res.Address = user.Address
	res.Username = user.Username
	return nil
}

func (u *User) DeleteById(ctx context.Context, req *pb.DeleteRequest, res *pb.Error) error {
	claims, err := u.getToken(ctx)
	if err != nil {
		return err
	}

	err = u.UserUsecase.DeleteById(ctx, claims.Id)
	if err != nil {
		return err
	}

	res.Message = ""
	return nil
}

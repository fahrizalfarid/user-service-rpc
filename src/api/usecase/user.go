package usecase

import (
	"context"
	"io"

	"log"
	"strings"

	"time"

	"github.com/fahrizalfarid/user-service-rpc/src/constant"
	"github.com/fahrizalfarid/user-service-rpc/src/model"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
	"github.com/fahrizalfarid/user-service-rpc/src/request"
	"github.com/fahrizalfarid/user-service-rpc/src/response"
	"github.com/fahrizalfarid/user-service-rpc/utils"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"gorm.io/gorm"
)

type userUsecase struct {
	UserValidatorSrv pb.UserValidatorService
	UserSrv          pb.UserService
	Authentication   utils.Authentication
}

func NewUserUsecase(auth utils.Authentication,
	userSrv pb.UserService,
	userValSrv pb.UserValidatorService) model.UserUsecase {
	return &userUsecase{
		UserValidatorSrv: userValSrv,
		UserSrv:          userSrv,
		Authentication:   auth,
	}
}

func (u *userUsecase) existsValidator(ctx context.Context, username, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	userFound, err := u.UserValidatorSrv.IsUsernameExists(ctx, &pb.UsernameRequest{
		Username: username,
	}, client.WithRetries(3))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return constant.ErrSrvNotAvailable
		}
		return err
	}

	if userFound.Found {
		return constant.ErrUsernameExists
	}

	ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	emailFound, err := u.UserValidatorSrv.IsEmailExists(ctx, &pb.EmailRequest{
		Email: email,
	})

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return constant.ErrSrvNotAvailable
		}
		return err
	}

	if emailFound.Found {
		return constant.ErrEmailExists
	}
	return nil
}

func (u *userUsecase) CreateUser(ctx context.Context, data *request.UserRequest) (string, int64, error) {
	err := u.existsValidator(ctx, data.Username, data.Email)
	if err != nil {
		return "", 0, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	rsp, err := u.UserSrv.Create(ctx, &pb.CreateRequest{
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Email:     data.Email,
		CreatedAt: time.Now().Unix(),
		Phone:     data.Phone,
		Address:   data.Address,
		DeletedAt: int64(0),
		Username:  data.Username,
		Password:  data.Password,
	}, client.WithRetries(3))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return "", 0, constant.ErrSrvNotAvailable
		}
		return "", 0, err
	}

	stringToken, err := u.Authentication.GenerateToken(rsp.Id, data.Username)
	if err != nil {
		return "", 0, err
	}

	return stringToken, rsp.Id, nil
}

func (u *userUsecase) GetUserById(ctx context.Context, id int64, token string) (*response.UserProfileResponse, error) {
	ctx = metadata.NewContext(ctx, map[string]string{
		"token": token,
	})

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	user, err := u.UserSrv.GetById(ctx, &pb.GetByIdRequest{Id: id},
		client.WithRetries(3))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return nil, constant.ErrSrvNotAvailable
		}
		return nil, err
	}

	return &response.UserProfileResponse{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Phone:     user.Phone,
		Address:   user.Address,
		Username:  user.Username,
	}, nil
}

func (u *userUsecase) Login(ctx context.Context, username, password string) (*response.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	user, err := u.UserSrv.Login(ctx, &pb.LoginRequest{
		UsernameOrEmail: username,
	}, client.WithRetries(3))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return nil, constant.ErrSrvNotAvailable
		}
		return nil, err
	}

	err = u.Authentication.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	token, err := u.Authentication.GenerateToken(user.Id, user.Username)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		UserId:   user.Id,
		Username: user.Username,
		Token:    token,
	}, nil
}

func (u *userUsecase) Find(ctx context.Context, username, token string) ([]*response.UserFound, error) {
	var data []*response.UserFound

	ctxMd := metadata.NewContext(ctx, map[string]string{
		"token": token,
	})

	ctxT, cancel := context.WithTimeout(ctxMd, 2*time.Second)
	defer cancel()

	user, err := u.UserSrv.Find(ctxT, &pb.FindRequest{Word: username}, client.WithBackoff(client.DefaultBackoff))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return nil, constant.ErrSrvNotAvailable
		}
		return nil, err
	}
	defer user.Close()

	for {
		rsp, err := user.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}

		data = append(data, &response.UserFound{
			Id:       rsp.Id,
			Username: rsp.Username,
			Fullname: rsp.Fullname,
			Email:    rsp.Email,
		})
		log.Println("api", rsp.Fullname, rsp.Email, rsp.Id)
	}

	if len(data) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return data, nil
}

func (u *userUsecase) FindWithArray(ctx context.Context, username, token string) (*pb.UserFoundArray, error) {
	ctxMd := metadata.NewContext(ctx, map[string]string{
		"token": token,
	})

	ctxT, cancel := context.WithTimeout(ctxMd, 2*time.Second)
	defer cancel()

	user, err := u.UserSrv.FindWithArray(ctxT, &pb.FindRequest{Word: username}, client.WithBackoff(client.DefaultBackoff))

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return nil, constant.ErrSrvNotAvailable
		}
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) UpdateById(ctx context.Context, data *request.UserUpdateRequest, token string) (*response.UserProfileResponse, error) {
	ctx = metadata.NewContext(ctx, map[string]string{
		"token": token,
	})

	if data.Password != "" {
		hashedPassword, errHash := u.Authentication.EncryptPassword(data.Password)
		if errHash != nil {
			return nil, errHash
		}

		data.Password = hashedPassword
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.UserSrv.UpdateById(ctx, &pb.UpdateRequest{
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		Username:  data.Username,
		Password:  data.Password,
	})

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return nil, constant.ErrSrvNotAvailable
		}
		return nil, err
	}

	return &response.UserProfileResponse{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Phone:     user.Phone,
		Address:   user.Address,
		Username:  user.Username,
	}, nil
}

func (u *userUsecase) DeleteById(ctx context.Context, token string) error {
	ctxMd := metadata.NewContext(ctx, map[string]string{
		"token": token,
	})

	ctxT, cancel := context.WithTimeout(ctxMd, 2*time.Second)
	defer cancel()

	_, err := u.UserSrv.DeleteById(ctxT, &pb.DeleteRequest{})
	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return constant.ErrSrvNotAvailable
		}
		return err
	}

	return nil
}

func (u *userUsecase) IsUserExists(ctx context.Context, emailOrUsername string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	found, err := u.UserValidatorSrv.IsUserExists(ctx, &pb.EmailOrUsernameRequest{EmailOrUsername: emailOrUsername},
		client.WithBackoff(client.DefaultBackoff))
	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			return constant.ErrSrvNotAvailable
		}
		return err
	}

	if !found.Found {
		return gorm.ErrRecordNotFound
	}

	return nil
}

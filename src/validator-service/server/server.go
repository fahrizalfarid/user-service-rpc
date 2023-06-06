package server

import (
	"time"

	"github.com/fahrizalfarid/user-service-rpc/conf"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
	"github.com/fahrizalfarid/user-service-rpc/src/validator-service/delivery"
	"github.com/fahrizalfarid/user-service-rpc/src/validator-service/repository"
	"github.com/fahrizalfarid/user-service-rpc/src/validator-service/usecase"
	"go-micro.dev/v4"
)

func RunUserValidatorSrv() micro.Service {

	db, err := conf.DatabaseConn()
	if err != nil {
		panic(err)
	}

	validator := repository.NewValidatorRepo(db)
	userValidatorUsecase := usecase.NewUserValidatorSvcUsecase(validator)

	service := micro.NewService(
		micro.Name("user-validator"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// // define your own port
	// service := micro.NewService(
	// 	micro.Name("user"),
	// 	micro.RegisterTTL(time.Second*30),
	// 	micro.RegisterInterval(time.Second*15),
	// 	micro.Server(
	// 		server.NewServer(
	// 			server.Address(os.Getenv("VALIDATOR_SRV")),
	// 		),
	// 	),
	// )

	service.Init()

	userValidator := delivery.UserValidator{
		UserValidator: userValidatorUsecase,
	}

	pb.RegisterUserValidatorHandler(service.Server(), &userValidator)
	return service
}

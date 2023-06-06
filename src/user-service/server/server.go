package server

import (
	"sync"

	"github.com/fahrizalfarid/user-service-rpc/conf"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
	"github.com/fahrizalfarid/user-service-rpc/src/user-service/delivery"
	"github.com/fahrizalfarid/user-service-rpc/src/user-service/repository"
	"github.com/fahrizalfarid/user-service-rpc/src/user-service/usecase"
	"github.com/fahrizalfarid/user-service-rpc/utils"
	"go-micro.dev/v4"
)

func RunUserSrv() micro.Service {
	db, err := conf.DatabaseConn()
	if err != nil {
		panic(err)
	}

	auth := utils.NewAuthentication()

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecase.NewUserSvcUsecase(userRepo)

	service := micro.NewService(
		micro.Name("user"),
	)

	// // define your own port
	// service := micro.NewService(
	// 	micro.Name("user"),
	// 	micro.WrapHandler(authWrapper.Authentication),
	// 	micro.RegisterTTL(time.Second*30),
	// 	micro.RegisterInterval(time.Second*10),
	// 	micro.Server(
	// 		server.NewServer(
	// 			server.Address(os.Getenv("USER_SRV")),
	// 		),
	// 	),
	// )

	service.Init()

	user := delivery.User{
		UserUsecase:    userUsecase,
		Authentication: auth,
		Mu:             sync.Mutex{},
	}

	pb.RegisterUserHandler(service.Server(), &user)
	return service
}

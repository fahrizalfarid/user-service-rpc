package server

import (
	"github.com/fahrizalfarid/user-service-rpc/conf"
	v1 "github.com/fahrizalfarid/user-service-rpc/src/api/delivery/v1"
	"github.com/fahrizalfarid/user-service-rpc/src/api/middleware"
	"github.com/fahrizalfarid/user-service-rpc/src/api/usecase"
	pb "github.com/fahrizalfarid/user-service-rpc/src/proto"
	"github.com/fahrizalfarid/user-service-rpc/utils"
	"github.com/labstack/echo/v4"

	mid "github.com/labstack/echo/v4/middleware"
	"go-micro.dev/v4"
)

func RunServer() *echo.Echo {
	conf.LoadEnv("./env")

	e := echo.New()
	e.Use(mid.Recover())

	service := micro.NewService()
	service.Init()

	userValSrv := pb.NewUserValidatorService("user-validator", service.Client())
	userSrv := pb.NewUserService("user", service.Client())

	auth := utils.NewAuthentication()

	authorization := middleware.NewAuthorizationMiddleware(auth)
	userUsecase := usecase.NewUserUsecase(auth, userSrv, userValSrv)

	v1.NewUserDelivery(e, userUsecase, authorization)

	return e
}

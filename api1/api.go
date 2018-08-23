package api1

import (
	"github.com/gin-gonic/gin"
	"goServer/core"
	"github.com/itsjamie/gin-cors"
	"time"
	"goServer/models"
)

const BASE_ROUTE = "api/v1"
const ROOT_ROUTE = "/root"
const ACCOUNT_ROUTE = "/account"
const GALLERY_ROUTE = "/gallery"
const CONTACT_ROUTE = "/contact"
const SCHEDULED_ROUTE = "/scheduled"
const SERVICE_ROUTE = "/service"
const SUPPORT_ROUTE = "/support"
const AWS_ROUTE = "/aws"

type Routes struct {
	Aws *gin.RouterGroup
	Index *gin.RouterGroup
	Account *gin.RouterGroup
	Service *gin.RouterGroup
	Scheduled *gin.RouterGroup
	Support *gin.RouterGroup
	Contact *gin.RouterGroup
	Gallery *gin.RouterGroup
	NoRoute *gin.RouterGroup
}

type API struct {
	core *core.Core
	svr  *gin.Engine
	BaseRoutes *Routes
	Err *models.AppError

}
func (a *API) StartRouters(port string) {
	a.svr.Run(port)
}

func InitRouters(ref *core.Core) *API {

	api := &API {
		core : ref,
		svr : ref.Srv.CoreRef,
		BaseRoutes : &Routes{},
	}

	api.svr.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Cache-Control",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
	}))

	api.svr.NoRoute() // 404 에러페이지
	api.InitRoot()
	api.InitAws()
	api.InitAccount()
	api.InitContact()
	api.InitScheduled()
	api.InitService()
	api.InitSupport()
	api.InitGallery()
		/*
		index.GET("/protected",ProtectedEndPoint)
		index.GET("/test", CheckMiddleWare(TestEndPoint))
		index.GET("/ping",pingTest)
		*/
		return api;
}


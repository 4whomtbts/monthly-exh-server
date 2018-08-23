package core

import (

	"github.com/gin-gonic/gin"
	CONSTANT "goServer/constants"

	"goServer/store"
)

type Server struct {

	CoreRef *gin.Engine
	PORT string
	Store store.Store
	RouterGroup map[string]*gin.RouterGroup

}

func InitServer() *Server {
	return &Server{
		CoreRef : gin.Default(),
		PORT : CONSTANT.SERVER_PORT,
	}
}

func (Core *Core) Start() {


}

func pingTest(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "AWS 연결성공",
	})
}
/*
func (core *Core) CtxPassMiddleWare(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("CtxPassMiddleWare")
		c.Set("srv",core.Srv)

		fn(c)

	}
*/




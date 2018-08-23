package api1

import (
	"goServer/mlog"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"goServer/models"
	"goServer/core"
)

const (
	HEADER_FORWARDED          = "X-Forwarded-For"
	HEADER_AUTH
)

type Response struct {
	StatusCode int `json:"status_code"`
	Data  map[string]interface{} `json:"data"`
	Err string `json:"err"`
	Header http.Header `json:"headers"`
}
/*
func NewResponse(status int, pass bool, data map[string]interface{},header http.Header) *Response{
	//return &Response{ status, data,header, pass}
}
*/

func Zstring(key string, value string) zap.Field {
	return zap.String(key, value)
}

func mergePath(path string) string {
	return BASE_ROUTE + path
}

func writeSimpleRequestLog(ctx *gin.Context, sensitivity int,userId string) {

	path := (ctx.Request.URL).String()

	if userId == "" {
		userId = "CUSTOMER"
	}

	mlog.Info("",
		Zstring("PATH", path),
		Zstring("USER_IP",ctx.ClientIP()),
		Zstring("USER_ID",userId),
		zap.Int("SENS",sensitivity),
	)
}


func writeVerboseRequestLog(routerName string, path string, ip string, sensitivity int,userId string){

}

func NewWebApiResponse(c *gin.Context, status int , message string, data map[string]string ){
	c.JSON(status, gin.H{
		"msg": message,
		"data" : data,
	})
}

func NewSuccessJsonResponse(c *gin.Context, status int , message string){
	c.JSON(status, gin.H{
		"message": message,
	})
}
func NewFailJsonResponse(c *gin.Context,status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
		"status_code": status,
	})
}

func WebApiErrorLog(where, id  string, severity int){

	//redLog := color.New(color.FgRed).FprintfFunc()

	mlog.Error(id  + "occured in " +where ,zap.Int("SEVERITY",severity))
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
/*
func (api *API) ReturnCoreContext(c *gin.Context) (*core.Core, *models.AppError) {
	ctx, err := c.Get("context")
	if err != true  {
		fmt.Println("제어1")
		return nil , models.NewAppError("api.ReturnCoreContext","internal.server.error",nil,"",500)
	}

	return ctx0, nil
}
*/

func GetContext(c *gin.Context) (*core.Core, *models.AppError){
	ctx, ext := c.Get("context")
	if ext != true{
		err := models.NewAppError("GetContext","failed.to.GetContext",nil,"",500)
		WebApiErrorLog(err.Where,err.Id,7)
		return nil , err
	}
	ctx0 := ctx.(*core.Core)
	return ctx0, nil
}
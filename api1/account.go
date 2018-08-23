package api1

import (
	_ "github.com/beego/bee/logger/colors"
	"github.com/gin-gonic/gin"
	"goServer/models"
	"fmt"
)

func (api *API) InitAccount(){
	api.BaseRoutes.Account = api.svr.Group(mergePath(ACCOUNT_ROUTE))
	{
		api.BaseRoutes.Account.GET("/test",api.ApiAuthRequired(testPath))
		api.BaseRoutes.Account.GET("/user/:userId",api.ApiHandler(api.GET_getUserExtById))

		api.BaseRoutes.Account.POST("/login",api.ApiHandler(api.POST_login))
		api.BaseRoutes.Account.POST("/logout",api.ApiHandler(api.POST_logout))
		api.BaseRoutes.Account.POST("/index")
		api.BaseRoutes.Account.POST("/register",api.ApiHandler(POST_register))
		api.BaseRoutes.Account.POST("/forgetId")
		api.BaseRoutes.Account.POST("/forgetPwd")
	}
}


func (api *API) POST_login(c *gin.Context) {
	fmt.Println("로그인 포스트트");
	fmt.Println("엑세스토큰있나? ",c.Request.Header.Get("access-token"))
	ctx0 , err := GetContext(c)
	if err != nil {

	}

	fmt.Println("로그인 포스트트");

	user := models.UserFromJson(c.Request.Body)
	result, AppErr := ctx0.LoginById(user.Id, user.Password)


	if AppErr != nil {
		WebApiErrorLog(err.Where,err.Id,7)
		fmt.Println("앱에러아이디 ",AppErr.Id)
		NewFailJsonResponse(c,403,AppErr.Id)
		return;
	}

	c.SetCookie("access-token",result["access-token"].(string),150000,"/","http://127.0.0.1",false,true)
	c.SetCookie("refresh-token",result["access-token"].(string),150000,"/","http://127.0.0.1",false,true)
	NewSuccessJsonResponse(c,200,"")
	//c.Writer.Header().Add("Authorization",result0["access-token"].(string))
	//c.Writer.Header().Add("Refresh-Token",result0["refresh-token"].(string))
	return;
}

func (api *API) POST_logout(c *gin.Context) {



}


func (api *API) GET_getUserExtById(c *gin.Context){
	ctx0 , err := GetContext(c)
	if err != nil {

	}

	userId := c.Param("userId")
	fmt.Println("파라미터 : ",userId)
	c.SetCookie("cookie","coo",15000,"/","127.0.0.1",false,true)
		if ctx0.GetUserExistenceById(userId) == true {
			NewWebApiResponse(c,200,"api.",nil)
			fmt.Println("유저있음")
		}else{
			NewWebApiResponse(c,200,"user.is.not.found",nil)
			fmt.Println("유저없음")
		}
	}



func POST_register(c *gin.Context){
/*



	fmt.Println(ctx0)

	fmt.Println("서버에 온 데이터 ",userDataForm)


	}

		if rst := <-ctx0.Store.User().Save(&userDataForm); rst.Err != nil {
			WebApiErrorLog(fmt.Sprintf("api.account.post.login.save.err : %s",rst.Err),8)
			NewFailJsonResponse(c,200,"잠시 후에 다시 시도해주세요")
	}else{
		c.JSON(200, gin.H{
			"ok" : false,
			"msg" : "가입완료",
		})
	}
*/
}
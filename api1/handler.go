package api1

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"goServer/web"
	"goServer/core"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"goServer/models"
)

func testPath(c *gin.Context) {
	fmt.Println("테스트패스 요청됨")
	writeSimpleRequestLog(c, 1, "")
	//mlog.Info("테스트패스 요청",zap.String("behavior","test"))

	fmt.Println(c.ClientIP())
	fmt.Println(c.Get("context"))

	ctx, err := c.Get("context")
	if err != true {

	}
	ctx0 := ctx.(core.Server)

	fmt.Println(ctx0)

	c.JSON(200, gin.H{
		"message": "test",
		"url ":    c.Request.RequestURI,
	})
	return;
}

func Refinery(c *gin.Context) {

}

func (api *API) ApiHandler(next gin.HandlerFunc) gin.HandlerFunc {
	// 인증이 필요없는 핸들러
	return gin.HandlerFunc(func(c *gin.Context) {

		context := &web.Context{
			Core: api.core,
		}
		c.Set("context", context.Core)
		next(c)
	})
}

func (api *API) ApiAuthRequired(next gin.HandlerFunc) gin.HandlerFunc {
	// 인증 된 유저전용,
	return gin.HandlerFunc(func(c *gin.Context) {

		context := &web.Context{
			Core: api.core,
		}
		c.Set("context", *context.Core)
//		accessToken := c.Request.Header.Get("access-token")
//		refreshToken :=c.Request.Header.Get("access-token")
		tokenHeader := c.Request.Header.Get(HEADER_AUTH)
		if tokenHeader == ""{
			NewFailJsonResponse(c,403,"unauthorized")
			return;
		}

		api.decryptToken(c,tokenHeader)
		//TODO 인증 안 됬으면 return
		next(c)
	})
}

func (api *API) ApiGalleryAuthRequired(next gin.HandlerFunc) gin.HandlerFunc {
	// 갤러리 관리자 전용
	return gin.HandlerFunc(func(c *gin.Context) {
		err := api.AuditToken(c, core.ROLE_ADMIN)
		if err != nil {
			NewFailJsonResponse(c,err.StatusCode,err.Id)
			return;
		}
		next(c)
	})
}

func ApiConfidential() {

}

func (api *API) AuditToken(c *gin.Context, permission int) (*models.AppError){

	context := &web.Context{
		Core: api.core,
	}
	c.Set("context", *context.Core)

	accessToken := c.Request.Header.Get("access-token")
//	refreshToken :=c.Request.Header.Get("refresh-token")
	tokenHeader := c.Request.Header.Get(HEADER_AUTH)

	if tokenHeader == ""{
		return models.NewAppError("api.AuditToken",models.TOKEN_INVALID,nil,"",401)
	}

	claims , err := api.InspectToken(context.Core, accessToken)
	if err != nil {
		return err
	}
	StructClaims := core.ClaimsToStruct(claims)

	if StructClaims.Role == core.ROLE_ADMIN {
		return nil
	}

	switch permission {
	case core.ROLE_COMMON:
		return nil
	case core.ROLE_GALLERY:
		return api.IsUserPermittedToGallary(context.Core, StructClaims,c.Request.Header.Get("galleryPath"))
	case core.ROLE_MANAGER:
		return nil
	default:
		return models.NewAppError("api.AuditToken","permission.is.not.valid",nil,"",401)
	}

}

func (api *API) IsUserPermittedToGallary(context *core.Core, claims *core.Claims , GallaryPath string) (*models.AppError) {

	if  rst := <-context.Srv.Store.User().Get(claims.Id); rst.Err != nil{
		return rst.Err
	}else{

		if rst.Data.(models.User).OwnGallery.GalleryPath != GallaryPath {
			return models.NewAppError("IsUserPermittedToGallery","user.is.not.permitted",nil,"",403)
		}

		if rst.Err != nil {
			return rst.Err
		}
		return nil
	}
}
func (api *API) decryptToken(gc *gin.Context, token string) (jwt.MapClaims) {

	bearerToken := strings.Split(token, " ")
	fmt.Sprintf("베어러토큰 : %s", token)
	if len(bearerToken) == 1 {
		token, err := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("token.signing.is.invalid")
			}
			return []byte(core.GLOBAL_SALT), nil
		})

		if err != nil {
			WebApiErrorLog("api.decryptToken","token.is.invalid", 10)
			NewFailJsonResponse(gc, http.StatusInternalServerError, models.TOKEN_INVALID)
			return nil
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			err := core.IsTokenExpired(claims["expired_at"].(int64), claims["type"].(string))
			NewFailJsonResponse(gc, http.StatusInternalServerError, "token.is.expired")
			if err == true {
				NewFailJsonResponse(gc, http.StatusInternalServerError, "token.is.expired")
				return nil
			}
			return claims
		}
	} else {
		NewFailJsonResponse(gc, http.StatusInternalServerError, models.TOKEN_INVALID)
	}
	return nil
}

func (api *API) InspectToken(context *core.Core , token string) (jwt.MapClaims , *models.AppError) {

	if token == "" {
		//		NewFailJsonResponse(gc, http.StatusUnauthorized, "unauthorized")
		fmt.Println("접근권한 없음")
		return nil , models.NewAppError("api.AnalyzeToken","token.is.not.exist",nil,"",401);
	}

	bearerToken := strings.Split(token, " ")
	fmt.Sprintf("베어러토큰 : %s", token)
	if len(bearerToken) == 1 {
		token, err := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("token.signing.is.invalid")
			}
			fmt.Println("암호키리턴")
			return []byte(core.GLOBAL_SALT), nil
		})

		if err != nil {
			fmt.Println("토큰에러 발생")
			//	NewFailJsonResponse(gc, http.StatusInternalServerError, models.TOKEN_INVALID)
			return nil, models.NewAppError("api.AnalyzeToken","token.signing.is.invalid",nil,"",401);
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			expired := core.IsTokenExpired(claims["expired_at"].(int64), claims["type"].(string))

			if expired {
				//				NewFailJsonResponse(gc, 401, "token.is.expired")
				return nil, models.NewAppError("api.AnalyzeToken","token.is.expired",nil,"",401);
			}
			return claims, nil
		} else {
			//NewFailJsonResponse(gc, http.StatusInternalServerError, models.TOKEN_INVALID)
			return nil, models.NewAppError("api.AnalyzeToken",models.TOKEN_INVALID,nil,"",401);
		}
	}
	return nil, models.NewAppError("api.AnalyzeToken","token.is.invalid",nil,"",401);
}



func (api *API) IsTokenDataCoincidence(context *core.Core, token string, key , value string) (*models.AppError) {

	claims, err  := api.InspectToken(context, token)
	if err != nil {
		return err
	}

	if  claims[key] == value || claims["Role"] == "admin" {
		return nil
	}
	return models.NewAppError("IsTokenRoleValid",key+".is.not.valid",nil,"",401)
}




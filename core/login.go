package core
/*
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/fatih/color"
	"fmt"
	_ "goServer/store"
	"goServer/models"
	"encoding/json"
)

type Claims struct {
	*jwt.StandardClaims
	Iss       string `json:"iss"`
	role      int    `json:"role"`
	Id        string `json:"id"`
	ExpiredAt int64  `json:"expired_at"`
	//ExpiredAt int64 `json:"expired_at"`
}

func (core *Core) LoginPostHandler(c *gin.Context) { // /login post
	color.Red("/login post")
	fmt.Println("컨텍스트 :")

	var authData models.User

	d := json.NewDecoder(c.Request.Body)
	d.UseNumber()
	if err := d.Decode(&authData); err != nil {
		panic(err)
	}

	if a := <-core.Srv.Store.User().Get(authData.Id); a.Data != nil {
		fmt.Println("SUCCESS")
		fmt.Println(authData)
	} else {
		fmt.Println("FAIL")
	}


	AccessToken := jwt.New(jwt.GetSigningMethod("HS256"))
	AccessToken.Claims = &Claims{
		&jwt.StandardClaims{
			//ExpiresAt :  models.GetMillis(),//models.GetExpiredTime()
		},
		"localhost:8080",
		777,
		authData.Id,
		models.GetExpiredTime(),
	}

	fmt.Println("아이디 : ", authData.Id)
	fmt.Println("패스워드 : ", authData.Password)

	fmt.Println("AccessToken :", *AccessToken)
	AccessTokenString, error := AccessToken.SignedString([]byte("secret"))
	fmt.Println("tokentString :", AccessTokenString)
	if error != nil {
		fmt.Println("/login 포스트 에러")
		fmt.Println(error)
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(JwtToken{Token: AccessTokenString})
	fmt.Println("/login 포스트 성공")
}
/*
		test := c.MustGet("model")
		test1 := test.(*store.UserStorage)
		if  a := <-test1.Get(authData.Id); a.Data != nil {
			fmt.Println("낫닐")
			fmt.Println(a.Data)
		}
		if b := <-test1.Insert(authData.Id,authData.Password); b.Data != nil {
			fmt.Println("낫닐#2")
			fmt.Println(b.Data)

		}

	*/
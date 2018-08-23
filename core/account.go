package core

import (
	"goServer/models"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Err  *models.AppError
	Data map[string]interface{}
}

type JwtToken struct {
	Token string `json:"token"`
}

type Claims struct {
	*jwt.StandardClaims
	Iss       string `json:"iss"`
	TokenType string `json:"type"`
	Role      int    `json:"role"`
	Id        string `json:"id"`
	ExpiredAt int64  `json:"expired_at"`

	//ExpiredAt int64 `json:"expired_at"`
}



func ClaimsToStruct(claims map[string]interface{}) *Claims{
	Claims := &Claims{}
	Claims.Id = claims["Id"].(string)
	Claims.Iss = claims["Iss"].(string)
	Claims.Iss = claims["TokenType"].(string)
	Claims.Role = claims["Role"].(int)
	Claims.ExpiredAt = claims["ExpiredAt"].(int64)
	return Claims
}


func NewResult(Err string, Data map[string]interface{}) *Result {
	return &Result{nil, make(map[string]interface{})}
}

func (c *Core) GetUserExistenceById(id string) bool {
	if rst := <-c.Srv.Store.User().IsUserIdExist(id); rst.Data == 0 {
		return false
	}
	return true

}
func (c *Core) LoginById(id, password string) ( map[string]interface{}, *models.AppError) {

	result := NewResult("", nil)
	fmt.Println("로그인 바이 아이디")
	if rst := <-c.Srv.Store.User().Get(id); rst.Err != nil {
		fmt.Println("유저 찾지못함")
		result.Err = models.NewAppError("loginById", "api.user.login.invalid", nil, "user.is.not.found", http.StatusNotFound)
		return nil, result.Err;
	} else {
		if models.ComparePassword(rst.Data.(*models.User).Password, password) {

			AccessToken, err := c.CreateToken(rst.Data.(*models.User),ACCESS_TOKEN )

			if err != nil {
				fmt.Sprintf(" 엑세스 토큰 생성 실패 : %s", result.Err)
				return nil, result.Err;
			}
			RefreshToken , err := c.CreateToken(rst.Data.(*models.User),ACCESS_TOKEN )

			if err != nil {
				fmt.Sprintf(" 리프레시 토큰 생성 실패 : %s", result.Err)
				return nil, result.Err;
			}



			result.Data["access-token"] = AccessToken
			result.Data["refresh-token"] = RefreshToken

		} else {
			fmt.Println("비밀번호 틀림")
			result.Err = models.NewAppError("loginById", "api.user.login.invalid", nil, "", http.StatusUnauthorized)
			return nil, result.Err
		}
	}
	return result.Data, nil
}

func (c *Core) CreateToken(user *models.User, tokenType string) (data interface{}, AppErr *models.AppError) {

	result := NewResult("", nil)
	fmt.Println("크리에잇 토큰")
	NewToken := jwt.New(jwt.GetSigningMethod("HS256"))
	NewToken.Claims = &Claims{
		&jwt.StandardClaims{
		},
		"localhost:8080",
		tokenType,
		user.Roles,
		user.Id,
		GetExpiredTime(tokenType),
	}
	NewTokenString, err := NewToken.SignedString([]byte(GLOBAL_SALT))

	if err != nil {
		fmt.Sprintf("%s 타입 토큰 생성 실패 : %s", tokenType, err)

		result.Err = models.NewAppError("loginById", "api.user.failed.to.create.token", nil, tokenType, http.StatusUnauthorized)

	} else {
		fmt.Sprintf("%s 타입 토큰 생성 성공 : %s", tokenType, NewTokenString)
	}

	return NewTokenString, nil

}
func IsTokenExpired(tokenCreatedTime int64, tokenType string) bool {

	if tokenType == ACCESS_TOKEN {
		if tokenCreatedTime+ACCESS_TOKEN_LIFE_TIME >= GetMillis() {
			return false
		} else {
			return true
		}
	} else {
		if tokenCreatedTime+REFRESH_TOKEN_LIFE_TIME >= GetMillis() {
			return false
		} else {
			return true
		}
	}
}

func GetExpiredTime(tokenType string) int64 {
	if tokenType == ACCESS_TOKEN {
		return GetMillis() + ToMillis(ACCESS_TOKEN_LIFE_TIME)
	} else {
		return GetMillis() + ToMillis(REFRESH_TOKEN_LIFE_TIME)
	}
}

func ToMillis(t int64) int64 {
	return t / int64(time.Millisecond)
}
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
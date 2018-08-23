package core
/*
import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"strings"
	"fmt"
	"goServer/models"
	"github.com/dgrijalva/jwt-go"
)


type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func IsTokenExpired(tokenTime interface{}) *models.AppError {
	if exp, ok := tokenTime.(float64); ok {
		if models.GetMillis() > int64(exp) {
			fmt.Println("토큰만료")

		}
	}
	return nil
}


func TestEndPoint(c *gin.Context) {
	fmt.Println("테스트 엔드포인트 도달")
	decoded, _ := c.Get("decoded")
	var user models.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(c.Writer).Encode(user)
}

func ProtectedEndPoint(c *gin.Context) {
	params := c.Request.URL.Query()

	token, err := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if err, ok := token.Method.(*jwt.SigningMethodRSA); !ok {

			if err != nil {
				fmt.Println("ProtectedEndPoint에러")
			}
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Errorf("%s", "error")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(c.Writer).Encode(user)
	} else {
		json.NewEncoder(c.Writer).Encode(Exception{Message: "Invalid authorization token"})
	}
}

func CheckMiddleWare(next gin.HandlerFunc) gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {

		authorizationHeader := c.Request.Header.Get("Authorization")

		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")
			fmt.Println("베어러토큰", bearerToken, len(bearerToken))
			if len(bearerToken) == 1 {
				token, err := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {

					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("THERE WAS ERROR")
					}
					fmt.Println("암호키리턴")
					return []byte("secret"), nil
				})

				if err != nil {
					fmt.Println("토큰에러")

					json.NewEncoder(c.Writer).Encode(Exception{Message: err.Error()})
					return
				}
				if token.Valid {
					claims := token.Claims.(jwt.MapClaims)
					err := IsTokenExpired(claims["expired_at"])
					if err != nil {
						json.NewEncoder(c.Writer).Encode(Exception{Message: "TOKEN EXPIRED"})
					}
				}
				c.Set("decoded", token.Claims)
				next(c)

			} else {
				fmt.Println("유효하지 않은 검증")
				json.NewEncoder(c.Writer).Encode(Exception{Message: "INVALID AUTH"})
			}

		} else {
			fmt.Println("에러")
			json.NewEncoder(c.Writer).Encode(Exception{Message: "AUTH HEADER REQUIRED"})
		}
	})
}

func LoginGetHandler(c *gin.Context) {

}

func LoginPostHandler(c *gin.Context) {
	/*
		var authData models.UserInfo
		//var timeNow  = time.Now().Unix()
		//var lifeTime int64 = 60 * 15
		//var expiredTime = timeNow + lifeTime
		fmt.Println("/login 포스트")
		f.Println("모델 데이터베이스 정의됨",models.Database)




		_ = json.NewDecoder(c.Request.Body).Decode(&authData)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"id":       authData.Id,
				"password": authData.Password,
				"expired_at" : time.Now().Unix()+CONSTANT.LIFE_TIME,
			})
		f.Println("token :",*token)
		tokenString, error := token.SignedString([]byte("secret"))
		f.Println("tokentString :",tokenString)
		if error != nil {
			f.Println("/login 포스트 에러")
			f.Println(error)


		}
		c.Writer.Header().Set("Content-Type","application/json")
		json.NewEncoder(c.Writer).Encode(JwtToken{Token: tokenString})
		f.Println("/login 포스트 성공")

}
	*/
package models

import (
	"encoding/json"
	"io"
	"strings"
	"net/mail"
	"fmt"
	"bytes"
	"encoding/base32"
	"github.com/pborman/uuid"
	"net/http"

	"time"
	"strconv"
)

type StringMap map[string]string
type StringArray []string
type EncryptStringMap map[string]string

const (
	TOKEN_INVALID = "token.is.invalid"
)

type AppError struct {
	Id 			  string `json:"id"`
	Message       string `json:"message"`  // Message to be displayed to the end user without debugging information
	DetailedError string `json:"detailed_error"` // internal error string to help the developer
	RequestId     string `json:"request_id"` // the requestId that's also set in the header
	StatusCode    int    `json:"status_code"` // the http status code
	Where         string `json:"-"`  // the function where it happend in the form of struct.func
	params 		  map[string]interface{}

}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26)
	return b.String()


}

func (er *AppError) Error() interface{} {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) ToJson() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}

}

func AppErrorFromJson(data io.Reader) *AppError {
	decoder := json.NewDecoder(data)
	var er AppError

	err := decoder.Decode(&er)
	if err == nil {
		return &er
	} else {
		return nil
	}
}



func NewAppError(where string, id string, params map[string]interface{},details string,status int ) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = status
	return ap
}


func NewMongoError(where string, method string,details string ) *AppError{
	Id := fmt.Sprintf("Database Error of method %s",method)
	return NewAppError(where,Id,nil,details,http.StatusInternalServerError)
}



func MapToJson(objmap map[string]string) string { // 테스트됨
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func MapFromJson(data io.Reader) map[string]string { // 테스트됨
	decoder := json.NewDecoder(data)
	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

func ArrayToJson(objmap []string) string { // 테스트됨
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}
func ArrayFromJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)
	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	} else {
		return objmap
	}
}
func IsLower(s string) bool { // 테스트  안함
	if strings.ToLower(s) == s {
		return true
	}
	return false
}

func IsValidEmail(email string) bool { // 테스트 됨

	if !IsLower(email) {
		return false
	}
	if _, err := mail.ParseAddress(email); err == nil {
		return true
	}
	return false

}
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GalleryToMap(data io.ReadCloser) (Gallery, *AppError){
	var gallery Gallery
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&gallery)

	if err != nil {
		return  gallery , NewAppError("GalleryToMap","failed.to.mapify",nil,"",500)
	}
	return gallery , nil
}

func GenerateImageHash(userId string, index int) string {

	var b bytes.Buffer
	encoder := base32.NewEncoder(base32.HexEncoding,&b)
	encoder.Write([]byte(userId+strconv.FormatInt(GetMillis(),10)))
	return b.String()+"="+strconv.Itoa(index)
}

func GenerateGalleryHash(userId string) string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(base32.HexEncoding,&b)
	encoder.Write([]byte("gallery:"+userId))
	return b.String()
}

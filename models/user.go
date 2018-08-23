package models

import (
	"regexp"
	"io"
	"encoding/json"
	"strings"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER_EMAIL_MAX_LENGTH = 128
	USER_NICKNAME_MAX_RUNES = 64
	USER_FIRST_NAME_MAX_RUNES = 64
	USER_LAST_NAME_MAX_RUNES = 64
	USER_NAME_MAX_LENGTH = 64
	USER_NAME_MIN_LENGTH = 1
	USER_PASSWORD_MAX_LENGTH = 72


)

type User struct {
	Id                 string    `json:"id"`
	CreateAt           int64     `json:"create_at,omitempty"`
	UpdateAt           int64     `json:"update_at,omitempty"`
	DeleteAt           int64     `json:"delete_at"`
	Username           string    `json:"username"`
	Password		   string    `json:"password,omitempty"`
	AuthData           string    `json:"auth_data,omitempty"`
	AuthService        string    `json:"auth_service"`
	Email              string    `json:"email"`
	EmailVerified      bool      `json:"email_verified,omitempty"`
	Nickname           string    `json:"nickname"`
	Organization	   string	 `json:"org"`
	Job				   string	 `json:"job"`
	Roles  		       int    	 `json:"roles"`
	OwnGallery		   Gallery  `json:"gallary"`
	NotifyProps        StringMap `json:"notify_props,omitempty"`
	LastPasswordUpdate int64     `json:"last_password_update,omitempty"`
	LastPictureUpdate  int64     `json:"last_picture_update,omitempty"`
	FailedAttempts     int       `json:"failed_attempts,omitempty"`
}


var validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)
var restrictedUsernames = []string {
	"creart",
	"artman",
	"admin",
}
func (u *User) IsValid() *AppError {
	if len(u.Id) != 26 {
		return InvalidUserError("id","")
	}

	if u.CreateAt == 0 {
		return InvalidUserError("create_at",u.Id)
	}

	if u.UpdateAt == 0 {
		return InvalidUserError("update_at",u.Id)
	}

	if !IsValidUsername(u.Username) {
		return InvalidUserError("uesrname", u.Id)
	}

	if len(u.Email) > USER_EMAIL_MAX_LENGTH || len(u.Email) == 0 {
		return InvalidUserError("email",u.Id)
	}

	if len(u.Password) > USER_PASSWORD_MAX_LENGTH {
		return InvalidUserError("password",u.Password)
	}
	return nil
}

func InvalidUserError(fieldName string, userId string) *AppError {
	id := fmt.Sprintf("model.user.is_valid.%s.app_error",fieldName)
	details := ""
	if userId != "" {
		details = "user_id="+userId
	}
	return NewAppError("User.IsValid",id, nil,details,http.StatusBadRequest)
}

func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = NewId()
	}

	if u.Username == "" {
		u.Username = NewId()
	}

	u.Username = NormalizeUsername(u.Username)
	u.Email = NormalizeEmail(u.Email)
	u.CreateAt = GetMillis()
	u.UpdateAt = u.CreateAt
	u.LastPasswordUpdate = u.CreateAt

	if len(u.Password) > 0 {
		u.Password = HashPassword(u.Password)
	}


}

func (u *User) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}else {
		return string(b)
	}
}



func IsValidUsername(s string) bool  {
	if len(s) < USER_NAME_MIN_LENGTH || len(s) > USER_NAME_MAX_LENGTH {
		return false
	}
	if !validUsernameChars.MatchString(s) {
		return false
	}

	for _, restrictedUsername := range restrictedUsernames {
		if s == restrictedUsername {
			return false
		}
	}
	return true
}



func NormalizeUsername(username string) string {
	return strings.ToLower(username)
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}

func UserFromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user User
	err := decoder.Decode(&user)
	if err == nil {
		return &user
	}else{
		return nil
	}
}


func UserMapToJson(u map[string]*User) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func UserMapFromJson(data io.Reader) map[string]*User {
	decoder := json.NewDecoder(data)
	var users map[string]*User
	err := decoder.Decode(&users)
	if err == nil {
		return users
	} else {
		return nil
	}
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password),10)
	fmt.Println("이게 돌아간다고 ?",hash)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ComparePassword(hash string, password string) bool {
	if len(password) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}

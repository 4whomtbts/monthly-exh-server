package models

func IsTokenExpired(tokenCreatedTime int64) bool {
	if tokenCreatedTime + TOKEN_LIFE_TIME >= GetMillis(){
		return true
	}else{
		return false
	}
}
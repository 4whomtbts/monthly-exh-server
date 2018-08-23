package models

type SystemSetting struct {
	EC3 string

}

const (
	EC3_INSTANCE = ""

)

var TOKEN_LIFE_TIME int64 = 15 * 60 * 1000

func(ss *SystemSetting) setDefault() {

}


package models

import (
	"testing"
)

func TestIsTokenExpired(t *testing.T) {
	fixedTime := GetMillis()
	shouldBeExpired := fixedTime + TOKEN_LIFE_TIME + 1
	shouldBeAlive  := fixedTime + TOKEN_LIFE_TIME - 1
	if !IsTokenExpired(shouldBeExpired) {

		t.Fatal("token should be expired")
	}

	if !IsTokenExpired(shouldBeAlive) {
		t.Log("shouldBeAlive :",shouldBeAlive)
		t.Log("expiredAt :", fixedTime+TOKEN_LIFE_TIME)
		t.Fatal("token should be alive")
	}

}

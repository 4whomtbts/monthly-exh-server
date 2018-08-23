package models

import (
	"testing"
	"net/http"
	"strings"
)

func TestAppError(t *testing.T) {
	err := NewAppError("util.go/TestAppError","message",nil,"",http.StatusInternalServerError)
	json := err.ToJson()
	rerr := AppErrorFromJson(strings.NewReader(json))
	if err.Message != rerr.Message {
		t.Fatal("Error jsonfiy error")
	}
	t.Log(err.Error())
}

func TestMapJson(t *testing.T){
	m := make(map[string]string)
	m["id"] = "test_id"
	json := MapToJson(m)
	rm := MapFromJson(strings.NewReader(json))

	if rm["id"] != "test_id" {
		t.Fatal("map should be valid")
	}

	rm2 := MapFromJson(strings.NewReader(""))
	if len(rm2) > 0 {
		t.Fatal("make should be invalid")
	}
}

func TestArrayJson(t *testing.T) {
	testArray := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}

	testJsontify := ArrayToJson(testArray)
	testArrayFromJson := ArrayFromJson(strings.NewReader(testJsontify))

	for i := range testArrayFromJson {
		if testArray[i] != testArrayFromJson[i] {
			t.Fatal("number should be equal")
		}
	}
}


func TestValidEmail(t *testing.T){
	if !IsValidEmail("kjun+1385@naver.com"){
		t.Error()
	}

	if IsValidEmail("@kjun+1385@naver.com"){
		t.Error("should be invalid")
	}
}
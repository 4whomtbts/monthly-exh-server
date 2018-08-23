package storetest

import (
	"testing"
	"goServer/store"
	"goServer/models"
	"strings"
)

func TestUserStore(t *testing.T, ss store.Store) {
	t.Run("Save", func(t *testing.T) { testUserStoreSave(t,ss)})
	t.Run("Get", func(t *testing.T) { testUserStoreGet(t, ss) })

}


func testUserStoreSave(t *testing.T, ss store.Store) {
	t.Log("세이브테스트시작")
	u1 := models.User{}
	u1.Email = models.NewId()
	u1.Username = models.NewId()
	t.Log("0")
	if err := (<-ss.User().Save(&u1)).Err; err != nil {
		t.Log("0-0")

		t.Fatal("couldn't save user",err)
	}

	u1.Id = ""
	if err := (<-ss.User().Save(&u1)).Err; err == nil {
		t.Fatal("should be unique email")
	}
	u1.Email = ""
	if err := (<-ss.User().Save(&u1)).Err; err == nil {
		t.Fatal("should be unique username")
	}
	u1.Email = strings.Repeat("012345679",20)
	u1.Username = ""
	if err := (<-ss.User().Save(&u1)).Err; err == nil {
		t.Logf("error %s",(<-ss.User().Save(&u1)).Err)
		t.Fatal("should be unique username")
	}
	for i := 0; i < 49; i++ {
		u1.Id = ""
		u1.Email = models.NewId()
		u1.Username = models.NewId()
		if err := (<-ss.User().Save(&u1)).Err; err != nil {
			t.Fatal("couldn't save item",err)
		}
	}
	u1.Id = ""
	u1.Email = models.NewId()
	u1.Username = models.NewId()
	if err := (<-ss.User().Save(&u1)).Err; err != nil {
		t.Fatal("couldn't save item",err)
	}
	t.Log("Save 테스트 엔드포인트")
}

func testUserStoreGet(t *testing.T, ss store.Store) {
	t.Log("4")
	u1 := &models.User{}
	u1.Email = models.NewId()
	store.Must(ss.User().Save(u1))

	if r1 := <-ss.User().Get(u1.Id); r1.Err != nil {
		t.Fatal(r1.Err)
	} else {
		if r1.Data.(*models.User).ToJson() != u1.ToJson(){
			t.Fatal("invalid returned user")
		}
	}
	if err := (<-ss.User().Get("")).Err; err == nil {
		t.Fatal("Missing id should have failed")
	}
}



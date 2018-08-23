package mocks

import (

	mock "github.com/stretchr/testify/mock"

	store "'/store"
	"fmt"
	"goServer/models"
)
type UserStore struct {
	mock.Mock
}

func (_m *UserStore) Save(user *models.User) store.StoreChannel {
	ret := _m.Called(user)
	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*models.User) store.StoreChannel); ok {
		r0 = rf(user)
	}else{
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}
	return r0
}

func(_m *UserStore) Get(id string) store.StoreChannel {
	fmt.Println("MOCK 겟 호출")
	ret := _m.Called(id)
	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	}else{
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}
	return r0
}
/*
func(_m *UserStore) NewUser(ud0 models.User) models.User {
	fmt.Println("MOCK <NewUser> 호출")
	ret := _m.Called()
	var r0 models.User
	r0 = ret.Get(0).(models.User)
	return r0
}
*/


package mocks

import (
	"github.com/stretchr/testify/mock"
	"goServer/store"
)

type Store struct {
	mock.Mock
}

func (_m *Store) User() store.UserStore {
	ret := _m.Called()
	var r0 store.UserStore
	if rf, ok := ret.Get(0).(func() store.UserStore); ok {
		r0 = rf()
	}else{
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.UserStore)
		}
	}
	return r0
}

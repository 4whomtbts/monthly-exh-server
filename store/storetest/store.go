package storetest

import (

	"goServer/store"
	"goServer/store/storetest/mocks"
	"github.com/stretchr/testify/mock"

)

func NewStoreChannel(result store.StoreResult) store.StoreChannel {
	ch := make(store.StoreChannel, 1)
	ch <- result
	return ch
}

type Store struct {
	UserStore mocks.UserStore
}

func(s *Store) User() store.UserStore { return &s.UserStore }

func (s *Store) AssertExpectations(t mock.TestingT) bool {
	return mock.AssertExpectationsForObjects(t,
		&s.UserStore,
			)
}
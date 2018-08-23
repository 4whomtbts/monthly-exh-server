package store

import "context"

type LayeredStoreDatabaseLayer interface {
	Store
}

type LayeredStore struct {
	TmpContext context.Context
	DatabaseLayer LayeredStoreDatabaseLayer


}

func NewLayeredStore(db LayeredStoreDatabaseLayer) Store {
	store := &LayeredStore{
		TmpContext : context.TODO(),
		DatabaseLayer : db,
	}
	return store
}

func (s *LayeredStore) User() UserStore {
	return s.DatabaseLayer.User()

}
func (s *LayeredStore) Post() PostStore {
	return s.DatabaseLayer.Post()
}
func (s *LayeredStore) Gallery() GalleryStore {
	return s.DatabaseLayer.Gallery()
}


type LayeredSchemeStore struct {
	*LayeredStore
}
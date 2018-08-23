package mongoStore

import (
	"goServer/store"
	"github.com/globalsign/mgo"
)

type MongoStore interface {

	User() store.UserStore
	Post() store.PostStore
}

type PublicStore interface{
	GetCollection() *mgo.Collection
}

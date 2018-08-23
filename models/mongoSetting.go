package models

import (
	"github.com/globalsign/mgo"
)

type MongoSetting struct {
	Signature string
	Collections map[string]string
	Url string
	Role mgo.Index
}

const (
	SIGNATURE_TEST = "test"
	COLLECTION_USER = "users"
	COLLECTION_POST = "posts"
	COLLECTION_GALLERY = "gallery"

)


func (ms *MongoSetting) SetDefaults() {
	ms.Signature = SIGNATURE_TEST

	ms.Collections = map[string]string {
		"users" :  COLLECTION_USER,
		"posts" :  COLLECTION_POST,
		"gallery" : COLLECTION_GALLERY,
	}
	ms.Role = mgo.Index{
		Unique : true,
	}

}



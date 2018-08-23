package models

import "github.com/globalsign/mgo/bson"

type Gallery struct {
	Id                 string    `json:"gallery_id"`
	OwnerId 		   string    `json:"id"`
	CreateAt           int64     `json:"create_at,omitempty"`
	UpdateAt           int64     `json:"update_at,omitempty"`
	DeleteAt           int64     `json:"delete_at"`
	GalleryPath        string    `json:"gallery_path"`
	Template		   string    `json:"template"`
	Articles		   []bson.ObjectId `json:"gallery_articles"`
	GalleryPlan		   string	 `json:"plan"`
	ExpiredAt		   int64	 `json:"expired_at, omitempty"`
}


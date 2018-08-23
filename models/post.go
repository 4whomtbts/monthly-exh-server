package models

import "github.com/globalsign/mgo/bson"

type GeneralArticle struct {
	ObjectId		   bson.ObjectId 	`json:"_id"`
	Id                 int 	     `json:"id"`
	CreateAt           int64     `json:"create_at,omitempty"`
	UpdateAt           int64     `json:"update_at,omitempty"`
	DeleteAt           int64     `json:"delete_at"`
	Contents		   string	 `json:"contents"`
	Username           string    `json:"username"`
	IsPinned           bool		 `json:"notify_props,omitempty"`
	ImageRef		   []string  `json:"images_ref"`
	GalleryRef		   string	 `json:"gallery_ref"`
	HitCount		   int 	     `json:"hit_count"`
}

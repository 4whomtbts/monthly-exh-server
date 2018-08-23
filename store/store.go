package store

import (
	"goServer/models"
	"time"
	"fmt"
	"github.com/globalsign/mgo/bson"
)
type StoreResult struct {
	Data interface{}
	Err *models.AppError

}

type StoreChannel chan StoreResult


func Do(f func(result *StoreResult)) StoreChannel {
	storeChannel := make(StoreChannel , 1)
	go func() {
		result := StoreResult{}
		f(&result)
		storeChannel <- result
		close(storeChannel)
	}()
	return storeChannel
}

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		time.Sleep(time.Second)
		fmt.Println("패닉 : ",r.Err)
		panic(r.Err)
	}
	return r.Data
}




type Store interface {
	User() UserStore
	Post() PostStore
	Gallery() GalleryStore
//	WebHookStore() WebHookStore
}

type WebHookStore interface {
//	Router() Route
}


type PostStore interface{
	//Save(post *models.Post) StoreChannel
	//Update(post *models.Post, newMessage string, newHashTags string) StoreChannel
	GetById(id int) StoreChannel
	GetByField(field, value string) StoreChannel
	GetByObjectId(objectId bson.ObjectId) StoreChannel
	Save(contents *models.GeneralArticle) StoreChannel
	Delete(postId string) StoreChannel

}

type UserStore interface {
	Get(id string) StoreChannel
	IsUserIdExist(id string) StoreChannel
	Save(user *models.User) StoreChannel
//	SignUp(data models.User) StoreChannel
}

type GalleryStore interface {
	GetByGalleryId(galleryId string) StoreChannel
	GetByOwnerId(galleryOwnerId string) StoreChannel
	GetByGalleryPath(galleryPath string) StoreChannel
	UpdateGalleryArticle(galleryPath string, article_Id bson.ObjectId) StoreChannel
	Save(user *models.Gallery) StoreChannel
	Delete(id string) StoreChannel
	//	SignUp(data models.User) StoreChannel
}

type SchemeStore interface {


}
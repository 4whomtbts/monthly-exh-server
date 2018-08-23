package mongoStore

import (
	"goServer/store"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"goServer/models"
)

type MongoGalleryStore struct {
	MongoStore
	p0 *mgo.Collection
	index *mgo.Index
}

func NewGalleryStore(mongoStore MongoStore, collection *mgo.Collection) store.GalleryStore {
	gs := &MongoGalleryStore{
		MongoStore : mongoStore,
		p0 : collection,
	}

	gs.index = &mgo.Index {

		Unique : true,
		Background : true,
		Sparse : true,
	}
	return gs
}

func (gs MongoGalleryStore) Save(gallery *models.Gallery) store.StoreChannel {
	return store.Do(func(result *store.StoreResult){
		err := gs.p0.Insert(gallery)
		if err != nil {
			result.Err = models.NewAppError("api.gallery_store.save","gallery_save_error",nil,"",500)
		}


	})
}

func (gs MongoGalleryStore) UpdateGalleryArticle(galleryPath string, article_Id bson.ObjectId) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		err := gs.p0.Update(  bson.M { "gallerypath" : bson.M{ "$eq" : galleryPath}}, bson.M { "$push" : bson.M{ "articles" : article_Id }})
		if err != nil {
			result.Err = models.NewAppError("gallery_store.UpdateGalleryArticle","failed.to.update.article",nil,err.Error(),500)

		}
})
}


func (gs MongoGalleryStore) GetByGalleryId(galleryId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := gs.p0.Find(bson.M{"id": galleryId}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("gallery_store.GetByGalleryId","Gallery.is.not.found",nil,"gallery_ownerId :"+galleryId, 404)
		} else if mu == nil {
			result.Err = models.NewAppError("gallery_store.GetByGalleryId", "Gallery.is.empty", nil,"gallery_ownerId :"+galleryId,404)
		} else {
			result.Data = mu
		}
	})
}
func (gs MongoGalleryStore) GetByOwnerId(galleryOwnerId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := gs.p0.Find(bson.M{"ownerid": galleryOwnerId }).One(mu)
		if err != nil {
			result.Err = models.NewAppError("gallery_store.GetByOwnerId","Gallery.is.not.found",nil,"gallery_ownerId :"+galleryOwnerId, 404)
		} else if mu == nil {
			result.Err = models.NewAppError("gallery_store.GetByOwnerId", "Gallery.is.empty", nil,"gallery_ownerId :"+galleryOwnerId,404)
		} else {
			result.Data = mu
		}
	})
}
func (gs MongoGalleryStore) GetByGalleryPath(galleryPath string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.Gallery{}
		err := gs.p0.Find(bson.M{"gallerypath": galleryPath}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("gallery_store.GetByGalleryPath", "gallery.is.not.found", nil,"gallery_path :"+galleryPath, 500)
		} else if mu == nil {
			result.Err = models.NewAppError("gallery_store.GetByGalleryPath", "gallery.is.empty", nil,"gallery_path :"+galleryPath, 500)
		} else {
			result.Data = mu
		}
	})
}

func (gs MongoGalleryStore) Delete(id string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := gs.p0.Find(bson.M{"id": id}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("gallery_store <GET>", "gallery.is.not.found", nil,"gallery_id :"+id, 500)
		} else if mu == nil {
			result.Err = models.NewAppError("gallery_store <GET> ", "data.is.not.found", nil,"gallery_data_doesnt_exist",404)
		} else {
			result.Data = mu
		}
	})
}

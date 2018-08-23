package mongoStore

import (
"goServer/store"
"github.com/globalsign/mgo"
"github.com/globalsign/mgo/bson"
"goServer/models"
	"fmt"
	"strconv"
)

type MongoPostStore struct {
	MongoStore
	p0 *mgo.Collection
	index *mgo.Index
}

func NewPostStore(mongoStore MongoStore, collection *mgo.Collection) store.PostStore {
	ps := &MongoPostStore{
		MongoStore : mongoStore,
		p0 : collection,
	}

	ps.index = &mgo.Index {

		Unique : true,
		Background : true,
		Sparse : true,
	}
	return ps
}

func (ps MongoPostStore) Save(post *models.GeneralArticle) store.StoreChannel {
	return store.Do(func(result *store.StoreResult){
		/*
				if len(post.Id) > 0 {
					result.Err = models.NewAppError("MongoPostStore <Save>","store.mongo_store->existing app error",nil,"post_id="+post.Id, http.StatpsBadRequest)
					return
				}
		*/
		/*
		if result.Err = post.IsValid(); result.Err != nil {
			return
		}
		*/
		//post.PreSave()

		err := ps.p0.Insert(post)

		if post.GalleryRef != "" {
			//mu := &models.GeneralArticle{}
			 num, err  := ps.p0.Find(bson.M{ "galleryref" : bson.M{"$ne" : "jun" },
			 }).Count()

			if err != nil {
				fmt.Println(err)
				result.Err = models.NewAppError("api.post_store.save","failed.to.getGalleryRef",nil,"",500)
			}
			fmt.Println("갤래프 ",num)
		}

		if err != nil {
			result.Err = models.NewAppError("api.post_store.save","post_save_error",nil,"",500)
		}
	})
}
func (ps MongoPostStore) GetById(id int) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := ps.p0.Find(bson.M{"id": id}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("api.post_Store.getbyid", "post.is.not.found", nil,"post_id :"+strconv.Itoa(id), 404)
		} else if mu == nil {
			result.Err = models.NewAppError("post_store <GET> ", "data.is.not.found", nil,"post_data_doesnt_exist",404)
		} else {
			result.Data = mu
		}
	})
}
func (ps MongoPostStore) GetByObjectId(objectId bson.ObjectId) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := ps.p0.Find(bson.M{"_id": objectId}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("post_store <GET>", "post.is.not.found", nil,"", 500)
		} else if mu == nil {
			result.Err = models.NewAppError("post_store <GET> ", "data.is.not.found", nil,"post_data_doesnt_exist",404)
		} else {
			result.Data = mu
		}
	})
}
func (ps MongoPostStore) GetByField(field, value string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := ps.p0.Find(bson.M{ field : value}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("post_store <GET>", "post.is.not.found", nil,"post_id :"+value, 500)
		} else if mu == nil {
			result.Err = models.NewAppError("post_store <GET> ", "data.is.not.found", nil,"post_data_doesnt_exist",404)
		} else {
			result.Data = mu
		}
	})
}
func (ps MongoPostStore) Delete(id string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.GeneralArticle{}
		err := ps.p0.Find(bson.M{"id": id}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("post_store <GET>", "post.is.not.found", nil,"post_id :"+id, 500)
		} else if mu == nil {
			result.Err = models.NewAppError("post_store <GET> ", "data.is.not.found", nil,"post_data_doesnt_exist",404)
		} else {
			result.Data = mu
		}
	})
}

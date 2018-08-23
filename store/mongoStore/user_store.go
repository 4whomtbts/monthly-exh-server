package mongoStore

import (
	"goServer/store"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"goServer/models"
	"net/http"
	"fmt"
)


 const (
 	DATA_IS_NOT_EXIST = "1500"
 	DATA_IS_ALREADY_EXIST = "1501"


 )

type MongoUserStore struct {
	MongoStore
	p0 *mgo.Collection
	index *mgo.Index
}

func NewUserStore(mongoStore MongoStore, collection *mgo.Collection) store.UserStore {
	us := &MongoUserStore{
		MongoStore: mongoStore,
		p0:         collection,
	}

	us.index  = &mgo.Index{
		Key : []string{"email","Email"},
		Unique:true,
		Background: true,
		Sparse : true,
	}
	return us
}

func (us MongoUserStore) Save(user *models.User) store.StoreChannel {
	return store.Do(func(result *store.StoreResult){
/*
		if len(user.Id) > 0 {
			result.Err = models.NewAppError("MongoUserStore <Save>","store.mongo_store->existing app error",nil,"user_id="+user.Id, http.StatusBadRequest)
			return
		}
*/
		//user.PreSave()
		/*
		if result.Err = user.IsValid(); result.Err != nil {
			return
		}
		*/
		user.PreSave()
		err := us.p0.Insert(user)
		if err != nil {
			result.Err = models.NewAppError("MongoStore/user_store.go","save.error",nil,"",http.StatusBadRequest)
			return
		}
	})
}

func (us MongoUserStore) Get(id string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.User{}
		err := us.p0.Find(bson.M{"id": id}).One(mu)
		if err != nil {
			result.Err = models.NewAppError("user_store <GET>", "user.is.not.found", nil,"user_id :"+id, http.StatusInternalServerError)
		} else if mu == nil {
			result.Err = models.NewAppError("user_store <GET> ", "data.is.not.found", nil,"user_data_doesnt_exist",http.StatusNoContent)
		} else {
			result.Data = mu
		}
	})
}
func (us MongoUserStore) IsUserIdExist(id string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		mu := &models.User{}
		err := us.p0.Find(bson.M{"id": id}).One(mu)
		if err != nil {
			// id 에 해당하는 유저가 존재하지 않을 때
			/*  ERROR : NOT FOUND  */
			result.Data = 0
		} else {
			fmt.Println("존재할 때  ",result.Data)
			fmt.Println("존재할 때 : ",mu)
			result.Data = 1
		}
	})
}

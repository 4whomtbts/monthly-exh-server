package mongoStore

import (
	"goServer/store"
	"github.com/globalsign/mgo"
	"goServer/models"
	"fmt"
	"net/http"
)

type MongoStores struct {
	user store.UserStore
	post store.PostStore
	gallery store.GalleryStore
}

type MongoSupplier struct {
	rrCounter int64
	session *mgo.Session
	collections map[string]*mgo.Collection
	settings *models.MongoSetting
	mongoStores MongoStores

}

func NewMongoStore(settings *models.MongoSetting) *MongoSupplier {
	supplier := &MongoSupplier{
		rrCounter : 0,
		settings : settings,
		collections : make(map[string]*mgo.Collection),
	}
	fmt.Println("서플라이어 초기화")
	supplier.initiateDatabaseConnection()


	supplier.mongoStores.user = NewUserStore(
			supplier,
			supplier.GetCol(models.COLLECTION_USER),
			)


	supplier.mongoStores.post = NewPostStore(
		supplier,
		supplier.GetCol(models.COLLECTION_POST),
		)

	supplier.mongoStores.gallery = NewGalleryStore(
		supplier,
		supplier.GetCol(models.COLLECTION_GALLERY),
	)

	fmt.Println("모델 컬렉션유저 :",models.COLLECTION_USER)
	fmt.Println(supplier.GetCol(models.COLLECTION_USER))
	return supplier


}

func (ss *MongoSupplier) initiateDatabaseConnection() {
		session, err := mgo.Dial(ss.settings.Url)
		if err != nil {
			models.NewAppError("store.supplier.go","DATABASE ERROR",nil,"session Error",http.StatusInternalServerError)
			panic(err)
		}
		ss.session = session

		ss.BindCollections()

}

func (ss *MongoSupplier) BindCollections() {
	signature := ss.settings.Signature
	session := ss.session
	for sign,col := range ss.settings.Collections {
		fmt.Println("ss.collections["+sign+"] = BindCollection(",col,signature,session)
		fmt.Println("시그니처 :",signature)
		fmt.Println("세션",session)
		ss.collections[sign] = BindCollection(col,signature,*session)

	}
}
func BindCollection(ColSignature string, Signature string, session mgo.Session) *mgo.Collection {
	return session.DB(Signature).C(ColSignature)
}
func (ss *MongoSupplier) User() store.UserStore {
	return ss.mongoStores.user
}
func (ss *MongoSupplier) Post() store.PostStore {
	return ss.mongoStores.post
}
func (ss *MongoSupplier) Gallery() store.GalleryStore {
	return ss.mongoStores.gallery
}
func (ss *MongoSupplier) GetCol(colSign string) *mgo.Collection {
	fmt.Println("컬렉션즈 :",ss.collections)
	fmt.Println("컬렉션즈 :",ss.collections["user"])
	fmt.Println("컬렉션즈 :",colSign)
	return ss.collections[colSign]
}



package mongoStore

import (
	"goServer/models"

	"goServer/store"
	"testing"
)

var storeTypes = []*struct {
	Name  string
	Func  func() (*models.MongoSetting, error)

	Store store.Store
}{
	{
	Name : "MongoDB",
	Store : store.NewLayeredStore(NewMongoStore(StoreInit())),
	},
}


func StoreTest(t *testing.T, f func(*testing.T,store.Store)){
	for _, st := range storeTypes {
		st := st // duplication
		t.Run(st.Name,func(t *testing.T) { f(t, st.Store)})
	}
}

func StoreInit() *models.MongoSetting {
	ms := &models.MongoSetting{}
	ms.SetDefaults()
	return ms

}

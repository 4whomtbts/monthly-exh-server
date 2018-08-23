package mongoStore

import (
	"testing"
	"goServer/store/storetest"
)

func TestUserStore(t *testing.T) {
	StoreTest(t, storetest.TestUserStore)

}


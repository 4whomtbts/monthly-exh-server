package core

import (
	"github.com/gin-gonic/gin"
	"goServer/models"
	"github.com/globalsign/mgo/bson"
)

func (c *Core) BuildGallery(gc *gin.Context) (*models.AppError) {

	gallery, err := models.GalleryToMap(gc.Request.Body)

	if err != nil {
		return err
	}

	gallery.Id = models.GenerateGalleryHash(gallery.OwnerId)
	gallery.CreateAt = models.GetMillis()

	if rst := <-c.Srv.Store.Gallery().Save(&gallery); rst.Err != nil {
		return rst.Err
	}
	return nil
}

func (c *Core) GetGallery(gc *gin.Context) (*models.Gallery, *models.AppError) {

	galleryPath := gc.Param("galleryPath")

	if rst := <-c.Srv.Store.Gallery().GetByGalleryPath(galleryPath); rst.Err != nil {
		return nil, rst.Err
	} else {

		return rst.Data.(*models.Gallery), nil
	}

}

func (c *Core) UpdateGalleryPost(galleryPath string, article bson.ObjectId) (*models.AppError) {

	if rst := <-c.Srv.Store.Gallery().UpdateGalleryArticle(galleryPath, article); rst.Err != nil {
		return rst.Err
	}
	return nil
}

func (c *Core) GetGalleryArticleWithGalleryPath(galleryPath string) ([]*models.GeneralArticle, *models.AppError) {

	if rst := <-c.Srv.Store.Gallery().GetByGalleryPath(galleryPath); rst.Err != nil {
		return nil, rst.Err
	} else {
		articles, err := c.GetGalleryArticles(rst.Data.(*models.Gallery).Articles)

		if err != nil {
			return nil, err
		} else {
			return articles, nil
		}

	}
}

func (c *Core) GetGalleryArticles(articleArray []bson.ObjectId) ([]*models.GeneralArticle, *models.AppError) {
	length := len(articleArray)
	var articles []*models.GeneralArticle

	for i := 0; i < length; i++ {

		rst, err := c.GetGalleryArticle(articleArray[i])

		if err != nil {
			return nil, err
		}

		articles = append(articles, rst)

	}
	return articles, nil
}


func (c *Core) GetGalleryArticle(articleObjectId bson.ObjectId) (*models.GeneralArticle, *models.AppError) {

	if rst := <-c.Srv.Store.Post().GetByObjectId(articleObjectId); rst.Err != nil {
		return nil, models.NewAppError("core.GetGalleryArticle", "failed.to.get.post", nil, rst.Err.DetailedError, 404)
	} else {
		return rst.Data.(*models.GeneralArticle) , nil
	}
}

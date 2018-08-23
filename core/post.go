package core

import (
	"goServer/models"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"path/filepath"
	"time"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (c *Core) WritePost(gc *gin.Context, galleryFlag bool) ([]string, *models.AppError) {

	var post models.GeneralArticle

	decoder := json.NewDecoder(gc.Request.Body)

	err := decoder.Decode(&post)
	if err != nil {
		fmt.Println(err)
		return nil, models.NewAppError("core.WritePost", "failed.to.jsonify", nil, "", 500)
	}

	newImageRef, receivedSignedUrl, GetSignedUrlErr := c.GetS3SignedUrls(post.Username, post.ImageRef)
	if GetSignedUrlErr != nil {
		return nil, GetSignedUrlErr
	}
	post.ImageRef = newImageRef
	post.ObjectId = bson.NewObjectId()
	if rst := <-c.Srv.Store.Post().Save(&post); rst.Err != nil {
		return nil, rst.Err
	} else {
		if galleryFlag == true {
			galleryPath := gc.Param("galleryPath")
			err := c.UpdateGalleryPost(galleryPath,post.ObjectId)
			if err != nil {
				return nil, err
			}
		}
	}
	return receivedSignedUrl, nil

}


func (c *Core) GetS3SignedUrls(userName string, images []string) ([]string, []string, *models.AppError) {

	var HashedImageName string
	var newImageRef []string
	var signedUrl [] string

	for i := 0; i < len(images); i++ {
		HashedImageName = models.GenerateImageHash(userName, i)
		url, err := c.GetS3SignedUrl(HashedImageName)
		newImageRef = append(newImageRef, HashedImageName)

		if err != nil {
			return nil, nil, err
		}
		signedUrl = append(signedUrl, url)
	}
	return newImageRef, signedUrl, nil
}

func (c *Core) GetS3SignedUrl(imageName string) (string, *models.AppError) {

	req, _ := c.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("spreadsheetbucket"),
		Key:    aws.String(filepath.Base(imageName)),
	})
	str, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", models.NewAppError("api.GetS3SignedUrl", "failed.to.GetS3SignedUrl", nil, "", 500)
	}
	return str, nil
}

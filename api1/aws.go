package api1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_"github.com/aws/aws-sdk-go/aws"
	_"path/filepath"
	_"os"
	"github.com/aws/aws-sdk-go/aws"
	"path/filepath"
	"goServer/core"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func (api *API)  InitAws() {


	api.BaseRoutes.Aws= api.svr.Group(mergePath(AWS_ROUTE))
	{
		api.BaseRoutes.Aws.GET("/test",api.ApiHandler(testPath))
		//api.BaseRoutes.Aws.POST("/s3/upload",api.ApiHandler(api.Img	))
		api.BaseRoutes.Aws.GET("/s3/download",api.ApiHandler(api.ImgDownload))
	}
}

func (api *API) GallaryImageUpdate(c *gin.Context) {
	ctx, _ := c.Get("context")
	var ctx0 core.Core = ctx.(core.Core)



	req , _  := ctx0.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket : aws.String("spreadsheetbucket"),
		Key : aws.String(filepath.Base("fuckingS3.jpeg")),
	})
	str, err := req.Presign(15 * time.Minute)
	if err != nil {
		return;
	}

	fmt.Println(str)
	return;

}

func (api *API) GallaryArticleImgUpload(c *gin.Context) {
	ctx, _ := c.Get("context")
	var ctx0 core.Core = ctx.(core.Core)

	//userId := c.Query("userId")


	req , _  := ctx0.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket : aws.String("spreadsheetbucket"),
		Key : aws.String(filepath.Base("fuckingS3.jpeg")),
	})
	str, err := req.Presign(15 * time.Minute)
	if err != nil {
		return;
	}

	fmt.Println(str)
	return;

}



func (api *API) ImgDownload(c *gin.Context) {
	ctx, _ := c.Get("context")
	var ctx0 core.Core = ctx.(core.Core)

	req , _  := ctx0.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket : aws.String("spreadsheetbucket"),
		Key : aws.String(filepath.Base("fuckingS3.jpeg")),
	})
	str, err := req.Presign(15 * time.Minute)
	if err != nil {
		return;
	}

	fmt.Println(str)
		return;

}


/*
ctx, _ := c.Get("context")
var ctx0 core.Core = ctx.(core.Core)

files  , err := c.FormFile("file")
if err != nil {
	fmt.Println(err)
}

file , err := files.Open()
defer file.Close()

 //var regExp = regexp.MustCompile(`/^data:text\/plain;base64,/`)
 //file = regExp.ReplaceAllString(file,"")
 svc := s3manager.NewUploader(ctx0.AwsSession)


 result, err:= svc.Upload(&s3manager.UploadInput{
	 Bucket : aws.String("spreadsheetbucket"),
	 Key : aws.String(filepath.Base("doraemong.jpeg")),
	 Body : file,
	 ContentEncoding : aws.String("base64"),
	 ContentType: aws.String(c.GetHeader("Content-Type")),
 })
 if err != nil {
	 fmt.Println(err)
	 return;
 }
 fmt.Println(result)
 return;
*/
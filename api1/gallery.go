package api1

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"goServer/core"
	"goServer/models"
	"strconv"
	"strings"
)

func (api *API) InitGallery() {
	api.BaseRoutes.Gallery = api.svr.Group(mergePath(GALLERY_ROUTE))
	{
		api.BaseRoutes.Gallery.POST("/:galleryPath/post", api.ApiHandler(api.GalleryWritePost))
		api.BaseRoutes.Gallery.GET("/:galleryPath/article/:postId", api.ApiHandler(api.ShowGalleryArticle))
		api.BaseRoutes.Gallery.POST("/:galleryPath", api.DetermineRouter(api.ApiHandler(api.BuildGallery), api.ApiHandler(api.GetGallery)))
		api.BaseRoutes.Gallery.GET("/:galleryPath", api.DetermineRouter(api.ApiHandler(api.BuildGallery), api.ApiHandler(api.GetGallery)))
		api.BaseRoutes.Gallery.GET("/:galleryPath/setting", api.ApiGalleryAuthRequired(api.GalleryGetSetting))
		api.BaseRoutes.Gallery.POST("/:galleryPath/setting", api.ApiGalleryAuthRequired(api.GalleryPostSetting))
		//api.BaseRoutes.Gallery.POST("build", api.ApiHandler(api.BuildGallery))

	}
}

func (api *API) DetermineRouter(former, latter gin.HandlerFunc) gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/api/v1/gallery/build") {
			former(c)
		} else {
			latter(c)
		}
	})
/*
	return gin.HandlerFunc(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/api/v1/gallery/build") {
			former
		} else {
			latter(c)
		}
	})*/
}

func (api *API) BuildGallery(c *gin.Context) {
	ctx, err := c.Get("context")
	if err == true {
		fmt.Println(err)
	}
	ctx0 := ctx.(*core.Core)
	buildError := ctx0.BuildGallery(c)

	if buildError != nil {
		NewFailJsonResponse(c, buildError.StatusCode, buildError.Id)
	}

	/*
		블로그식별번호, 게시글들
	 */
}

func (api *API) GetGallery(c *gin.Context) {
	ctx ,err := GetContext(c)
	if err != nil {
		return
	}
	galleryData, renderError := ctx.GetGallery(c)

	if renderError != nil {
		NewFailJsonResponse(c, renderError.StatusCode, renderError.Id)
		return;
	}

	articles, getArticleErr := ctx.GetGalleryArticleWithGalleryPath(galleryData.GalleryPath)

	if getArticleErr != nil {
		NewFailJsonResponse(c, getArticleErr.StatusCode, getArticleErr.Id +","+getArticleErr.DetailedError)
	}

	c.JSON(200, gin.H{
		"articles": articles,
		"template": galleryData.Template,
	})
}

func (api *API) GalleryWritePost(c *gin.Context) {
	ctx, err := c.Get("context")
	if err == true {
		fmt.Println(err)
	}
	ctx0 := ctx.(*core.Core)
	rst, WritePostError := ctx0.WritePost(c, true)
	if WritePostError != nil {
		NewFailJsonResponse(c, 500, WritePostError.Id)
		return
	}
	c.JSON(200, gin.H{
		"signedUrl": rst,
	})
}

func (api *API) ShowGalleryArticle(c *gin.Context) {
	ctx, err := GetContext(c)

	if err != nil {
		NewFailJsonResponse(c,500,err.Id)
		return
	}
	postId, parseIntError := strconv.Atoi(c.Param("postId"))

	if parseIntError != nil {
		NewFailJsonResponse(c,500,"post.is.not.found")
	}

	if rst := <-ctx.Srv.Store.Post().GetById(postId); rst.Err != nil {
		NewFailJsonResponse(c,rst.Err.StatusCode, rst.Err.Id)
		return
	}else{

		c.JSON(200, gin.H{
			"article" : rst.Data.(*models.GeneralArticle),
		})
	}


}

func (api *API) GalleryGetSetting(c *gin.Context) {

}

func (api *API) GalleryPostSetting(c *gin.Context) {

}

package api1

func (api *API)  InitRoot() {


	api.BaseRoutes.Index = api.svr.Group(mergePath(ROOT_ROUTE))
	{
		api.BaseRoutes.Index.GET("/test",api.ApiHandler(testPath))
		api.BaseRoutes.Index.GET("/index")

}
}



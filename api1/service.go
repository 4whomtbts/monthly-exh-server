package api1

func (api *API) InitService() {
	Service := api.svr.Group(mergePath(SERVICE_ROUTE))
	{
		Service.GET("/test",testPath)
		Service.GET("/index")
	}

}
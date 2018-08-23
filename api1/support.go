package api1

func (api *API) InitSupport() {
	Support := api.svr.Group(mergePath(SUPPORT_ROUTE))
	{
		Support.GET("/test",testPath)
		Support.GET("/index")
	}
}
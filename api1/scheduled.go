package api1

func (api *API)  InitScheduled() {
	api.BaseRoutes.Scheduled = api.svr.Group(mergePath(SCHEDULED_ROUTE))
	{
		api.BaseRoutes.Scheduled.GET("/test",testPath)
		api.BaseRoutes.Scheduled.GET("/index")

	}

}
package api1


func (api *API) InitContact() {
	api.BaseRoutes.Contact = api.svr.Group(mergePath(CONTACT_ROUTE))
	{
		api.BaseRoutes.Contact.GET("/test",testPath)
		api.BaseRoutes.Contact.GET("/index")
	}


}
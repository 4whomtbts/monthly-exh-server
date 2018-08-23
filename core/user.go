package core

import "goServer/models"

func (c *Core) CreateUserDefault(data *models.User) {
	c.Srv.Store.User().Save(data)
}
package core
/*
import (
	"goServer/models"
)

func (a *Core) Config() *models.Config {
	if cfg := a.config.Load(); cfg != nil {
		return cfg.(*models.Config)
	}
	return &models.Config{}
}

// Registers a function with a given to be called when the config is reloaded and may have changed. The function
// will be called with two arguments: the old config and the new config. AddConfigListener returns a unique ID
// for the listener thast can later be used to remove it.
func (a *Core) AddConfigListener(listener func(*models.Config, *models.Config)) string {
	id := models.NewId()
	a.configListeners[id] = listener
	return id
}

*/
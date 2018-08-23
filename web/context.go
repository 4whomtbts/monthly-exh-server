package web

import (
	"goServer/mlog"
	"goServer/core"
	"goServer/models"
)

type Context struct {
	Core          *core.Core
	Log           *mlog.Logger
	Params        *Params
	Err           *models.AppError
	RequestId     string
	IpAddress     string
	Path          string
}
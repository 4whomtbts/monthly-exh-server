package main

import (
	"goServer/core"
	"goServer/api1"
	"goServer/constants"
)

func main(){
	_core  := core.StartCore()
	api := api1.InitRouters(_core)
	api.StartRouters(constants.SERVER_PORT)

}


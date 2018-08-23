package web

import (
	"net/http"
	"goServer/core"
)

type Handler struct {
	App            *core.Core
	HandleFunc     func(*Context, http.ResponseWriter, *http.Request)
	RequireToken bool
	TrustRequester bool
	RequireMfa     bool
	IsStatic       bool
}

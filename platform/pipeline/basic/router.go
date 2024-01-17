package basic

import (
	"platform/http/handling"
	"platform/pipeline"
	"strings"
)

type RouterComponent struct {
	Routes []handling.Route
}

func (r *RouterComponent) Init() {}
func (r *RouterComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	for _, route := range r.Routes {
		if ctx.Request.Method == route.HTTPMethod && strings.Contains(ctx.Request.URL.Path, route.Pattern) {
			route.ExecuteHandler(ctx)
		}
	}

	next(ctx)
}

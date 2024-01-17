package handling

import (
	"platform/pipeline"
)

type Route struct {
	Pattern    string
	HTTPMethod string
	Handler    Handler
}

func (route *Route) ExecuteHandler(ctx *pipeline.ComponentContext) {
	route.Handler.Execute(ctx)
}

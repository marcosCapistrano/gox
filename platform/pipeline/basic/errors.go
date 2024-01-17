package basic

import (
	"fmt"
	"net/http"
	"platform/logging"
	"platform/pipeline"
)

type ErrorComponent struct {
	Logger logging.Logger
}

func recoveryFunc(ctx *pipeline.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ErrorComponent) Init() {}
func (c *ErrorComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	defer recoveryFunc(ctx, c.Logger)
	next(ctx)
	if ctx.GetError() != nil {
		c.Logger.Debugf("Error: %v", ctx.GetError())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

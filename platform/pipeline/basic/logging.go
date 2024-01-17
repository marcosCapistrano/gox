package basic

import (
	"net/http"
	"platform/logging"
	"platform/pipeline"
)

type LoggingResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (w *LoggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

type LoggingComponent struct {
	Logger logging.Logger
}

func (lc *LoggingComponent) Init() {}
func (lc *LoggingComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	loggingWriter := LoggingResponseWriter{0, ctx.ResponseWriter}
	ctx.ResponseWriter = &loggingWriter
	lc.Logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	lc.Logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL)
}

package main

import (
	"platform/config"
	"platform/http"
	"platform/http/handling"
	"platform/logging"
	"platform/pipeline"
	"platform/pipeline/basic"
	"tmotos/handlers"
)

func createPipeline(cfg config.Configuration, logger logging.Logger) pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.LoggingComponent{Logger: logger},
		&basic.ErrorComponent{Logger: logger},
		&basic.StaticFileComponent{Config: cfg},
		&basic.RouterComponent{
			Routes: []handling.Route{
				{Pattern: "/loja", HTTPMethod: "GET", Handler: handlers.StoreHandler{}},
			},
		},
	)
}

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	logger := logging.NewDefaultLogger(cfg)

	http.Serve(createPipeline(cfg, logger), cfg, logger).Wait()
}

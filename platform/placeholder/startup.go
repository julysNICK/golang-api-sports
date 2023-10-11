package placeholder

import (
	"platform/http"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline( //do
		&basic.ServicesComponent{}, //do
		&basic.LoggingComponent{},  // do
		// &basic.ErrorComponent{}, do
		&basic.StaticFileComponent{}, //do
		&SimpleMessageComponent{},    // do
	)
}

func Start() {
	results, err := services.Call(http.Serve, createPipeline()) //do
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}

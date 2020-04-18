package templates

// Main is a template for main go file.
const Main = `package main

import (
	"fmt"
	"log"
	"net/http"


	"go.uber.org/zap"


	"{{.Name}}/pkg/config"
	"{{.Name}}/pkg/logic"
	"{{.Name}}/pkg/service"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()


	logger.Info("{{.Name}} started")


	env, err := config.ProcessEnv()
	if err != nil {
		log.Fatal(err.Error())
	}


	api := logic.NewAPI(logger)
	httpWrapper := service.NewHTTPWrapper(api)
	router := service.NewRouter(httpWrapper.Paths(), logger)


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), router))
}`

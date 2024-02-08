package main

import (
	"context"
	"github.com/rahmaninsani/dns-api/app"
	"github.com/rahmaninsani/dns-api/config"
	"github.com/rahmaninsani/dns-api/exception"
	"github.com/rahmaninsani/dns-api/handler"
	"github.com/rahmaninsani/dns-api/repository"
	"github.com/rahmaninsani/dns-api/route"
	"github.com/rahmaninsani/dns-api/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	config.SetupConfig()
}

func main() {
	db := app.NewDB(config.Config.Subcommand.Serve.File)

	router := app.NewRouter()
	router.PanicHandler = exception.ErrorHandler

	securityCaseRepository := repository.NewSecurityCaseRepository(db)
	securityCaseUseCase := usecase.NewSecurityCaseUseCase(securityCaseRepository)
	securityCaseHandler := handler.NewSecurityCaseHandler(securityCaseUseCase)
	route.NewSecurityCaseRouter(router, securityCaseHandler)

	server := http.Server{
		Addr:    config.Config.Subcommand.Serve.Address,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started at:", config.Config.Subcommand.Serve.Address)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

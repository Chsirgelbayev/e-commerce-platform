package app

import (
	"app/internal/config"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/swag/cmd/swag"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) (App, error) {
	log.Print("Router initializing")

	router := httprouter.New()

	log.Println("Swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler(httpSwagger.WrapHandler))

	app := App{
		config: config,
	}
	return app, nil
}

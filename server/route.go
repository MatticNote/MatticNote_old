package server

import (
	"github.com/atreugo/cors"
	"github.com/savsgio/atreugo/v11"
)

func ConfigureRoute(app *atreugo.Atreugo) {
	app.GET("/", func(ctx *atreugo.RequestCtx) error {
		// WIP
		return ctx.TextResponse("GET")
	})
	apiPath := app.NewGroupPath("/api")

	apiConfigureRoute(apiPath)
}

func apiConfigureRoute(p *atreugo.Router) {
	p.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: false,
	}))
	p.GET("/", func(ctx *atreugo.RequestCtx) error {
		// WIP
		return ctx.TextResponse("API")
	})
}

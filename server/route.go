package server

import (
	apiV1 "github.com/MatticNote/MatticNote/server/api/v1"
	"github.com/atreugo/cors"
	"github.com/savsgio/atreugo/v11"
)

func ConfigureRoute(app *atreugo.Atreugo) {
	app.GET("/", func(ctx *atreugo.RequestCtx) error {
		// WIP
		return ctx.TextResponse("GET")
	})

	apiConfigureRoute(app.NewGroupPath("/api"))
}

func apiConfigureRoute(p *atreugo.Router) {
	p.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: false,
	}))
	apiV1ConfigureRoute(p.NewGroupPath("/v1"))
}

func apiV1ConfigureRoute(p *atreugo.Router) {
	p.GET("/user/{uuid}", apiV1.GetEntryUser)
}

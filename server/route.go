package server

import (
	"github.com/MatticNote/MatticNote/config"
	apiV1 "github.com/MatticNote/MatticNote/server/api/v1"
	"github.com/MatticNote/MatticNote/server/view"
	_ "github.com/MatticNote/MatticNote/server/view"
	"github.com/atreugo/cors"
	"github.com/gorilla/csrf"
	"github.com/savsgio/atreugo/v11"
)

func ConfigureRoute(app *atreugo.Atreugo) {
	app.GET("/", func(ctx *atreugo.RequestCtx) error {
		// WIP
		return ctx.TextResponse("GET")
	})
	internalPath := app.NewGroupPath("/i")
	internalPath.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   config.Config.Server.Endpoint,
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: true,
	}))
	csrfProtect := csrf.Protect(
		[]byte(config.Config.Server.CsrfSecret),
		csrf.Secure(config.Config.Server.CsrfSecure),
	)

	internalPath.NetHTTPPath("GET", "/signup", csrfProtect(view.InternalSignup{}))
	internalPath.NetHTTPPath("POST", "/signup", csrfProtect(view.InternalSignupPost{}))

	apiConfigureRoute(app.NewGroupPath("/api"))
}

func apiConfigureRoute(p *atreugo.Router) {
	p.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: false,
	}))
	apiV1.ConfigureRouteV1(p.NewGroupPath("/v1"))
}

package server

import (
	"github.com/MatticNote/MatticNote/config"
	apiV1 "github.com/MatticNote/MatticNote/server/api/v1"
	"github.com/MatticNote/MatticNote/server/view"
	_ "github.com/MatticNote/MatticNote/server/view"
	"github.com/atreugo/cors"
	"github.com/gorilla/csrf"
	"github.com/savsgio/atreugo/v11"
	"net/http"
)

func ConfigureRoute(app *atreugo.Atreugo) {
	app.GET("/", func(ctx *atreugo.RequestCtx) error {
		// WIP
		return ctx.TextResponse("GET")
	})

	internalConfigureRoute(app.NewGroupPath("/i"))
	apiConfigureRoute(app.NewGroupPath("/api"))
}

func internalConfigureRoute(internalPath *atreugo.Router) {
	internalPath.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   config.Config.Server.Endpoint,
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: true,
	}))
	csrfProtect := csrf.Protect(
		[]byte(config.Config.Server.CsrfSecret),
		csrf.Secure(config.Config.Server.CsrfSecure),
		csrf.ErrorHandler(http.HandlerFunc(view.CSRFTokenErrorView)),
	)

	internalPath.NetHTTPPath("GET", "/signup", csrfProtect(http.HandlerFunc(view.InternalSignup)))
	internalPath.NetHTTPPath("POST", "/signup", csrfProtect(http.HandlerFunc(view.InternalSignupPost)))
}

func apiConfigureRoute(r *atreugo.Router) {
	r.UseAfter(cors.New(cors.Config{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization", "Origin"},
		AllowCredentials: false,
	}))
	apiV1.ConfigureRouteV1(r.NewGroupPath("/v1"))
}

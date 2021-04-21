package server

import "github.com/savsgio/atreugo/v11"

func ConfigureRoute(app *atreugo.Atreugo) {
	app.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("GET")
	})
}

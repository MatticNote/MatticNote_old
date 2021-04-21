package view

import "github.com/savsgio/atreugo/v11"

func NotFoundErrorView(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(404)
	return nil
}

func MethodNotAllowedView(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(405)
	return nil
}

func ErrorView(ctx *atreugo.RequestCtx, _ error, _ int) {
	ctx.SetStatusCode(500)
}

func PanicView(ctx *atreugo.RequestCtx, _ interface{}) {
	_ = ctx.TextResponse("FATAL INTERNAL SERVER ERROR", 500)
}

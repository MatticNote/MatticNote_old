package view

import (
	"github.com/MatticNote/MatticNote/mnEmbed"
	"github.com/savsgio/atreugo/v11"
	"html/template"
	"log"
	"net/http"
)

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

func PanicView(ctx *atreugo.RequestCtx, errTrace interface{}) {
	log.Println("It has occurred panic error: ")
	log.Println(errTrace)
	_ = ctx.TextResponse("FATAL INTERNAL SERVER ERROR", 500)
}

func CSRFTokenErrorView(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	tmpl := template.Must(template.ParseFS(mnEmbed.Templates, "template/csrf_error.html"))
	_ = tmpl.Execute(w, nil)
}

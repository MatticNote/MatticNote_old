package view

import (
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/savsgio/atreugo/v11"
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

func CSRFTokenErrorView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte(fmt.Sprintf("CSRF error: %v", csrf.FailureReason(r))))
}

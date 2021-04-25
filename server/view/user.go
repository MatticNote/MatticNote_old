package view

import (
	mnEmbed "github.com/MatticNote/MatticNote/embed"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
)

type (
	InternalSignup     struct{}
	InternalSignupPost struct{}
)

func (is InternalSignup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(mnEmbed.Templates, "template/signup.html"))
	_ = tmpl.Execute(w, map[string]interface{}{
		"_CSRF": csrf.TemplateField(r),
	})
}

func (isp InternalSignupPost) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("POST"))
}

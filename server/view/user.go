package view

import (
	mnEmbed "github.com/MatticNote/MatticNote/mnEmbed"
	"github.com/MatticNote/MatticNote/server/internal"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
)

func writeSignupForm(w *http.ResponseWriter, r *http.Request, errorMsg interface{}) {
	tmpl := template.Must(template.ParseFS(mnEmbed.Templates, "template/signup.html"))
	_ = tmpl.Execute(*w, map[string]interface{}{
		"_CSRF": csrf.TemplateField(r),
		"error": errorMsg,
	})
}

func InternalSignup(w http.ResponseWriter, r *http.Request) {
	writeSignupForm(&w, r, nil)
}

func InternalSignupPost(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if username == "" || email == "" || password == "" {
		writeSignupForm(&w, r, "Invalid form")
		return
	}

	err := internal.CreateUser(
		email,
		username,
		password,
	)

	if err != nil {
		if err == internal.ErrAccountAlreadyExists {
			writeSignupForm(&w, r, err.Error())
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}

	_, _ = w.Write([]byte("POST"))
}

package view

import (
	"context"
	"database/sql"
	"github.com/MatticNote/MatticNote/db"
	"github.com/MatticNote/MatticNote/mnEmbed"
	"github.com/MatticNote/MatticNote/server/internal"
	"github.com/gorilla/csrf"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
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

func writeLoginForm(w *http.ResponseWriter, r *http.Request, errorMsg interface{}) {
	tmpl := template.Must(template.ParseFS(mnEmbed.Templates, "template/login.html"))
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
		false,
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

func InternalLogin(w http.ResponseWriter, r *http.Request) {
	writeLoginForm(&w, r, nil)
}

func InternalLoginPost(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := []byte(r.PostFormValue("password"))

	var targetUser struct {
		uuid           string
		email          sql.NullString
		username       string
		password       sql.NullString
		isMailVerified bool
		isActive       bool
		isSuspend      bool
		u2fIsEnable    sql.NullBool
	}

	err := db.DB.QueryRow(
		context.Background(),
		"SELECT \"user\".uuid, email, username, password, is_mail_verified, is_active, is_suspend, u2f.is_enable FROM \"user\""+
			" LEFT JOIN user_2fa u2f on \"user\".two_fa = u2f.uuid WHERE (email ILIKE $1 OR username ILIKE $1) AND host IS NULL",
		login,
	).Scan(
		&targetUser.uuid,
		&targetUser.email,
		&targetUser.username,
		&targetUser.password,
		&targetUser.isMailVerified,
		&targetUser.isActive,
		&targetUser.isSuspend,
		&targetUser.u2fIsEnable,
	)
	if err != nil && err == pgx.ErrNoRows {
		writeLoginForm(&w, r, "invalid login or password")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Database Connection error"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(targetUser.password.String), password)
	if err != nil {
		writeLoginForm(&w, r, "invalid login or password")
		return
	}

	if !targetUser.isActive {
		writeLoginForm(&w, r, "Account is no longer available.")
		return
	}

	if targetUser.isSuspend {
		writeLoginForm(&w, r, "Account is suspended.")
		return
	}

	if !targetUser.isMailVerified {
		writeLoginForm(&w, r, "Email is not verified yet. Please verify first")
		return
	}

	if targetUser.u2fIsEnable.Valid && targetUser.u2fIsEnable.Bool {
		// TODO: 2fa authentication
	} else {
		// TODO: Login system
		_, _ = w.Write([]byte("Login OK"))
		return
	}
}

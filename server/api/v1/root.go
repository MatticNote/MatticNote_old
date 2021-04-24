package apiV1

import (
	"github.com/savsgio/atreugo/v11"
)

func ConfigureRouteV1(r *atreugo.Router) {
	r.GET("/user/{uuid}", GetEntryUser)
}

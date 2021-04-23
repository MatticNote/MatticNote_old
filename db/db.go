package db

import (
	"embed"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/lib/pq"
	"net/http"
)

//go:embed migrations/*.sql
var migrations embed.FS

func GetMigrateInstance() (source.Driver, error) {
	hfs, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return nil, err
	}
	return hfs, nil
}

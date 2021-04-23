package db

import (
	"context"
	"embed"
	"fmt"
	"github.com/MatticNote/MatticNote/config"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"net/http"
)

//go:embed migrations/*.sql
var migrations embed.FS

var DB *pgxpool.Pool

func GetMigrateInstance() (source.Driver, error) {
	hfs, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return nil, err
	}
	return hfs, nil
}

func InitDB(cfg *config.MatticNoteConfig) error {
	pgx, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s pool_max_conns=%d",
			cfg.Database.Address,
			cfg.Database.Port,
			cfg.Database.Name,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Sslmode,
			cfg.Database.MaxConnect,
		),
	)

	if err == nil {
		DB = pgx
	}

	return err
}

func CloseDB() {
	DB.Close()
}

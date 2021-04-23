package main

import (
	"database/sql"
	"fmt"
	"github.com/MatticNote/MatticNote/config"
	"github.com/MatticNote/MatticNote/db"
	"github.com/MatticNote/MatticNote/server"
	"github.com/MatticNote/MatticNote/server/view"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/savsgio/atreugo/v11"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	DefaultPort = 3000
	DefaultAddr = "127.0.0.1"
)

var mnAppCli = &cli.App{
	Name:        "MatticNote",
	Description: "ActivityPub compatible SNS that aims to be easy for everyone to use",
	Commands: []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start server",
			Action:  startServer,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:        "port",
					Usage:       "Specifies the port number for listening to the server",
					Aliases:     []string{"p"},
					EnvVars:     []string{"MN_PORT"},
					Value:       DefaultPort,
					DefaultText: "3000",
				},
				&cli.StringFlag{
					Name:        "address",
					Usage:       "Specified the address for listening to the server",
					Aliases:     []string{"a"},
					EnvVars:     []string{"MN_ADDR"},
					Value:       DefaultAddr,
					DefaultText: "127.0.0.1",
				},
				&cli.BoolFlag{
					Name:    "skip-migration",
					Usage:   "Start the server without the migration process. Specify when all migrations are applicable.",
					Aliases: []string{"m"},
					EnvVars: []string{"MN_SKIP_MIGRATION"},
				},
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create configuration file",
			Action:  initConfig,
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Migrate the database",
			Action:  migrateAction,
		},
	},
}

func main() {
	if err := mnAppCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer(c *cli.Context) error {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		return err
	}
	if !c.Bool("skip-migration") {
		log.Print("Migrating database")
		if err := processMigration(cfg); err != nil {
			if err.Error() != "no change" { // Ignore errors due to no change
				return err
			}
		}
	}

	var (
		addr     = c.String("address")
		addrPort = c.Uint("port")
	)

	if c.String("address") == DefaultAddr {
		addr = cfg.Server.Address
	}
	if c.Uint("port") == DefaultPort {
		addrPort = cfg.Server.Port
	}

	mnServer := atreugo.New(atreugo.Config{
		Name:                 "MatticNote",
		Addr:                 fmt.Sprintf("%s:%d", addr, addrPort),
		NotFoundView:         view.NotFoundErrorView,
		MethodNotAllowedView: view.MethodNotAllowedView,
		ErrorView:            view.ErrorView,
		PanicView:            view.PanicView,
	})
	server.ConfigureRoute(mnServer)
	return mnServer.ListenAndServe()
}

func initConfig(_ *cli.Context) error {
	// WIP
	return nil
}

func migrateAction(_ *cli.Context) error {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		return err
	}

	return processMigration(cfg)
}

func processMigration(cfg *config.MatticNoteConfig) error {
	dbMigrations, err := db.GetMigrateInstance()
	if err != nil {
		return err
	}

	dbCon, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			cfg.Database.Address,
			cfg.Database.Port,
			cfg.Database.Name,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Sslmode,
		),
	)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(dbCon, &postgres.Config{})
	if err != nil {
		return err
	}

	dbMigrate, err := migrate.NewWithInstance(
		"httpfs", dbMigrations,
		"postgresql", driver)
	if err != nil {
		return err
	}

	err = dbMigrate.Up()
	if err != nil {
		return err
	}

	err, err2 := dbMigrate.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}

	return nil
}

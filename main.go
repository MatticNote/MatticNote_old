package main

import (
	"fmt"
	"github.com/MatticNote/MatticNote/server"
	"github.com/MatticNote/MatticNote/server/view"
	"github.com/savsgio/atreugo/v11"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
					Value:       3000,
					DefaultText: "3000",
				},
				&cli.StringFlag{
					Name:        "address",
					Usage:       "Specified the address for listening to the server",
					Aliases:     []string{"a"},
					EnvVars:     []string{"MN_ADDR"},
					Value:       "127.0.0.1",
					DefaultText: "127.0.0.1",
				},
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create configuration file",
			Action:  initConfig,
		},
	},
}

func main() {
	if err := mnAppCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer(c *cli.Context) error {
	mnServer := atreugo.New(atreugo.Config{
		Name:                 "MatticNote",
		Addr:                 fmt.Sprintf("%s:%d", c.String("address"), c.Uint("port")),
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

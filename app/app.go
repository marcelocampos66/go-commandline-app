package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func getIps(context *cli.Context) {
	host := context.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(context *cli.Context) {
	host := context.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

// Generate will be return the linecommand app ready to be executed
func Generate() *cli.App {
	var app *cli.App = cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search IPs and Names of internet servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Search the ip of a web address",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "servers",
			Usage:  "get servers name on web",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

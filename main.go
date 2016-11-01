package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func cmdShare(c *cli.Context) error {
	host := c.String("b")
	sp := c.Int("sp")
	jp := c.Int("jp")
	token := c.String("t")

	if len(token) == 0 {
		token = generateToken()
	}

	serve(host, sp, jp)
	share(host, sp, token)
	return nil
}

func cmdJoin(c *cli.Context) error {
	host := c.String("s")
	port := c.Int("p")
	token := c.String("t")

	if len(token) == 0 {
		fmt.Println("ERROR:", "session token is required")
		return nil
	}

	join(host, port, token)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "raptor"
	app.Usage = "share terminal sessions with ease"

	app.Commands = []cli.Command{
		{
			Name:    "share",
			Aliases: []string{"s"},
			Usage:   "share a terminal session",
			Action:  cmdShare,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "b", Value: "0.0.0.0", Usage: "bind address for listeners"},
				cli.IntFlag{Name: "sp", Value: 8080, Usage: "listen on port for sharing requests"},
				cli.IntFlag{Name: "jp", Value: 8081, Usage: "listen on port for joining requests"},
				cli.StringFlag{Name: "t", Value: "", Usage: "set session token"},
			},
		},
		{
			Name:    "join",
			Aliases: []string{"j"},
			Usage:   "join a shared session",
			Action:  cmdJoin,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "s", Value: "0.0.0.0", Usage: "address to connect to"},
				cli.IntFlag{Name: "p", Value: 8081, Usage: "port to connect to"},
				cli.StringFlag{Name: "t", Value: "", Usage: "session token"},
			},
		},
	}

	app.Run(os.Args)
}

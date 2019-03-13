package main

import (
	"log"
	"os"

	"github.com/samlazrak/networkcli/lib"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Net Probe CLI"
	app.Usage = "Query IP, CNAMEs, MX records, and Name Servers"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "samlazrak.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Look up name servers of a host.",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				lib.Nameserver(c.String("host"))
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Look up IP addresses of a host.",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				lib.IP(c.String("host"))
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Look up CNAME of a host.",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				lib.CNAME(c.String("host"))
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Look up MX records of a host.",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				lib.MX(c.String("host"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

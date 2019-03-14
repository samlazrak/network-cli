package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	app := cli.NewApp()
	app.Name = "Octoprint Go-Cli"
	app.Usage = "Control Octoprint and your 3D printer from the command line"

	Flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "octopi.local",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "auth",
			Usage: "Supply auth api key",
			Flags: Flags,
			Action: func(c *cli.Context) error {

				err := ioutil.WriteFile("authkey.txt", []byte(c.String("host")), 0644)
				if err != nil {
					panic(err)
				}

				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
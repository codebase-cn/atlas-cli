package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

// Declare package level vairables

func main() {

	// Start the "real" program

	app := cli.NewApp()
	app.Name = "atlas-cli"
	app.Usage = "Atlas commandline API!"
	app.Version = "0.0.2-alpha"
	app.Action = func(c *cli.Context) {
		fmt.Println("Nothing to do.  Try `help` or `-h` to see what's possible.")
	}

	atlas_user := &Credentials{}
	atlas_user.Login()

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Provide your login/API credentials",
			Action: func(c *cli.Context) {
				atlas_user.Query()
				atlas_user.Save()
			},
		},
		{
			Name:  "whoami",
			Usage: "Display your login/API credentials",
			Action: func(c *cli.Context) {
				fmt.Printf("You are %s.\n", atlas_user.User)
			},
		},
		{
			Name:  "build",
			Usage: "Build a project",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "pdf",
					Usage: " build a pdf",
				},
				cli.BoolFlag{
					Name:  "html",
					Usage: " build html format",
				},
				cli.BoolFlag{
					Name:  "epub",
					Usage: " build epub format",
				},
				cli.BoolFlag{
					Name:  "mobi",
					Usage: " build mobi format",
				},
				cli.StringFlag{
					Name:  "branch, b",
					Value: "master",
					Usage: "branch to build",
				},
			},
			Action: func(c *cli.Context) {
				args := &BuildArgs{}
				args.Parse(c)
				build := &Builds{}
				build.Build(*atlas_user, *args)
			},
		},
	}

	app.Run(os.Args)

}

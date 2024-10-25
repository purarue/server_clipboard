package main

import (
	"fmt"
	"github.com/purarue/server_clipboard"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func checkServerAddress(addr string) string {
	if strings.TrimSpace(addr) == "" {
		log.Fatalln("--server_address or envvar $CLIPBOARD_ADDRESS is required")
	}
	return addr
}

func main() {
	app := &cli.App{
		Name:  "server_clipboard",
		Usage: "share clipboard between devices using a server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   5025,
				Usage:   "port to listen on",
				EnvVars: []string{"CLIPBOARD_PORT"},
			},
			&cli.StringFlag{
				Name:     "password",
				Value:    "",
				Usage:    "password to use",
				Required: true,
				EnvVars:  []string{"CLIPBOARD_PASSWORD"},
			},
			&cli.StringFlag{
				Name:     "server_address",
				Usage:    "server address to connect to",
				Required: false,
				EnvVars:  []string{"CLIPBOARD_ADDRESS"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "start server",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "debug",
						Aliases: []string{"d"},
						Value:   false,
						Usage:   "enable debug logging",
					},
					&cli.IntFlag{
						Name: "clear-after",
						// 0 means never clear
						Aliases: []string{"c"},
						Value:   0,
						Usage:   "clear clipboard after this many seconds [0 means never clear]",
					},
				},
				Action: func(c *cli.Context) error {
					return server_clipboard.Server(c.String("password"), c.Int("port"), c.Bool("debug"), c.Int("clear-after"))
				},
			},
			{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "copy to server clipboard",
				Action: func(c *cli.Context) error {
					text, err := server_clipboard.Copy(c.String("password"), checkServerAddress(c.String("server_address")), server_clipboard.FetchClipboard(c.String("clipboard")))
					if err != nil {
						return err
					}
					if strings.TrimSpace(text) != "" {
						fmt.Fprintln(os.Stderr, text)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "clipboard",
						EnvVars:  []string{"CLIPBOARD_CONTENTS"},
						Usage:    "clipboard data to upload to server",
						Required: false,
					},
				},
			},
			{
				Name:    "paste",
				Aliases: []string{"p"},
				Usage:   "paste from server clipboard",
				Action: func(c *cli.Context) error {
					text, err := server_clipboard.Paste(c.String("password"), checkServerAddress(c.String("server_address")))
					if err != nil {
						return err
					}

					if strings.TrimSpace(text) != "" {
						err := server_clipboard.SetClipboard(text)
						if err != nil {
							// if we have text, print text regardless of if there was an error
							fmt.Println(text)
							return err
						} else {
							fmt.Fprintln(os.Stderr, "pasted into local clipboard")
						}
					} else {
						fmt.Fprintln(os.Stderr, "server returned empty clipboard")
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

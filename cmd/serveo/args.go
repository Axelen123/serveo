package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Axelen123/serveo"
	"github.com/Axelen123/serveo/internal"
	"github.com/akamensky/argparse"
)

func parse() *serveo.Args {
	parser := argparse.NewParser("serveo", "Unofficial client for serveo.net")

	usr, err := user.Current()
	if err != nil {
		internal.Error("cannot get user", err)
	}

	config := parser.String("c", "config", &argparse.Options{Required: false, Help: "Custom config path", Default: serveo.ConfigName})
	domain := parser.String("d", "domain", &argparse.Options{Required: false, Help: "Sets domain/alias. Usage: -d myalias or --domain mydomain.com", Default: usr.Username})

	init := parser.NewCommand("init", "Generate config file")

	ssh := parser.NewCommand("ssh", "Expose SSH Server")

	http := parser.NewCommand("http", "Expose HTTP Server")
	port := http.Int("p", "port", &argparse.Options{Required: false, Help: "Set port", Default: 80})
	args := os.Args
	// Append empty string if os.Args only contains the command name.
	// This is because argparse will give [sub]command required error otherwise.
	if len(args) == 1 {
		args = append(args, "")
	}
	err = parser.Parse(args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	return &serveo.Args{Commands: serveo.Commands{Init: init.Happened(), SSH: ssh.Happened(), HTTP: http.Happened()}, Flags: serveo.Flags{Config: *config, Port: *port, Domain: *domain}}
}

package main

import (
	"fmt"
	"os"
	"os/user"

	serveo "github.com/Axelen123/serveo/pkg"
	internal "github.com/Axelen123/serveo/internal"
	"github.com/akamensky/argparse"
)

func parse() *serveo.Args {
	parser := argparse.NewParser("serveo", "Unofficial client for serveo.net")

	usr, err := user.Current()
	if err != nil {
		internal.Error(err)
	}

	config := parser.String("c", "config", &argparse.Options{Required: false, Help: "Custom config path", Default: serveo.ConfigName})
	domain := parser.String("d", "domain", &argparse.Options{Required: false, Help: "Sets domain/alias. Usage: -d myalias or --domain mydomain.com", Default: usr.Username})

	init := parser.NewCommand("init", "Generate config file")

	ssh := parser.NewCommand("ssh", "Expose SSH")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	return &serveo.Args{Commands: serveo.Commands{Init: init.Happened(), SSH: ssh.Happened()}, Flags: serveo.Flags{Config: *config, Domain: *domain}}}
}
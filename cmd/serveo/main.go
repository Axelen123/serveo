package main

import (
	"fmt"
	"os/user"

	"github.com/Axelen123/serveo"
	"github.com/Axelen123/serveo/internal"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	args := parse()
	if args.Commands.Init {
		fmt.Println("Generating config file...")
		err := genConf(args.Flags.Config)
		if err != nil {
			internal.Error("Error while generating config file", err)
		}
		fmt.Println("Successfully generated config file")
		return
	} else if args.Commands.SSH {
		serveo.Start(&serveo.Config{HTTP: -1, SSH: true, Domain: args.Flags.Domain, TCP: []serveo.TCP{}})
	} else if args.Commands.HTTP {
		serveo.Start(&serveo.Config{HTTP: args.Flags.Port, SSH: false, Domain: args.Flags.Domain, TCP: []serveo.TCP{}})
	}
	conf, err := serveo.GetConfig(args.Flags.Config)
	if err != nil {
		internal.Error("Error while parsing/reading config", err)
	}
	if conf.Domain == "" {
		usr, err := user.Current()
		if err != nil {
			internal.Error("Error while getting user", err)
		}
		conf.Domain = usr.Username
	}
	serveo.Start(conf)
}

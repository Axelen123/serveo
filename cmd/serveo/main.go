package main

import (
	"fmt"
	"os/user"

	internal "github.com/Axelen123/serveo/internal"
	serveo "github.com/Axelen123/serveo/pkg"
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
		fmt.Println("Exposing SSH server...")
		
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

package serveo

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/Axelen123/serveo/internal"
)

// ConfigName contains the name of the config file
const ConfigName = "serveo.config.json"

// Start connects to serveo.net and starts forwarding with specified config
func Start(config *Config) {
	forwards := config.TCP

	if config.SSH {
		forwards = append(forwards, TCP{Local: Endpoint{Host: "localhost", Port: 22}, Remote: Endpoint{Host: config.Domain, Port: 22}})
	}

	if config.HTTP != -1 {
		forwards = append(forwards, TCP{Local: Endpoint{Host: "localhost", Port: config.HTTP}, Remote: Endpoint{Host: config.Domain, Port: 80}})
	}

	if len(forwards) < 1 {
		internal.Error("error while trying to start", errors.New("serveo: nothing to forward"))
	}
	args := []string{}
	for _, f := range forwards {
		args = append(args, "-R", forwardStr(f))
	}
	args = append(args, "serveo.net")
	cmd := exec.Command("ssh", args...)
	var waitStatus syscall.WaitStatus
	// TODO: add log output
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			internal.Error("error while connecting to serveo.net", fmt.Errorf("exit status %d", waitStatus.ExitStatus()))
		}
		if err != nil {
			internal.Error("error while connecting to serveo.net", err)
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		if waitStatus.ExitStatus() != 0 {
			internal.Error("cannot connect to serveo.net", fmt.Errorf("exit status %d", waitStatus.ExitStatus()))
		}
		fmt.Println("Done")
	}
}

func forwardStr(f TCP) string {

	if f.Remote.Port == 80 || f.Remote.Port == 443 {
		// TODO: add http output parsing
	}

	local := f.Local.String()
	remote := f.Remote.String()

	return fmt.Sprintf("%s:%s", remote, local)
}

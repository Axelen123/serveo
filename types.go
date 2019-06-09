package serveo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// Config holds the config
type Config struct {
	HTTP     int        `json:"http"`
	SSH      bool       `json:"ssh"`
	Domain   string     `json:"domain"`
	Forwards []Forwards `json:"forwards"`
}

// GetConfig reads and marshals the config file
func GetConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)

	conf := new(Config)
	err = json.Unmarshal(data, conf)
	return conf, err
}

// Endpoint holds endpoint information
type Endpoint struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (e *Endpoint) String() string {
	return e.Host + ":" + strconv.Itoa(e.Port)
}

// Forwards holds JSON data for forwarded ports
type Forwards struct {
	Local  Endpoint `json:"local"`
	Remote Endpoint `json:"remote"`
}

// Commands holds parsed commands
type Commands struct {
	Init bool
	SSH bool
}

// Flags holds parsed flags
type Flags struct {
	Config string
	Domain string
}

// Args holds parsed command line arguments
type Args struct {
	Commands Commands
	Flags    Flags
}

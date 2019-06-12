package serveo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

// Config holds the config
type Config struct {
	HTTP   int    `json:"http"`
	SSH    bool   `json:"ssh"`
	Domain string `json:"domain"`
	TCP    []TCP  `json:"tcp"`
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
	if err != nil {
		return nil, err
	}
	for i, f := range conf.TCP {
		if f.Local.Host == "" {
			conf.TCP[i].Local.Host = "localhost"
		}
	}
	return conf, nil
}

// Endpoint holds endpoint information
type Endpoint struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (e Endpoint) String() string {
	port := strconv.Itoa(e.Port)

	if e.Host != "" {
		return e.Host + ":" + port
	}
	return port
}

// TCP holds JSON data for forwarded ports
type TCP struct {
	Local  Endpoint `json:"local"`
	Remote Endpoint `json:"remote"`
}

// Commands holds parsed commands
type Commands struct {
	Init bool
	SSH  bool
	HTTP bool
}

// Flags holds parsed flags
type Flags struct {
	Config string
	Domain string
	Port   int
}

// Args holds parsed command line arguments
type Args struct {
	Commands Commands
	Flags    Flags
}

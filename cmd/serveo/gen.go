package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Axelen123/serveo"
)

func genConf(path string) error {
	data, err := json.MarshalIndent(serveo.Config{SSH: false, HTTP: 80, Domain: "", TCP: []serveo.TCP{}}, "", "\t")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return ioutil.WriteFile(path, data, 0644)
}

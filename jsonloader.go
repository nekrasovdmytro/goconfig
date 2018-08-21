package goconfig

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfigLoader struct {}

func (l JsonConfigLoader) load(file string) map[string]string {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var result map[string]string

	err = json.Unmarshal(data, &result)

	if err != nil {
		panic(err)
	}

	return result
}

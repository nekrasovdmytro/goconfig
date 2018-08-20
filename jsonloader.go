package goconfig

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfigLoader struct {}

func (l JsonConfigLoader) load(file string) interface{} {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var result interface{}

	json.Unmarshal(data, result)

	return result
}

package main

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFile() []Serve {
	data, _ := ioutil.ReadFile("exercises1.json")
	var ser []Serve
	json.Unmarshal([]byte(data), &ser)
	return ser
}

package main

import "strings"

func filter(arrServe []Serve) []Serve {
	var data []Serve
	for i := range arrServe {
		if strings.Contains(strings.ToLower(arrServe[i].Class), "admin") {
			data = append(data, arrServe[i])
		}
	}
	return data
}
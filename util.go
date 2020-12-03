package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func readFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(b), "\n")
}

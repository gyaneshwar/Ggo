package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type JsonType struct {
	Delay time.Duration
}

func main() {
	var jt JsonType
	d, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(d, &jt)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(jt.Delay)
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Configuration Config

type Config struct {
	Port      int64  `json:"port"`
	Domain    string `json:"domain"`
	KeyLenght int    `json:"keylenght"`
}

var ConfDummy = Config{
	Port:      5000,
	Domain:    "localhost",
	KeyLenght: 8,
}

func config() {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		Configuration = ConfDummy
		create, err := os.Create("config.json")
		handleError(err)
		marshal, err := json.MarshalIndent(ConfDummy, "", "  ")
		handleError(err)
		_, err = create.Write(marshal)
		handleError(err)
		err = create.Close()
		handleError(err)
		return
	}
	err = json.Unmarshal(raw, &Configuration)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

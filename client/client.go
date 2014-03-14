package main

import (
	"github.com/ggitau/mdapi"
	"log"
	"os"
)

func main() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, err := mdapi.NewMdcli("tcp://localhost:5555", verbose)
	if err != nil {
		log.Panic(err.Error())
	}
	reply, err := session.Send("reverse", "hello", "world")
	if err != nil {
		log.Panic(err.Error())
	}
	for _, val := range reply {
		log.Println(val)
	}
}

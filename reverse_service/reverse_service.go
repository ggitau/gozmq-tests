package main

import (
	"github.com/ggitau/commands"
	"github.com/ggitau/mdapi"
	"log"
	"os"
)

type ReverseService struct {
	Name string
}

var commands map[string]*Command

func init() {
	commands = make(map[string]*Command)
	commands["reverse"] = &ReverseCommand{Name: "reverse"}
}
func (r *ReverseService) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (r *ReverseService) Start() {
	var verbose bool
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		verbose = true
	}
	session, _ := mdapi.NewMdwrk("tcp://localhost:5555", "reverse", verbose)

	var err error
	var request, reply []string

	for {
		request, err = session.Recv(reply)
		if err != nil {
			break //  Worker was interrupted
		}

		if verbose {
			log.Println("Reversing..")
			log.Println(request)
		}

		for i := 0; i < len(request); i++ {
			request[i] = r.Reverse(request[i])
		}

		reply = request //  Echo is complex... :-)
	}

	log.Println(err)
}

func main() {
	service := &ReverseService{Name: "reverse"}
	service.Start()
}

package main

import (
	"github.com/boggydigital/clo"
	"github.com/boggydigital/glo/cmd"
	"log"
	"os"
)

func main() {

	defs, err := clo.LoadDefinitions("clo.json")
	if err != nil {
		log.Fatal("error loading definitions: ", err.Error())
	}

	req, err := clo.Parse(os.Args[1:], defs)
	if err != nil {
		log.Fatal("error parsing command-line objectives: ", err.Error())
	}

	if err := cmd.Route(req, defs); err != nil {
		log.Fatal("error routing request to handler: ", err.Error())
	}
}

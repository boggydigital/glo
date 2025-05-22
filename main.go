package main

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/glo/cli"
	"github.com/boggydigital/nod"
	"log"
	"os"
)

var (
	//go:embed "cli-commands.txt"
	cliCommands []byte
	//go:embed "cli-help.txt"
	cliHelp []byte
)

func main() {

	nod.EnableStdOutPresenter()

	ns := nod.Begin("grams (g) to pounds (lb) and ounces (oz)")
	defer ns.Done()

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"convert": cli.ConvertHandler,
	})

	if err = defs.AssertCommandsHaveHandlers(); err != nil {
		log.Fatalln(err)
	}

	if err = defs.Serve(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}

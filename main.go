package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		aliasName = flag.String("a", "", "alias shortcut")
		cmdName   = flag.String("s", "", "alias command name")
	)
	flag.Parse()
	args := os.Args
	switch {
	case (len(args) < 2) || (args[1] == "-h"):
		flag.Usage()
		os.Exit(1)
	case len(*aliasName) == 0 || len(*cmdName) == 0:
		flag.Usage()
		os.Exit(1)
	}
	fmt.Printf("Creating alias as: %s, from:%s\n", *aliasName, *cmdName)
	err := alias{command: *cmdName, short: *aliasName}.Create()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// need to fix this
	fmt.Println("Refreshing bash shell ENV\nYou shall open a new shell window to use your new alias")
	err = exec.Command("source", "~/.bash_profile").Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
	os.Exit(0)
}

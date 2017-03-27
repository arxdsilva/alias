package main

import (
	"flag"
	"fmt"
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
	err := alaias(*aliasName, *cmdName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// need to fix this
	fmt.Println("Refreshing bash shell ENV")
	cmd := exec.Command("source", "~/.bash_profile")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

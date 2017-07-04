package main

import (
	"fmt"
	"os"
	"os/user"
)

type alias struct {
	command string
	short   string
}

func (a alias) Create() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	err = os.Chdir(usr.HomeDir)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(".bash_profile", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("\nalias %s=\"%s\"", a.short, a.command))
	return err
}

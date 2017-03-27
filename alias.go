package main

import (
	"fmt"
	"os"
	"os/user"
)

func alaias(aliasName, cmdName string) error {
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
	text := fmt.Sprintf("\nalias %s=\"%s\"", aliasName, cmdName)
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

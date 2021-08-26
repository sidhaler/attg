package main

import (
	"fmt"
	"log"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/sidhaler/attg/Util"

	at "github.com/sidhaler/attg/attconf"
	"github.com/spf13/viper"
)

var cfpath string = "./attg.toml"

func main() {
	perms, err := os.Stat(cfpath)
	// Just to make sure everything will work
	if err != nil {
		fmt.Println("Configuration files wasn't found, creating new one....")
		os.Create(cfpath)
		fmt.Println("Now re-run application")
		os.Exit(00)
	}
	fmt.Print("\033[H\033[2J")
	mode := perms.Mode()
	fel, err := os.OpenFile(cfpath, os.O_RDWR, mode)
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigType("toml")
	err = viper.ReadConfig(fel)
	if err != nil {
		log.Fatal(err)
	}
	var e at.Atcfg
	e.Getconf()
	e.Warns()
	fmt.Println(e)

	t := prompt.New(
		Util.ExeCommand,
		Util.Comp,
	)
	t.Run()
	fel.Close()

}

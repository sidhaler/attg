package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/alexflint/go-arg"
	"github.com/c-bata/go-prompt"
	"github.com/sidhaler/attg/Util"
	at "github.com/sidhaler/attg/attconf"
	"github.com/spf13/viper"
)

var args struct {
	Install bool   `arg:"-i" default:"false" help:"Moves app to binary folder, and creates config file"`
	SrcFile string `arg:"positional" `
	Copy    bool   `arg:"-c" default:"false" help:"Copy config"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	arg.MustParse(&args)
	if args.Copy {
		if args.SrcFile == "" {
			fmt.Println("Provide .toml file as argument")
		}

		at.CopyConfig(args.SrcFile)
	}

	if args.Install {
		kernel := runtime.GOOS
		switch kernel {
		case "linux":
			fmt.Println("Linux detected...")
			err := os.Rename(at.Defbinpath, at.Binpathlinux)
			check(err)
			err = os.Mkdir(at.FOLDERlinux, os.FileMode(os.O_RDWR))
			check(err)
			_, err = os.Create(at.Cfpathlinux)
			check(err)
		case "darwin":
			fmt.Println("Macos detected...")
			err := os.Rename(at.Defbinpath, at.Binpathdarwin)
			check(err)
			err = os.Mkdir(at.FOLDERdarwin, os.FileMode(os.O_RDWR))
			check(err)
			_, err = os.Create(at.Cfpathdarwin)
			check(err)
		}
		fmt.Println("OKAY BOSS, everything is on correct place, re-run app.")
		os.Exit(00)
	}
	var cfpath string
	kernel := runtime.GOOS
	if kernel == "darwin" {
		cfpath = at.Cfpathdarwin
	} else {
		cfpath = at.Cfpathlinux
	}

	perms, err := os.Stat(cfpath)
	// Just to make sure everything will work
	if err != nil {
		fmt.Println("Configuration files wasn't found, creating new one....")
		os.Create(cfpath)
		fmt.Println("Now re-run application, with -i[install] flag.")
		os.Exit(00)
	}
	fmt.Print("\033[H\033[2J")
	mode := perms.Mode()
	fel, err := os.OpenFile(cfpath, os.O_RDWR, mode)
	check(err)
	viper.SetConfigType("toml")
	err = viper.ReadConfig(fel)
	check(err)
	var e at.Atcfg
	e.Warns()
	fmt.Println(e)
	defer fel.Close()
	t := prompt.New(
		Util.ExeCommand,
		Util.Comp,
	)
	t.Run()

}

package Util

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	at "github.com/sidhaler/attg/attconf"
	"github.com/sidhaler/attg/dbUtil"
)

//
var cfpath string

// const (
// 	cfpathlinux  string = "/usr/bin/attg/attg.toml"
// 	cfpathdarwin string = "/usr/local/bin/attg/attg.toml"
// )

func Comp(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		// DATABASE calls
		{Text: "fetch", Description: "Gets data from db (kinda bugged)"},
		{Text: "add", Description: "Add data to table (in progress)"},
		{Text: "remove", Description: "Remove record from table (in progress)"},
		{Text: "edit", Description: "Edit record in table (in progress)"},
		{Text: "import", Description: "Import data from table (kinda bugged)"},
		{Text: "↓", Description: ""},
		// configuration calls, don't use "" while writing data
		{Text: "SETPASSWORD", Description: "Set password of db"},
		{Text: "SETUSERNAME", Description: "Set username of db"},
		{Text: "SETDATABASE", Description: "Set database to operate with"},
		{Text: "SETTABLE", Description: "Table to get data from"},
		{Text: "SETPORT", Description: "Default (mysql): 3306"},
		{Text: "shc", Description: "Show config file"},
		{Text: "↓", Description: ""},
		// other commands
		{Text: "new", Description: "Remove old configuration file and create new one"},
		{Text: "clear", Description: "Clears terminal ; )"},
		{Text: "testcon", Description: "Test connection"},
		{Text: "…………………………………………", Description: ""},
		//{Text: "install", Description: "Move application to binary dir."},
	}
}

var s at.Atcfg

func ExeCommand(str string) {
	str = strings.TrimSpace(str)
	block := strings.Split(str, " ")
	block[0] = strings.ToUpper(block[0])
	switch block[0] {
	// CONFIGURATION
	case "SETPASSWORD":
		if len(block) < 2 {
			fmt.Println("please set password after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "SETDATABASE":
		if len(block) < 2 {
			fmt.Println("please set database after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "SETUSERNAME":
		if len(block) < 2 {
			fmt.Println("please set username after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "SETTABLE":
		if len(block) < 2 {
			fmt.Println("please set table after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "SETHOST":
		if len(block) < 2 {
			fmt.Println("please set host after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "SETPORT":
		if len(block) < 2 {
			fmt.Println("please set port after 'space'")
			return
		}
		s.SetConfig(block[0], block[1])
		fmt.Println("Success")
	case "FETCH":
		if len(block) < 2 {
			dbUtil.Fetchall()
			return
		}
		if len(block) < 3 {
			dbUtil.FetchwithID(block[1])
		}
	case "IMPORT":
		if len(block) < 2 {
			block = append(block, "DATA.txt")
		}
		dbUtil.ImportAll(block[1])
	case "SHC":
		fmt.Println("Address =>", s.Getconf().Host)
		fmt.Println("Database =>", s.Getconf().Database)
		fmt.Println("Username =>", s.Getconf().Usr)
		fmt.Println("Password =>", s.Getconf().Passwd)
		fmt.Println("Port =>", s.Getconf().Port)
		fmt.Println("Table =>", s.Getconf().Table)
	case "TESTCON":
		dbUtil.TestOpen()
	case "CLEAR":
		/*
			cmd := exec.Command("clear")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		*/
		fmt.Print("\033[H\033[2J")
	case "NEW":
		kernel := runtime.GOOS
		switch kernel {
		case "linux":
			cfpath = at.Binpathlinux
		case "darwin":
			cfpath = at.Binpathdarwin
		default:
			cfpath = at.Cfpathlinux
		}
		os.Remove(cfpath)
		os.Create(cfpath)
		fmt.Println("New config file succesfully created!")
	default:
		return
	}
}

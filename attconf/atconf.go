package attconf

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var cfpath string = "./attg.toml"

type Configuration struct {
	Passwd, Database, Usr, Table, Host string
	Port                               int
}
type Atcfg struct {
	configuration Configuration
}

func (s *Atcfg) SetConfig(key string, value string) {
	switch key {
	case "SETDATABASE":
		viper.Set("DATABASE", value)
	case "SETPORT":
		val, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		viper.Set("PORT", val)
	case "SETTABLE":
		viper.Set("TABLE", value)
	case "SETHOST":
		viper.Set("HOST", value)
	case "SETPASSWORD":
		viper.Set("PASSWD", value)
	case "SETUSERNAME":
		viper.Set("USERNAME", value)
	}
	err := viper.WriteConfigAs(cfpath)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Atcfg) Getconf() Configuration {

	s.configuration.Database = viper.GetString("DATABASE")
	s.configuration.Passwd = viper.GetString("PASSWD")
	s.configuration.Usr = viper.GetString("USERNAME")
	s.configuration.Port = viper.GetInt("PORT")
	s.configuration.Table = viper.GetString("TABLE")
	s.configuration.Host = viper.GetString("HOST")

	if s.configuration.Port == 0 {
		s.configuration.Port = 3306
	}

	return s.configuration

}

func (s *Atcfg) Warns() {
	s.Getconf()
	if s.configuration.Host == "" {
		fmt.Println("Please provide host address")
	}
	if s.configuration.Database == "" {
		fmt.Println("Please provide db name")
	}
	if s.configuration.Passwd == "" {
		fmt.Println("Please provide password")
	}
	if s.configuration.Usr == "" {
		fmt.Println("Please provide username")
	}
	if s.configuration.Table == "" {
		fmt.Println("Please provide table")
	}

}

func (s *Atcfg) FatalWarns() {
	if s.configuration.Host == "" {
		fmt.Println("Please provide host address")
		os.Exit(01)
	}
	if s.configuration.Database == "" {
		fmt.Println("Please provide db name")
		os.Exit(01)
	}
	if s.configuration.Passwd == "" {
		fmt.Println("Please provide password")
		os.Exit(01)
	}
	if s.configuration.Usr == "" {
		fmt.Println("Please provide username")
		os.Exit(01)
	}
	if s.configuration.Table == "" {
		fmt.Println("Please provide table")
		os.Exit(01)
	}

}

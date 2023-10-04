package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

// TODO: add a config file option and load from it

type Config struct {
	port    int
	version bool
	help    bool
}

var _conf Config

func init() {
	flag.BoolVar(&_conf.help, "help", false, "get help")
	flag.IntVar(&_conf.port, "port", 8080, "the port to run the server on")
	flag.BoolVar(&_conf.version, "version", false, "get the current version of the server")
}

func loadConfig() Config {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalf("Unable to read build info for package")
	}
	flag.Parse()
	if _conf.help {
		flag.PrintDefaults()
		os.Exit(0)
	} else if _conf.version {
		fmt.Printf("%s %s (%s) build opts: [", info.Main.Path, info.Main.Version, info.GoVersion)
		for i, setting := range info.Settings {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s=%s", setting.Key, setting.Value)
		}
		fmt.Print("]\n")
		os.Exit(0)
	}
	return _conf
}

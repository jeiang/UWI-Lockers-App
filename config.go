package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/charmbracelet/log"
)

// TODO: add a config file option and load from it

type Config struct {
	addr string
}

var _rawFlags struct {
	addr     string
	help     bool
	logLevel string
	port     int
	version  bool
}

func init() {
	flag.StringVar(&_rawFlags.addr, "addr", "", "set the local address of the server that data will be served from")
	flag.StringVar(&_rawFlags.logLevel, "level", "", "how much to log")
	flag.BoolVar(&_rawFlags.help, "help", false, "get help")
	flag.IntVar(&_rawFlags.port, "port", 80, "the port to run the server on")
	flag.BoolVar(&_rawFlags.version, "version", false, "get the current version of the server")
}

func loadConfig() Config {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalf("Unable to read build info for package")
	}
	flag.Parse()
	if _rawFlags.help {
		flag.PrintDefaults()
		os.Exit(0)
	} else if _rawFlags.version {
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

	conf := Config{
		addr: fmt.Sprintf("%s:%d", _rawFlags.addr, _rawFlags.port),
	}

	var logLevel log.Level
	if _rawFlags.logLevel != "" {
		logLevel = log.ParseLevel(_rawFlags.logLevel)
	} else {
		lv, ok := os.LookupEnv("LOG")
		if ok {
			logLevel = log.ParseLevel(lv)
		}
	}
	logger := log.NewWithOptions(os.Stdout, log.Options{Level: logLevel})
	log.SetDefault(logger)

	return conf
}

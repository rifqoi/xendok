package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg Config

func ProcessArgs() Args {
	var arg Args
	f := flag.NewFlagSet("Server usage", 1)
	f.StringVar(&arg.ConfigPath, "c", "", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		fmt.Fprintln(f.Output())
	}

	f.Parse(os.Args[1:])
	if arg.ConfigPath == "" {
		fmt.Println("Usage of Server usage:")
		f.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(arg.ConfigPath); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("no such file or directory: %s", arg.ConfigPath))
		os.Exit(1)

	}

	return arg
}

func Init(args Args) error {
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Get() Config {
	return cfg
}

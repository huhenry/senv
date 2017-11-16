package main

import (
	"flag"
	"strings"
	"os"
	"fmt"
	"github.com/jamowei/senv/senv"
)

const (
	host_default    string = "127.0.0.1"
	port_default    string = "8888"
	name_default    string = "application"
	profile_default string = "default"
	label_default   string = "master"
)

var (
	host    string
	port    string
	name    string
	profile string
	label   string
)

func main() {
	params()
	cfg := senv.NewConfig(host, port, name, profile, label,
		&EnvKeyFormatter{"_", true},
		&EnvValFormatter{})
	cfg.Fetch()
	cfg.Process()
	for k, v := range cfg.Properties {
		fmt.Println(k, "=", v)
	}
}

func params() {
	flag.StringVar(&host, "host", host_default, "host")
	flag.StringVar(&port, "port", port_default, "port")
	flag.StringVar(&name, "name", name_default, "name")
	flag.StringVar(&profile, "profile", profile_default, "profile")
	flag.StringVar(&label, "label", label_default, "port")
	if help := flag.Arg(0); strings.HasSuffix(help, "h") || strings.HasSuffix(help, "help") {
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()
}

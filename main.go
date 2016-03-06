package main

import (
	"flag"
	"fmt"
)

var (
	sock *string
	port *string
)

func main() {
	port = flag.String("p", "8081", "listen port")
	sock = flag.String("sock", "unix:///var/run/docker.sock", "docker sock position")
	ServerMod := flag.Bool("servermod", true, "open server module")
	FileMod := flag.Bool("filemod", false, "open file module")
	WatchContainers := flag.String("wc", "", "watch containers split name by ,")
	FileOutputPosition := flag.String("o", "/tmp/log.out", "choose output file name")
	flag.Parse()
	fmt.Println("launch server")

	if *ServerMod == true {
		// New thread run server
		go OpenServer()
	}

	if *FileMod == true {
		go FileModuleLoad()
	}
}

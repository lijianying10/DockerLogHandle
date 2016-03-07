package main

import (
	"DockerLogHandle/global"
	"DockerLogHandle/tcp"
	"DockerLogHandle/watch"
	"flag"
	"fmt"
)

func main() {
	port := flag.String("p", "8081", "listen port")
	sock := flag.String("sock", "unix:///var/run/docker.sock", "docker sock position")
	ServerMod := flag.Bool("servermod", true, "open server module")
	FileMod := flag.Bool("filemod", false, "open file module")
	WatchContainers := flag.String("wc", "", "watch containers split name by ,")
	FileOutputPosition := flag.String("o", "/tmp/log.out", "choose output file name")
	flag.Parse()
	global.Port = *port
	global.Sock = *sock
	fmt.Println("launch server")

	if *FileMod == true {
		// New thread watch containers
		c := make(chan bool)
		go watch.FileModuleLoad(*WatchContainers, *FileOutputPosition, c)
		res := <-c
		if res {
			fmt.Println("Finish serve")
		}
	}

	if *ServerMod == true {
		// New thread run server
		c := make(chan bool)
		go tcp.OpenServer(c)
		res := <-c
		if res {
			fmt.Println("tcp server closed")
		}
	}

}

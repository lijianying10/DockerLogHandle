package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
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
	//FileOutputPosition:=flag.String("o", ""
	flag.Parse()
	fmt.Println("launch server")

	if *ServerMod == true {
		// New thread run server
		go OpenServer()
	}

	if *FileMod == true {

	}
}

func OpenServer() {
	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		fmt.Println("err listen", err.Error())
		os.Exit(2)
	}

	for {
		conn, _ := ln.Accept()
		go ConnHandler(conn)
	}

}

// ConnHandler test godoc
func ConnHandler(conn net.Conn) {
	defer conn.Close()
	for {
		conn.Write([]byte("Input container what you want"))
		message, err := bufio.NewReader(conn).ReadString('\n')
		// handle client close connection
		if err != nil {
			fmt.Println(err.Error() + "Connection going to close , thread close!")
			conn.Close()
			return
		}
		fmt.Print("Message Received:", string(message))
		conn.Write([]byte("you are watching container :" + message))
		Dc(conn, message[:len(message)-2])
	}
}

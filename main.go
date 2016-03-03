package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("launch server")

	ln, err := net.Listen("tcp", ":8081")
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

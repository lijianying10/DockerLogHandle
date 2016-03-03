package main

import (
	"fmt"
	"io"
	"net"

	"github.com/samalba/dockerclient"
)

func Dc(conn net.Conn, id string) {
	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	r, err := docker.ContainerLogs(id, &dockerclient.LogOptions{
		Follow: true,
		Stdout: true,
		Stderr: true,
		Tail:   100,
	})
	if err != nil {
		conn.Write([]byte(err.Error()))
		conn.Close()
		return
	}

	_, err = io.Copy(conn, r)
	if err != nil {
		fmt.Println(err.Error() + "error output")
	}

}

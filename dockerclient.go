package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/samalba/dockerclient"
)

func DockerReadLog() io.ReadCloser {
	docker, err := dockerclient.NewDockerClient(*sock, nil)
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
		// Output Fatal log
		fmt.Println("Fatal Initial Docker client error")
		os.Exit(1)
		return nil
	}
	return r
}

// Dc return log stream to a socket connection
func Dc(conn net.Conn, id string) {
	r := DockerReadLog()
	_, err = io.Copy(conn, r)
	if err != nil {
		fmt.Println(err.Error() + "error output")
	}
}

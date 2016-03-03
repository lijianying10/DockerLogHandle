package main

import (
	"fmt"
	"io"
	"os"

	"github.com/samalba/dockerclient"
)

func main() {
	docker, err := dockerclient.NewDockerClient("unix:///tmp/docker.sock", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	r, err := docker.ContainerLogs("foo", &dockerclient.LogOptions{
		Follow: true,
		Stdout: true,
		Stderr: true,
		Tail:   100,
	})

	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println(err.Error() + "error output")
	}

}

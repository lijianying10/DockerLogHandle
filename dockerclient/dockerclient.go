package dockerclient

import (
	"DockerLogHandle/global"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/samalba/dockerclient"
)

func DockerReadLog(id string, opt dockerclient.LogOptions) io.ReadCloser {
	docker, err := dockerclient.NewDockerClient(global.Sock, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	r, err := docker.ContainerLogs(id, &opt)
	if err != nil {
		// Output Fatal log
		fmt.Println("Fatal Initial Docker client error")
		os.Exit(1)
		return nil
	}
	return r
}

// DataWriteToTCPConnection return log stream to a socket connection Write log via tcp connection
func DataWriteToTCPConnection(conn net.Conn, id string) {
	r := DockerReadLog(id, dockerclient.LogOptions{
		Follow: true,
		Stdout: true,
		Stderr: true,
		Tail:   100,
	})
	_, err := io.Copy(conn, r)
	if err != nil {
		fmt.Println(err.Error() + "error output")
	}
}

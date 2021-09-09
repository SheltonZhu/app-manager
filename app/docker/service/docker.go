package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
	"time"
)

var (
	timeout = time.Second * 60
)

func getCli() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}

func List(option types.ContainerListOptions) []types.Container {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()
	cli := getCli()
	containers, err := cli.ContainerList(ctx, option)
	if err != nil {
		panic(err)
	}
	return containers

}

func Stop(containerId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()
	cli := getCli()
	if err := cli.ContainerStop(ctx, containerId, &timeout); err != nil {
		panic(err)
	}

}

func Restart(containerId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()
	cli := getCli()
	if err := cli.ContainerRestart(ctx, containerId, &timeout); err != nil {
		panic(err)
	}
}

func logs(containerId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)
	defer cancel()
	cli := getCli()
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerId, options)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

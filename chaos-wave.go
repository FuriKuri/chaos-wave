package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewClient(
		"unix:///var/run/docker.sock",
		"v1.22",
		nil,
		map[string]string{"User-Agent": "chaos-wave-0.1"})
	if err != nil {
		panic(err)
	}

	for _, c := range containers(cli) {
		if c.State == "running" {
			fmt.Println(c.ID + " " + c.State + " " + c.Status)
		}
	}
}

func containers(cli *client.Client) []types.Container {
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	return containers
}

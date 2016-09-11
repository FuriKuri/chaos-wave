package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

const (
	imageName = "furikuri/chaos-wave"
)

func main() {
	cli, err := client.NewClient(
		"unix:///var/run/docker.sock",
		"v1.22",
		nil,
		map[string]string{"User-Agent": "chaos-wave-0.1"})
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Duration: %s\n", duration())
	fmt.Printf("Interval: %s\n", interval())
	endTime := time.Now().Add(duration())
	for endTime.After(time.Now()) {
		time.Sleep(interval())
		stopRandomContainer(cli, randomGen)
	}
}

func validContainer(c types.Container) bool {
	if c.State == "running" && c.Image != imageName {
		return true
	}
	return false
}

func stopRandomContainer(cli *client.Client, randomGen *rand.Rand) {
	var validContainers []types.Container
	for _, c := range containers(cli) {
		if validContainer(c) {
			validContainers = append(validContainers, c)
		}
	}
	if len(validContainers) == 0 {
		fmt.Println("No running container. Nothing to kill :(")
	}
	if len(validContainers) > 0 {
		containerToStop := validContainers[randomGen.Intn(len(validContainers))]
		fmt.Printf("Stop and remove container with name '%s' \n", containerToStop.Names[0])
		stopAndRemoveContainer(containerToStop, cli)
	}
}

func stopAndRemoveContainer(container types.Container, cli *client.Client) {
	cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})
}

func containers(cli *client.Client) []types.Container {
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	return containers
}

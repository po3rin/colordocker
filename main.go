package main

import (
	"colordocker/colorize"
	"colordocker/docker"
	"context"

	"github.com/docker/docker/api/types"
)

func main() {
	cli := docker.Client

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	color := colorize.New()
	color.Container(containers)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	color = colorize.New()
	color.Image(images)
	if err != nil {
		panic(err)
	}
}

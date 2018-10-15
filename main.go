package main

import (
	"colordocker/colorize"
	"colordocker/docker"
	"context"
	"flag"
	"fmt"

	"github.com/docker/docker/api/types"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments.")
		return
	}

	switch args[0] {
	case "ps":
		colorizeContainer()
	case "image":
		colorizeImage()
	default:
		fmt.Println("argument is invalid.")
	}
}

func colorizeContainer() {
	cli := docker.Client
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	colorize.Container(containers)
}

func colorizeImage() {
	cli := docker.Client
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	colorize.Image(images)
	if err != nil {
		panic(err)
	}
}

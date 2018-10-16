package main

import (
	"colordocker/colorize"
	"colordocker/docker"
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "colordocker"
	app.Usage = "colordocker is commnad line tool that color output docker command."
	app.Version = "0.0.1"
	app.Action = func(context *cli.Context) error {
		switch context.Args().Get(0) {
		case "ps":
			colorizeContainer()
		case "image":
			colorizeImage()
		case "":
			fmt.Println("argument is empty.")
		default:
			fmt.Println("argument is invalid.")
		}
		return nil
	}
	app.Run(os.Args)
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

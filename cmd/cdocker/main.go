package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/po3rin/colordocker"
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
		case "images":
			colorizeImage()
		case "":
			fmt.Println("msg: argument is empty.\nplease exec 'cdocker help'")
		default:
			fmt.Println("msg: argument is invalid.\nplease exec 'cdocker help'")
		}
		return nil
	}
	app.Run(os.Args)
}

func colorizeContainer() {
	cli := colordocker.Client
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	colordocker.Container(containers)
}

func colorizeImage() {
	cli := colordocker.Client
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	colordocker.Image(images)
	if err != nil {
		panic(err)
	}
}

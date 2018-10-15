package colorize

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

var (
	id      string
	names   string
	repos   string
	tag     string
	image   string
	created string
	status  string
	color   string
	size    string
	data    [][]string
)

const (
	evenLineColor = "32m"
	oddLineColor  = "36m"
)

func Container(containers []types.Container) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "NAMES", "IMAGE", "STATUS"})

	for i, container := range containers {

		if i%2 == 0 {
			color = evenLineColor
		} else {
			color = oddLineColor
		}

		id = fmt.Sprintf("\x1b[%s%s\x1b[0m", color, container.ID[:10])
		names = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			strings.Replace(strings.Replace(container.Names[0], " ", "", -1), "/", "", -1))
		image = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			container.Image)
		status = fmt.Sprintf(
			"\x1b[%s%v\x1b[0m",
			color,
			strings.Replace(container.Status, " ", "-", -1))

		data = append(data, []string{id, names, image, status})
	}

	table.AppendBulk(data)
	table.Render()
}

func Image(images []types.ImageSummary) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"REPOSITORY", "TAG", "IMAGE ID", "SIZE"})

	for i, image := range images {
		if i%2 == 0 {
			color = evenLineColor
		} else {
			color = oddLineColor
		}

		repoTags := strings.Split(strings.Replace(image.RepoTags[0], " ", "", -1), ":")

		repos = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			repoTags[0],
		)
		tag = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			repoTags[1],
		)
		id = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			image.ID[7:19])
		size = fmt.Sprintf(
			"\x1b[%s%s\x1b[0m",
			color,
			strings.Replace(humanize.Bytes(uint64(image.Size)), " ", "", -1))

		data = append(data, []string{repos, tag, id, size})
	}

	table.AppendBulk(data)
	table.Render()
}

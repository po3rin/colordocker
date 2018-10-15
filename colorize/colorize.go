package colorize

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

type Colorize struct {
	ItemNum int
}

var (
	id      string
	names   string
	labels  string
	image   string
	created string
	status  string
	color   string
	size    string
)

func New() Colorize {
	return Colorize{}
}

func (c *Colorize) Container(containers []types.Container) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "NAMES", "IMAGE", "STATUS"})

	var data [][]string
	for _, container := range containers {
		c.ItemNum++
		if c.ItemNum%2 == 0 {
			color = "32m"
		} else {
			color = "36m"
		}

		id = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, container.ID[:10])
		names = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, container.Image)
		image = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, container.Names)
		status = fmt.Sprintf("\x1b[%s%v\x1b[0m\n", color, strings.Split(container.Status, " ")[0])

		data = append(data, []string{id, names, image, status})
	}

	table.AppendBulk(data)
	table.Render()
}

func (c *Colorize) Image(images []types.ImageSummary) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"REPOSITORY", "LABELS", "SIZE"})

	var data [][]string
	for _, image := range images {
		c.ItemNum++
		if c.ItemNum%2 == 0 {
			color = "32m"
		} else {
			color = "36m"
		}

		id = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, image.ID[7:19])
		labels = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, image.RepoTags)
		size = fmt.Sprintf("\x1b[%s%s\x1b[0m\n", color, strings.Replace(humanize.Bytes(uint64(image.Size)), " ", "", -1))

		data = append(data, []string{id, labels, size})
	}

	table.AppendBulk(data)
	table.Render()
}

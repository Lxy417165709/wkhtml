package wkhtml

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ImgBuilder struct {
	toolPath string
	input    string
	output   string
	width    uint32
	height   uint32
	html     string
}

type Option func(*ImgBuilder)

func NewImgBuilder(options ...Option) *ImgBuilder {
	imgBuilder := &ImgBuilder{
		toolPath: "wkhtmltoimage",
		input:    "-",
		output:   "-",
		width:    0,
		height:   0,
		html:     "",
	}
	for _, option := range options {
		option(imgBuilder)
	}
	return imgBuilder
}

func (i *ImgBuilder) Build() ([]byte, error) {
	if i.toolPath == "" {
		return nil, fmt.Errorf("toolPath blank")
	}
	cmd := exec.Command(i.toolPath, i.buildArgs()...)
	if i.input == "-" {
		cmd.Stdin = strings.NewReader(i.html)
	}
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("CombinedOutput fail. err=[%+v].", err)
		return nil, err
	}
	return data, nil
}

func (i *ImgBuilder) buildArgs() []string {
	args := make([]string, 0)
	args = append(args, "--format")
	args = append(args, "jpg")
	args = append(args, i.input)
	args = append(args, i.output)
	return args
}

func Input(input string) Option {
	return func(i *ImgBuilder) {
		i.input = input
	}
}

func Output(output string) Option {
	return func(i *ImgBuilder) {
		i.output = output
	}
}

func Html(html string) Option {
	return func(i *ImgBuilder) {
		i.html = html
	}
}

func ToolPath(path string) Option {
	return func(i *ImgBuilder) {
		i.toolPath = path
	}
}

func Width(width uint32) Option {
	return func(i *ImgBuilder) {
		i.width = width
	}
}

func Height(height uint32) Option {
	return func(i *ImgBuilder) {
		i.height = height
	}
}

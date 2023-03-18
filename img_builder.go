package wkhtml

import (
	"fmt"
	"os/exec"
	"strings"
)

// imgBuilder 图片构建器。
type imgBuilder struct {
	toolPath string // wkhtmltoimage 工具的路径，默认为 "wkhtmltoimage"。

	input  string // 输入，默认为 "-"(表示标准输入)。
	html   string // html数据，默认为 ""，仅当 input == '-' 时生效。
	output string // 输出，默认为 "-"(表示标准输出)。

	width   uint32 // 图片宽度，默认为 0(自适应)。
	height  uint32 // 图片高度，默认为 0(自适应)。
	format  string // 图片类型，默认为 "png"。
	quality uint32 // 图片品质，取值越大图片品质越高，取值区间 [0, 100]，默认为 94。

	attachArgs []string // 附加参数，当以上参数无法满足要求时，可利用该参数进行自定义。
}

// NewImgBuilder 新建图片构建器。
func NewImgBuilder(options ...ImgBuilderOption) *imgBuilder {
	imgBuilder := &imgBuilder{
		toolPath:   "wkhtmltoimage",
		input:      "-",
		html:       "",
		output:     "-",
		width:      0,
		height:     0,
		format:     "png",
		quality:    94,
		attachArgs: nil,
	}
	for _, option := range options {
		option(imgBuilder)
	}
	return imgBuilder
}

// Exec 执行，返回标准输出。
func (i *imgBuilder) Exec() ([]byte, error) {
	// 1. 校验。
	if i.toolPath == "" {
		return nil, fmt.Errorf("toolPath blank")
	}

	// 2. 执行。
	cmd := exec.Command(i.toolPath, i.buildArgs()...)
	if i.input == "-" {
		cmd.Stdin = strings.NewReader(i.html)
	}
	data, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// 3. 返回。
	return data, nil
}

// buildArgs 构建参数。
func (i *imgBuilder) buildArgs() []string {
	args := make([]string, 0)
	if i.format != "" {
		args = append(args, "--format")
		args = append(args, i.format)
	}
	if i.width != 0 {
		args = append(args, "--width")
		args = append(args, fmt.Sprintf("%d", i.width))
		args = append(args, "--disable-smart-width")
	}
	if i.height != 0 {
		args = append(args, "--height")
		args = append(args, fmt.Sprintf("%d", i.height))
	}
	if i.quality != 0 {
		args = append(args, "--quality")
		args = append(args, fmt.Sprintf("%d", i.quality))
	}
	args = append(args, i.attachArgs...)
	args = append(args, i.input)
	args = append(args, i.output)
	return args
}

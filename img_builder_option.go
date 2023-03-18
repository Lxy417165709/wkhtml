package wkhtml

// ImgBuilderOption 图片构建器参数。
type ImgBuilderOption func(*imgBuilder)

// Input 输入参数。
func Input(input string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.input = input
	}
}

// Output 输出参数。
func Output(output string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.output = output
	}
}

// Html html数据参数。
func Html(html string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.html = html
	}
}

// ToolPath 工具路径参数。
func ToolPath(path string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.toolPath = path
	}
}

// Width 图片宽度参数。
func Width(width uint32) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.width = width
	}
}

// Height 图片高度参数。
func Height(height uint32) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.height = height
	}
}

// Format 图片类型参数。
func Format(format string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.format = format
	}
}

// Quality 图片品质参数。
func Quality(quality uint32) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.quality = quality
	}
}

// AttachArgs 附加参数集。
func AttachArgs(args []string) ImgBuilderOption {
	return func(i *imgBuilder) {
		i.attachArgs = args
	}
}

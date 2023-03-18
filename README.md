# 1. 背景

由于工作中需要用到`html 转 图像`功能，经过多番搜索，没有找到特别满意的开源库，因此自己手动开发、并开源共享。

# 2. 语言

`Go`

# 3. 原理

1. 底层: 借助`wkhtmltoimage`工具，实现`html 转 图像`的功能。
2. 顶层: 利用`Go语言`进行封装，暴露易用接口。

# 4. 使用

## 4.1 步骤

1. 安装`wkhtmltoimage`工具，并将其路径加入`PATH`环境变量。 ([wkhtmltoimage工具官网](https://wkhtmltopdf.org/))
2. 安装本库依赖。
    ```shell
    go get -u github.com/Lxy417165709/wkhtml
    ```
3. 开始使用。

## 4.2 示例代码

### 4.2.1 读取网站`html`，形成图片。

运行以下代码，将在运行目录下生成名为`example_out.png`的文件。

```go
package main

import "github.com/Lxy417165709/wkhtml"

func main() {
	_, err := wkhtml.NewImgBuilder(wkhtml.Input("www.baidu.com"), wkhtml.Output("example_out.png")).Exec()
	if err != nil {
		panic(err)
	}
}
```

### 4.2.2 直接读取`html`，形成图片。

运行以下代码，将在运行目录下生成名为`example_out.png`的文件。

```go
package main

import "github.com/Lxy417165709/wkhtml"

func main() {
	_, err := wkhtml.NewImgBuilder(wkhtml.Html("<h1>hello world<h1>"), wkhtml.Output("example_out.png")).Exec()
	if err != nil {
		panic(err)
	}
}
```

### 4.2.3 执行后，读取标准输出、提取图像数据、保存图像文件。

运行以下代码，将在运行目录下生成名为`example_out.png`的文件。

```go
package main

import (
	"github.com/Lxy417165709/wkhtml"
	"os"
)

func main() {
	// 1. 执行，获取标准输出。
	data, err := wkhtml.NewImgBuilder(wkhtml.Html("<h1>hello world<h1>")).Exec()
	if err != nil {
		panic(err)
	}

	// 2. 提取图像数据。
	pngData := wkhtml.ExtractData(data, &wkhtml.DataExtractorPng{})

	// 3. 写入图像文件。
	pngFile, err := os.Create("example_out.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()
	if _, err := pngFile.Write(pngData); err != nil {
		panic(err)
	}
}
```

# 5. 常见错误

## 5.1 执行时报错: `exec: "wkhtmltoimage": executable file not found in %PATH%`

原因: 未安装`wkhtmltoimage`工具，或是已安装，但未将其路径加入`PATH`环境变量。

解决: 参考上文 `步骤4.1`。
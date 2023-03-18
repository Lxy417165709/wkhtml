# 背景

由于工作中用到了 `html 转 图像` 的功能，经过多番搜索，没有找到特别满意的开源库，因此自己手动开发。

# 语言

Go

# 原理

1. 底层: 借助`wkhtmltoimage`工具，实现 `html 转 图像` 的功能。
2. 顶层: 利用`Go语言`进行封装，暴露易用接口。

# 使用
## 步骤
1. 安装 `wkhtmltoimage` 工具，并加入环境变量的`Path`。 ([wkhtmltoimage工具官网](https://wkhtmltopdf.org/))
2. 安装本库依赖。
    ```shell
    go get -u github.com/Lxy417165709/wkhtml
    ```
3. 开始使用。

## 示例代码
```go

```
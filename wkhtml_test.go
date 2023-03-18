package wkhtml

import (
	"os"
	"testing"
)

func TestExec1(t *testing.T) {
	_, err := NewImgBuilder(Input("www.baidu.com"), Output("example_out.png")).Exec()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestExec2(t *testing.T) {
	_, err := NewImgBuilder(Input("www.baidu.com"), Format("jpg"),
		Output("example_out.jpg"), Width(200), Height(300),
		Quality(100), AttachArgs([]string{"-q"}),
	).Exec()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestExecAndExtract(t *testing.T) {
	data, err := NewImgBuilder(Input("www.baidu.com")).Exec()
	if err != nil {
		t.Error(err)
		return
	}
	imgData := ExtractData(data, &DataExtractorPng{})
	file, err := os.Create("example.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	if _, err = file.Write(imgData); err != nil {
		t.Error(err)
		return
	}
}

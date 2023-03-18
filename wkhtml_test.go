package wkhtml

import (
	"os"
	"testing"
)

func TestExec1(t *testing.T) {
	_, err := NewImgBuilder(Input("www.baidu.com"), Output("example_out1.png")).Exec()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestExec2(t *testing.T) {
	_, err := NewImgBuilder(Input("www.baidu.com"), Format("jpg"),
		Output("example_out2.jpg"), Width(200), Height(300),
		Quality(100), AttachArgs([]string{"-q"}),
	).Exec()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestExec3(t *testing.T) {
	data, err := NewImgBuilder(Input("www.baidu.com")).Exec()
	if err != nil {
		t.Error(err)
		return
	}
	imgData := ExtractData(data, &DataExtractorPng{})
	file, err := os.Create("example_out3.png")
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

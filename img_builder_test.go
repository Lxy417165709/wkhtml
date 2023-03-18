package wkhtml

import (
	"reflect"
	"testing"
)

func Test_imgBuilder_buildArgs(t *testing.T) {
	type fields struct {
		toolPath   string
		input      string
		html       string
		output     string
		width      uint32
		height     uint32
		format     string
		quality    uint32
		attachArgs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "",
			fields: fields{
				toolPath:   "wkhtmltoimage",
				input:      "www.baidu.com",
				html:       "abc",
				output:     "test.jpg",
				width:      100,
				height:     200,
				format:     "jpg",
				quality:    50,
				attachArgs: []string{"-q"},
			},
			want: []string{
				"--width", "100", "--disable-smart-width", "--height", "200", "--format", "jpg", "--quality", "50", "-q", "www.baidu.com",
				"test.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &imgBuilder{
				toolPath:   tt.fields.toolPath,
				input:      tt.fields.input,
				html:       tt.fields.html,
				output:     tt.fields.output,
				width:      tt.fields.width,
				height:     tt.fields.height,
				format:     tt.fields.format,
				quality:    tt.fields.quality,
				attachArgs: tt.fields.attachArgs,
			}
			if got := i.buildArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

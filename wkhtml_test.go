package wkhtml

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"os"
	"testing"
)

func TestBuild(t *testing.T) {
	data, err := NewImgBuilder(Input("www.baidu.com")).Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	//decoded, err := jpeg.Decode(bytes.NewReader(data))
	//for err != nil {
	//	data = data[1:]
	//	if len(data) == 0 {
	//		break
	//	}
	//	decoded, err = jpeg.Decode(bytes.NewReader(data))
	//}
	//fmt.Println(data)
	//fmt.Println(decoded.At(0, 0))

	//jpeg.Decode()
	//jpgHead := bytes.Index(data, []byte("\x89PNG\r\n\x1a\n"))

	jpgHead := bytes.Index(data, []byte("\xff\xd8"))
	jpgTail := bytes.Index(data, []byte("\xff\xd9"))
	jpegData := data[jpgHead : jpgTail+2]
	img, err := jpeg.Decode(bytes.NewReader(jpegData))
	if err != nil {
		fmt.Println(err)
		return
	}
	newJpegDataBuf := new(bytes.Buffer)
	err = jpeg.Encode(newJpegDataBuf, img, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(jpegData), len(newJpegDataBuf.Bytes()))
	pwd := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\wkhtml\`
	f, err := os.Create(pwd + "example1.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Write(jpegData)
	if err != nil {
		fmt.Println(err)
		return
	}

	f2, err := os.Create(pwd + "example2.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	_, err = f2.Write(newJpegDataBuf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	//jpegData := data[jpgHead : jpgTail+2]
	//pwd := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\wkhtml\`
	//f, err := os.Create(pwd + "example.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(jpgHead, jpgTail)
	//
	//defer f.Close()
	//_, err = f.Write(jpegData)
	//if err != nil {
	//	fmt.Println(err)
	//	return

	//img, err := jpeg.Decode(bytes.NewReader(jpegData))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//newJpegDataBuf := new(bytes.Buffer)
	//
	//err = jpeg.Encode(newJpegDataBuf, img, nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(len(data), len(jpegData), len(newJpegDataBuf.Bytes()))
	//fmt.Println(1, string(data[len(data)-50:]))
	//fmt.Println(1, string(data[len(data)-30:]))
	//fmt.Println(1, string(data[len(data)-20:]))
	//fmt.Println(1, string(data[len(data)-10:]))
	//fmt.Println(1, string(data[len(data)-80:]))
	//fmt.Println(1, string(data[len(data)-101:]))
	//x := data[len(newJpegDataBuf.Bytes())+jpgHead:]
	//fmt.Println(x)
	//fmt.Println(string(x[0:]))

	//jpgHead := bytes.Index(data, []byte("\xff\xd8"))
	//pwd := `C:\Users\李学悦\Desktop\all\github\Lxy417165709\wkhtml\`
	//f, err := os.Create(pwd + "example.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//
	//_, err = f.Write(data[jpgHead:])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}

// 经过  deCode、enCode 后，图像会失真...

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/snowlyg/czlib"
)

// 相关文档 http://www.zlib.net/manual.html#Basic

func Gzip(body []byte) ([]byte, error) {
	outb := make([]byte, 0, 16*1024)
	out := bytes.NewBuffer(outb)
	writer, err := czlib.NewWriterLevel2(out, -1, 8, -9, 8, 0)
	if err != nil {
		return []byte{}, err
	}
	n, err := writer.Write(body)
	if n != len(body) {
		return []byte{}, fmt.Errorf("compressed %d, expected %d", n, len(body))
	}
	if err != nil {
		return []byte{}, err
	}
	err = writer.Close()
	if err != nil {
		return []byte{}, err
	}
	return out.Bytes(), nil
}

// 使用参考 https://github.com/madler/zlib/blob/master/examples/zran.c
func Gunzip(body []byte) ([]byte, error) {
	reader, err := czlib.NewReaderLevel2(bytes.NewBuffer(body), -15)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(reader)
}

func main() {
	gzip, err := Gzip([]byte("fdsfsdfdsf"))
	if err != nil {
		panic(err)
	}
	gunzip, err := Gunzip(gzip)
	if err != nil {
		panic(err)
	}
	println(string(gunzip))
}

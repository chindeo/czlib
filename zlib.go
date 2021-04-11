package czlib

//go:generate go run -tags zlib_generate zlib_generate.go

// #cgo CPPFLAGS:
// #cgo LDFLAGS:
import "C"
import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// level method  windowBits  memLevel  strategy
// GzipLevel2 相关文档 http://www.zlib.net/manual.html#Basic
func GzipLevel2(body []byte) ([]byte, error) {
	outb := make([]byte, 0, 16*1024)
	out := bytes.NewBuffer(outb)
	// 兼容 nodejs zlib.createInflateRaw zlib.createDeflateRaw
	writer, err := NewWriterLevel2(out, -1, 8, -9, 8, 0)
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

// 兼容 nodejs zlib.createInflateRaw zlib.createDeflateRaw level -15
// Gunzip 使用参考 https://github.com/madler/zlib/blob/master/examples/zran.c
func GunzipLevel2(body []byte, level int) ([]byte, error) {
	reader, err := NewReaderLevel2(bytes.NewBuffer(body), level)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(reader)
}

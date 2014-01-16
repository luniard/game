package util

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	// compressor, err := zlib.NewWriterLevelDict(&buf, zlib.DefaultCompression, data)
	compressor := zlib.NewWriter(&buf)
	// defer compressor.Close()
	// if err != nil {
	// 	return nil, err
	// }
	compressor.Write(data)
	compressor.Close()

	return buf.Bytes(), nil

}

func Decompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(data)
	// fmt.Println("buf", buf.Bytes())
	decompressor, err := zlib.NewReader(&buf)
	if err != nil {
		fmt.Println("error at decompress", err)
		return nil, err
	}
	result := make([]byte, len(data))
	n, _ := decompressor.Read(result)
	decompressor.Close()
	return result[0:n], nil

}

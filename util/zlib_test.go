package util

import (
	// "bytes"
	// "compress/zlib"
	"fmt"
	// "io"
	// "os"
	"testing"
)

func TestZlib(t *testing.T) {
	// var b bytes.Buffer
	// w := zlib.NewWriter(&b)
	// w.Write([]byte("hello, world"))
	// w.Close()
	// // fmt.Println("compress", b.Bytes())

	// // fmt.Println("buf", b.Bytes())
	// r, err := zlib.NewReader(&b)
	// if err != nil {
	// 	fmt.Println("ig error")
	// }
	// io.Copy(os.Stdout, r)
	// r.Close()

	var input = []byte("Your Hero Game")
	fmt.Println("原始数据", input)
	test, _ := Compress(input)
	fmt.Println("压缩数据", test)
	test, _ = Decompress(test)
	fmt.Println("解压数据", test)

}

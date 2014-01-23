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

	var input = []byte("kobe system")
	fmt.Println("原始数据", input)
	test, _ := Compress(input)
	fmt.Println("压缩数据", test)
	fix := test
	fix[2] += 1
	fix = fix[0 : len(fix)-9]
	fix = append(fix, 0, 0, 0, 0, 0)
	fmt.Println(fix)
	copy(fix[len(fix)-4:], test[len(test)-4:])
	fmt.Println("调整数据", fix)
	test, _ = Decompress(test)
	fmt.Println("解压数据", test)

}

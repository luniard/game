package util

import (
	"fmt"
	"testing"
)

func TestTea(t *testing.T) {
	// data := make([]int, 5)
	// test := intToByte(data, 1)
	// fmt.Println("test ok", test)
	KEY := []int32{0x183f5c45, 0x426f739e, 0x719b3ffa, 0x358fac52}

	// info := []byte("UGG GAME GOOD")
	info := []byte{85, 71, 71, -1, 16, 1, 16, 1, -4, -9}
	fmt.Println("原数据: ", info)

	secretInfo := Encrypt32(info, KEY)
	fmt.Println("密数据: ", secretInfo)

	decryptInfo := Decrypt32(secretInfo, KEY)
	fmt.Println("解数据: ", decryptInfo)

}

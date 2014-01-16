package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/luniard/game/net/message"
	"github.com/luniard/game/util"
)

const (
	flash_policy  string = "<policy-file-request/>"
	MASK_COMPRESS uint32 = 1
	MASK_ENCRYPT  uint32 = 2
)

var (
	ENCRYPT_KEY []int32 = []int32{0x183f5c45, 0x426f739e, 0x719b3ffa, 0x358fac52}
)

type Decoder struct {
}

func (decoder *Decoder) Decode(b []byte) interface{} {
	fmt.Println("decode")
	fmt.Println(b)
	if len(b) < 22 {
		return nil
	}

	fmt.Println("read magic number")
	magic := b[0:22]

	if string(magic) == flash_policy {
		return "flash_policy"
	}

	//todo check magic number

	buf := bytes.NewBuffer(b) // b is []byte
	//skip magic
	bs := make([]byte, 3)
	readBuf(buf, &bs)

	var version int8
	// binary.Read(buf, binary.LittleEndian, &version)
	readBuf(buf, &version)
	fmt.Println(version)

	var srcId uint16
	readBuf(buf, &srcId)
	fmt.Println(srcId)

	var dstId uint16
	readBuf(buf, &dstId)
	fmt.Println(dstId)

	var flag uint32
	var bodylen int16

	readBuf(buf, &flag)
	readBuf(buf, &bodylen)
	fmt.Println("bodylen ", bodylen)
	if buf.Len() < int(bodylen) {
		return nil
	}
	body := make([]byte, bodylen)
	readBuf(buf, &body)
	fmt.Println(body)
	if (flag & MASK_ENCRYPT) == MASK_ENCRYPT {
		fmt.Println("MASK_ENCRYPT")
		body = util.Encrypt32(body, ENCRYPT_KEY)
	}

	if (flag & MASK_COMPRESS) == MASK_COMPRESS {
		fmt.Println("MASK_COMPRESS")
		body, _ = util.Decompress(body)
	}

	var req message.Request

	return req

}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func readBuf(r io.Reader, data interface{}) {
	binary.Read(r, binary.BigEndian, data)
}

package net

import (
	"bytes"
	"encoding/binary"
	// "fmt"
	"io"

	"github.com/luniard/game/net/message"
	"github.com/luniard/game/util"
)

type Encoder struct {
}

func (encoder *Encoder) Encode(data interface{}) []byte {
	// fmt.Println("encode")
	// response := []byte{}

	buf := bytes.NewBuffer([]byte{})

	switch response := data.(type) {
	case string:
		writeBuf(buf, response)
		writeBuf(buf, 0)

	case message.Message:
		// fmt.Println("encode response...", response)
		writeBuf(buf, message.MagicCode)
		writeBuf(buf, byte(1))
		writeBuf(buf, response.MessageHeader.DstId)
		writeBuf(buf, response.MessageHeader.SrcId)

		bodylen := 8 + len(response.Body)
		// println("body length: ", bodylen)
		// body := make([]byte, bodylen)
		var body []byte
		bufBody := bytes.NewBuffer(body)
		writeBuf(bufBody, response.MessageHeader.SeqNum)
		writeBuf(bufBody, response.MessageHeader.MsgCode)
		if bodylen > 8 {
			writeBuf(bufBody, response.Body)
		}

		body = bufBody.Bytes()

		// fmt.Println("test body before compress", body)
		gzipedBody, _ := util.Compress(body)

		//TODO fix byte for ugg client
		// fix := gzipedBody
		// fix[2] += 1
		// fix = fix[0 : len(fix)-9]
		// fix = append(fix, 0, 0, 0, 0, 0)
		// copy(fix[len(fix)-4:], gzipedBody[len(gzipedBody)-4:])

		// fmt.Println("after compress and fix", fix)

		encryptByte := util.Encrypt32(gzipedBody, ENCRYPT_KEY)
		// encryptByte := util.Encrypt32(fix, ENCRYPT_KEY)

		flag := MASK_ENCRYPT | MASK_COMPRESS | uint32(0)
		writeBuf(buf, flag)
		writeBuf(buf, int16(len(encryptByte)))
		writeBuf(buf, encryptByte)

		// fmt.Println("test encode flag encryptByte", flag)
		// fmt.Println(encryptByte)

	}
	return buf.Bytes()
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func writeBuf(w io.Writer, data interface{}) {
	// binary.Read(r, binary.BigEndian, data)
	binary.Write(w, binary.BigEndian, data)
}

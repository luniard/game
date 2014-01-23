package net

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	message "github.com/luniard/game/net/message"
)

type Handler struct {
}

func (hanlder *Handler) Handle(data interface{}) interface{} {
	fmt.Println("handle..")

	// switch request := data.(type) {
	// case message.UARequest:
	// 	fmt.Println("get ua request...", request)
	// 	fmt.Println("TEST PASS UA")
	// 	authCode := makeSessionId()
	// 	fmt.Println("authCode", authCode)

	// 	resp := message.UAResponse{}
	// 	resp.AuthCode = authCode
	// 	resp.Message.MessageHeader = request.Message.MessageHeader
	// 	fmt.Println("resp--0", resp)
	// 	resp.Encode()
	// 	fmt.Println("resp--1", resp)
	// 	return resp

	// }

	// fmt.Println("common resp")
	// var resp message.Response

	if msg, ok := data.(message.Message); ok {
		println(msg.MessageHeader.MsgCode)
		if msg.MessageHeader.MsgCode == 0x1001 {
			//todo module
			// fmt.Println("get here")
			authCode := makeSessionId()
			// authCode := "wahaha"
			fmt.Println("generate authCode", authCode)
			// req := message.UARequest{}
			// req.Message = message.Message{header, messageBody}
			// req.Decode()

			resp := message.UAResponse{}
			resp.AuthCode = authCode
			resp.Message.MessageHeader = msg.MessageHeader
			resp.Message.MessageHeader.MsgCode += 1
			resp.Encode()
			return resp.Message
		}
	}

	return nil
}

func NewHandler() *Handler {
	return &Handler{}
}

func makeSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

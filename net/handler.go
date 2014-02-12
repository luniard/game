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

type handler func(msg message.Message) message.Message

var Handlers map[uint32]handler

func init() {
	Handlers = make(map[uint32]handler)
}

func (hanlder *Handler) Handle(data interface{}) interface{} {
	fmt.Println("handle..")

	if msg, ok := data.(message.Message); ok {
		println(msg.MessageHeader.MsgCode)
		if msg.MessageHeader.MsgCode == 0x1001 {
			authCode := makeSessionId()
			fmt.Println("generate authCode", authCode)

			resp := message.UAResponse{}
			resp.AuthCode = authCode
			resp.Message.MessageHeader = msg.MessageHeader
			resp.Message.MessageHeader.MsgCode += 1
			resp.Encode()
			return resp.Message
		} else {
			fmt.Println(Handlers)
			f := Handlers[msg.MessageHeader.MsgCode]
			return f(msg)
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

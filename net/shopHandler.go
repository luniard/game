package net

import (
	"fmt"
	"github.com/luniard/game/net/message"
)

func init() {
	fmt.Println("register shop handler")
	Handlers[0x3603] = handleShop
}

func handleShop(msg message.Message) message.Message {
	fmt.Println("get shop request")
	req := message.ShopRequest{}
	req.Message = msg
	req.Decode()

	resp := message.ShopResponse{}
	// resp.AuthCode = authCode
	resp.Message.MessageHeader = msg.MessageHeader
	resp.Message.MessageHeader.MsgCode += 1
	resp.Encode()
	fmt.Println("send shop response")
	return resp.Message
}

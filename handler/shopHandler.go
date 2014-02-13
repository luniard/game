package handler

import (
	"fmt"
	"github.com/luniard/game/model"
	"github.com/luniard/game/net/message"
)

func init() {
	fmt.Println("register shop handler")
	Handlers[0x3603] = handleShop
}

func handleShop(msg message.Message) message.Message {
	req := message.ShopRequest{}
	req.Message = msg
	req.Decode()

	model.AddShop("test1", "test2")

	resp := message.ShopResponse{}
	// resp.AuthCode = authCode
	resp.Message.MessageHeader = msg.MessageHeader
	resp.Message.MessageHeader.MsgCode += 1
	resp.Encode()
	fmt.Println("ShopHandler>>> shop handled")
	return resp.Message
}

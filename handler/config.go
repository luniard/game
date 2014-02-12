package handler

import (
	message "github.com/luniard/game/net/message"
)

type handler func(msg message.Message) message.Message

var Handlers map[uint32]handler

func init() {
	Handlers = make(map[uint32]handler)
}

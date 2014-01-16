package net

import (
	"fmt"

	message "github.com/luniard/game/net/message"
)

type Handler struct {
}

func (hanlder *Handler) Handle(data interface{}) interface{} {
	fmt.Println("handle..")

	var resp message.Response

	return resp
}

func NewHandler() *Handler {
	return &Handler{}
}

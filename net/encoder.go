package net

import (
	"fmt"
)

type Encoder struct {
}

func (encoder *Encoder) Encode(data interface{}) []byte {
	fmt.Println("encode")
	response := []byte{}
	return response
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

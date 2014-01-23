package net

import (
	"fmt"
)

type Pipe struct {
	Decoder *Decoder
	Handler *Handler
	Encoder *Encoder
}

func NewPipe() *Pipe {
	fmt.Println("Pipe>>> create a pipe")
	return &Pipe{
		Decoder: NewDecoder(),
		Handler: NewHandler(),
		Encoder: NewEncoder(),
	}
}

func (p *Pipe) Handle(buf []byte) []byte {
	data := p.Decoder.Decode(buf)
	fmt.Println("Pipe>>> decode success")
	response := p.Handler.Handle(data)
	fmt.Println("Pipe>>> handle success")
	out := p.Encoder.Encode(response)
	fmt.Println("Pipe>>> encode success")

	return out
}

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
	fmt.Println("create a pipe")
	return &Pipe{
		Decoder: NewDecoder(),
		Handler: NewHandler(),
		Encoder: NewEncoder(),
	}
}

func (p *Pipe) Handle(buf []byte) {
	data := p.Decoder.Decode(buf)
	response := p.Handler.Handle(data)
	p.Encoder.Encode(response)

}

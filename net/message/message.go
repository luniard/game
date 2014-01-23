package message

import (
	"encoding/json"
	"fmt"
)

var (
	MagicCode []byte = []byte{0x55, 0x47, 0x47}
)

type Message struct {
	MessageHeader
	Body []byte
}

type Decode interface {
	// Message
	Decode()
	GetMsgCode() uint32
}

type Encode interface {
	Encode()
}

type DecodeRequestI interface {
	DecodeBody() map[string]*json.RawMessage
}

type EncodeResponseI interface {
	EncodeJson(m map[string]interface{}) []byte
}

// type ProxiedRequest struct {
// 	Message
// }

// type ProxiedResponse struct {
// 	Message
// }

// func (p *ProxiedRequest) GetHeader() MessageHeader {
// 	return p.Header
// }
// func (p *ProxiedRequest) GetBody() []byte {
// 	return p.Body
// }

// func (p *ProxiedResponse) GetHeader() MessageHeader {
// 	return p.Header
// }
// func (p *ProxiedResponse) GetBody() []byte {
// 	return p.Body
// }

// type DecodeRequest struct {
// 	// header MessageHeader
// 	Body []byte
// }

// func (d *DecodeRequest) Decode() {

// 	d.DecodeBody(objmap)
// }

func (msg *Message) DecodeBody() map[string]*json.RawMessage {
	fmt.Println("decode body")
	var objmap map[string]*json.RawMessage
	// fmt.Println("body", d.Body)
	json.Unmarshal(msg.Body, &objmap)
	// fmt.Println("json unmarshal", objmap)
	return objmap
}

func (msg *Message) EncodeJson(m map[string]interface{}) []byte {
	fmt.Println("encode body")

	// fmt.Println("body", d.Body)
	data, _ := json.Marshal(m)
	// fmt.Println("json unmarshal", objmap)
	return data
}

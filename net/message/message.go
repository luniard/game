package message

import (
	"encoding/json"
	// "fmt"
)

var (
	MagicCode []byte = []byte{0x55, 0x47, 0x47}
)

type Message struct {
	MessageHeader
	Body []byte
	msg  string //返回客户端的消息
	code int    //返回客户端的消息码
}

type Decode interface {
	// Message
	Decode()
}

type Encode interface {
	Encode()
}

type MsgCode interface {
	GetMsgCode() uint32
}

type DecodeRequestI interface {
	DecodeBody() map[string]*json.RawMessage
}

type EncodeResponseI interface {
	EncodeJson(m map[string]interface{}) []byte
}

func (msg *Message) DecodeBody() map[string]*json.RawMessage {
	// fmt.Println("decode body")
	var objmap map[string]*json.RawMessage
	// fmt.Println("body", d.Body)
	json.Unmarshal(msg.Body, &objmap)
	// fmt.Println("json unmarshal", objmap)
	return objmap
}

func (msg *Message) EncodeJson(m map[string]interface{}) []byte {
	// fmt.Println("encode body")

	// fmt.Println("body", d.Body)
	m["msg"] = msg.msg
	m["code"] = msg.code
	data, _ := json.Marshal(m)
	// fmt.Println("json unmarshal", objmap)
	return data
}

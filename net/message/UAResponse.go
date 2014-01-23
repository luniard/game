package message

// import (
// 	// "encoding/json"
// 	"fmt"
// )

type UAResponse struct {
	Message
	AuthCode string
}

func (u *UAResponse) Encode() {
	var m = make(map[string]interface{})
	m["authcode"] = u.AuthCode
	data := u.Message.EncodeJson(m)
	// fmt.Println("encode ua body")
	u.Message.Body = data
}

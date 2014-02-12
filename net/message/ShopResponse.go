package message

// import (
// 	""
// )

type ShopResponse struct {
	Message
	AuthCode string
}

func (u *ShopResponse) Encode() {
	var m = make(map[string]interface{})
	// m["authcode"] = u.AuthCode
	u.Message.code = 200
	data := u.Message.EncodeJson(m)
	// fmt.Println("encode ua body")
	u.Message.Body = data
}

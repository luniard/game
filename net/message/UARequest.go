package message

import (
	"encoding/json"
	"fmt"
)

type UARequest struct {
	Message
	userName string
	token    string
	authCode string
}

func (u *UARequest) Decode() {
	m := u.Message.DecodeBody()
	fmt.Println("get ua decode body")
	json.Unmarshal(*m["username"], &u.userName)
	json.Unmarshal(*m["token"], &u.token)
	// json.Unmarshal(*(m["authcode"]), &u.authCode)
	fmt.Println(u.userName)
}

func (u *UARequest) GetMsgCode() int32 {
	return 0x1001
}

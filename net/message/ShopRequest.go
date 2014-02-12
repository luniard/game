package message

// import (
// 	"encoding/json"
// )

type ShopRequest struct {
	Message
}

func (u *ShopRequest) Decode() {

}

func (u *ShopRequest) GetMsgCode() int32 {
	return 0x3603
}

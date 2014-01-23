package message

type MessageHeader struct {
	Version  int8
	SrcId    uint16
	DstId    uint16
	SeqNum   uint32
	MsgCode  uint32
	AuthCode string
}

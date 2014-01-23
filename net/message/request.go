package message

type Request interface {
	Decode(b []byte) interface{}
}

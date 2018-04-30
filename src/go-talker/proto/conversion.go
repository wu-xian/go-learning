package proto

import (
	proto "github.com/golang/protobuf/proto"
)

//MessageToBytes
func MessageToBytes(msg *Message) ([]byte, error) {
	messageBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return messageBytes, nil
}

//BytesToMessage
func BytesToMessage(bytess []byte) (*Message, error) {
	msg := &Message{}
	err := proto.Unmarshal(bytess, msg)
	return msg, err
}

package proto

import (
	_ "github.com/golang/protobuf/proto"
)

// //Marshal Message
// func Marshal(msg proto.Message) ([]byte, error) {
// 	messageBytes, err := proto.Marshal(msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return messageBytes, nil
// }

// //Unmarshal Message
// func Unmarshal(bytess []byte) (msg proto.Message, err error) {
// 	err = proto.Unmarshal(bytess, msg)
// 	return msg, err
// }

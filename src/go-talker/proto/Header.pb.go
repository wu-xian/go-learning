// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package proto

import (
	p "github.com/golang/protobuf/proto"
)

func (self *Header) Marshal(t Header) ([]byte, error) {
	return p.Marshal(&t)
}

func (self *Header) Unmarshal(bytes []byte) (t Header) {
	p.Unmarshal(bytes, &t)
	return t
}

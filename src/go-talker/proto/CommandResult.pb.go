// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package proto

import (
	p "github.com/golang/protobuf/proto"
)

func (self *CommandResult) Marshal(t CommandResult) ([]byte, error) {
	return p.Marshal(&t)
}

func (self *CommandResult) Unmarshal(bytes []byte) (t CommandResult) {
	p.Unmarshal(bytes, &t)
	return t
}

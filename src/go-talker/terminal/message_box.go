package terminal

import (
	"sync"

	"learn/src/go-talker/proto"

	"github.com/cjbassi/termui"
)

type MessageList struct {
	*termui.Table
	Text        string
	TextLine    int
	TextFgColor int
	TextBgColor int
	WrapLength  int
}

var (
	messageLocker sync.Mutex
	messageChan   chan *proto.MessageWarpper = make(chan *proto.MessageWarpper, 0)
)

func NewMessageList() *MessageList {
	return &MessageList{
		Table: termui.NewTable(),
	}
}

func (self *MessageList) Buffer() *termui.Buffer {
	buf := self.Table.Buffer()
	//c1 := termui.NewCell()
	return buf
}

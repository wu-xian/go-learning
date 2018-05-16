package terminal

import (
	"sync"

	"learn/src/go-talker/proto"

	"github.com/cjbassi/termui"
)

type MessageList struct {
	*termui.Table
	termui.Block
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

func (self *MessageList) Buffer() *termui.Buffer {
	buf := self.Buffer()
	//c1 := termui.NewCell()
	return buf
}

package terminal

import (
	"sync"

	"learn/src/go-talker/proto"

	"github.com/cjbassi/termui"
)

type MessageList struct {
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

// ListenToServer ?
// func ListenToServer() {
// 	go func() {
// 		select {
// 		case _ = <-inChan:
// 			{
// 				termui.Render(termui.Body)
// 			}
// 		case _ = <-outChan:
// 			{
// 				termui.Render(termui.Body)
// 			}
// 		}
// 	}()
// }

func (self *MessageList) Buffer() *termui.Buffer {
	buf := self.Buffer()
	for i, v := range clients {
		buf.SetString(3, 3*i, v.Name, 35, 47)
	}
	return buf
}

package terminal

import (
	"fmt"
	"sync"

	"github.com/cjbassi/termui"
)

type MessageBox struct {
	*termui.Block
	Messages    []Message
	Boxes       []Box
	MessageChan chan *Message
	locker      sync.Mutex
	showIndex   int32
	lastIndex   int32
}

func (self *MessageBox) AddMessage(message Message) {
	self.locker.Lock()
	self.Messages = append(self.Messages, message)
	self.locker.Unlock()
}

func NewMessageBox() *MessageBox {
	return &MessageBox{
		Block: termui.NewBlock(),
		Messages: []Message{Message{Content: "1111"},
			Message{Content: "22"},
			Message{Content: "3333333"},
			Message{Content: "44444444"},
			Message{Content: "5555"},
			Message{Content: "666"},
			Message{Content: "777"},
			Message{Content: "888"},
			Message{Content: "999"},
			Message{Content: "000"},
			Message{Content: "111"},
			Message{Content: "222"},
			Message{Content: "333"},
			Message{Content: "444"},
			Message{Content: "5555"},
		},
		Boxes:       make([]Box, 0),
		MessageChan: make(chan *Message, 0),
	}
}

type Message struct {
	Content string //消息内容
	Name    string //消息发送者名称
	Time    int64  //发送消息时间
}

type Box struct {
	message *Message
	index   int32
}

func NewMessageList() *MessageBox {
	return &MessageBox{
		Block: termui.NewBlock(),
	}
}

func (self *MessageBox) Buffer() *termui.Buffer {
	self.locker.Lock()
	buf := self.Block.Buffer()
	y := self.Block.Y
	y2 := y/2 - 1
	if len(self.Messages) > y2 {
		i := len(self.Messages) - y2
		shown := self.Messages[i:]
		for _i, _v := range shown {
			buf.SetString(2, 2+2*_i, messageFormatter(_v.Name, _v.Content), 35, 47)
		}
	} else {
		for _i, _v := range self.Messages {
			buf.SetString(2, 2+2*_i, messageFormatter(_v.Name, _v.Content), 35, 47)
		}
	}
	self.locker.Unlock()
	return buf
}

func messageFormatter(name, content string) string {
	return fmt.Sprintf("[%s]: %s", name, content)
}

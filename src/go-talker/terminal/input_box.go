package terminal

import "github.com/cjbassi/termui"

type InputBox struct {
	termui.Block
	Text        string
	TextFgColor int
	TextBgColor int
	WrapLength  int
}

// NewInputBox return a new input box pointer
func NewInputBox() *InputBox {
	return &InputBox{}
}

// Buffer return current input box buffer
func (self *InputBox) Buffer() *termui.Buffer {
	buf := self.Block.Buffer()
	buf.SetString(0, 0, "wu-xian", 35, 47)
	return buf
}

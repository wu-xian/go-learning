package terminal

import _ "github.com/wu-xian/termui"

type Par struct {
	Block
	Text        string
	TextFgColor Attribute
	TextBgColor Attribute
	WrapLength  int
}

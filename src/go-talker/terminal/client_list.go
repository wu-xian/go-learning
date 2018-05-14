package terminal

import (
	"github.com/cjbassi/termui"
)

type ClientList struct {
	termui.Block
	Text        string
	TextLine    int
	TextFgColor int
	TextBgColor int
	WrapLength  int
}

type Client struct {
	Id   int32
	Name string
	Flag int32 //0 :logout  , 1 login
}

// ListenToServer ?
func ListenToServer(c chan Client) {

}

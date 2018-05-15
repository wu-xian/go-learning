package terminal

import (
	"sync"

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
}

var (
	clients      []Client
	clientLocker sync.Mutex
	inChan       chan *Client = make(chan *Client, 0)
	outChan      chan *Client = make(chan *Client, 0)
)

func Add(client *Client) {
	clientLocker.Lock()
	defer clientLocker.Unlock()
	clients = append(clients, *client)
	inChan <- client
}

func Remove(client *Client) {
	clientLocker.Lock()
	defer clientLocker.Unlock()
	index := -1
	for i, v := range clients {
		if v.Id == client.Id {
			index = i
			break
		}
	}

	clients = append(clients[:index], clients[index+1:]...)
	outChan <- client
}

// ListenToServer ?
func ListenToServer() {
	go func() {
		select {
		case _ = <-inChan:
			{
				termui.Render(termui.Body)
			}
		case _ = <-outChan:
			{
				termui.Render(termui.Body)
			}
		}
	}()
}

func (self *ClientList) Buffer() *termui.Buffer {
	buf := self.Buffer()
	for i, v := range clients {
		buf.SetString(3, 3*i, v.Name, 35, 47)
	}
	return buf
}

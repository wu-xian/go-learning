package terminal

import (
	"fmt"
	"learn/src/go-talker/proto"
	"sync"

	"learn/src/go-talker/log"

	ui "github.com/cjbassi/termui"
)

var UILocker sync.Mutex

func LoopClientUI(messageChan chan *proto.MessageWarpper, messagePublishChan chan string) {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()
	defer ui.Clear()

	ui.Body.Cols = 12
	ui.Body.Rows = 12

	inputBox := NewInputBox()
	inputBox.XOffset = 9
	inputBox.X = 12
	inputBox.BorderBg = 7

	clientList := NewClientList()
	clientList.BorderBg = 7

	messageBox := NewMessageBox()
	messageBox.BorderBg = 7

	ui.Body.Set(0, 0, 4, 12, clientList)
	ui.Body.Set(4, 0, 12, 8, messageBox)
	ui.Body.Set(4, 8, 12, 12, inputBox)

	ui.Render(ui.Body)

	ui.On("<C-c>", func(e ui.Event) {
		ui.StopLoop()
	})

	go func() {
		for {
			select {
			case _ = <-clientList.InChan:
				{
					UILocker.Lock()
					ui.Render(ui.Body)
					UILocker.Unlock()
				}
			case _ = <-clientList.OutChan:
				{
					UILocker.Lock()
					ui.Render(ui.Body)
					UILocker.Unlock()
				}
			case _ = <-messageBox.InChan:
				{
					UILocker.Lock()
					ui.Render(ui.Body)
					UILocker.Unlock()
				}
			}
		}
	}()

	go func() {
		for {
			message := <-messageChan
			log.Logger.Info("get message , type:", message.Type)
			switch message.Type {
			case proto.COMMUNICATION_TYPE_ClientLogin:
				{
					clientList.Add(&Client{
						Id:   message.MessageClientLogin.Id,
						Name: message.MessageClientLogin.Name,
					})
				}
			case proto.COMMUNICATION_TYPE_ClientLogout:
				{
					clientList.Remove(&Client{
						Id: message.MessageClientLogout.Id,
					})
				}
			case proto.COMMUNICATION_TYPE_ClientReceived:
				{
					messageBox.AddMessage(Message{
						Name:    message.MessageClientReceived.Name,
						Content: message.MessageClientReceived.Content,
					})
				}
			default:
				{

					log.Logger.Info("received:", message)
				}
			}
		}
	}()

	inputBox.ListenInput(messagePublishChan)

	ui.On("<resize>", func(e ui.Event) {
		UILocker.Lock()
		ui.Clear()
		ui.Body.Width, ui.Body.Height = e.Width, e.Height
		ui.Body.Resize()
		ui.Render(ui.Body)
		UILocker.Unlock()
	})

	ui.Loop()
	fmt.Println("application stopping")
	UILocker.Lock()
	UILocker.Unlock()
}

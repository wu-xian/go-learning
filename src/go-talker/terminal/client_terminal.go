package terminal

import (
	ui "github.com/cjbassi/termui"
)

func LoopClientUI() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()
	defer ui.Clear()

	ui.Body.Cols = 12
	ui.Body.Rows = 12
	input_box := NewInputBox()
	input_box.XOffset = 9
	input_box.X = 12
	input_box.BorderBg = 7
	input_box.Label = "Message"

	client_list := ui.NewBlock()
	client_list.BorderBg = 7

	message_list := ui.NewBlock()
	message_list.BorderBg = 7

	ui.Body.Set(0, 0, 4, 12, client_list)
	ui.Body.Set(4, 0, 12, 8, message_list)
	ui.Body.Set(4, 8, 12, 12, input_box)

	ui.Render(ui.Body)

	ui.On("<C-c>", func(e ui.Event) {
		ui.StopLoop()
	})

	resizeChan := make(chan bool, 1)
	go func(c chan bool) {
		for {
			_ = <-c
			ui.Clear()
			ui.Body.Resize()
			ui.Render(ui.Body)
		}
	}(resizeChan)
	ui.On("<resize>", func(e ui.Event) {
		ui.Body.Width, ui.Body.Height = e.Width, e.Height
		resizeChan <- true
	})

	ui.Loop()
}

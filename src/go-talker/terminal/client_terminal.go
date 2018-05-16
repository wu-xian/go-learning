package terminal

import (
	ui "github.com/cjbassi/termui"
)

func LoopClientUI(message chan string) {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()
	defer ui.Clear()

	ui.Body.Cols = 12
	ui.Body.Rows = 12
	//input_box := NewInputBox()
	input_box := ui.NewBlock()
	input_box.XOffset = 9
	input_box.X = 12
	input_box.BorderBg = 7
	//input_box.Label = "Message"
	//input_box.ListenInput(message)

	client_list := NewClientList()
	//client_list := ui.NewBlock()
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

	//resizeChan := make(chan bool, 0)
	// go func(c chan bool) {
	// 	for {
	// 		_ = <-c
	// 		ui.Render(ui.Body)
	// 	}
	// }(resizeChan)

	go func() {
		for {
			select {
			case _ = <-client_list.InChan:
				{
					ui.Render(ui.Body)
				}
			case _ = <-client_list.OutChan:
				{
					ui.Render(ui.Body)
				}
			}
		}
	}()

	var ii int32 = 4
	ui.On("a", func(e ui.Event) {
		client_list.Add(&Client{
			Id:   ii,
			Name: "23333",
		})
		ii++
	})

	ui.On("<resize>", func(e ui.Event) {
		ui.Clear()
		ui.Body.Width, ui.Body.Height = e.Width, e.Height
		ui.Body.Resize()
		ui.Render(ui.Body)
		//resizeChan <- true
	})

	ui.Loop()
}

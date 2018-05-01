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
	p := ui.NewBlock()
	p.BorderBg = 7
	p.X = 2
	p.Y = 2
	p.XOffset = 0
	p.YOffset = 0
	p.Label = "Message"

	ui.Body.Set(6, 6, 12, 12, p)

	ui.Render(ui.Body)

	ui.On("<C-c>", func(e ui.Event) {
		ui.StopLoop()
	})

	resizeChan := make(chan bool, 1)
	go func(c chan bool) {
		for {
			_ = <-c
			ui.Clear()
			ui.Render(ui.Body)
		}
	}(resizeChan)
	ui.On("<resize>", func(e ui.Event) {
		ui.Body.Width, ui.Body.Height = e.Width, e.Height
		ui.Body.Resize()
		resizeChan <- true
	})

	ui.Loop()
}

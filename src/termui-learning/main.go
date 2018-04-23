package main

import (
	ui "github.com/ttacon/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar("enter ctrl+x to quite")
	p.Height = 10
	p.Width = 150
	p.X = 0
	p.TextBgColor = ui.ColorWhite
	p.TextFgColor = ui.ColorBlack
	p.BorderLabel = "everything is ok"
	p.BorderFg = ui.ColorCyan

	ui.Render(p)

	ui.Handle("/sys/kbd/q", func(e ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-x", func(e ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}

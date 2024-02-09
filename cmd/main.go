package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pomodoro-go/timer"
	"strconv"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			counter := 0
			timer.NewTimer(time.Second * 1).
				Subscribe(func() {
					hello.SetText("Counter: " + strconv.Itoa(counter) + " Timer: " + time.Now().String())
					counter++
				}).
				Periodic().
				Start()
		}),
	))

	w.ShowAndRun()
}

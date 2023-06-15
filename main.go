package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	fileDialog := dialog.NewFileOpen(func(uriReadCloser fyne.URIReadCloser, err error) {
		filePath := uriReadCloser.URI().Path()

		fmt.Println(filePath)
	}, w)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			fileDialog.Show()
		}),
	))

	w.ShowAndRun()
}

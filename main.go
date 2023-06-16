package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	fileDialog := dialog.NewFileOpen(func(uriReadCloser fyne.URIReadCloser, err error) {
		filePath := uriReadCloser.URI().Path()

		fmt.Println(filePath)

		dat, err := os.ReadFile(filePath)

		fmt.Println(decimalToHex(dat[0x00B9F6]), " ", decimalToHex(dat[0x00B9C4]), " ", decimalToHex(dat[0x00B992]))

	}, w)
	fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".smc"}))

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

func decimalToHex(decimal byte) string {
	return strconv.FormatInt(int64(decimal), 16)
}

func decompressLZ2(data []byte) {

}

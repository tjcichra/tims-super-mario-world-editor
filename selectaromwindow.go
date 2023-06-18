package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func createSelectARomWindow(app fyne.App) fyne.Window {
	selectARomWindow := app.NewWindow("Select a ROM")
	selectARomWindow.Resize(fyne.NewSize(300, 300))

	openRomFileDialog := dialog.NewFileOpen(func(uriReadCloser fyne.URIReadCloser, err error) {
		if uriReadCloser == nil {
			return
		}

		filePath := uriReadCloser.URI().Path()

		fmt.Println(filePath)

		dat, err := os.ReadFile(filePath)

		fmt.Println(decimalToHex(dat[0x00B9F6]), " ", decimalToHex(dat[0x00B9C4]), " ", decimalToHex(dat[0x00B992]))

		// decompressLZ2(dat[0x0881FD:])
		// gfx0file := decompressLZ2(dat[((0x08D9F9&0x7F0000)>>1)|(0x08D9F9&0x7FFF):])

		// var opts jpeg.Options
		// opts.Quality = 1

		// fmt.Println(i)
	}, selectARomWindow)
	openRomFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".smc"}))

	openRomButton := widget.NewButton("Open ROM", func() {
		openRomFileDialog.Show()
	})

	mostRecentLabel := widget.NewLabel("Most Recent:")

	selectARomContainer1 := container.NewVBox(openRomButton, mostRecentLabel)

	selectARomWindow.SetContent(selectARomContainer1)

	return selectARomWindow
}

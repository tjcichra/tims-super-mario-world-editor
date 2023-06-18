package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	selectARomWindow := createSelectARomWindow(a)

	selectARomWindow.Show()

	w := a.NewWindow("Hello")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			fmt.Println("Display help")
		}),
	)

	// imgBytes, _ := os.ReadFile("photo")
	// img := canvas.NewImageFromReader(bytes.NewReader(imgBytes), "photo")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewBorder(toolbar, nil, nil, nil, container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		})),
	// container.NewMax(img),
	))

	// w.Show()
	a.Run()
}

func decimalToHex(decimal byte) string {
	return strconv.FormatInt(int64(decimal), 16)
}

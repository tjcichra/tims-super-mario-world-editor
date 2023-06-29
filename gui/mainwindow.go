package gui

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createMainWindow(app fyne.App, image *image.Paletted) fyne.Window {
	mainWindow := app.NewWindow("Hello")

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

	imgComponent := canvas.NewImageFromImage(image)
	imgComponent.ScaleMode = canvas.ImageScalePixels

	hello := widget.NewLabel("Hello Fyne!")
	mainWindow.SetContent(container.NewBorder(toolbar, nil, nil, nil, container.NewVBox(
		hello),
		imgComponent,
	))

	return mainWindow
}

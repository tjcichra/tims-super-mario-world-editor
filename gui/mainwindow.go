package gui

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createMainWindow(app fyne.App, image *image.Paletted) fyne.Window {
	// graphicsWindow := createGraphicsWindow(app)

	mainWindow := app.NewWindow("Hello")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			fmt.Println("Display help")
			// graphicsWindow.Show()
		}),
	)

	tabs := container.NewAppTabs()

	tree := createMainTree(tabs, toolbar)

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// )

	// tabs := container.NewAppTabs([]*container.TabItem{
	// 	// container.NewTabItem("Tab 1", container.NewBorder(toolbar, nil, nil, nil, nil)),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!"))}...,
	// )

	// tabs.SetTabLocation(container.TabLocationLeading)

	// imgComponent := canvas.NewImageFromImage(image)
	// imgComponent.ScaleMode = canvas.ImageScalePixels

	// hello := widget.NewLabel("Hello Fyne!")
	// mainWindow.SetContent(container.NewBorder(toolbar, nil, nil, nil, container.NewVBox(
	// 	hello),
	// // imgComponent,
	// ))

	mainWindow.SetContent(container.NewBorder(toolbar, nil, tree, nil, tabs)) // imgComponent,

	return mainWindow
}

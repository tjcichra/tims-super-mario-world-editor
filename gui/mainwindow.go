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
			// graphicsWindow.Show()
		}),
	)
	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"levels", "b", "c"}
			case "levels":
				return []widget.TreeNodeID{"a1", "a2"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "levels"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		})

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// )

	tabs := container.NewAppTabs([]*container.TabItem{
		// container.NewTabItem("Tab 1", container.NewBorder(toolbar, nil, nil, nil, nil)),
		container.NewTabItem("Tab 2", widget.NewLabel("World!"))}...,
	)

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

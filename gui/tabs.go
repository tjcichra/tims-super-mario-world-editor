package gui

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createLevelTab(levelId int64) *container.TabItem {
	tabText := fmt.Sprintf("Level %X", levelId)
	return container.NewTabItem(tabText, widget.NewLabel("Level"))
}

func createGraphicsTab(graphicsId int64) *container.TabItem {
	tabText := fmt.Sprintf("Graphics %X", graphicsId)
	return container.NewTabItem(tabText, widget.NewLabel("Graphics"))
}

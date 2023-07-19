package gui

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createLevelToolbar() []widget.ToolbarItem {
	return []widget.ToolbarItem{
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	}
}

package main

import (
	"strconv"
	"tims-super-mario-world-editor/gui"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	selectARomWindow := gui.CreateSelectARomWindow(app)
	selectARomWindow.ShowAndRun()
}

func decimalToHex(decimal byte) string {
	return strconv.FormatInt(int64(decimal), 16)
}

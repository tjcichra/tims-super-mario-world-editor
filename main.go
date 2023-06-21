package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	selectARomWindow := createSelectARomWindow(app)
	selectARomWindow.ShowAndRun()
}

func decimalToHex(decimal byte) string {
	return strconv.FormatInt(int64(decimal), 16)
}

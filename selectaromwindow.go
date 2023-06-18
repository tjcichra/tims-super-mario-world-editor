package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

const TSMWE_DATA_FILE string = "tsmwe-data.json"

type TSMWEData struct {
	RecentRoms []string `json:"recent_roms"`
}

func createSelectARomWindow(app fyne.App, mainWindow fyne.Window) fyne.Window {
	tsmweData := readTSMWEData()
	recentRoms := tsmweData.RecentRoms
	fmt.Println("recentRoms", recentRoms)
	// tsmweData.RecentRoms[0] = "test"

	// filed, _ = json.MarshalIndent(tsmweData, "", " ")

	// os.WriteFile(TSMWE_DATA_FILE, filed, 0644)

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
		mainWindow.Show()
		selectARomWindow.Hide()
	}, selectARomWindow)
	openRomFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".smc"}))

	openRomButton := widget.NewButton("Open ROM", func() {
		openRomFileDialog.Show()
	})

	mostRecentLabel := widget.NewLabel("Most Recent:")

	mostRecentList := widget.NewList(
		func() int {
			return len(recentRoms)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			listButton := o.(*widget.Button)
			romPath := recentRoms[i]

			listButton.SetText(romPath)
			listButton.OnTapped = func() {
				mainWindow.Show()
				selectARomWindow.Hide()
			}
		})

	selectARomContainer1 := container.NewVBox(openRomButton, mostRecentLabel, mostRecentList)

	selectARomWindow.SetContent(selectARomContainer1)

	return selectARomWindow
}

func readTSMWEData() TSMWEData {
	var tsmweData TSMWEData

	file, err := os.ReadFile(TSMWE_DATA_FILE)
	if err != nil {
		tsmweData.RecentRoms = []string{}
		return tsmweData
	}

	json.Unmarshal(file, &tsmweData)

	return tsmweData
}

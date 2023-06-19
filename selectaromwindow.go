package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

const TSMWE_DATA_FILE string = "tsmwe-data.json"

type RecentRomEntry struct {
	RomPath      string `json:"romPath"`
	LastAccessed int64  `json:"lastAccessed"`
}

type TSMWEData struct {
	RecentRoms []RecentRomEntry `json:"recent_roms"`
}

func createSelectARomWindow(app fyne.App, mainWindow fyne.Window) fyne.Window {
	tsmweData := readTSMWEData()
	recentRoms := tsmweData.RecentRoms

	// Sort recent roms by when they were last accessed.
	sort.Slice(recentRoms, func(i, j int) bool {
		return recentRoms[i].LastAccessed > recentRoms[j].LastAccessed
	})

	selectARomWindow := app.NewWindow("Select a ROM")
	selectARomWindow.Resize(fyne.NewSize(300, 300))

	openRomFileDialog := dialog.NewFileOpen(func(uriReadCloser fyne.URIReadCloser, err error) {
		if uriReadCloser == nil {
			return
		}

		romPath := uriReadCloser.URI().Path()

		fmt.Println(romPath)

		dat, err := os.ReadFile(romPath)

		fmt.Println(decimalToHex(dat[0x00B9F6]), " ", decimalToHex(dat[0x00B9C4]), " ", decimalToHex(dat[0x00B992]))

		// decompressLZ2(dat[0x0881FD:])
		// gfx0file := decompressLZ2(dat[((0x08D9F9&0x7F0000)>>1)|(0x08D9F9&0x7FFF):])

		// var opts jpeg.Options
		// opts.Quality = 1

		// fmt.Println(i)
		updateRecentlyOpenedRoms(romPath, tsmweData)
		openRom(mainWindow, selectARomWindow)
	}, selectARomWindow)
	openRomFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".smc"}))

	openRomButton := widget.NewButton("Open ROM", func() {
		openRomFileDialog.Show()
	})

	recentlyOpenedRomsLabel := widget.NewLabel("Recently Opened ROMs:")

	recentlyOpenedRomsList := widget.NewList(
		func() int {
			return len(recentRoms)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", nil)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			listButton := o.(*widget.Button)
			romPath := recentRoms[i].RomPath

			listButton.SetText(romPath)
			listButton.OnTapped = func() {
				updateRecentlyOpenedRoms(romPath, tsmweData)
				openRom(mainWindow, selectARomWindow)
			}
		})

	selectARomContainer1 := container.NewVBox(openRomButton, recentlyOpenedRomsLabel, recentlyOpenedRomsList)

	selectARomWindow.SetContent(selectARomContainer1)

	return selectARomWindow
}

func readTSMWEData() TSMWEData {
	var tsmweData TSMWEData

	file, err := os.ReadFile(TSMWE_DATA_FILE)
	if err != nil {
		tsmweData.RecentRoms = []RecentRomEntry{}
		return tsmweData
	}

	json.Unmarshal(file, &tsmweData)

	return tsmweData
}

func updateRecentlyOpenedRoms(romPath string, tsmweData TSMWEData) {
	recentRoms := tsmweData.RecentRoms
	romPathIndex := slices.IndexFunc(recentRoms, func(e RecentRomEntry) bool { return e.RomPath == romPath })

	if romPathIndex == -1 {
		recentRoms = append(recentRoms, RecentRomEntry{
			RomPath:      romPath,
			LastAccessed: time.Now().Unix(),
		})
	} else {
		recentRoms[romPathIndex].LastAccessed = time.Now().Unix()
	}

	tsmweData.RecentRoms = recentRoms

	fileJson, _ := json.MarshalIndent(tsmweData, "", " ")

	os.WriteFile(TSMWE_DATA_FILE, fileJson, 0644)
}

func openRom(mainWindow fyne.Window, selectARomWindow fyne.Window) {
	mainWindow.Show()
	selectARomWindow.Hide()
}

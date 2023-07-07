package gui

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const NUMBER_OF_LEVELS = 0x1FF
const NUMBER_OF_GRAPHICS_FILES = 0x33

func createMainTree(tabs *container.AppTabs) *widget.Tree {
	var displayNameMap = make(map[string]string)

	levelNumbers := make([]widget.TreeNodeID, NUMBER_OF_LEVELS+1)
	for i := range levelNumbers {
		hexNumber := strings.ToUpper(strconv.FormatInt(int64(i), 16))
		levelId := "level" + hexNumber

		levelNumbers[i] = levelId
		displayNameMap[levelId] = hexNumber
	}

	graphicsNumbers := make([]widget.TreeNodeID, NUMBER_OF_GRAPHICS_FILES+1)
	for i := range graphicsNumbers {
		hexNumber := strings.ToUpper(strconv.FormatInt(int64(i), 16))
		graphicsId := "graphics" + hexNumber

		graphicsNumbers[i] = graphicsId
		displayNameMap[graphicsId] = hexNumber
	}

	overworldNumbers := []widget.TreeNodeID{"overworld0", "overworld1", "overworld2", "overworld3", "overworld4", "overworld5", "overworld6"}
	displayNameMap["overworld0"] = "Main Map"
	displayNameMap["overworld1"] = "Yoshi's Island"
	displayNameMap["overworld2"] = "Vanilla Dome"
	displayNameMap["overworld3"] = "Forest of Illusion"
	displayNameMap["overworld4"] = "Valley of Bowser"
	displayNameMap["overworld5"] = "Star World"
	displayNameMap["overworld6"] = "Special World"

	return widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"Levels", "Overworld Maps", "Graphics"}
			case "Levels":
				return levelNumbers
			case "Overworld Maps":
				return overworldNumbers
			case "Graphics":
				return graphicsNumbers
			}
			return []widget.TreeNodeID{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "Levels" || id == "Overworld Maps" || id == "Graphics"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewButton("Leaf template", nil)
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			if branch {
				o.(*widget.Label).SetText(id)
			} else {
				button := o.(*widget.Button)

				button.SetText(displayNameMap[id])
				button.OnTapped = func() {
					if strings.HasPrefix(id, "level") {
						levelId, _ := strconv.ParseInt(id[5:], 16, 64)

						tabs.Append(createLevelTab(levelId))
					} else if strings.HasPrefix(id, "overworld") {
						overworldId, _ := strconv.ParseInt(id[9:], 16, 64)

						fmt.Println("overworld", overworldId)
					} else if strings.HasPrefix(id, "graphics") {
						graphicsId, _ := strconv.ParseInt(id[8:], 16, 64)

						tabs.Append(createGraphicsTab(graphicsId))
					}
				}
			}
		})
}

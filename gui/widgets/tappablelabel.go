package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TappableLabel struct {
	widget.Label
	OnTapped func()
}

func NewTappableLabel(text string) *TappableLabel {
	label := &TappableLabel{}
	label.ExtendBaseWidget(label)
	label.SetText(text)

	return label
}

func (c *TappableLabel) Tapped(_ *fyne.PointEvent) {
	c.OnTapped()
}

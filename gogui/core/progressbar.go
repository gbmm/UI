package core

import (
	"strconv"
)

type ProgressBar struct {
	Wnd
	color    *Color
	progress int
	finish   bool
	Finish   func()
}

func (b *ProgressBar) Draw() {
	if b.color == nil {
		b.color = &Color{}
	}
	b.draw(b.color)
}

func (b *ProgressBar) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *ProgressBar) draw(color *Color) {
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{255, 255, 255})
	if b.Hidden {
		return
	}
	length := b.X + int(float32(b.progress)/100*float32(b.Width-60))
	if b.progress >= 100 {
		length = b.X + b.Width - 60
		b.progress = 100
		b.finish = true
		c := Color{0, 255, 0}
		b.Phy.FillRect(b.X, b.Y, length, b.Height+b.Y, c)
		b.Phy.DrawText(length, b.Y+(b.Height-FONT_HEIGHT)/2, strconv.Itoa(b.progress), c)
		if b.Finish != nil {
			b.Finish()
		}
	} else {
		b.Phy.FillRect(b.X, b.Y, length, b.Height+b.Y, *b.color)
		b.Phy.DrawText(length, b.Y+(b.Height-FONT_HEIGHT)/2, strconv.Itoa(b.progress), *b.color)
	}
}

func (b *ProgressBar) SetColor(c *Color) {
	b.color = c
}

func (b *ProgressBar) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *ProgressBar) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}

func (b *ProgressBar) SetProgress(progress int) {
	b.progress = progress
	b.Draw()
}

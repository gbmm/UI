package example

import (
	"UI/gogui/core"
	"strconv"
)

type Calculator struct {
	Widget *core.Wnd
	Btns   [15]*core.Button
	Input  *core.Input
	str1   string
	str2   string
	oper   string
}

func NewUICalculator(window *core.Window) *Calculator {
	ui := &Calculator{}
	ui.Widget = core.NewWnd(window, 0, 0, window.Width, window.Height)
	for i := 0; i < len(ui.Btns); i++ {
		ui.Btns[i] = &core.Button{}
	}
	ui.Input = &core.Input{}
	layout := core.NewGridLayout(ui.Widget, 5, 4)
	layout.AddChild(ui.Input, 0, 0, 1, 4)
	for i := 1; i <= 9; i++ {
		layout.AddChild(ui.Btns[i], (i-1)/3+1, (i-1)%3, 1, 1)
		ui.Btns[i].SetText(strconv.Itoa(i))
	}
	layout.AddChild(ui.Btns[0], 4, 0, 1, 1)
	ui.Btns[0].SetText("0")
	layout.AddChild(ui.Btns[10], 1, 3, 1, 1)
	ui.Btns[10].SetText("+")
	layout.AddChild(ui.Btns[11], 2, 3, 1, 1)
	ui.Btns[11].SetText("-")
	layout.AddChild(ui.Btns[12], 3, 3, 1, 1)
	ui.Btns[12].SetText("*")
	layout.AddChild(ui.Btns[13], 4, 3, 1, 1)
	ui.Btns[13].SetText("/")
	layout.AddChild(ui.Btns[14], 4, 1, 1, 2)
	ui.Btns[14].SetText("=")

	ui.Btns[0].Click = ui.Click0
	ui.Btns[1].Click = ui.Click1
	ui.Btns[2].Click = ui.Click2
	ui.Btns[3].Click = ui.Click3
	ui.Btns[4].Click = ui.Click4
	ui.Btns[5].Click = ui.Click5
	ui.Btns[6].Click = ui.Click6
	ui.Btns[7].Click = ui.Click7
	ui.Btns[8].Click = ui.Click8
	ui.Btns[9].Click = ui.Click9
	ui.Btns[10].Click = ui.Click10
	ui.Btns[11].Click = ui.Click11
	ui.Btns[12].Click = ui.Click12
	ui.Btns[13].Click = ui.Click13
	ui.Btns[14].Click = ui.Click14

	return ui
}

func (ui *Calculator) Click0() {
	if ui.oper == "" {
		ui.str1 += "0"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "0"
		ui.Input.SetText(ui.str2)
	}

}

func (ui *Calculator) Click1() {
	if ui.oper == "" {
		ui.str1 += "1"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "1"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click2() {
	if ui.oper == "" {
		ui.str1 += "2"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "2"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click3() {
	if ui.oper == "" {
		ui.str1 += "3"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "3"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click4() {
	if ui.oper == "" {
		ui.str1 += "4"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "4"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click5() {
	if ui.oper == "" {
		ui.str1 += "5"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "5"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click6() {
	if ui.oper == "" {
		ui.str1 += "6"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "6"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click7() {
	if ui.oper == "" {
		ui.str1 += "7"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "7"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click8() {
	if ui.oper == "" {
		ui.str1 += "8"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "8"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click9() {
	if ui.oper == "" {
		ui.str1 += "9"
		ui.Input.SetText(ui.str1)
	} else {
		ui.str2 += "9"
		ui.Input.SetText(ui.str2)
	}
}

func (ui *Calculator) Click10() {
	ui.oper = ""
	if ui.str1 != "" {
		ui.oper = "+"
	}
}

func (ui *Calculator) Click11() {
	ui.oper = ""
	if ui.str1 != "" {
		ui.oper = "-"
	}
}

func (ui *Calculator) Click12() {
	ui.oper = ""
	if ui.str1 != "" {
		ui.oper = "*"
	}
}

func (ui *Calculator) Click13() {
	ui.oper = ""
	if ui.str1 != "" {
		ui.oper = "/"
	}
}

func (ui *Calculator) Click14() {
	switch ui.oper {
	case "+":
		x1, _ := strconv.Atoi(ui.str1)
		x2, _ := strconv.Atoi(ui.str2)
		x := x1 + x2
		ui.oper = ""
		ui.str1 = ""
		ui.str2 = ""
		ui.Input.SetText(strconv.Itoa(x))
	case "-":
		x1, _ := strconv.Atoi(ui.str1)
		x2, _ := strconv.Atoi(ui.str2)
		x := x1 - x2
		ui.oper = ""
		ui.str1 = ""
		ui.str2 = ""
		ui.Input.SetText(strconv.Itoa(x))
	case "*":
		x1, _ := strconv.Atoi(ui.str1)
		x2, _ := strconv.Atoi(ui.str2)
		x := x1 * x2
		ui.oper = ""
		ui.str1 = ""
		ui.str2 = ""
		ui.Input.SetText(strconv.Itoa(x))
	case "/":
		x1, _ := strconv.Atoi(ui.str1)
		x2, _ := strconv.Atoi(ui.str2)
		x := x1 / x2
		ui.oper = ""
		ui.str1 = ""
		ui.str2 = ""
		ui.Input.SetText(strconv.Itoa(x))
	}
}

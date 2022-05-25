package main

import (
	"UI/gogui/core"
	"UI/gogui/example"
)

func main() {
	//go example.NewUIWidgets(core.NewApp(600, 480))
	//go example.NewUICalculator(core.NewApp(350, 480))
	// go example.NewUIWeather(core.NewApp(600, 480))
	//go example.NewUISnake(core.NewApp(600, 480))
	//go example.NewUISnake2(core.NewApp(600, 480))
	go example.NewUIQuicksort(core.NewApp(600, 480))
	go core.LoopEvent()
	select {}
}

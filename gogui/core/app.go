package core

import (
	"fmt"
	"reflect"
)

func NewApp(width, height int32) *Window {
	window := CreateWindow(width, height, "window")
	go UpdateUI(window)
	return window
}

func RunUI(ui interface{}) {
	var typeInfo = reflect.TypeOf(ui)
	var valInfo = reflect.ValueOf(ui)
	fmt.Println(typeInfo, valInfo)
	num := typeInfo.NumField()
	fmt.Println(num, typeInfo)
	for i := 0; i < num; i++ {
		k := typeInfo.Field(i)
		key := typeInfo.Field(i).Name
		val := valInfo.Field(i)
		fmt.Println(key, k.Type, val.Kind())

		switch val.Kind() {
		case reflect.Ptr:
			fmt.Println("-----------")
			if key == "Timer" || key == "Widget" {
				continue
			}
			fmt.Println(val.Elem().Addr().Interface().(WndInterface))
			// ui.Widget.AddChild(val.Elem().Addr().Interface().(core.WndInterface))
			// fmt.Println(val.Addr())
			// fmt.Println(val.Interface().(core.WndInterface))

		case reflect.Int:
			fmt.Println(val.Int())
		}
	}
}

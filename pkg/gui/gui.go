package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	hook "github.com/robotn/gohook"
	"screencarry/pkg/shot"
)


func Create() {
	myApp := app.New()
	myWin := myApp.NewWindow("截图工具")

	myWin.CenterOnScreen()
	myWin.SetFixedSize(true)
	label := widget.NewLabel("designed by mirrorlied - v1.0")
	btn1 := widget.NewButton("Shot!!!", func() {
		if l := shot.Shot(); l != "" {
			label.SetText("save: " + l)
		} else {
			label.SetText("designed by mirrorlied - v1.0")
		}
	})
	// Todo: 音效
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btn1, label)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(300, 80))
	add(myWin, label)
	myWin.ShowAndRun()
}

func add(win fyne.Window, label *widget.Label) {
	hook.Register(hook.KeyDown, []string{"q", "alt"}, func(e hook.Event) {
		fmt.Println("隐藏")
		win.Hide()
	})

	hook.Register(hook.KeyDown, []string{"w", "alt"}, func(e hook.Event) {
		fmt.Println("显示")
		win.CenterOnScreen()
		win.Show()
	})

	hook.Register(hook.KeyDown, []string{"a", "alt"}, func(e hook.Event) {
		fmt.Println("截图")
		if l := shot.Shot(); l != "" {
			label.SetText("save: " + l)
		} else {
			label.SetText("designed by mirrorlied - v1.0")
		}
	})

	s := hook.Start()
	go func() {
		<-hook.Process(s)
	}()
}
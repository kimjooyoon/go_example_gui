package exfyne

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_ "unicode/utf8"
)

func RunApp() {
	a := app.New()

	title := "Test Title"
	w := a.NewWindow(title)

	hello := widget.NewLabel("테디창 테스트")

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("출력", func() {
			hello.SetText("출력 버튼 클릭 됨")
		}),
	))

	w.ShowAndRun()
}

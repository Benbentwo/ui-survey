package main

import (
	"fyne.io/fyne/layout"

	binary "github.com/Benbentwo/ui-survey/app"
	cobraUi "github.com/Benbentwo/ui-survey/cobra-ui"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var data_struct struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const preferenceCurrentTab = "currentTab"

func main() {
	ex := app.New()
	Show(ex)
	ex.Run()
}

func Show(app fyne.App) {
	cmd := binary.Get(nil)
	window := app.NewWindow("Command Runner")
	cui := cobraUi.CobraUi{
		Command: cmd,
		// Window: window,
	}
	form := cui.NewForm()

	submit := cui.CreateRunCommandButton(false)
	buttons := widget.NewHBox(
		widget.NewButton("Cancel", func() {
			window.Close()
		}),
		layout.NewSpacer(),
		submit)
	window.SetContent(fyne.NewContainerWithLayout(
		nil, form, buttons, nil))
}

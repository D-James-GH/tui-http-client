package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = 0

	mainSection := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	methodSelect := newMethodDropdown()
	inputField := tview.NewInputField().
		SetLabel("Url:").
		SetFieldWidth(0)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			_, method := methodSelect.GetCurrentOption()
			url := inputField.GetText()
			req, err := http.NewRequest(method, url, nil)
			res, err := http.DefaultClient.Do(req)
			if err == nil {
				if b, err := io.ReadAll(res.Body); err == nil {
					b := string(b)
					mainSection.SetText(b)
				}
			}

		}
	})

	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(-2, -5).
		SetBorders(true).
		AddItem(methodSelect, 0, 0, 1, 1, 0, 0, true).
		AddItem(inputField, 0, 1, 1, 1, 0, 0, false).
		AddItem(mainSection, 1, 0, 1, 2, 0, 0, false)
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			app.SetFocus(inputField)
			return event
		}
		return event
	})

	if err := app.SetRoot(grid, true).SetFocus(grid).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func fetch(method string, url string, out *tview.TextView, app *tview.Application) {
	res, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if b, err := io.ReadAll(res.Body); err == nil {
		b := string(b)
		out.SetText(b)
	}
}

func newMethodDropdown() *tview.DropDown {
	selected := new(tcell.Style).
		Foreground(tcell.Color253).
		Background(tcell.Color235).
		Bold(true)
	unselected := new(tcell.Style).
		Foreground(tcell.Color235).
		Background(tcell.Color253).
		Bold(true)

	methodSelect := tview.NewDropDown().
		SetLabel("Method (hit Enter): ").
		SetOptions([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}, nil).
		SetCurrentOption(0).
		SetListStyles(selected, unselected).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorWhite)

	methodSelect.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			return tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone)
		case 'k':
			return tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone)
		}
		return event
	})

	return methodSelect
}

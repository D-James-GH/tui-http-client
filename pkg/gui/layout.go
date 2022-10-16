package gui

import (
	"github.com/d-james-gh/tui-http-client/pkg/state"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	View       tview.Primitive
	app        *tview.Application
	state      *state.State
	resultArea *tview.TextView
}

func NewLayout(state *state.State, app *tview.Application) *Layout {
	l := &Layout{
		state: state,
		app:   app,
	}
	main := l.MainArea()
	sidebar := l.SideBar()
	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(-1, -3).
		SetBorders(true).
		AddItem(sidebar, 0, 0, 2, 1, 0, 0, false).
		AddItem(main, 0, 1, 2, 1, 0, 0, true)

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})
	l.View = grid

	return l
}
func (l *Layout) MainArea() tview.Primitive {
	main := tview.NewFlex()
	main.SetDirection(tview.FlexColumnCSS)
	header := l.Header()
	main.AddItem(header, 0, 1, false)
	main.AddItem(l.ResultArea(), 0, 1, false)
	return main
}

func (l *Layout) SideBar() tview.Primitive {
	s := tview.NewBox()
	return s
}

func (l *Layout) Header() tview.Primitive {
	h := tview.NewFlex()
	url := l.UrlInput()

	h.AddItem(l.MethodSelect(), 0, 1, true)
	h.AddItem(url, 0, 1, false)
	return h

}

// An input for a url
func (l *Layout) UrlInput() *tview.InputField {
	inputField := tview.NewInputField().
		SetLabel("Url:").
		SetFieldWidth(0)

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			l.state.Url = inputField.GetText()
			l.state.SendRequest()
			l.UpdateResult()
		}
	})
	return inputField
}

// A dropdown to select a http method.
func (l *Layout) MethodSelect() tview.Primitive {
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
func (l *Layout) ResultArea() tview.Primitive {
	t := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).SetText(l.state.Result).SetChangedFunc(func() {
		l.app.Draw()
	})
	l.resultArea = t
	return t
}

func (l *Layout) UpdateResult() {
	l.resultArea.SetText(l.state.Result)
}

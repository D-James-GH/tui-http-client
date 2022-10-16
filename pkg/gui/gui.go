package gui

import (
	"log"

	"github.com/d-james-gh/tui-http-client/pkg/state"
	"github.com/rivo/tview"
)

type Gui struct {
	App    *tview.Application
	Layout *Layout
	State  *state.State
}

func NewGui() *Gui {
	app := tview.NewApplication()
	tview.Styles.PrimitiveBackgroundColor = 0
	s := new(state.State)
	l := NewLayout(s, app)

	app.SetRoot(l.View, true).SetFocus(l.View).EnableMouse(true)

	return &Gui{
		App:    app,
		Layout: l,
	}
}

func (g *Gui) Run() {
	err := g.App.Run()
	if err != nil {
		log.Fatalln(err)
	}

}

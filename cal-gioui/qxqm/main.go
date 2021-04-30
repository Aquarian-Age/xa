package main

import (
	"context"
	"gioui.org/unit"
	"os"
	"os/signal"
	"sync"

	"gioui.org/app"
	"gioui.org/widget/material"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		a := NewApplication(ctx)
		letters := NewBFYQWindow()
		a.NewWindow("演禽", letters,
			app.Size(unit.Dp(900), unit.Dp(600)))
		a.Wait()
		os.Exit(0)
	}()

	app.Main()
}

type Application struct {
	Context  context.Context
	Shutdown func()
	Theme    *material.Theme
	active   sync.WaitGroup
}

func NewApplication(ctx context.Context) *Application {
	ctx, cancel := context.WithCancel(ctx)
	return &Application{
		Context:  ctx,
		Shutdown: cancel,
		Theme:    utf8Font(),
	}
}

func (a *Application) Wait() {
	a.active.Wait()
}

func (a *Application) NewWindow(title string, view View, opts ...app.Option) {
	opts = append(opts, app.Title(title))
	w := &Window{
		App:    a,
		Window: app.NewWindow(opts...),
	}
	a.active.Add(1)
	go func() {
		defer a.active.Done()
		_ = view.Run(w)
	}()
}

type Window struct {
	App *Application
	*app.Window
}

type View interface {
	Run(w *Window) error
}

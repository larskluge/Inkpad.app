package main

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/alexflint/gallium"
	"github.com/atotto/clipboard"
)

type app struct {
	ui     *gallium.App
	window *gallium.Window
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Inkpad.log"))
	gallium.Loop(os.Args, onReady)
}

func (a *app) handleMenuLogin() {
	log.Println("login clicked")
	url, err := clipboard.ReadAll()
	check(err)

	if strings.HasPrefix(url, "https://www.inkpad.io/login") {
		a.window.LoadURL(url)
	} else {
		a.ui.Post(gallium.Notification{
			Title:    "Login Failed",
			Subtitle: "To login, please copy the login url you should have received via email into your clipboard and try again.",
		})
	}
}

func (a *app) handleMenuQuit() {
	log.Println("quit clicked")
	os.Exit(0)
}

func onReady(ui *gallium.App) {
	opts := gallium.FramedWindow
	opts.Title = "Inkpad"
	opts.Shape = gallium.Rect{
		Width:  1058,
		Height: 675,
		Bottom: 0,
		Left:   0,
	}
	window, err := ui.OpenWindow("https://www.inkpad.io/login", opts)
	check(err)

	app := app{
		ui:     ui,
		window: window,
	}

	ui.SetMenu([]gallium.Menu{
		{
			Title: "Inkpad",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:   "Login",
					OnClick: app.handleMenuLogin,
				},
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: gallium.MustParseKeys("cmd q"),
					OnClick:  app.handleMenuQuit,
				},
			},
		},
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

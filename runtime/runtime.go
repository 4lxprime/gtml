package runtime

import (
	"syscall/js"

	"github.com/4lxprime/gtml"
	"github.com/4lxprime/gtml/elements"
)

func setFunc(name string, fn func()) {
	js.Global().Set(
		name,
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				fn()
				return nil
			},
		),
	)
}

func Runtime(appElement *gtml.App) {
	// ---------------- Runtime Events ---->

	stopch := make(chan struct{})
	loadch := make(chan struct{})

	// javascript stop function, should be called by the wasm page
	setFunc("stop", func() {
		close(stopch)
	})

	// javascript end loading function, should be called by the wasm page
	setFunc("loaded", func() {
		close(loadch)
	})

	// ---------------- App ---->

	js.Global().Set(
		"app",
		elements.Build(
			appElement.Element,
		),
	)

	// ---------------- State Manager ---->

	// state manager start function
	setFunc("stateManagerStart", func() {
		go appElement.StateManager.Start()
	})
	// state manager stop function
	setFunc("stateManagerStop", func() {
		appElement.StateManager.Stop()
	})

	<-stopch
	js.Global().Call("stateManagerStop")
}

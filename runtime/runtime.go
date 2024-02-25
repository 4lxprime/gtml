package runtime

import (
	"syscall/js"

	"github.com/4lxprime/gtml"
	"github.com/4lxprime/gtml/elements"
)

func SetFunc(name string, fn func()) {
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

// todo: test this one
func CallFunc(name string, args ...interface{}) js.Value {
	jsFunc := js.Global().Get(name)
	return jsFunc.Invoke(args...)
}

func Runtime(appElement *gtml.App) {
	// ---------------- Runtime Events ---->

	stopch := make(chan struct{})
	loadch := make(chan struct{})

	// javascript stop function, should be called by the wasm page
	SetFunc("stop", func() {
		close(stopch)
	})

	// javascript end loading function, should be called by the wasm page
	SetFunc("loaded", func() {
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
	SetFunc("stateManagerStart", func() {
		go appElement.StateManager.Start()
	})
	// state manager stop function
	SetFunc("stateManagerStop", func() {
		appElement.StateManager.Stop()
	})

	<-stopch
	js.Global().Call("stateManagerStop")
}

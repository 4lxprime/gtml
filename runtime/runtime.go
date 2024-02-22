package runtime

import (
	"syscall/js"

	"github.com/4lxprime/agit/elements"
)

// func compile(output string) error {
// 	cmd := exec.Command("go", "build", "-o", output, "runtime/runtime.go")

// 	cmd.Env = append(cmd.Env, "GOOS=js")
// 	cmd.Env = append(cmd.Env, "GOARCH=wasm")

// 	if err := cmd.Run(); err != nil {
// 		return err
// 	}

// 	return nil
// }

func UseState(value any) chan any {
	datach := make(chan any)

	datach <- value

	return datach
}

func Runtime(appElement elements.Element) {
	stopch := make(chan struct{})
	loadch := make(chan struct{})

	// javascript stop function, should be called by the wasm page
	js.Global().Set("stop", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		close(stopch)
		return nil
	}))

	// javascript end loading function, should be called by the wasm page
	js.Global().Set("loaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		close(loadch)
		return nil
	}))

	js.Global().Set("app", elements.Build(appElement))

	<-stopch
}

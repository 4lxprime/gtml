package main

import (
	"fmt"

	"github.com/4lxprime/gtml"
	. "github.com/4lxprime/gtml/elements"
	"github.com/4lxprime/gtml/runtime"
)

func HeaderCmp(title string) Element {
	return Div()()
}

func App(app *gtml.App) *gtml.App {
	paddingState := app.UseState(20)

	return app.Use(
		Div(
			Style(
				"background-color: cyan;",
				"padding: 10px;",
			),
		)(
			gtml.If(1 == 3)(
				Text("mmmmmhh"),
			).Elif(3 == 3)(
				Text("super cool"),
			).Value(),

			gtml.For(0, 4)(
				func(i int) Element {
					return HeaderCmp(fmt.Sprintf("AGIT Wasm %d", i))
				},
			),

			gtml.Each([]string{"test", "Super test"})(
				func(i int, a any) Element {
					return P()(
						Text(a.(string)),
					)
				},
			),

			P(
				Style("color: blue;"),
			)(
				Text("Click on the button:"),
			),
			Button(
				Style(
					fmt.Sprintf(
						"padding: %vpx;",
						paddingState.Get(),
					),
					"margin: 5px;",
					"color: red;",
				),
				OnClick(func() {
					fmt.Println("button get clicked")
					paddingState.Set(paddingState.Get().(int) + 1)
				}),
			)(
				Text("Hello World!"),
			),
		),
	)
}

func main() {
	runtime.Runtime(
		App(gtml.NewApp()),
	)
}

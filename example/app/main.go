package main

import (
	"fmt"

	"github.com/4lxprime/gtml"
	. "github.com/4lxprime/gtml/elements"
	"github.com/4lxprime/gtml/runtime"
)

func HeaderCmp(title string) Element {
	return Div(
		Style(
			"position: absolute;",
			"top: 0; left: 0;",
			"margin-bottom: 10px;",
		),
	)(
		Div(
			Style("float: left;"),
		)(
			P(
				Style(
					"font-size: 100px;",
					"font-weight: bold;",
				),
				OnMouseOver(func() {
					fmt.Println("someone hover the logo")
				}),
			)(
				Text(title),
			),
		),
		Div(
			Style("float: right;"),
		)(
			Input(
				Type("text"),
				Placeholder("This must be cool"),
				Required,
			)(),
			Button(
				Style(
					"padding: 5px; margin: 5px;",
					"background-color: blue;",
				),
				OnClick(func() {
					fmt.Println("Login!!")
				}),
			)(
				Text("Login"),
			),
		),
	)
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
			HeaderCmp("AGIT Wasm"),
			P(
				Style("color: blue;"),
			)(
				Text("Click on the button:"),
			),
			Button(
				Style(fmt.Sprintf("padding: %vpx;", paddingState.Get())),
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

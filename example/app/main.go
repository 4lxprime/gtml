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

func Index(app *gtml.App) *gtml.App {
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

			gtml.Each2[string]([]string{
				"test",
				"Super test",
			})(func(i int, a string) Element {
				return P()(
					Text(a),
				)
			}),

			P(
				Style("color: blue;"),
			)(
				Text("Click on the button:"),
			),
			// here is an example of a custom element with custom attribute
			// we may need a way to block unusable attributes with syntax
			CustomElem[struct {
				Href             string
				MyCustomAttribut EventHandler
			}](
				"test",
				CustomAttr[EventHandler]("MyCustomAttribut", func() {
					fmt.Println("test")
				}),
				Href("https://google.com"),
			)(),
			Button(
				Style(
					fmt.Sprintf(
						"padding: %vpx;",
						paddingState.Get(),
					),
					"margin: 5px;",
					"color: red;",
				),
				Type("submit"),
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
		Index(
			gtml.NewApp(),
		),
	)
}

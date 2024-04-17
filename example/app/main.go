package main

import (
	"fmt"
	"time"

	"github.com/4lxprime/gtml"
	. "github.com/4lxprime/gtml/elements"
	"github.com/4lxprime/gtml/runtime"
)

func Index(app *gtml.App) *gtml.App {
	paddingState := app.UseState(20)

	clickP := P(
		Style("color: blue;"),
	)(
		Text("Click on the button:"),
	)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("should change")
		clickP.AppendChild(Text("tada"))
		Update(clickP)
	}()

	return app.Use(
		Div(
			Style(
				"background-color: cyan;",
				fmt.Sprintf("padding: %vpx;", paddingState.Get()),
			),
		)(
			gtml.If(1 == 3)(
				Text("mmmmmhh"),
			).Elif(3 == 3)(
				Text("super cool"),
			).Value(),

			gtml.Each2[string]([]string{
				"test",
				"super test",
				"super mega test",
			})(func(i int, a string) Element {
				return P()(
					Text(a),
				)
			}),

			clickP,
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

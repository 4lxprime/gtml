# GTML (GolangText Markup Language)
 golang reactive component frontend framwork using wasm
 ```golang
func App(app *gtml.App) *gtml.App {
	return app.Use(
		Div(
			Style(
				"background-color: cyan;",
				"padding: 10px;",
			),
		)(
			gtml.If(1 == 3)(
				Text("ok ... but there is maybe a bug"),
			).Elif(3 == 3)(
				Text("welcome!"),
			).Value(),

			gtml.Each([]string{"fast", "simple", "the best language"})(
				func(i int, a any) Element {
					return P()(
						Textf("golang is %s", a.(string)),
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
					"padding: 10px;",
					"color: red;",
				),
				OnClick(func() {
					fmt.Println("button get clicked!")
				}),
			)(
				Text("Hello World!"),
			),
		),
	)
}
 ```
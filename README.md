# GTML (GolangText Markup Language)
 golang reactive component frontend framwork using wasm
 ```golang
func Index(app *gtml.App) *gtml.App {
	return app.Use(
		Div(
			Style(
				"background-color: cyan;",
				"padding: 10px;",
			),
		)(
			P()(
				Text("Hello World")
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
 ```
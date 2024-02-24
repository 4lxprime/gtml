# GTML (GolangText Markup Language)
 golang reactive component frontend framwork using wasm
 ```golang
func App() Element {
	padding := 20

	return Div(
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
			Style(fmt.Sprintf("padding: %dpx;", padding)),
			OnClick(func() {
				fmt.Println("button get clicked")
				padding++
			}),
		)(
			Text("Hello World!"),
		),
	)
}
 ```
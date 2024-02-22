package main

import (
	"fmt"

	. "github.com/4lxprime/agit/elements"
	"github.com/4lxprime/agit/runtime"
)

func Menu(title string) Element {
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

func App() Element {
	padding := 20

	return Div(
		Style(
			"background-color: cyan;",
			"padding: 10px;",
		),
	)(
		Menu("AGIT Wasm"),
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
		Div()(),
	)
}

func main() {
	runtime.Runtime(App())
}

package elements

import (
	"fmt"
	"reflect"
	"strings"
	"syscall/js"
)

// Element interface represents our custom element and is a
// superset of DOMElement interface, the difference between
// Element and DOMElement is that Element represents a complete
// and buildable element with child elements and DOMElement just
// represents the standard and minimum required DOM element
type Element interface {
	GetChilds() []Element
	AppendChild(Element)
	GetElName() string
	// and so every dom element property
	DOMElement
}

// DOMElement interface represents every attributes of a
// standard D.O.M. element and will be used in every element
type DOMElement interface {
	GetAccessKey() string
	GetClass() string
	GetContentEditable() string
	GetContextMenu() string
	GetDir() string
	GetDraggable() string
	GetDropZone() string
	GetHidden() bool
	GetID() string
	GetInputMode() string
	GetLang() string
	GetSpellCheck() string
	GetStyle() string
	GetTabIndex() int64
	GetTitle() string
}

// BasicElement represents every attributes of a global
// D.O.M. element and will be used in every element.
//
// NOTE: this will implement the DOMElement interface
type BasicElement struct {
	AccessKey       string
	Class           string
	ContentEditable string
	ContextMenu     string
	Data            map[string]interface{} // data-*
	Dir             string
	Draggable       string
	DropZone        string
	Hidden          bool
	ID              string
	InputMode       string
	Lang            string
	SpellCheck      string
	Style           string
	TabIndex        int64
	Title           string
	// event handlers:

	OnClick     EventHandler
	OnDblClick  EventHandler
	OnMouseDown EventHandler
	OnMouseUp   EventHandler
	OnMouseMove EventHandler
	OnMouseOver EventHandler
	OnMouseOut  EventHandler
	OnKeyDown   EventHandler
	OnKeyUp     EventHandler
	OnFocus     EventHandler
	OnBlur      EventHandler
	OnChange    EventHandler
	OnSubmit    EventHandler
	OnReset     EventHandler
}

func (e BasicElement) GetAccessKey() string       { return e.AccessKey }
func (e BasicElement) GetClass() string           { return e.Class }
func (e BasicElement) GetContentEditable() string { return e.ContentEditable }
func (e BasicElement) GetContextMenu() string     { return e.ContextMenu }
func (e BasicElement) GetDir() string             { return e.Dir }
func (e BasicElement) GetDraggable() string       { return e.Draggable }
func (e BasicElement) GetDropZone() string        { return e.DropZone }
func (e BasicElement) GetHidden() bool            { return e.Hidden }
func (e BasicElement) GetID() string              { return e.ID }
func (e BasicElement) GetInputMode() string       { return e.InputMode }
func (e BasicElement) GetLang() string            { return e.Lang }
func (e BasicElement) GetSpellCheck() string      { return e.SpellCheck }
func (e BasicElement) GetStyle() string           { return e.Style }
func (e BasicElement) GetTabIndex() int64         { return e.TabIndex }
func (e BasicElement) GetTitle() string           { return e.Title }

func UseState(v any) (func() any, func(any)) {
	datach := make(chan any)
	datach <- v

	getFn := func() any {
		return <-datach
	}

	setFn := func(v any) {
		datach <- v
	}

	return getFn, setFn
}

// DOM builder
func Build(element Element) js.Func {
	return js.FuncOf(func(this js.Value, vals []js.Value) any {
		document := js.Global().Get("document")
		body := document.Get("body")

		// internal function that will spawn element
		// in the dom, attached to the given parent js element
		var buildElement func(elem Element, parent js.Value) js.Value
		buildElement = func(elem Element, parent js.Value) js.Value {
			switch el := elem.(type) {
			case TextEl: // fake element for text
				parent.Set("innerText", el.InnerText)
				fmt.Println(el.InnerText)

				// return parrent because we just added text content
				// in parent element and we don't spawn a new element
				return parent

			default:
				elementMap := fieldsToMap(el)
				fmt.Println(elementMap, el.GetElName(), el.GetChilds())
				jsElement := document.Call("createElement", el.GetElName())

				for attributeName, attributeValue := range elementMap {
					fmt.Println(attributeName, reflect.TypeOf(attributeValue))
					// with this we can do a specific logic for events
					switch attr := attributeValue.(type) {
					case EventHandler: // event handler case
						// event listener with custom action for the event handler
						jsElement.Call(
							"addEventListener",
							strings.ToLower(
								attributeName[2:], // after On (e.g. OnClick -> Click -> click)
							),
							js.FuncOf(func(this js.Value, vals []js.Value) any {
								attr()
								return nil
							}),
						)

					// and the normal attribute logic here
					default:
						jsElement.Set(
							strings.ToLower(attributeName),
							fmt.Sprintf("%v", attributeValue), // parsing interface{}
						)
					}
				}

				// loop over each child element and create the tree
				for _, child := range el.GetChilds() {
					buildElement(child, jsElement)
				}

				// spawn (append to the dom) the new element
				parent.Call("appendChild", jsElement)

				return jsElement
			}
		}

		buildElement(element, body)

		// runtime loaded function
		js.Global().Call("loaded")

		// runtime start the state manager
		js.Global().Call("stateManagerStart")

		return nil
	})
}

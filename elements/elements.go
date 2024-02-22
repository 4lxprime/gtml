package elements

import (
	"fmt"
	"log"
)

// this element is just an implementation of raw text
// and should not have neither children nor attributes
type TextEl struct {
	BasicElement
	InnerText string
	elName    string
}

func (e TextEl) GetChilds() []Element { return []Element{} }
func (e TextEl) GetElName() string    { return e.elName }

func Text(text string) TextEl {
	el := TextEl{elName: "rawtext"}

	el.InnerText = text

	return el
}

func Textf(format string, a ...any) TextEl {
	el := TextEl{elName: "rawtext"}

	el.InnerText = fmt.Sprintf(format, a...)

	return el
}

// ---------------- Div Element ----->

type DivEl struct {
	BasicElement
	childs []Element
	elName string
}

func (e DivEl) GetChilds() []Element { return e.childs }
func (e DivEl) GetElName() string    { return e.elName }

func Div(attributes ...Attribute) func(...Element) DivEl {
	el := DivEl{elName: "div"}

	// this will with reflection add the attribute to the good struct
	for _, attribute := range attributes {
		if err := setField(&el, attribute.Name, attribute.Value); err != nil {
			log.Println(err)
		}
	}

	return func(elements ...Element) DivEl {
		for _, element := range elements {
			el.childs = append(el.childs, element)
		}

		return el
	}
}

// ---------------- P Element ---->

type PEl struct {
	BasicElement
	childs []Element
	elName string
}

func (e PEl) GetChilds() []Element { return e.childs }
func (e PEl) GetElName() string    { return e.elName }

func P(attributes ...Attribute) func(...Element) Element {
	el := PEl{elName: "p"}

	// this will with reflection add the attribute to the good struct
	for _, attribute := range attributes {
		if err := setField(&el, attribute.Name, attribute.Value); err != nil {
			log.Println(err)
		}
	}

	return func(elements ...Element) Element {
		for _, element := range elements {
			el.childs = append(el.childs, element)
		}

		return el
	}
}

// ---------------- Button Element ---->

type ButtonEl struct {
	BasicElement
	Type           string
	Value          string
	Disabled       bool
	Form           string
	FormAction     string
	FormEncType    string
	FormMethod     string
	FormNoValidate bool
	FormTarget     string
	childs         []Element
	elName         string
}

func (e ButtonEl) GetChilds() []Element { return e.childs }
func (e ButtonEl) GetElName() string    { return e.elName }

func Button(attributes ...Attribute) func(...Element) ButtonEl {
	el := ButtonEl{elName: "button"}

	// this will with reflection add the attribute to the good struct
	for _, attribute := range attributes {
		if err := setField(&el, attribute.Name, attribute.Value); err != nil {
			log.Println(err)
		}
	}

	return func(elements ...Element) ButtonEl {
		for _, element := range elements {
			el.childs = append(el.childs, element)
		}

		return el
	}
}

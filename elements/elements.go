package elements

import (
	"fmt"
	"log"
	"syscall/js"
)

// this will help implementing every elements
// that can have childrens elements
func elementImpl(
	el interface{},
	attributes []Attribute,
) func(elements ...Element) Element {
	for _, attribute := range attributes {
		if err := setFieldValue(el, attribute.Name, attribute.Value); err != nil {
			log.Println(err)
		}
	}

	return func(elements ...Element) Element {
		for _, child := range elements {
			el.(Element).AppendChild(child)
		}

		return el.(Element)
	}
}

// this will help implementing every elements without childs
func elementAutoCloseImpl(
	el Element,
	attributes []Attribute,
) Element {
	// this will with reflection add the attribute to the good struct
	for _, attribute := range attributes {
		if err := setFieldValue(el, attribute.Name, attribute.Value); err != nil {
			log.Println(err)
		}
	}

	return el
}

// ---------------- Custom Elements ----->

// custom element that should be used in conditions
type EmptyEl struct {
	BasicElement
	elName  string
	ElValue js.Value
}

func (e *EmptyEl) GetChilds() []Element   { return []Element{} }
func (e *EmptyEl) AppendChild(el Element) {}
func (e *EmptyEl) GetElName() string      { return e.elName }
func (e *EmptyEl) GetElValue() js.Value   { return e.ElValue }

// empty value, will not be rendered in the DOM
var None = EmptyEl{
	elName: "none",
}

// todo: impl this at compile time
type CustomEl[T interface{}] struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
	Custom  T
}

func (e *CustomEl[T]) GetChilds() []Element   { return e.childs }
func (e *CustomEl[T]) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *CustomEl[T]) GetElName() string      { return e.elName }
func (e *CustomEl[T]) GetElValue() js.Value   { return e.ElValue }

func CustomElem[T interface{}](name string, attributes ...Attribute) func(...Element) Element {
	el := &CustomEl[T]{elName: name}
	return elementImpl(el, attributes)
}

type SliceEl struct {
	BasicElement
	childs  []Element
	ElValue js.Value
}

func (e *SliceEl) GetChilds() []Element   { return e.childs }
func (e *SliceEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SliceEl) GetElName() string      { return "slice" }
func (e *SliceEl) GetElValue() js.Value   { return e.ElValue }

// this element is just an implementation of raw text
// and should not have neither children nor attributes
type TextEl struct {
	BasicElement
	InnerText string
	elName    string
	ElValue   js.Value
}

func (e *TextEl) GetChilds() []Element   { return []Element{} }
func (e *TextEl) AppendChild(el Element) {}
func (e *TextEl) GetElName() string      { return e.elName }
func (e *TextEl) GetElValue() js.Value   { return e.ElValue }

func Text(text string) *TextEl {
	el := &TextEl{elName: "rawtext"}

	el.InnerText = text

	return el
}

func Textf(format string, a ...any) *TextEl {
	el := &TextEl{elName: "rawtext"}

	el.InnerText = fmt.Sprintf(format, a...)

	return el
}

// ---------------- Standard Elements ----->
// folowing elements struct are maybe wrong or not correcly implemented
// because i don't wanted to waste days looking for every elements attributes
// and specifications, so thanks GPT
// ----------------------------------------

type AEl struct {
	BasicElement
	Href    string
	Target  string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *AEl) GetChilds() []Element   { return e.childs }
func (e *AEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *AEl) GetElName() string      { return e.elName }
func (e *AEl) GetElValue() js.Value   { return e.ElValue }

func A(attributes ...Attribute) func(...Element) Element {
	el := &AEl{elName: "a"}
	return elementImpl(el, attributes)
}

type AbbrEl struct {
	BasicElement
	Title   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *AbbrEl) GetChilds() []Element   { return e.childs }
func (e *AbbrEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *AbbrEl) GetElName() string      { return e.elName }
func (e *AbbrEl) GetElValue() js.Value   { return e.ElValue }

func Abbr(attributes ...Attribute) func(...Element) Element {
	el := &AbbrEl{elName: "abbr"}
	return elementImpl(el, attributes)
}

type AddressEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *AddressEl) GetChilds() []Element   { return e.childs }
func (e *AddressEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *AddressEl) GetElName() string      { return e.elName }
func (e *AddressEl) GetElValue() js.Value   { return e.ElValue }

func Address(attributes ...Attribute) func(...Element) Element {
	el := &AddressEl{elName: "address"}
	return elementImpl(el, attributes)
}

type ArticleEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *ArticleEl) GetChilds() []Element   { return e.childs }
func (e *ArticleEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ArticleEl) GetElName() string      { return e.elName }
func (e *ArticleEl) GetElValue() js.Value   { return e.ElValue }

func Article(attributes ...Attribute) func(...Element) Element {
	el := &ArticleEl{elName: "article"}
	return elementImpl(el, attributes)
}

type AsideEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *AsideEl) GetChilds() []Element   { return e.childs }
func (e *AsideEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *AsideEl) GetElName() string      { return e.elName }
func (e *AsideEl) GetElValue() js.Value   { return e.ElValue }

func Aside(attributes ...Attribute) func(...Element) Element {
	el := &AsideEl{elName: "aside"}
	return elementImpl(el, attributes)
}

type AudioEl struct {
	BasicElement
	Src      string
	Controls bool
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *AudioEl) GetChilds() []Element   { return e.childs }
func (e *AudioEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *AudioEl) GetElName() string      { return e.elName }
func (e *AudioEl) GetElValue() js.Value   { return e.ElValue }

func Audio(attributes ...Attribute) func(...Element) Element {
	el := &AudioEl{elName: "audio"}
	return elementImpl(el, attributes)
}

type BaseEl struct {
	BasicElement
	Href    string
	Target  string
	elName  string
	ElValue js.Value
}

func (e *BaseEl) GetChilds() []Element   { return []Element{} }
func (e *BaseEl) AppendChild(el Element) {}
func (e *BaseEl) GetElName() string      { return e.elName }
func (e *BaseEl) GetElValue() js.Value   { return e.ElValue }

func Base(attributes ...Attribute) Element {
	el := &BaseEl{elName: "base"}
	return elementAutoCloseImpl(el, attributes)
}

type BdiEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *BdiEl) GetChilds() []Element   { return e.childs }
func (e *BdiEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *BdiEl) GetElName() string      { return e.elName }
func (e *BdiEl) GetElValue() js.Value   { return e.ElValue }

func Bdi(attributes ...Attribute) func(...Element) Element {
	el := &BdiEl{elName: "bdi"}
	return elementImpl(el, attributes)
}

type BdoEl struct {
	BasicElement
	Dir     string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *BdoEl) GetChilds() []Element   { return e.childs }
func (e *BdoEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *BdoEl) GetElName() string      { return e.elName }
func (e *BdoEl) GetElValue() js.Value   { return e.ElValue }

func Bdo(attributes ...Attribute) func(...Element) Element {
	el := &BdoEl{elName: "bdo"}
	return elementImpl(el, attributes)
}

type BlockquoteEl struct {
	BasicElement
	Cite    string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *BlockquoteEl) GetChilds() []Element   { return e.childs }
func (e *BlockquoteEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *BlockquoteEl) GetElName() string      { return e.elName }
func (e *BlockquoteEl) GetElValue() js.Value   { return e.ElValue }

func Blockquote(attributes ...Attribute) func(...Element) Element {
	el := &BlockquoteEl{elName: "blockquote"}
	return elementImpl(el, attributes)
}

type BrEl struct {
	BasicElement
	elName  string
	ElValue js.Value
}

func (e *BrEl) GetChilds() []Element   { return []Element{} }
func (e *BrEl) AppendChild(el Element) {}
func (e *BrEl) GetElName() string      { return e.elName }
func (e *BrEl) GetElValue() js.Value   { return e.ElValue }

func Br(attributes ...Attribute) Element {
	el := &BrEl{elName: "br"}
	return elementAutoCloseImpl(el, attributes)
}

type CanvasEl struct {
	BasicElement
	Width   int64
	Height  int64
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *CanvasEl) GetChilds() []Element   { return e.childs }
func (e *CanvasEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *CanvasEl) GetElName() string      { return e.elName }
func (e *CanvasEl) GetElValue() js.Value   { return e.ElValue }

func Canvas(attributes ...Attribute) func(...Element) Element {
	el := &CanvasEl{elName: "canvas"}
	return elementImpl(el, attributes)
}

type CaptionEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *CaptionEl) GetChilds() []Element   { return e.childs }
func (e *CaptionEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *CaptionEl) GetElName() string      { return e.elName }
func (e *CaptionEl) GetElValue() js.Value   { return e.ElValue }

func Caption(attributes ...Attribute) func(...Element) Element {
	el := &CaptionEl{elName: "caption"}
	return elementImpl(el, attributes)
}

type CiteEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *CiteEl) GetChilds() []Element   { return e.childs }
func (e *CiteEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *CiteEl) GetElName() string      { return e.elName }
func (e *CiteEl) GetElValue() js.Value   { return e.ElValue }

func Cite(attributes ...Attribute) func(...Element) Element {
	el := &CiteEl{elName: "cite"}
	return elementImpl(el, attributes)
}

type CodeEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *CodeEl) GetChilds() []Element   { return e.childs }
func (e *CodeEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *CodeEl) GetElName() string      { return e.elName }
func (e *CodeEl) GetElValue() js.Value   { return e.ElValue }

func Code(attributes ...Attribute) func(...Element) Element {
	el := &CodeEl{elName: "code"}
	return elementImpl(el, attributes)
}

type ColEl struct {
	BasicElement
	Span    int64
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *ColEl) GetChilds() []Element   { return e.childs }
func (e *ColEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ColEl) GetElName() string      { return e.elName }
func (e *ColEl) GetElValue() js.Value   { return e.ElValue }

func Col(attributes ...Attribute) func(...Element) Element {
	el := &ColEl{elName: "col"}
	return elementImpl(el, attributes)
}

type ColgroupEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *ColgroupEl) GetChilds() []Element   { return e.childs }
func (e *ColgroupEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ColgroupEl) GetElName() string      { return e.elName }
func (e *ColgroupEl) GetElValue() js.Value   { return e.ElValue }

func Colgroup(attributes ...Attribute) func(...Element) Element {
	el := &ColgroupEl{elName: "colgroup"}
	return elementImpl(el, attributes)
}

type DatalistEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DatalistEl) GetChilds() []Element   { return e.childs }
func (e *DatalistEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DatalistEl) GetElName() string      { return e.elName }
func (e *DatalistEl) GetElValue() js.Value   { return e.ElValue }

func Datalist(attributes ...Attribute) func(...Element) Element {
	el := &DatalistEl{elName: "datalist"}
	return elementImpl(el, attributes)
}

type DataEl struct {
	BasicElement
	Value   string
	elName  string
	ElValue js.Value
}

type DdEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DdEl) GetChilds() []Element   { return e.childs }
func (e *DdEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DdEl) GetElName() string      { return e.elName }
func (e *DdEl) GetElValue() js.Value   { return e.ElValue }

func Dd(attributes ...Attribute) func(...Element) Element {
	el := &DdEl{elName: "dd"}
	return elementImpl(el, attributes)
}

type DelEl struct {
	BasicElement
	Cite     string
	DateTime string
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *DelEl) GetChilds() []Element   { return e.childs }
func (e *DelEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DelEl) GetElName() string      { return e.elName }
func (e *DelEl) GetElValue() js.Value   { return e.ElValue }

func Del(attributes ...Attribute) func(...Element) Element {
	el := &DelEl{elName: "del"}
	return elementImpl(el, attributes)
}

type DetailsEl struct {
	BasicElement
	Open    bool
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DetailsEl) GetChilds() []Element   { return e.childs }
func (e *DetailsEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DetailsEl) GetElName() string      { return e.elName }
func (e *DetailsEl) GetElValue() js.Value   { return e.ElValue }

func Details(attributes ...Attribute) func(...Element) Element {
	el := &DetailsEl{elName: "details"}
	return elementImpl(el, attributes)
}

type DfnEl struct {
	BasicElement
	Title   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DfnEl) GetChilds() []Element   { return e.childs }
func (e *DfnEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DfnEl) GetElName() string      { return e.elName }
func (e *DfnEl) GetElValue() js.Value   { return e.ElValue }

func Dfn(attributes ...Attribute) func(...Element) Element {
	el := &DfnEl{elName: "dfn"}
	return elementImpl(el, attributes)
}

type DialogEl struct {
	BasicElement
	Open    bool
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DialogEl) GetChilds() []Element   { return e.childs }
func (e *DialogEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DialogEl) GetElName() string      { return e.elName }
func (e *DialogEl) GetElValue() js.Value   { return e.ElValue }

func Dialog(attributes ...Attribute) func(...Element) Element {
	el := &DialogEl{elName: "dialog"}
	return elementImpl(el, attributes)
}

type DlEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DlEl) GetChilds() []Element   { return e.childs }
func (e *DlEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DlEl) GetElName() string      { return e.elName }
func (e *DlEl) GetElValue() js.Value   { return e.ElValue }

func Dl(attributes ...Attribute) func(...Element) Element {
	el := &DlEl{elName: "dl"}
	return elementImpl(el, attributes)
}

type DtEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DtEl) GetChilds() []Element   { return e.childs }
func (e *DtEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DtEl) GetElName() string      { return e.elName }
func (e *DtEl) GetElValue() js.Value   { return e.ElValue }

func Dt(attributes ...Attribute) func(...Element) Element {
	el := &DtEl{elName: "dt"}
	return elementImpl(el, attributes)
}

type EmEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *EmEl) GetChilds() []Element   { return e.childs }
func (e *EmEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *EmEl) GetElName() string      { return e.elName }
func (e *EmEl) GetElValue() js.Value   { return e.ElValue }

func Em(attributes ...Attribute) func(...Element) Element {
	el := &EmEl{elName: "em"}
	return elementImpl(el, attributes)
}

type EmbedEl struct {
	BasicElement
	Src     string
	Type    string
	Width   int64
	Height  int64
	elName  string
	ElValue js.Value
}

func (e *EmbedEl) GetChilds() []Element   { return []Element{} }
func (e *EmbedEl) AppendChild(el Element) {}
func (e *EmbedEl) GetElName() string      { return e.elName }
func (e *EmbedEl) GetElValue() js.Value   { return e.ElValue }

func Embed(attributes ...Attribute) func(...Element) Element {
	el := &EmbedEl{elName: "embed"}
	return elementImpl(el, attributes)
} // todo

type FieldsetEl struct {
	BasicElement
	Disabled bool
	Form     string
	Name     string
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *FieldsetEl) GetChilds() []Element   { return e.childs }
func (e *FieldsetEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *FieldsetEl) GetElName() string      { return e.elName }
func (e *FieldsetEl) GetElValue() js.Value   { return e.ElValue }

func Fieldset(attributes ...Attribute) func(...Element) Element {
	el := &FieldsetEl{elName: "fieldset"}
	return elementImpl(el, attributes)
}

type FigcaptionEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *FigcaptionEl) GetChilds() []Element   { return e.childs }
func (e *FigcaptionEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *FigcaptionEl) GetElName() string      { return e.elName }
func (e *FigcaptionEl) GetElValue() js.Value   { return e.ElValue }

func Figcaption(attributes ...Attribute) func(...Element) Element {
	el := &FigcaptionEl{elName: "figcaption"}
	return elementImpl(el, attributes)
}

type FigureEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *FigureEl) GetChilds() []Element   { return e.childs }
func (e *FigureEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *FigureEl) GetElName() string      { return e.elName }
func (e *FigureEl) GetElValue() js.Value   { return e.ElValue }

func Figure(attributes ...Attribute) func(...Element) Element {
	el := &FigureEl{elName: "figure"}
	return elementImpl(el, attributes)
}

type FooterEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *FooterEl) GetChilds() []Element   { return e.childs }
func (e *FooterEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *FooterEl) GetElName() string      { return e.elName }
func (e *FooterEl) GetElValue() js.Value   { return e.ElValue }

func Footer(attributes ...Attribute) func(...Element) Element {
	el := &FooterEl{elName: "footer"}
	return elementImpl(el, attributes)
}

type FormEl struct {
	BasicElement
	Action       string
	Autocomplete string
	Method       string
	Name         string
	NoValidate   bool
	Target       string
	childs       []Element
	elName       string
	ElValue      js.Value
}

func (e *FormEl) GetChilds() []Element   { return e.childs }
func (e *FormEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *FormEl) GetElName() string      { return e.elName }
func (e *FormEl) GetElValue() js.Value   { return e.ElValue }

func Form(attributes ...Attribute) func(...Element) Element {
	el := &FormEl{elName: "form"}
	return elementImpl(el, attributes)
}

type H1El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H1El) GetChilds() []Element   { return e.childs }
func (e *H1El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H1El) GetElName() string      { return e.elName }
func (e *H1El) GetElValue() js.Value   { return e.ElValue }

func H1(attributes ...Attribute) func(...Element) Element {
	el := &H1El{elName: "h1"}
	return elementImpl(el, attributes)
}

type H2El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H2El) GetChilds() []Element   { return e.childs }
func (e *H2El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H2El) GetElName() string      { return e.elName }
func (e *H2El) GetElValue() js.Value   { return e.ElValue }

func H2(attributes ...Attribute) func(...Element) Element {
	el := &H2El{elName: "h2"}
	return elementImpl(el, attributes)
}

type H3El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H3El) GetChilds() []Element   { return e.childs }
func (e *H3El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H3El) GetElName() string      { return e.elName }
func (e *H3El) GetElValue() js.Value   { return e.ElValue }

func H3(attributes ...Attribute) func(...Element) Element {
	el := &H3El{elName: "h3"}
	return elementImpl(el, attributes)
}

type H4El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H4El) GetChilds() []Element   { return e.childs }
func (e *H4El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H4El) GetElName() string      { return e.elName }
func (e *H4El) GetElValue() js.Value   { return e.ElValue }

func H4(attributes ...Attribute) func(...Element) Element {
	el := &H4El{elName: "h4"}
	return elementImpl(el, attributes)
}

type H5El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H5El) GetChilds() []Element   { return e.childs }
func (e *H5El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H5El) GetElName() string      { return e.elName }
func (e *H5El) GetElValue() js.Value   { return e.ElValue }

func H5(attributes ...Attribute) func(...Element) Element {
	el := &H5El{elName: "h5"}
	return elementImpl(el, attributes)
}

type H6El struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *H6El) GetChilds() []Element   { return e.childs }
func (e *H6El) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *H6El) GetElName() string      { return e.elName }
func (e *H6El) GetElValue() js.Value   { return e.ElValue }

func H6(attributes ...Attribute) func(...Element) Element {
	el := &H6El{elName: "h6"}
	return elementImpl(el, attributes)
}

type HrEl struct {
	BasicElement
	elName  string
	ElValue js.Value
}

func (e *HrEl) GetChilds() []Element   { return []Element{} }
func (e *HrEl) AppendChild(el Element) {}
func (e *HrEl) GetElName() string      { return e.elName }
func (e *HrEl) GetElValue() js.Value   { return e.ElValue }

func Hr(attributes ...Attribute) func(...Element) Element {
	el := &HrEl{elName: "hr"}
	return elementImpl(el, attributes)
} // todo

type IEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *IEl) GetChilds() []Element   { return e.childs }
func (e *IEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *IEl) GetElName() string      { return e.elName }
func (e *IEl) GetElValue() js.Value   { return e.ElValue }

func I(attributes ...Attribute) func(...Element) Element {
	el := &IEl{elName: "i"}
	return elementImpl(el, attributes)
}

type IframeEl struct {
	BasicElement
	Src                 string
	SrcDoc              string
	Width               string
	Height              string
	Name                string
	Sandbox             string
	AllowFullscreen     bool
	AllowPaymentRequest bool
	ReferrerPolicy      string
	elName              string
	ElValue             js.Value
}

func (e *IframeEl) GetChilds() []Element   { return []Element{} }
func (e *IframeEl) AppendChild(el Element) {}
func (e *IframeEl) GetElName() string      { return e.elName }
func (e *IframeEl) GetElValue() js.Value   { return e.ElValue }

func Iframe(attributes ...Attribute) func(...Element) Element {
	el := &IframeEl{elName: "iframe"}
	return elementImpl(el, attributes)
} // todo

type ImgEl struct {
	BasicElement
	Src         string
	Alt         string
	Width       string
	Height      string
	Loading     string
	CrossOrigin string
	UseMap      string
	IsMap       bool
	Sizes       string
	SrcSet      string
	elName      string
	ElValue     js.Value
}

func (e *ImgEl) GetChilds() []Element   { return []Element{} }
func (e *ImgEl) AppendChild(el Element) {}
func (e *ImgEl) GetElName() string      { return e.elName }
func (e *ImgEl) GetElValue() js.Value   { return e.ElValue }

func Img(attributes ...Attribute) func(...Element) Element {
	el := &ImgEl{elName: "img"}
	return elementImpl(el, attributes)
} // todo

type InsEl struct {
	BasicElement
	Cite     string
	DateTime string
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *InsEl) GetChilds() []Element   { return e.childs }
func (e *InsEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *InsEl) GetElName() string      { return e.elName }
func (e *InsEl) GetElValue() js.Value   { return e.ElValue }

func Ins(attributes ...Attribute) func(...Element) Element {
	el := &InsEl{elName: "ins"}
	return elementImpl(el, attributes)
}

type KbdEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *KbdEl) GetChilds() []Element   { return e.childs }
func (e *KbdEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *KbdEl) GetElName() string      { return e.elName }
func (e *KbdEl) GetElValue() js.Value   { return e.ElValue }

func Kbd(attributes ...Attribute) func(...Element) Element {
	el := &KbdEl{elName: "kbd"}
	return elementImpl(el, attributes)
}

type LegendEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *LegendEl) GetChilds() []Element   { return e.childs }
func (e *LegendEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *LegendEl) GetElName() string      { return e.elName }
func (e *LegendEl) GetElValue() js.Value   { return e.ElValue }

func Legend(attributes ...Attribute) func(...Element) Element {
	el := &LegendEl{elName: "legend"}
	return elementImpl(el, attributes)
}

type LiEl struct {
	BasicElement
	Value   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *LiEl) GetChilds() []Element   { return e.childs }
func (e *LiEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *LiEl) GetElName() string      { return e.elName }
func (e *LiEl) GetElValue() js.Value   { return e.ElValue }

func Li(attributes ...Attribute) func(...Element) Element {
	el := &LiEl{elName: "li"}
	return elementImpl(el, attributes)
}

type LinkEl struct {
	BasicElement
	Href           string
	Rel            string
	HrefLang       string
	Media          string
	Sizes          string
	Type           string
	ReferrerPolicy string
	As             string
	CrossOrigin    string
	Integrity      string
	Nonce          string
	childs         []Element
	elName         string
	ElValue        js.Value
}

func (e *LinkEl) GetChilds() []Element   { return e.childs }
func (e *LinkEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *LinkEl) GetElName() string      { return e.elName }
func (e *LinkEl) GetElValue() js.Value   { return e.ElValue }

func Link(attributes ...Attribute) func(...Element) Element {
	el := &LinkEl{elName: "link"}
	return elementImpl(el, attributes)
}

type MainEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *MainEl) GetChilds() []Element   { return e.childs }
func (e *MainEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *MainEl) GetElName() string      { return e.elName }
func (e *MainEl) GetElValue() js.Value   { return e.ElValue }

func Main(attributes ...Attribute) func(...Element) Element {
	el := &MainEl{elName: "main"}
	return elementImpl(el, attributes)
}

type MapEl struct {
	BasicElement
	Name    string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *MapEl) GetChilds() []Element   { return e.childs }
func (e *MapEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *MapEl) GetElName() string      { return e.elName }
func (e *MapEl) GetElValue() js.Value   { return e.ElValue }

func Map(attributes ...Attribute) func(...Element) Element {
	el := &MapEl{elName: "map"}
	return elementImpl(el, attributes)
}

type MarkEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *MarkEl) GetChilds() []Element   { return e.childs }
func (e *MarkEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *MarkEl) GetElName() string      { return e.elName }
func (e *MarkEl) GetElValue() js.Value   { return e.ElValue }

func Mark(attributes ...Attribute) func(...Element) Element {
	el := &MarkEl{elName: "mark"}
	return elementImpl(el, attributes)
}

type MenuEl struct {
	BasicElement
	Type    string
	Label   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *MenuEl) GetChilds() []Element   { return e.childs }
func (e *MenuEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *MenuEl) GetElName() string      { return e.elName }
func (e *MenuEl) GetElValue() js.Value   { return e.ElValue }

func Menu(attributes ...Attribute) func(...Element) Element {
	el := &MenuEl{elName: "menu"}
	return elementImpl(el, attributes)
}

type MenuItemEl struct {
	BasicElement
	Type    string
	Label   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *MenuItemEl) GetChilds() []Element   { return e.childs }
func (e *MenuItemEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *MenuItemEl) GetElName() string      { return e.elName }
func (e *MenuItemEl) GetElValue() js.Value   { return e.ElValue }

func MenuItem(attributes ...Attribute) func(...Element) Element {
	el := &MenuItemEl{elName: "menuitem"}
	return elementImpl(el, attributes)
}

type MeterEl struct {
	BasicElement
	Value   int64
	Min     int64
	Max     int64
	Low     int64
	High    int64
	Optimum int64
	elName  string
	ElValue js.Value
}

func (e *MeterEl) GetChilds() []Element   { return []Element{} }
func (e *MeterEl) AppendChild(el Element) {}
func (e *MeterEl) GetElName() string      { return e.elName }
func (e *MeterEl) GetElValue() js.Value   { return e.ElValue }

func Meter(attributes ...Attribute) func(...Element) Element {
	el := &MeterEl{elName: "meter"}
	return elementImpl(el, attributes)
} // todo

type NavEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *NavEl) GetChilds() []Element   { return e.childs }
func (e *NavEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *NavEl) GetElName() string      { return e.elName }
func (e *NavEl) GetElValue() js.Value   { return e.ElValue }

func Nav(attributes ...Attribute) func(...Element) Element {
	el := &NavEl{elName: "nav"}
	return elementImpl(el, attributes)
}

type ObjectEl struct {
	BasicElement
	Data    string
	Type    string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *ObjectEl) GetChilds() []Element   { return e.childs }
func (e *ObjectEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ObjectEl) GetElName() string      { return e.elName }
func (e *ObjectEl) GetElValue() js.Value   { return e.ElValue }

func Object(attributes ...Attribute) func(...Element) Element {
	el := &ObjectEl{elName: "object"}
	return elementImpl(el, attributes)
}

type OlEl struct {
	BasicElement
	Type     string
	Reversed bool
	Start    int64
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *OlEl) GetChilds() []Element   { return e.childs }
func (e *OlEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *OlEl) GetElName() string      { return e.elName }
func (e *OlEl) GetElValue() js.Value   { return e.ElValue }

func Ol(attributes ...Attribute) func(...Element) Element {
	el := &OlEl{elName: "ol"}
	return elementImpl(el, attributes)
}

type OptGroupEl struct {
	BasicElement
	Label   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *OptGroupEl) GetChilds() []Element   { return e.childs }
func (e *OptGroupEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *OptGroupEl) GetElName() string      { return e.elName }
func (e *OptGroupEl) GetElValue() js.Value   { return e.ElValue }

func OptGroup(attributes ...Attribute) func(...Element) Element {
	el := &OptGroupEl{elName: "optgroup"}
	return elementImpl(el, attributes)
}

type OptionEl struct {
	BasicElement
	Value    string
	Selected bool
	Disabled bool
	elName   string
	ElValue  js.Value
}

func (e *OptionEl) GetChilds() []Element   { return []Element{} }
func (e *OptionEl) AppendChild(el Element) {}
func (e *OptionEl) GetElName() string      { return e.elName }
func (e *OptionEl) GetElValue() js.Value   { return e.ElValue }

func Option(attributes ...Attribute) func(...Element) Element {
	el := &OptionEl{elName: "option"}
	return elementImpl(el, attributes)
} // todo

type OutputEl struct {
	BasicElement
	For     string
	Name    string
	Value   string
	elName  string
	ElValue js.Value
}

func (e *OutputEl) GetChilds() []Element   { return []Element{} }
func (e *OutputEl) AppendChild(el Element) {}
func (e *OutputEl) GetElName() string      { return e.elName }
func (e *OutputEl) GetElValue() js.Value   { return e.ElValue }

func Output(attributes ...Attribute) func(...Element) Element {
	el := &OutputEl{elName: "output"}
	return elementImpl(el, attributes)
} // todo

type ParamEl struct {
	BasicElement
	Name    string
	Value   string
	elName  string
	ElValue js.Value
}

func (e *ParamEl) GetChilds() []Element   { return []Element{} }
func (e *ParamEl) AppendChild(el Element) {}
func (e *ParamEl) GetElName() string      { return e.elName }
func (e *ParamEl) GetElValue() js.Value   { return e.ElValue }

func Param(attributes ...Attribute) func(...Element) Element {
	el := &ParamEl{elName: "param"}
	return elementImpl(el, attributes)
} // todo

type PictureEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *PictureEl) GetChilds() []Element   { return e.childs }
func (e *PictureEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *PictureEl) GetElName() string      { return e.elName }
func (e *PictureEl) GetElValue() js.Value   { return e.ElValue }

func Picture(attributes ...Attribute) func(...Element) Element {
	el := &PictureEl{elName: "picture"}
	return elementImpl(el, attributes)
}

type PreEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *PreEl) GetChilds() []Element   { return e.childs }
func (e *PreEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *PreEl) GetElName() string      { return e.elName }
func (e *PreEl) GetElValue() js.Value   { return e.ElValue }

func Pre(attributes ...Attribute) func(...Element) Element {
	el := &PreEl{elName: "pre"}
	return elementImpl(el, attributes)
}

type ProgressEl struct {
	BasicElement
	Value   int64
	Max     int64
	elName  string
	ElValue js.Value
}

func (e *ProgressEl) GetChilds() []Element   { return []Element{} }
func (e *ProgressEl) AppendChild(el Element) {}
func (e *ProgressEl) GetElName() string      { return e.elName }
func (e *ProgressEl) GetElValue() js.Value   { return e.ElValue }

func Progress(attributes ...Attribute) func(...Element) Element {
	el := &ProgressEl{elName: "progress"}
	return elementImpl(el, attributes)
} // todo

type QEl struct {
	BasicElement
	Cite    string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *QEl) GetChilds() []Element   { return e.childs }
func (e *QEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *QEl) GetElName() string      { return e.elName }
func (e *QEl) GetElValue() js.Value   { return e.ElValue }

func Q(attributes ...Attribute) func(...Element) Element {
	el := &QEl{elName: "q"}
	return elementImpl(el, attributes)
}

type RpEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *RpEl) GetChilds() []Element   { return e.childs }
func (e *RpEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *RpEl) GetElName() string      { return e.elName }
func (e *RpEl) GetElValue() js.Value   { return e.ElValue }

func Rp(attributes ...Attribute) func(...Element) Element {
	el := &RpEl{elName: "rp"}
	return elementImpl(el, attributes)
}

type RtEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *RtEl) GetChilds() []Element   { return e.childs }
func (e *RtEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *RtEl) GetElName() string      { return e.elName }
func (e *RtEl) GetElValue() js.Value   { return e.ElValue }

func Rt(attributes ...Attribute) func(...Element) Element {
	el := &RtEl{elName: "rt"}
	return elementImpl(el, attributes)
}

type RubyEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *RubyEl) GetChilds() []Element   { return e.childs }
func (e *RubyEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *RubyEl) GetElName() string      { return e.elName }
func (e *RubyEl) GetElValue() js.Value   { return e.ElValue }

func Ruby(attributes ...Attribute) func(...Element) Element {
	el := &RubyEl{elName: "ruby"}
	return elementImpl(el, attributes)
}

type SEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SEl) GetChilds() []Element   { return e.childs }
func (e *SEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SEl) GetElName() string      { return e.elName }
func (e *SEl) GetElValue() js.Value   { return e.ElValue }

func S(attributes ...Attribute) func(...Element) Element {
	el := &SEl{elName: "s"}
	return elementImpl(el, attributes)
}

type SampEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SampEl) GetChilds() []Element   { return e.childs }
func (e *SampEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SampEl) GetElName() string      { return e.elName }
func (e *SampEl) GetElValue() js.Value   { return e.ElValue }

func Samp(attributes ...Attribute) func(...Element) Element {
	el := &SampEl{elName: "samp"}
	return elementImpl(el, attributes)
}

type SectionEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SectionEl) GetChilds() []Element   { return e.childs }
func (e *SectionEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SectionEl) GetElName() string      { return e.elName }
func (e *SectionEl) GetElValue() js.Value   { return e.ElValue }

func Section(attributes ...Attribute) func(...Element) Element {
	el := &SectionEl{elName: "section"}
	return elementImpl(el, attributes)
}

type SelectEl struct {
	BasicElement
	Name     string
	Size     int64
	Multiple bool
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *SelectEl) GetChilds() []Element   { return e.childs }
func (e *SelectEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SelectEl) GetElName() string      { return e.elName }
func (e *SelectEl) GetElValue() js.Value   { return e.ElValue }

func Select(attributes ...Attribute) func(...Element) Element {
	el := &SelectEl{elName: "select"}
	return elementImpl(el, attributes)
}

type SmallEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SmallEl) GetChilds() []Element   { return e.childs }
func (e *SmallEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SmallEl) GetElName() string      { return e.elName }
func (e *SmallEl) GetElValue() js.Value   { return e.ElValue }

func Small(attributes ...Attribute) func(...Element) Element {
	el := &SmallEl{elName: "small"}
	return elementImpl(el, attributes)
}

type SourceEl struct {
	BasicElement
	Src     string
	Type    string
	Srcset  string
	Sizes   string
	Media   string
	elName  string
	ElValue js.Value
}

func (e *SourceEl) GetChilds() []Element   { return []Element{} }
func (e *SourceEl) AppendChild(el Element) {}
func (e *SourceEl) GetElName() string      { return e.elName }
func (e *SourceEl) GetElValue() js.Value   { return e.ElValue }

func Source(attributes ...Attribute) func(...Element) Element {
	el := &SourceEl{elName: "source"}
	return elementImpl(el, attributes)
} // todo

type SpanEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SpanEl) GetChilds() []Element   { return e.childs }
func (e *SpanEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SpanEl) GetElName() string      { return e.elName }
func (e *SpanEl) GetElValue() js.Value   { return e.ElValue }

func Span(attributes ...Attribute) func(...Element) Element {
	el := &SpanEl{elName: "span"}
	return elementImpl(el, attributes)
}

type StrongEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *StrongEl) GetChilds() []Element   { return e.childs }
func (e *StrongEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *StrongEl) GetElName() string      { return e.elName }
func (e *StrongEl) GetElValue() js.Value   { return e.ElValue }

func Strong(attributes ...Attribute) func(...Element) Element {
	el := &StrongEl{elName: "strong"}
	return elementImpl(el, attributes)
}

type SubEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SubEl) GetChilds() []Element   { return e.childs }
func (e *SubEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SubEl) GetElName() string      { return e.elName }
func (e *SubEl) GetElValue() js.Value   { return e.ElValue }

func Sub(attributes ...Attribute) func(...Element) Element {
	el := &SubEl{elName: "sub"}
	return elementImpl(el, attributes)
}

type SummaryEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SummaryEl) GetChilds() []Element   { return e.childs }
func (e *SummaryEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SummaryEl) GetElName() string      { return e.elName }
func (e *SummaryEl) GetElValue() js.Value   { return e.ElValue }

func Summary(attributes ...Attribute) func(...Element) Element {
	el := &SummaryEl{elName: "summary"}
	return elementImpl(el, attributes)
}

type SupEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *SupEl) GetChilds() []Element   { return e.childs }
func (e *SupEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *SupEl) GetElName() string      { return e.elName }
func (e *SupEl) GetElValue() js.Value   { return e.ElValue }

func Sup(attributes ...Attribute) func(...Element) Element {
	el := &SupEl{elName: "sup"}
	return elementImpl(el, attributes)
}

type TableEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TableEl) GetChilds() []Element   { return e.childs }
func (e *TableEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TableEl) GetElName() string      { return e.elName }
func (e *TableEl) GetElValue() js.Value   { return e.ElValue }

func Table(attributes ...Attribute) func(...Element) Element {
	el := &TableEl{elName: "table"}
	return elementImpl(el, attributes)
}

type TBodyEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TBodyEl) GetChilds() []Element   { return e.childs }
func (e *TBodyEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TBodyEl) GetElName() string      { return e.elName }
func (e *TBodyEl) GetElValue() js.Value   { return e.ElValue }

func TBody(attributes ...Attribute) func(...Element) Element {
	el := &TBodyEl{elName: "tbody"}
	return elementImpl(el, attributes)
}

type TdEl struct {
	BasicElement
	Colspan int64
	Rowspan int64
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TdEl) GetChilds() []Element   { return e.childs }
func (e *TdEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TdEl) GetElName() string      { return e.elName }
func (e *TdEl) GetElValue() js.Value   { return e.ElValue }

func Td(attributes ...Attribute) func(...Element) Element {
	el := &TdEl{elName: "td"}
	return elementImpl(el, attributes)
}

type TextareaEl struct {
	BasicElement
	Name     string
	Cols     int64
	Rows     int64
	Disabled bool
	Readonly bool
	Required bool
	Value    string
	elName   string
	ElValue  js.Value
}

func (e *TextareaEl) GetChilds() []Element   { return []Element{} }
func (e *TextareaEl) AppendChild(el Element) {}
func (e *TextareaEl) GetElName() string      { return e.elName }
func (e *TextareaEl) GetElValue() js.Value   { return e.ElValue }

func Textarea(attributes ...Attribute) func(...Element) Element {
	el := &TextareaEl{elName: "textarea"}
	return elementImpl(el, attributes)
} // todo

type TFootEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TFootEl) GetChilds() []Element   { return e.childs }
func (e *TFootEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TFootEl) GetElName() string      { return e.elName }
func (e *TFootEl) GetElValue() js.Value   { return e.ElValue }

func TFoot(attributes ...Attribute) func(...Element) Element {
	el := &TFootEl{elName: "tfoot"}
	return elementImpl(el, attributes)
}

type ThEl struct {
	BasicElement
	Colspan int64
	Rowspan int64
	Scope   string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *ThEl) GetChilds() []Element   { return e.childs }
func (e *ThEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ThEl) GetElName() string      { return e.elName }
func (e *ThEl) GetElValue() js.Value   { return e.ElValue }

func Th(attributes ...Attribute) func(...Element) Element {
	el := &ThEl{elName: "th"}
	return elementImpl(el, attributes)
}

type TheadEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TheadEl) GetChilds() []Element   { return e.childs }
func (e *TheadEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TheadEl) GetElName() string      { return e.elName }
func (e *TheadEl) GetElValue() js.Value   { return e.ElValue }

func Thead(attributes ...Attribute) func(...Element) Element {
	el := &TheadEl{elName: "thead"}
	return elementImpl(el, attributes)
}

type TimeEl struct {
	BasicElement
	DateTime string
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *TimeEl) GetChilds() []Element   { return e.childs }
func (e *TimeEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TimeEl) GetElName() string      { return e.elName }
func (e *TimeEl) GetElValue() js.Value   { return e.ElValue }

func Time(attributes ...Attribute) func(...Element) Element {
	el := &TimeEl{elName: "time"}
	return elementImpl(el, attributes)
}

type TrEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *TrEl) GetChilds() []Element   { return e.childs }
func (e *TrEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *TrEl) GetElName() string      { return e.elName }
func (e *TrEl) GetElValue() js.Value   { return e.ElValue }

func Tr(attributes ...Attribute) func(...Element) Element {
	el := &TrEl{elName: "tr"}
	return elementImpl(el, attributes)
}

type TrackEl struct {
	BasicElement
	Kind    string
	Src     string
	Srclang string
	Label   string
	Default bool
	elName  string
	ElValue js.Value
}

func (e *TrackEl) GetChilds() []Element   { return []Element{} }
func (e *TrackEl) AppendChild(el Element) {}
func (e *TrackEl) GetElName() string      { return e.elName }
func (e *TrackEl) GetElValue() js.Value   { return e.ElValue }

func Track(attributes ...Attribute) func(...Element) Element {
	el := &TrackEl{elName: "track"}
	return elementImpl(el, attributes)
} // todo

type UEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *UEl) GetChilds() []Element   { return e.childs }
func (e *UEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *UEl) GetElName() string      { return e.elName }
func (e *UEl) GetElValue() js.Value   { return e.ElValue }

func U(attributes ...Attribute) func(...Element) Element {
	el := &UEl{elName: "u"}
	return elementImpl(el, attributes)
}

type UlEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *UlEl) GetChilds() []Element   { return e.childs }
func (e *UlEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *UlEl) GetElName() string      { return e.elName }
func (e *UlEl) GetElValue() js.Value   { return e.ElValue }

func Ul(attributes ...Attribute) func(...Element) Element {
	el := &UlEl{elName: "ul"}
	return elementImpl(el, attributes)
}

type VarEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *VarEl) GetChilds() []Element   { return e.childs }
func (e *VarEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *VarEl) GetElName() string      { return e.elName }
func (e *VarEl) GetElValue() js.Value   { return e.ElValue }

func Var(attributes ...Attribute) func(...Element) Element {
	el := &VarEl{elName: "var"}
	return elementImpl(el, attributes)
}

type VideoEl struct {
	BasicElement
	Src      string
	Poster   string
	Width    int64
	Height   int64
	AutoPlay bool
	Controls bool
	Loop     bool
	Muted    bool
	Preload  string
	childs   []Element
	elName   string
	ElValue  js.Value
}

func (e *VideoEl) GetChilds() []Element   { return e.childs }
func (e *VideoEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *VideoEl) GetElName() string      { return e.elName }
func (e *VideoEl) GetElValue() js.Value   { return e.ElValue }

func Video(attributes ...Attribute) func(...Element) Element {
	el := &VideoEl{elName: "video"}
	return elementImpl(el, attributes)
}

type WbrEl struct {
	BasicElement
	elName  string
	ElValue js.Value
}

func (e *WbrEl) GetChilds() []Element   { return []Element{} }
func (e *WbrEl) AppendChild(el Element) {}
func (e *WbrEl) GetElName() string      { return e.elName }
func (e *WbrEl) GetElValue() js.Value   { return e.ElValue }

func Wbr(attributes ...Attribute) func(...Element) Element {
	el := &WbrEl{elName: "wbr"}
	return elementImpl(el, attributes)
} // todo

// ---------------- Div Element ----->

type DivEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *DivEl) GetChilds() []Element   { return e.childs }
func (e *DivEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *DivEl) GetElName() string      { return e.elName }
func (e *DivEl) GetElValue() js.Value   { return e.ElValue }

func Div(attributes ...Attribute) func(...Element) Element {
	el := &DivEl{elName: "div"}

	return elementImpl(el, attributes)
}

// ---------------- P Element ---->

type PEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *PEl) GetChilds() []Element   { return e.childs }
func (e *PEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *PEl) GetElName() string      { return e.elName }
func (e *PEl) GetElValue() js.Value   { return e.ElValue }

func P(attributes ...Attribute) func(...Element) Element {
	el := &PEl{elName: "p"}

	return elementImpl(el, attributes)
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
	ElValue        js.Value
}

func (e *ButtonEl) GetChilds() []Element   { return e.childs }
func (e *ButtonEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *ButtonEl) GetElName() string      { return e.elName }
func (e *ButtonEl) GetElValue() js.Value   { return e.ElValue }

func Button(attributes ...Attribute) func(...Element) Element {
	el := &ButtonEl{elName: "button"}

	return elementImpl(el, attributes)
}

// ---------------- Label Element ---->

type LabelEl struct {
	BasicElement
	For     string
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *LabelEl) GetChilds() []Element   { return e.childs }
func (e *LabelEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *LabelEl) GetElName() string      { return e.elName }
func (e *LabelEl) GetElValue() js.Value   { return e.ElValue }

func Label(attributes ...Attribute) func(...Element) Element {
	el := &LabelEl{elName: "label"}

	return elementImpl(el, attributes)
}

// ---------------- Input Element ---->

type InputEl struct {
	BasicElement
	Type        string
	Name        string
	Value       string
	Placeholder string
	Required    bool
	childs      []Element
	elName      string
	ElValue     js.Value
}

func (e *InputEl) GetChilds() []Element   { return e.childs }
func (e *InputEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *InputEl) GetElName() string      { return e.elName }
func (e *InputEl) GetElValue() js.Value   { return e.ElValue }

func Input(attributes ...Attribute) func(...Element) Element {
	el := &InputEl{elName: "input"}

	return elementImpl(el, attributes)
}

// ---------------- Header Element ---->

type HeaderEl struct {
	BasicElement
	childs  []Element
	elName  string
	ElValue js.Value
}

func (e *HeaderEl) GetChilds() []Element   { return e.childs }
func (e *HeaderEl) AppendChild(el Element) { e.childs = append(e.childs, el) }
func (e *HeaderEl) GetElName() string      { return e.elName }
func (e *HeaderEl) GetElValue() js.Value   { return e.ElValue }

func Header(attributes ...Attribute) func(...Element) Element {
	el := &HeaderEl{elName: "header"}

	return elementImpl(el, attributes)
}

package elements

import "fmt"

type Attribute struct {
	Name  string
	Value interface{} // can be a function also
}

// every boole arguments with default false
var (
	Async = Attribute{
		Name:  "Async",
		Value: true,
	}

	AutoFocus = Attribute{
		Name:  "Autofocus",
		Value: true,
	}

	AutoPlay = Attribute{
		Name:  "Autoplay",
		Value: true,
	}

	Checked = Attribute{
		Name:  "Checked",
		Value: true,
	}

	Controls = Attribute{
		Name:  "Controls",
		Value: true,
	}

	Defer = Attribute{
		Name:  "Defer",
		Value: true,
	}

	Disabled = Attribute{
		Name:  "Disabled",
		Value: true,
	}

	Loop = Attribute{
		Name:  "Loop",
		Value: true,
	}

	Multiple = Attribute{
		Name:  "Multiple",
		Value: true,
	}

	Muted = Attribute{
		Name:  "Muted",
		Value: true,
	}

	Open = Attribute{
		Name:  "Open",
		Value: true,
	}

	PlaysInline = Attribute{
		Name:  "PlaysInline",
		Value: true,
	}

	ReadOnly = Attribute{
		Name:  "ReadOnly",
		Value: true,
	}

	Required = Attribute{
		Name:  "Required",
		Value: true,
	}

	Selected = Attribute{
		Name:  "Selected",
		Value: true,
	}
)

// specifies the type of file the server accepts
//
// example:
//
//	<input type="file" name="poster" accept="image/png, image/jpeg" />
func Accept(v ...any) Attribute {
	return Attribute{
		Name:  "Accept",
		Value: fmt.Sprint(v...),
	}
}

// action used only in Form element
//
// example:
//
//	<form action="/action_page.php">
func Action(value string) Attribute {
	return Attribute{
		Name:  "Action",
		Value: value,
	}
}

// alternative text string to display on browsers that do not display images.
//
// example:
//
//	<img src="https://via.placeholder.com/350x150" alt="Text that represents the image">
//
// NOTE: the alt is required if the href attribute is used.
func Alt(v ...any) Attribute {
	return Attribute{
		Name:  "Alt",
		Value: fmt.Sprint(v...),
	}
}

// Aria name is the aria-<name>
//
// example:
//
//	<div role="progressbar" aria-valuemin="0" aria-valuemax="100" />
func Aria(name string, value string) Attribute {
	return Attribute{
		Name:  "Aria",
		Value: map[string]interface{}{name: value},
	}
}

// specifies the type of content for the link
//
// example:
//
//	<link rel="preload" href="style.css" as="style">
//
// NOTE: can only be used if the content is being preloaded
func As(value string) Attribute {
	return Attribute{
		Name:  "As",
		Value: value,
	}
}

// specifies if form elements should be automatically completed by the browser
//
// example:
//
//	<input name="email" id="email" type="email" autocomplete="off" />
//
// NOTE: should only be on/off
func AutoComplete(value string) Attribute {
	return Attribute{
		Name:  "AutoComplete",
		Value: value,
	}
}

// specifies the character encoding used for the document
//
// example:
//
//	<meta charset="utf-8">
func Charset(value string) Attribute {
	return Attribute{
		Name:  "Charset",
		Value: value,
	}
}

// link css style to the element
//
// example:
//
//	<p class="note editorial">JS is sooo slow</p>
func Class(v ...any) Attribute {
	return Attribute{
		Name:  "Class",
		Value: fmt.Sprint(v...),
	}
}

// specifies the width of the textarea
//
// example:
//
//	<textarea rows="4" cols="50">
//
// NOTE: default value is 20 and this work only on textarea elements
func Cols(value int64) Attribute {
	return Attribute{
		Name:  "Cols",
		Value: value,
	}
}

// specifies the amount of columns the cell extends
//
// example:
//
//	<td colspan="2">Golang is better</td>
//
// NOTE: work only on td elements
func ColSpan(value int64) Attribute {
	return Attribute{
		Name:  "ColSpan",
		Value: value,
	}
}

// specifies the content for the element
//
// example:
//
//	<meta name="description" content="react ?... nahhh">
func Content(v ...any) Attribute {
	return Attribute{
		Name:  "Content",
		Value: fmt.Sprint(v...),
	}
}

// data name is the data-<name>.
// form a class of attributes called custom data attributes,
// that allow proprietary information to be exchanged between
// the HTML and its DOM representation by scripts
//
// example:
//
//	<li data-id="97865"></li>
//
// NOTE: work only on td elements
func Data(name, value string) Attribute {
	return Attribute{
		Name:  "Data",
		Value: map[string]interface{}{name: value},
	}
}

// specifies the ID of the element that the label element is related to
//
// example:
//
//	<label for="golang">Golang is simple</label>
func For(value string) Attribute {
	return Attribute{
		Name:  "For",
		Value: value,
	}
}

// specifies the form(s) that the element belongs to
//
// example:
//
//	<button type="submit" form="form1" value="Submit">Submit</button>
//
// NOTE: this will work on: button, dieldset, input, label, meter,
// object, output, select, textarea
func FormAttr(value string) Attribute {
	return Attribute{
		Name:  "Form",
		Value: value,
	}
}

// specifies the height of the element, default is measured in pixels
//
// example:
//
//	<img src="https://via.placeholder.com/350x150" height="42" width="42">
//
// NOTE: this will work on: canvas, embed, iframe, img, input, object, video
func Height(value int64) Attribute {
	return Attribute{
		Name:  "Height",
		Value: value,
	}
}

// determines the URL to where the link will point to or the name of the anchor
//
// example:
//
//	<a href="https://go.dev/">Golang is Fast</a>
func Href(value string) Attribute {
	return Attribute{
		Name:  "Href",
		Value: value,
	}
}

// defines an identifier (ID) which must be unique in the whole document
//
// example:
//
//	<p id="golang">Golang is blazinglyfast</p>
//
// NOTE: will apply on every element
func ID(value string) Attribute {
	return Attribute{
		Name:  "ID",
		Value: value,
	}
}

// define the language of an element
//
// example:
//
//	<p lang="fr">Golang est incroyable</p>
//
// NOTE: will apply on every element
func Lang(value string) Attribute {
	return Attribute{
		Name:  "Lang",
		Value: value,
	}
}

// specifies whether a browser should load an image immediately or
// to defer loading of off-screen images until
// for example the user scrolls near them
//
// example:
//
//	<img src="https://via.placeholder.com/350x150" loading="lazy">
//
// NOTE: will apply only on img elements
func Loading(v ...any) Attribute {
	return Attribute{
		Name:  "Loading",
		Value: fmt.Sprint(v...),
	}
}

// specifies the maximum value for this element
//
// example:
//
//	<input type="number" name="quantity" min="1" max="5">
//
// NOTE: will apply only on: input, meter, progress
func Max(value int64) Attribute {
	return Attribute{
		Name:  "Max",
		Value: value,
	}
}

// specifies the maximum number of characters a user can enter
//
// example:
//
//	<input type="text" name="usrname" maxlength="10">
//
// NOTE: will apply only on: input, textarea
func MaxLength(value int64) Attribute {
	return Attribute{
		Name:  "MaxLength",
		Value: value,
	}
}

// specifies the HTTP method used when submitting the form
//
// example:
//
//	<input type="text" name="usrname" maxlength="10">
//
// NOTE: will apply only on form elements
func Method(value string) Attribute {
	return Attribute{
		Name:  "Method",
		Value: value,
	}
}

// specifies the minimum value for this element
//
// example:
//
//	<input type="number" name="quantity" min="1" max="5">
//
// NOTE: will apply only on: input, meter, progress
func Min(value int64) Attribute {
	return Attribute{
		Name:  "Min",
		Value: value,
	}
}

// specifies the minimum number of characters a user can enter
//
// example:
//
//	<input type="text" name="usrname" maxlength="10">
//
// NOTE: will apply only on: input, textarea
func MinLength(value int64) Attribute {
	return Attribute{
		Name:  "MinLength",
		Value: value,
	}
}

// specifies the name of the element
//
// example:
//
//	<button name="subject">HTML</button>
//
// NOTE: will apply only on: button, fieldset, form, iframe, input,
// map, meta, object, output, param, select, textarea
func Name(value string) Attribute {
	return Attribute{
		Name:  "Name",
		Value: value,
	}
}

// specifies the regular expression that the element is checked against
//
// example:
//
//	<input type="text" pattern="[^@\s]+@[^@\s]+">
//
// NOTE: will apply only on input elements
func Pattern(value string) Attribute {
	return Attribute{
		Name:  "Pattern",
		Value: value,
	}
}

// specifies a hint that describes the expected value
//
// example:
//
//	<input type="text" name="username" placeholder="Your Name">
func Placeholder(v ...any) Attribute {
	return Attribute{
		Name:  "Placeholder",
		Value: fmt.Sprint(v...),
	}
}

// specifies the image to display before the user plays the video
//
// example:
//
//	<video controls src="movie.mp4" poster="/images/w3html5.gif" />
//
// NOTE: will apply only on video elements
func Poster(value string) Attribute {
	return Attribute{
		Name:  "Poster",
		Value: value,
	}
}

// specifies how the author will think how the media should be loaded
//
// example:
//
//	<audio controls preload="none">
//
// NOTE: will apply only on: video, audio
func Preload(value string) Attribute {
	return Attribute{
		Name:  "Preload",
		Value: value,
	}
}

// specifies the relationship between the current and target documents
//
// example:
//
//	<a rel="nofollow" href="https://go.dev/">Golang is light</a>
//
// NOTE: will apply only on: a, area, link, form
func Rel(value string) Attribute {
	return Attribute{
		Name:  "Rel",
		Value: value,
	}
}

// defines a role for an element. This role provides information about the type of widget being created
//
// example:
//
//	<div role="navigation">...</div>
//
// NOTE: Applies to every element
func Role(value string) Attribute {
	return Attribute{
		Name:  "Role",
		Value: value,
	}
}

// specifies the number of rows in a text area
//
// example:
//
//	<textarea rows="4">...</textarea>
//
// NOTE: Applies only to textarea elements
func Rows(value int64) Attribute {
	return Attribute{
		Name:  "Rows",
		Value: value,
	}
}

// specifies the number of rows a cell should span
//
// example:
//
//	<td rowspan="2">...</td>
//
// NOTE: Applies to td and th elements
func RowSpan(value int64) Attribute {
	return Attribute{
		Name:  "RowSpan",
		Value: value,
	}
}

// epecifies the source URL for an image
//
// example:
//
//	<img src="image.jpg">
//
// NOTE: Applies only to img elements
func Src(value string) Attribute {
	return Attribute{
		Name:  "Src",
		Value: value,
	}
}

// specifies one or more sources for an image
//
// example:
//
//	<img srcset="image-320w.jpg  320w, image-480w.jpg  480w">
//
// NOTE: Applies only to img elements
func SrcSet(value string) Attribute {
	return Attribute{
		Name:  "SrcSet",
		Value: value,
	}
}

// specifies the legal number intervals for an input field
//
// example:
//
//	<input type="number" step="2">
//
// NOTE: Applies only to input elements with type="number"
func Step(value float64) Attribute {
	return Attribute{
		Name:  "Step",
		Value: value,
	}
}

// specifies inline CSS styles for an element
//
// example:
//
//	<div style="color: blue;">...</div>
//
// NOTE: Applies to every element
func Style(v ...any) Attribute {
	return Attribute{
		Name:  "Style",
		Value: fmt.Sprint(v...),
	}
}

// specifies the tab order of an element
//
// example:
//
//	<input type="text" tabindex="1">
//
// NOTE: applies to every element
func TabIndex(value int64) Attribute {
	return Attribute{
		Name:  "TabIndex",
		Value: value,
	}
}

// specifies where to open the linked document
//
// example:
//
//	<a href="https://example.com" target="_blank">...</a>
//
// NOTE: Applies to a and area elements
func Target(value string) Attribute {
	return Attribute{
		Name:  "Target",
		Value: value,
	}
}

// provides advisory information about the element
//
// example:
//
//	<abbr title="Hypertext Markup Language">HTML</abbr>
//
// NOTE: Applies to every element
func Title(v ...any) Attribute {
	return Attribute{
		Name:  "Title",
		Value: fmt.Sprint(v...),
	}
}

// specifies the type of input element
//
// example:
//
//	<input type="text">
//
// NOTE: Applies to input elements
func Type(value string) Attribute {
	return Attribute{
		Name:  "Type",
		Value: value,
	}
}

// specifies the initial value for an input field
//
// example:
//
//	<input type="text" value="Default">
//
// NOTE: Applies to input, textarea, and select elements
func Value(v ...any) Attribute {
	return Attribute{
		Name:  "Value",
		Value: fmt.Sprint(v...),
	}
}

// specifies the width of an element.
//
// example:
//
//	<div style="width:  200px;">...</div>
//
// NOTE: Applies to every element
func Width(value int64) Attribute {
	return Attribute{
		Name:  "Width",
		Value: value,
	}
}

// specifies how the form-data should be encoded when submitting it to the server
//
// example:
//
//	<form enctype="multipart/form-data">...</form>
//
// NOTE: Applies to form elements
func EncType(value string) Attribute {
	return Attribute{
		Name:  "EncType",
		Value: value,
	}
}

// ---------------- Event Handlers ----->

type EventHandler func()

func OnClick(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnClick",
		Value: handler,
	}
}

func OnDblClick(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnDblClick",
		Value: handler,
	}
}

func OnMouseDown(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnMouseDown",
		Value: handler,
	}
}

func OnMouseUp(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnMouseUp",
		Value: handler,
	}
}

func OnMouseMove(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnMouseMove",
		Value: handler,
	}
}

func OnMouseOut(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnMouseOut",
		Value: handler,
	}
}

func OnMouseOver(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnMouseOut",
		Value: handler,
	}
}

func OnKeyDown(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnKeyDown",
		Value: handler,
	}
}

func OnKeyUp(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnKeyUp",
		Value: handler,
	}
}

func OnFocus(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnFocus",
		Value: handler,
	}
}

func OnBlur(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnBlur",
		Value: handler,
	}
}

func OnChange(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnChange",
		Value: handler,
	}
}

func OnSubmit(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnSubmit",
		Value: handler,
	}
}

func OnReset(handler EventHandler) Attribute {
	return Attribute{
		Name:  "OnReset",
		Value: handler,
	}
}

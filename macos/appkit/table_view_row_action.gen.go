// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableViewRowAction] class.
var TableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}

type _TableViewRowActionClass struct {
	objc.Class
}

// An interface definition for the [TableViewRowAction] class.
type ITableViewRowAction interface {
	objc.IObject
	Style() TableViewRowActionStyle
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	Title() string
	SetTitle(value string)
	Image() Image
	SetImage(value IImage)
}

// A single action to present when the user swipes horizontally on a table row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction?language=objc
type TableViewRowAction struct {
	objc.Object
}

func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TableViewRowActionClass) RowActionWithStyleTitleHandler(style TableViewRowActionStyle, title string, handler func(action TableViewRowAction, row int)) TableViewRowAction {
	rv := objc.Call[TableViewRowAction](tc, objc.Sel("rowActionWithStyle:title:handler:"), style, title, handler)
	return rv
}

// Creates and returns a new table view row action object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401994-rowactionwithstyle?language=objc
func TableViewRowAction_RowActionWithStyleTitleHandler(style TableViewRowActionStyle, title string, handler func(action TableViewRowAction, row int)) TableViewRowAction {
	return TableViewRowActionClass.RowActionWithStyleTitleHandler(style, title, handler)
}

func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := objc.Call[TableViewRowAction](tc, objc.Sel("alloc"))
	return rv
}

func TableViewRowAction_Alloc() TableViewRowAction {
	return TableViewRowActionClass.Alloc()
}

func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := objc.Call[TableViewRowAction](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableViewRowAction() TableViewRowAction {
	return TableViewRowActionClass.New()
}

func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := objc.Call[TableViewRowAction](t_, objc.Sel("init"))
	return rv
}

// The style applied to the action button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401982-style?language=objc
func (t_ TableViewRowAction) Style() TableViewRowActionStyle {
	rv := objc.Call[TableViewRowActionStyle](t_, objc.Sel("style"))
	return rv
}

// The background color of the action button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401984-backgroundcolor?language=objc
func (t_ TableViewRowAction) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the action button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401984-backgroundcolor?language=objc
func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The title of the action button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401992-title?language=objc
func (t_ TableViewRowAction) Title() string {
	rv := objc.Call[string](t_, objc.Sel("title"))
	return rv
}

// The title of the action button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/1401992-title?language=objc
func (t_ TableViewRowAction) SetTitle(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setTitle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/2177311-image?language=objc
func (t_ TableViewRowAction) Image() Image {
	rv := objc.Call[Image](t_, objc.Sel("image"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewrowaction/2177311-image?language=objc
func (t_ TableViewRowAction) SetImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setImage:"), objc.Ptr(value))
}
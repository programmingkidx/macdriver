// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextLayoutFragment] class.
var TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}

type _TextLayoutFragmentClass struct {
	objc.Class
}

// An interface definition for the [TextLayoutFragment] class.
type ITextLayoutFragment interface {
	objc.IObject
	InvalidateLayout()
	FrameForTextAttachmentAtLocation(location PTextLocation) coregraphics.Rect
	FrameForTextAttachmentAtLocationObject(locationObject objc.IObject) coregraphics.Rect
	DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef)
	TextLayoutManager() TextLayoutManager
	BottomMargin() float64
	TextLineFragments() []TextLineFragment
	State() TextLayoutFragmentState
	RangeInElement() TextRange
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	TopMargin() float64
	LeadingPadding() float64
	LayoutFragmentFrame() coregraphics.Rect
	RenderingSurfaceBounds() coregraphics.Rect
	TextAttachmentViewProviders() []TextAttachmentViewProvider
	TrailingPadding() float64
	TextElement() TextElement
}

// A class that represents the layout fragment typically corresponding to a rendering surface such as layer or view subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment?language=objc
type TextLayoutFragment struct {
	objc.Object
}

func TextLayoutFragmentFrom(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextLayoutFragment) InitWithTextElementRange(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](t_, objc.Sel("initWithTextElement:range:"), objc.Ptr(textElement), objc.Ptr(rangeInElement))
	return rv
}

// Create a new layout fragment using the provided text element and range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809966-initwithtextelement?language=objc
func NewTextLayoutFragmentWithTextElementRange(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	instance := TextLayoutFragmentClass.Alloc().InitWithTextElementRange(textElement, rangeInElement)
	instance.Autorelease()
	return instance
}

func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](tc, objc.Sel("alloc"))
	return rv
}

func TextLayoutFragment_Alloc() TextLayoutFragment {
	return TextLayoutFragmentClass.Alloc()
}

func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutFragment() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](t_, objc.Sel("init"))
	return rv
}

// Invalidates any layout information associated with the text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809967-invalidatelayout?language=objc
func (t_ TextLayoutFragment) InvalidateLayout() {
	objc.Call[objc.Void](t_, objc.Sel("invalidateLayout"))
}

// Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809964-framefortextattachmentatlocation?language=objc
func (t_ TextLayoutFragment) FrameForTextAttachmentAtLocation(location PTextLocation) coregraphics.Rect {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("frameForTextAttachmentAtLocation:"), po0)
	return rv
}

// Returns the frame in the text layout fragment coordinate system for the attachment at the location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809964-framefortextattachmentatlocation?language=objc
func (t_ TextLayoutFragment) FrameForTextAttachmentAtLocationObject(locationObject objc.IObject) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("frameForTextAttachmentAtLocation:"), objc.Ptr(locationObject))
	return rv
}

// Renders the visual representation of this element in the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3824748-drawatpoint?language=objc
func (t_ TextLayoutFragment) DrawAtPointInContext(point coregraphics.Point, context coregraphics.ContextRef) {
	objc.Call[objc.Void](t_, objc.Sel("drawAtPoint:inContext:"), point, context)
}

// Returns the layout manager for this text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809975-textlayoutmanager?language=objc
func (t_ TextLayoutFragment) TextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("textLayoutManager"))
	return rv
}

// Returns the amount of space reserved during paragraph layout between the bottom of the last line in the paragraph and the bottom of the text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3852574-bottommargin?language=objc
func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := objc.Call[float64](t_, objc.Sel("bottomMargin"))
	return rv
}

// Returns an array of text line fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809976-textlinefragments?language=objc
func (t_ TextLayoutFragment) TextLineFragments() []TextLineFragment {
	rv := objc.Call[[]TextLineFragment](t_, objc.Sel("textLineFragments"))
	return rv
}

// The layout information state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809972-state?language=objc
func (t_ TextLayoutFragment) State() TextLayoutFragmentState {
	rv := objc.Call[TextLayoutFragmentState](t_, objc.Sel("state"))
	return rv
}

// Returns the range inside the text element relative to the document origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809970-rangeinelement?language=objc
func (t_ TextLayoutFragment) RangeInElement() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("rangeInElement"))
	return rv
}

// The queue on which the framework dispatches layout operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809969-layoutqueue?language=objc
func (t_ TextLayoutFragment) LayoutQueue() foundation.OperationQueue {
	rv := objc.Call[foundation.OperationQueue](t_, objc.Sel("layoutQueue"))
	return rv
}

// The queue on which the framework dispatches layout operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809969-layoutqueue?language=objc
func (t_ TextLayoutFragment) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutQueue:"), objc.Ptr(value))
}

// Returns the amount of space reserved during paragraph layout between the top of the text layout fragment and the top of the first line in the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3852576-topmargin?language=objc
func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := objc.Call[float64](t_, objc.Sel("topMargin"))
	return rv
}

// Returns the amount of margin space reserved during paragraph layout between the leading edge of the text layout fragment and the start of the lines in the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3852575-leadingpadding?language=objc
func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := objc.Call[float64](t_, objc.Sel("leadingPadding"))
	return rv
}

// Returns the rectangle the framework uses for tiling the layout fragment inside the target layout coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809968-layoutfragmentframe?language=objc
func (t_ TextLayoutFragment) LayoutFragmentFrame() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("layoutFragmentFrame"))
	return rv
}

// Returns the bounds defining the area required for rendering the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809971-renderingsurfacebounds?language=objc
func (t_ TextLayoutFragment) RenderingSurfaceBounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("renderingSurfaceBounds"))
	return rv
}

// Returns the attachment view provider associated with the text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809973-textattachmentviewproviders?language=objc
func (t_ TextLayoutFragment) TextAttachmentViewProviders() []TextAttachmentViewProvider {
	rv := objc.Call[[]TextAttachmentViewProvider](t_, objc.Sel("textAttachmentViewProviders"))
	return rv
}

// Returns the amount of margin space reserved during paragraph layout between the end of the lines in the paragraph and the trailing edge of the text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3852577-trailingpadding?language=objc
func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := objc.Call[float64](t_, objc.Sel("trailingPadding"))
	return rv
}

// Returns the parent text element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutfragment/3809974-textelement?language=objc
func (t_ TextLayoutFragment) TextElement() TextElement {
	rv := objc.Call[TextElement](t_, objc.Sel("textElement"))
	return rv
}
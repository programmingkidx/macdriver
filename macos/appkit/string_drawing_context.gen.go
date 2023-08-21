// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StringDrawingContext] class.
var StringDrawingContextClass = _StringDrawingContextClass{objc.GetClass("NSStringDrawingContext")}

type _StringDrawingContextClass struct {
	objc.Class
}

// An interface definition for the [StringDrawingContext] class.
type IStringDrawingContext interface {
	objc.IObject
	TotalBounds() coregraphics.Rect
	ActualScaleFactor() float64
	MinimumScaleFactor() float64
	SetMinimumScaleFactor(value float64)
}

// An object that manages metrics for drawing attributed strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingcontext?language=objc
type StringDrawingContext struct {
	objc.Object
}

func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StringDrawingContextClass) Alloc() StringDrawingContext {
	rv := objc.Call[StringDrawingContext](sc, objc.Sel("alloc"))
	return rv
}

func StringDrawingContext_Alloc() StringDrawingContext {
	return StringDrawingContextClass.Alloc()
}

func (sc _StringDrawingContextClass) New() StringDrawingContext {
	rv := objc.Call[StringDrawingContext](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStringDrawingContext() StringDrawingContext {
	return StringDrawingContextClass.New()
}

func (s_ StringDrawingContext) Init() StringDrawingContext {
	rv := objc.Call[StringDrawingContext](s_, objc.Sel("init"))
	return rv
}

// The most recent bounding rectangle that the system used to draw the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingcontext/1530525-totalbounds?language=objc
func (s_ StringDrawingContext) TotalBounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](s_, objc.Sel("totalBounds"))
	return rv
}

// The actual scale factor that the system applied to the font during drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingcontext/1531498-actualscalefactor?language=objc
func (s_ StringDrawingContext) ActualScaleFactor() float64 {
	rv := objc.Call[float64](s_, objc.Sel("actualScaleFactor"))
	return rv
}

// The scale factor that determines the smallest font size to use during drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingcontext/1534020-minimumscalefactor?language=objc
func (s_ StringDrawingContext) MinimumScaleFactor() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minimumScaleFactor"))
	return rv
}

// The scale factor that determines the smallest font size to use during drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsstringdrawingcontext/1534020-minimumscalefactor?language=objc
func (s_ StringDrawingContext) SetMinimumScaleFactor(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumScaleFactor:"), value)
}
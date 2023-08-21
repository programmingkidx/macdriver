// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ATSTypesetter] class.
var ATSTypesetterClass = _ATSTypesetterClass{objc.GetClass("NSATSTypesetter")}

type _ATSTypesetterClass struct {
	objc.Class
}

// An interface definition for the [ATSTypesetter] class.
type IATSTypesetter interface {
	ITypesetter
	GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, paragraphSeparatorGlyphRange foundation.Range, lineOrigin foundation.Point)
}

// A concrete typesetter object that places glyphs during the text layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter?language=objc
type ATSTypesetter struct {
	Typesetter
}

func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		Typesetter: TypesetterFrom(ptr),
	}
}

func (ac _ATSTypesetterClass) Alloc() ATSTypesetter {
	rv := objc.Call[ATSTypesetter](ac, objc.Sel("alloc"))
	return rv
}

func ATSTypesetter_Alloc() ATSTypesetter {
	return ATSTypesetterClass.Alloc()
}

func (ac _ATSTypesetterClass) New() ATSTypesetter {
	rv := objc.Call[ATSTypesetter](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewATSTypesetter() ATSTypesetter {
	return ATSTypesetterClass.New()
}

func (a_ ATSTypesetter) Init() ATSTypesetter {
	rv := objc.Call[ATSTypesetter](a_, objc.Sel("init"))
	return rv
}

// Calculates the line fragment rectangle and the portion of the rectangle that contains marks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/1528343-getlinefragmentrect?language=objc
func (a_ ATSTypesetter) GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, paragraphSeparatorGlyphRange foundation.Range, lineOrigin foundation.Point) {
	objc.Call[objc.Void](a_, objc.Sel("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}

// Returns a shared instance of the typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/1530993-sharedtypesetter?language=objc
func (ac _ATSTypesetterClass) SharedTypesetter() ATSTypesetter {
	rv := objc.Call[ATSTypesetter](ac, objc.Sel("sharedTypesetter"))
	return rv
}

// Returns a shared instance of the typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/1530993-sharedtypesetter?language=objc
func ATSTypesetter_SharedTypesetter() ATSTypesetter {
	return ATSTypesetterClass.SharedTypesetter()
}
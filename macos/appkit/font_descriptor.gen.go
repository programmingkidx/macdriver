// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FontDescriptor] class.
var FontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}

type _FontDescriptorClass struct {
	objc.Class
}

// An interface definition for the [FontDescriptor] class.
type IFontDescriptor interface {
	objc.IObject
	ObjectForKey(attribute FontDescriptorAttributeName) objc.Object
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor
	FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor
	FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor
	FontDescriptorWithFace(newFace string) FontDescriptor
	FontDescriptorWithSize(newPointSize float64) FontDescriptor
	FontDescriptorWithFamily(newFamily string) FontDescriptor
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor
	FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor
	FontAttributes() map[FontDescriptorAttributeName]objc.Object
	RequiresFontAssetRequest() bool
	PointSize() float64
	Matrix() foundation.AffineTransform
	SymbolicTraits() FontDescriptorSymbolicTraits
	PostscriptName() string
}

// A dictionary of attributes that describe a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor?language=objc
type FontDescriptor struct {
	objc.Object
}

func FontDescriptorFrom(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FontDescriptor) InitWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("initWithFontAttributes:"), attributes)
	return rv
}

// Initializes and returns a new font descriptor with the specified attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469991-initwithfontattributes?language=objc
func NewFontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	instance := FontDescriptorClass.Alloc().InitWithFontAttributes(attributes)
	instance.Autorelease()
	return instance
}

func (f_ FontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithDesign:"), design)
	return rv
}

// Returns a new font descriptor based on the current object, but with the specified design style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/3152380-fontdescriptorwithdesign?language=objc
func FontDescriptor_FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	instance := FontDescriptorClass.Alloc().FontDescriptorWithDesign(design)
	instance.Autorelease()
	return instance
}

func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := objc.Call[FontDescriptor](fc, objc.Sel("alloc"))
	return rv
}

func FontDescriptor_Alloc() FontDescriptor {
	return FontDescriptorClass.Alloc()
}

func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := objc.Call[FontDescriptor](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

func (f_ FontDescriptor) Init() FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("init"))
	return rv
}

// Returns the font attribute specified by the given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469837-objectforkey?language=objc
func (f_ FontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("objectForKey:"), attribute)
	return rv
}

// Returns a font descriptor with the name and size attributes set to the given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469912-fontdescriptorwithname?language=objc
func (fc _FontDescriptorClass) FontDescriptorWithNameSize(fontName string, size float64) FontDescriptor {
	rv := objc.Call[FontDescriptor](fc, objc.Sel("fontDescriptorWithName:size:"), fontName, size)
	return rv
}

// Returns a font descriptor with the name and size attributes set to the given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469912-fontdescriptorwithname?language=objc
func FontDescriptor_FontDescriptorWithNameSize(fontName string, size float64) FontDescriptor {
	return FontDescriptorClass.FontDescriptorWithNameSize(fontName, size)
}

// Returns a normalized font descriptor whose specified attributes match those of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469839-matchingfontdescriptorwithmandat?language=objc
func (f_ FontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("matchingFontDescriptorWithMandatoryKeys:"), objc.Ptr(mandatoryKeys))
	return rv
}

// Returns a font descriptor with a dictionary of attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469856-fontdescriptorwithfontattributes?language=objc
func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.Call[FontDescriptor](fc, objc.Sel("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}

// Returns a font descriptor with a dictionary of attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469856-fontdescriptorwithfontattributes?language=objc
func FontDescriptor_FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	return FontDescriptorClass.FontDescriptorWithFontAttributes(attributes)
}

// Returns a new font descriptor based on the current object, but with the specified symbolic traits taking precedence over the existing ones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469843-fontdescriptorwithsymbolictraits?language=objc
func (f_ FontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithSymbolicTraits:"), symbolicTraits)
	return rv
}

// Returns a new font descriptor based on the current object, but with the specified attributes taking precedence over the existing ones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469987-fontdescriptorbyaddingattributes?language=objc
func (f_ FontDescriptor) FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorByAddingAttributes:"), attributes)
	return rv
}

// Returns a font descriptor that contains the text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/3553196-preferredfontdescriptorfortextst?language=objc
func (fc _FontDescriptorClass) PreferredFontDescriptorForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	rv := objc.Call[FontDescriptor](fc, objc.Sel("preferredFontDescriptorForTextStyle:options:"), style, options)
	return rv
}

// Returns a font descriptor that contains the text style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/3553196-preferredfontdescriptorfortextst?language=objc
func FontDescriptor_PreferredFontDescriptorForTextStyleOptions(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	return FontDescriptorClass.PreferredFontDescriptorForTextStyleOptions(style, options)
}

// Returns a new font descriptor based on the current object, but with the specified face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469928-fontdescriptorwithface?language=objc
func (f_ FontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithFace:"), newFace)
	return rv
}

// Returns a new font descriptor based on the current object, but with the specified point size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469835-fontdescriptorwithsize?language=objc
func (f_ FontDescriptor) FontDescriptorWithSize(newPointSize float64) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithSize:"), newPointSize)
	return rv
}

// Returns a new font descriptor based on the current object, but with the specified font family. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469866-fontdescriptorwithfamily?language=objc
func (f_ FontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithFamily:"), newFamily)
	return rv
}

// Returns all the fonts available on the system whose specified attributes match those of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469985-matchingfontdescriptorswithmanda?language=objc
func (f_ FontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("matchingFontDescriptorsWithMandatoryKeys:"), objc.Ptr(mandatoryKeys))
	return rv
}

// Returns a new font descriptor based on the current object, but with the specified font matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469983-fontdescriptorwithmatrix?language=objc
func (f_ FontDescriptor) FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor {
	rv := objc.Call[FontDescriptor](f_, objc.Sel("fontDescriptorWithMatrix:"), objc.Ptr(matrix))
	return rv
}

// The receiver’s dictionary of attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469831-fontattributes?language=objc
func (f_ FontDescriptor) FontAttributes() map[FontDescriptorAttributeName]objc.Object {
	rv := objc.Call[map[FontDescriptorAttributeName]objc.Object](f_, objc.Sel("fontAttributes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/2890793-requiresfontassetrequest?language=objc
func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.Call[bool](f_, objc.Sel("requiresFontAssetRequest"))
	return rv
}

// The point size of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469829-pointsize?language=objc
func (f_ FontDescriptor) PointSize() float64 {
	rv := objc.Call[float64](f_, objc.Sel("pointSize"))
	return rv
}

// The current transform matrix of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469950-matrix?language=objc
func (f_ FontDescriptor) Matrix() foundation.AffineTransform {
	rv := objc.Call[foundation.AffineTransform](f_, objc.Sel("matrix"))
	return rv
}

// A bit mask that describes the traits of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469858-symbolictraits?language=objc
func (f_ FontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	rv := objc.Call[FontDescriptorSymbolicTraits](f_, objc.Sel("symbolicTraits"))
	return rv
}

// The PostScript name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/1469948-postscriptname?language=objc
func (f_ FontDescriptor) PostscriptName() string {
	rv := objc.Call[string](f_, objc.Sel("postscriptName"))
	return rv
}
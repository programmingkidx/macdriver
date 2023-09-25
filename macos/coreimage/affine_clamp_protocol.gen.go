// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a affine clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp?language=objc
type PAffineClamp interface {
	// optional
	SetTransform(value coregraphics.AffineTransform)
	HasSetTransform() bool

	// optional
	Transform() coregraphics.AffineTransform
	HasTransform() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PAffineClamp = (*AffineClampObject)(nil)

// A concrete type for the [PAffineClamp] protocol.
type AffineClampObject struct {
	objc.Object
}

func (a_ AffineClampObject) HasSetTransform() bool {
	return a_.RespondsToSelector(objc.Sel("setTransform:"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228055-transform?language=objc
func (a_ AffineClampObject) SetTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("setTransform:"), value)
}

func (a_ AffineClampObject) HasTransform() bool {
	return a_.RespondsToSelector(objc.Sel("transform"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228055-transform?language=objc
func (a_ AffineClampObject) Transform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("transform"))
	return rv
}

func (a_ AffineClampObject) HasSetInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228054-inputimage?language=objc
func (a_ AffineClampObject) SetInputImage(value Image) {
	objc.Call[objc.Void](a_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (a_ AffineClampObject) HasInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228054-inputimage?language=objc
func (a_ AffineClampObject) InputImage() Image {
	rv := objc.Call[Image](a_, objc.Sel("inputImage"))
	return rv
}
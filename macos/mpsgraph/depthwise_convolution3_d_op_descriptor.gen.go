// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DepthwiseConvolution3DOpDescriptor] class.
var DepthwiseConvolution3DOpDescriptorClass = _DepthwiseConvolution3DOpDescriptorClass{objc.GetClass("MPSGraphDepthwiseConvolution3DOpDescriptor")}

type _DepthwiseConvolution3DOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [DepthwiseConvolution3DOpDescriptor] class.
type IDepthwiseConvolution3DOpDescriptor interface {
	objc.IObject
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	ChannelDimensionIndex() int
	SetChannelDimensionIndex(value int)
	PaddingValues() []foundation.Number
	SetPaddingValues(value []foundation.INumber)
	DilationRates() []foundation.Number
	SetDilationRates(value []foundation.INumber)
	Strides() []foundation.Number
	SetStrides(value []foundation.INumber)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor?language=objc
type DepthwiseConvolution3DOpDescriptor struct {
	objc.Object
}

func DepthwiseConvolution3DOpDescriptorFrom(ptr unsafe.Pointer) DepthwiseConvolution3DOpDescriptor {
	return DepthwiseConvolution3DOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DepthwiseConvolution3DOpDescriptorClass) DescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides []foundation.INumber, dilationRates []foundation.INumber, paddingValues []foundation.INumber, paddingStyle PaddingStyle) DepthwiseConvolution3DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution3DOpDescriptor](dc, objc.Sel("descriptorWithStrides:dilationRates:paddingValues:paddingStyle:"), strides, dilationRates, paddingValues, paddingStyle)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750684-descriptorwithstrides?language=objc
func DepthwiseConvolution3DOpDescriptor_DescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides []foundation.INumber, dilationRates []foundation.INumber, paddingValues []foundation.INumber, paddingStyle PaddingStyle) DepthwiseConvolution3DOpDescriptor {
	return DepthwiseConvolution3DOpDescriptorClass.DescriptorWithStridesDilationRatesPaddingValuesPaddingStyle(strides, dilationRates, paddingValues, paddingStyle)
}

func (dc _DepthwiseConvolution3DOpDescriptorClass) DescriptorWithPaddingStyle(paddingStyle PaddingStyle) DepthwiseConvolution3DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution3DOpDescriptor](dc, objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750683-descriptorwithpaddingstyle?language=objc
func DepthwiseConvolution3DOpDescriptor_DescriptorWithPaddingStyle(paddingStyle PaddingStyle) DepthwiseConvolution3DOpDescriptor {
	return DepthwiseConvolution3DOpDescriptorClass.DescriptorWithPaddingStyle(paddingStyle)
}

func (dc _DepthwiseConvolution3DOpDescriptorClass) Alloc() DepthwiseConvolution3DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution3DOpDescriptor](dc, objc.Sel("alloc"))
	return rv
}

func DepthwiseConvolution3DOpDescriptor_Alloc() DepthwiseConvolution3DOpDescriptor {
	return DepthwiseConvolution3DOpDescriptorClass.Alloc()
}

func (dc _DepthwiseConvolution3DOpDescriptorClass) New() DepthwiseConvolution3DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution3DOpDescriptor](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDepthwiseConvolution3DOpDescriptor() DepthwiseConvolution3DOpDescriptor {
	return DepthwiseConvolution3DOpDescriptorClass.New()
}

func (d_ DepthwiseConvolution3DOpDescriptor) Init() DepthwiseConvolution3DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution3DOpDescriptor](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750686-paddingstyle?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](d_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750686-paddingstyle?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3787589-channeldimensionindex?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) ChannelDimensionIndex() int {
	rv := objc.Call[int](d_, objc.Sel("channelDimensionIndex"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3787589-channeldimensionindex?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) SetChannelDimensionIndex(value int) {
	objc.Call[objc.Void](d_, objc.Sel("setChannelDimensionIndex:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750687-paddingvalues?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) PaddingValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](d_, objc.Sel("paddingValues"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750687-paddingvalues?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) SetPaddingValues(value []foundation.INumber) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingValues:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750685-dilationrates?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) DilationRates() []foundation.Number {
	rv := objc.Call[[]foundation.Number](d_, objc.Sel("dilationRates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750685-dilationrates?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) SetDilationRates(value []foundation.INumber) {
	objc.Call[objc.Void](d_, objc.Sel("setDilationRates:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750688-strides?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) Strides() []foundation.Number {
	rv := objc.Call[[]foundation.Number](d_, objc.Sel("strides"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution3dopdescriptor/3750688-strides?language=objc
func (d_ DepthwiseConvolution3DOpDescriptor) SetStrides(value []foundation.INumber) {
	objc.Call[objc.Void](d_, objc.Sel("setStrides:"), value)
}
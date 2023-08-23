// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSpatialNormalizationGradient] class.
var CNNSpatialNormalizationGradientClass = _CNNSpatialNormalizationGradientClass{objc.GetClass("MPSCNNSpatialNormalizationGradient")}

type _CNNSpatialNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNSpatialNormalizationGradient] class.
type ICNNSpatialNormalizationGradient interface {
	ICNNGradientKernel
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// A gradient spatial normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient?language=objc
type CNNSpatialNormalizationGradient struct {
	CNNGradientKernel
}

func CNNSpatialNormalizationGradientFrom(ptr unsafe.Pointer) CNNSpatialNormalizationGradient {
	return CNNSpatialNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNSpatialNormalizationGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNSpatialNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalizationGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942461-initwithdevice?language=objc
func NewCNNSpatialNormalizationGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNSpatialNormalizationGradient {
	instance := CNNSpatialNormalizationGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNSpatialNormalizationGradientClass) Alloc() CNNSpatialNormalizationGradient {
	rv := objc.Call[CNNSpatialNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNSpatialNormalizationGradient_Alloc() CNNSpatialNormalizationGradient {
	return CNNSpatialNormalizationGradientClass.Alloc()
}

func (cc _CNNSpatialNormalizationGradientClass) New() CNNSpatialNormalizationGradient {
	rv := objc.Call[CNNSpatialNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSpatialNormalizationGradient() CNNSpatialNormalizationGradient {
	return CNNSpatialNormalizationGradientClass.New()
}

func (c_ CNNSpatialNormalizationGradient) Init() CNNSpatialNormalizationGradient {
	rv := objc.Call[CNNSpatialNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSpatialNormalizationGradient) InitWithDevice(device metal.PDevice) CNNSpatialNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNSpatialNormalizationGradientWithDevice(device metal.PDevice) CNNSpatialNormalizationGradient {
	instance := CNNSpatialNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNSpatialNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSpatialNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSpatialNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSpatialNormalizationGradient {
	instance := CNNSpatialNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942470-beta?language=objc
func (c_ CNNSpatialNormalizationGradient) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942470-beta?language=objc
func (c_ CNNSpatialNormalizationGradient) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942486-delta?language=objc
func (c_ CNNSpatialNormalizationGradient) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942486-delta?language=objc
func (c_ CNNSpatialNormalizationGradient) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942478-alpha?language=objc
func (c_ CNNSpatialNormalizationGradient) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradient/2942478-alpha?language=objc
func (c_ CNNSpatialNormalizationGradient) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
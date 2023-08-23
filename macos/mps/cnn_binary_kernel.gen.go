// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBinaryKernel] class.
var CNNBinaryKernelClass = _CNNBinaryKernelClass{objc.GetClass("MPSCNNBinaryKernel")}

type _CNNBinaryKernelClass struct {
	objc.Class
}

// An interface definition for the [CNNBinaryKernel] class.
type ICNNBinaryKernel interface {
	IKernel
	BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) uint
	TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State
	TemporaryResultStateForCommandBufferObjectPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State
	ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State
	AppendBatchBarrier() bool
	TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	TemporaryResultStateBatchForCommandBufferObjectPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBufferObject objc.IObject, primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer metal.PCommandBuffer, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationImages *foundation.Array)
	EncodeBatchToCommandBufferObjectPrimaryImagesSecondaryImagesDestinationImages(commandBufferObject objc.IObject, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationImages *foundation.Array)
	ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) uint
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, outState IState, isTemporary bool) Image
	EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, outState IState, isTemporary bool) Image
	IsResultStateReusedAcrossBatch() bool
	SecondaryKernelHeight() uint
	PrimarySourceFeatureChannelOffset() uint
	SetPrimarySourceFeatureChannelOffset(value uint)
	IsStateModified() bool
	SecondaryEdgeMode() ImageEdgeMode
	SetSecondaryEdgeMode(value ImageEdgeMode)
	PrimaryKernelHeight() uint
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
	SecondarySourceFeatureChannelOffset() uint
	SetSecondarySourceFeatureChannelOffset(value uint)
	SecondaryDilationRateY() uint
	PrimaryDilationRateY() uint
	SecondaryKernelWidth() uint
	PrimaryEdgeMode() ImageEdgeMode
	SetPrimaryEdgeMode(value ImageEdgeMode)
	PrimaryDilationRateX() uint
	Padding() NNPaddingWrapper
	SetPadding(value PNNPadding)
	SetPaddingObject(valueObject objc.IObject)
	SecondarySourceFeatureChannelMaxCount() uint
	SetSecondarySourceFeatureChannelMaxCount(value uint)
	SecondaryDilationRateX() uint
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	PrimarySourceFeatureChannelMaxCount() uint
	SetPrimarySourceFeatureChannelMaxCount(value uint)
	PrimaryStrideInPixelsY() uint
	SetPrimaryStrideInPixelsY(value uint)
	PrimaryOffset() Offset
	SetPrimaryOffset(value Offset)
	PrimaryStrideInPixelsX() uint
	SetPrimaryStrideInPixelsX(value uint)
	IsBackwards() bool
	SecondaryOffset() Offset
	SetSecondaryOffset(value Offset)
	ClipRect() metal.Region
	SetClipRect(value metal.Region)
	DestinationImageAllocator() ImageAllocatorWrapper
	SetDestinationImageAllocator(value PImageAllocator)
	SetDestinationImageAllocatorObject(valueObject objc.IObject)
	PrimaryKernelWidth() uint
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
}

// A convolution neural network kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel?language=objc
type CNNBinaryKernel struct {
	Kernel
}

func CNNBinaryKernelFrom(ptr unsafe.Pointer) CNNBinaryKernel {
	return CNNBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (c_ CNNBinaryKernel) InitWithDevice(device metal.PDevice) CNNBinaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryKernel](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice?language=objc
func NewCNNBinaryKernelWithDevice(device metal.PDevice) CNNBinaryKernel {
	instance := CNNBinaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryKernelClass) Alloc() CNNBinaryKernel {
	rv := objc.Call[CNNBinaryKernel](cc, objc.Sel("alloc"))
	return rv
}

func CNNBinaryKernel_Alloc() CNNBinaryKernel {
	return CNNBinaryKernelClass.Alloc()
}

func (cc _CNNBinaryKernelClass) New() CNNBinaryKernel {
	rv := objc.Call[CNNBinaryKernel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBinaryKernel() CNNBinaryKernel {
	return CNNBinaryKernelClass.New()
}

func (c_ CNNBinaryKernel) Init() CNNBinaryKernel {
	rv := objc.Call[CNNBinaryKernel](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBinaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryKernel](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBinaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryKernel {
	instance := CNNBinaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237261-batchencodingstoragesizeforprima?language=objc
func (c_ CNNBinaryKernel) BatchEncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) uint {
	rv := objc.Call[uint](c_, objc.Sel("batchEncodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstateforcommandbu?language=objc
func (c_ CNNBinaryKernel) TemporaryResultStateForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), po0, objc.Ptr(primaryImage), objc.Ptr(secondaryImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947938-temporaryresultstateforcommandbu?language=objc
func (c_ CNNBinaryKernel) TemporaryResultStateForCommandBufferObjectPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947935-resultstateforprimaryimage?language=objc
func (c_ CNNBinaryKernel) ResultStateForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("resultStateForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942642-appendbatchbarrier?language=objc
func (c_ CNNBinaryKernel) AppendBatchBarrier() bool {
	rv := objc.Call[bool](c_, objc.Sel("appendBatchBarrier"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNBinaryKernel) TemporaryResultStateBatchForCommandBufferPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), po0, primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947933-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNBinaryKernel) TemporaryResultStateBatchForCommandBufferObjectPrimaryImageSecondaryImageSourceStatesDestinationImage(commandBufferObject objc.IObject, primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:primaryImage:secondaryImage:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942686-destinationimagedescriptorforsou?language=objc
func (c_ CNNBinaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](c_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942670-encodebatchtocommandbuffer?language=objc
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationImages(commandBuffer metal.PCommandBuffer, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationImages *foundation.Array) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationImages:"), po0, primaryImages, secondaryImages, destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942670-encodebatchtocommandbuffer?language=objc
func (c_ CNNBinaryKernel) EncodeBatchToCommandBufferObjectPrimaryImagesSecondaryImagesDestinationImages(commandBufferObject objc.IObject, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationImages *foundation.Array) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationImages:"), objc.Ptr(commandBufferObject), primaryImages, secondaryImages, destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947930-resultstatebatchforprimaryimage?language=objc
func (c_ CNNBinaryKernel) ResultStateBatchForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage *foundation.Array, secondaryImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("resultStateBatchForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), primaryImage, secondaryImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/3237262-encodingstoragesizeforprimaryima?language=objc
func (c_ CNNBinaryKernel) EncodingStorageSizeForPrimaryImageSecondaryImageSourceStatesDestinationImage(primaryImage IImage, secondaryImage IImage, sourceStates []IState, destinationImage IImage) uint {
	rv := objc.Call[uint](c_, objc.Sel("encodingStorageSizeForPrimaryImage:secondaryImage:sourceStates:destinationImage:"), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947936-encodetocommandbuffer?language=objc
func (c_ CNNBinaryKernel) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, outState IState, isTemporary bool) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:"), po0, objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(outState), isTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2947936-encodetocommandbuffer?language=objc
func (c_ CNNBinaryKernel) EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationStateDestinationStateIsTemporary(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, outState IState, isTemporary bool) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationStateIsTemporary:"), objc.Ptr(commandBufferObject), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(outState), isTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942659-isresultstatereusedacrossbatch?language=objc
func (c_ CNNBinaryKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942664-secondarykernelheight?language=objc
func (c_ CNNBinaryKernel) SecondaryKernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryKernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse?language=objc
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("primarySourceFeatureChannelOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942656-primarysourcefeaturechanneloffse?language=objc
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimarySourceFeatureChannelOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942660-isstatemodified?language=objc
func (c_ CNNBinaryKernel) IsStateModified() bool {
	rv := objc.Call[bool](c_, objc.Sel("isStateModified"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode?language=objc
func (c_ CNNBinaryKernel) SecondaryEdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](c_, objc.Sel("secondaryEdgeMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865631-secondaryedgemode?language=objc
func (c_ CNNBinaryKernel) SetSecondaryEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryEdgeMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942648-primarykernelheight?language=objc
func (c_ CNNBinaryKernel) PrimaryKernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryKernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy?language=objc
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865639-secondarystrideinpixelsy?language=objc
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff?language=objc
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondarySourceFeatureChannelOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942654-secondarysourcefeaturechanneloff?language=objc
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondarySourceFeatureChannelOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942667-secondarydilationratey?language=objc
func (c_ CNNBinaryKernel) SecondaryDilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryDilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942662-primarydilationratey?language=objc
func (c_ CNNBinaryKernel) PrimaryDilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryDilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942658-secondarykernelwidth?language=objc
func (c_ CNNBinaryKernel) SecondaryKernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryKernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode?language=objc
func (c_ CNNBinaryKernel) PrimaryEdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](c_, objc.Sel("primaryEdgeMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865646-primaryedgemode?language=objc
func (c_ CNNBinaryKernel) SetPrimaryEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryEdgeMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942677-primarydilationratex?language=objc
func (c_ CNNBinaryKernel) PrimaryDilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryDilationRateX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding?language=objc
func (c_ CNNBinaryKernel) Padding() NNPaddingWrapper {
	rv := objc.Call[NNPaddingWrapper](c_, objc.Sel("padding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding?language=objc
func (c_ CNNBinaryKernel) SetPadding(value PNNPadding) {
	po0 := objc.WrapAsProtocol("MPSNNPadding", value)
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865630-padding?language=objc
func (c_ CNNBinaryKernel) SetPaddingObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax?language=objc
func (c_ CNNBinaryKernel) SecondarySourceFeatureChannelMaxCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondarySourceFeatureChannelMaxCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951918-secondarysourcefeaturechannelmax?language=objc
func (c_ CNNBinaryKernel) SetSecondarySourceFeatureChannelMaxCount(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondarySourceFeatureChannelMaxCount:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942645-secondarydilationratex?language=objc
func (c_ CNNBinaryKernel) SecondaryDilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryDilationRateX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx?language=objc
func (c_ CNNBinaryKernel) SecondaryStrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865649-secondarystrideinpixelsx?language=objc
func (c_ CNNBinaryKernel) SetSecondaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco?language=objc
func (c_ CNNBinaryKernel) PrimarySourceFeatureChannelMaxCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("primarySourceFeatureChannelMaxCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2951919-primarysourcefeaturechannelmaxco?language=objc
func (c_ CNNBinaryKernel) SetPrimarySourceFeatureChannelMaxCount(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimarySourceFeatureChannelMaxCount:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy?language=objc
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865656-primarystrideinpixelsy?language=objc
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset?language=objc
func (c_ CNNBinaryKernel) PrimaryOffset() Offset {
	rv := objc.Call[Offset](c_, objc.Sel("primaryOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865645-primaryoffset?language=objc
func (c_ CNNBinaryKernel) SetPrimaryOffset(value Offset) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx?language=objc
func (c_ CNNBinaryKernel) PrimaryStrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865658-primarystrideinpixelsx?language=objc
func (c_ CNNBinaryKernel) SetPrimaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865652-isbackwards?language=objc
func (c_ CNNBinaryKernel) IsBackwards() bool {
	rv := objc.Call[bool](c_, objc.Sel("isBackwards"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset?language=objc
func (c_ CNNBinaryKernel) SecondaryOffset() Offset {
	rv := objc.Call[Offset](c_, objc.Sel("secondaryOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865628-secondaryoffset?language=objc
func (c_ CNNBinaryKernel) SetSecondaryOffset(value Offset) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect?language=objc
func (c_ CNNBinaryKernel) ClipRect() metal.Region {
	rv := objc.Call[metal.Region](c_, objc.Sel("clipRect"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865641-cliprect?language=objc
func (c_ CNNBinaryKernel) SetClipRect(value metal.Region) {
	objc.Call[objc.Void](c_, objc.Sel("setClipRect:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator?language=objc
func (c_ CNNBinaryKernel) DestinationImageAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](c_, objc.Sel("destinationImageAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator?language=objc
func (c_ CNNBinaryKernel) SetDestinationImageAllocator(value PImageAllocator) {
	po0 := objc.WrapAsProtocol("MPSImageAllocator", value)
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865651-destinationimageallocator?language=objc
func (c_ CNNBinaryKernel) SetDestinationImageAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2942666-primarykernelwidth?language=objc
func (c_ CNNBinaryKernel) PrimaryKernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryKernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset?language=objc
func (c_ CNNBinaryKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865643-destinationfeaturechanneloffset?language=objc
func (c_ CNNBinaryKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}
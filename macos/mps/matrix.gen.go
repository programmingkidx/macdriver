// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Matrix] class.
var MatrixClass = _MatrixClass{objc.GetClass("MPSMatrix")}

type _MatrixClass struct {
	objc.Class
}

// An interface definition for the [Matrix] class.
type IMatrix interface {
	objc.IObject
	ResourceSize() uint
	SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject)
	Device() metal.DeviceWrapper
	RowBytes() uint
	Columns() uint
	Data() metal.BufferWrapper
	Offset() uint
	DataType() DataType
	Rows() uint
	Matrices() uint
	MatrixBytes() uint
}

// A 2D array of data that stores the data's values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix?language=objc
type Matrix struct {
	objc.Object
}

func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ Matrix) InitWithDeviceDescriptor(device metal.PDevice, descriptor IMatrixDescriptor) Matrix {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Matrix](m_, objc.Sel("initWithDevice:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942567-initwithdevice?language=objc
func NewMatrixWithDeviceDescriptor(device metal.PDevice, descriptor IMatrixDescriptor) Matrix {
	instance := MatrixClass.Alloc().InitWithDeviceDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (m_ Matrix) InitWithBufferOffsetDescriptor(buffer metal.PBuffer, offset uint, descriptor IMatrixDescriptor) Matrix {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[Matrix](m_, objc.Sel("initWithBuffer:offset:descriptor:"), po0, offset, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3229863-initwithbuffer?language=objc
func NewMatrixWithBufferOffsetDescriptor(buffer metal.PBuffer, offset uint, descriptor IMatrixDescriptor) Matrix {
	instance := MatrixClass.Alloc().InitWithBufferOffsetDescriptor(buffer, offset, descriptor)
	instance.Autorelease()
	return instance
}

func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Call[Matrix](mc, objc.Sel("alloc"))
	return rv
}

func Matrix_Alloc() Matrix {
	return MatrixClass.Alloc()
}

func (mc _MatrixClass) New() Matrix {
	rv := objc.Call[Matrix](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrix() Matrix {
	return MatrixClass.New()
}

func (m_ Matrix) Init() Matrix {
	rv := objc.Call[Matrix](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942569-resourcesize?language=objc
func (m_ Matrix) ResourceSize() uint {
	rv := objc.Call[uint](m_, objc.Sel("resourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942571-synchronizeoncommandbuffer?language=objc
func (m_ Matrix) SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("synchronizeOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2942571-synchronizeoncommandbuffer?language=objc
func (m_ Matrix) SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("synchronizeOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}

// The device on which the matrix will be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143209-device?language=objc
func (m_ Matrix) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](m_, objc.Sel("device"))
	return rv
}

// The stride, in bytes, between corresponding elements of consecutive rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143208-rowbytes?language=objc
func (m_ Matrix) RowBytes() uint {
	rv := objc.Call[uint](m_, objc.Sel("rowBytes"))
	return rv
}

// The number of columns in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143207-columns?language=objc
func (m_ Matrix) Columns() uint {
	rv := objc.Call[uint](m_, objc.Sel("columns"))
	return rv
}

// The buffer that stores the matrix data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143205-data?language=objc
func (m_ Matrix) Data() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](m_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/3375740-offset?language=objc
func (m_ Matrix) Offset() uint {
	rv := objc.Call[uint](m_, objc.Sel("offset"))
	return rv
}

// The type of the values in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143197-datatype?language=objc
func (m_ Matrix) DataType() DataType {
	rv := objc.Call[DataType](m_, objc.Sel("dataType"))
	return rv
}

// The number of rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2143210-rows?language=objc
func (m_ Matrix) Rows() uint {
	rv := objc.Call[uint](m_, objc.Sel("rows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873334-matrices?language=objc
func (m_ Matrix) Matrices() uint {
	rv := objc.Call[uint](m_, objc.Sel("matrices"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrix/2873344-matrixbytes?language=objc
func (m_ Matrix) MatrixBytes() uint {
	rv := objc.Call[uint](m_, objc.Sel("matrixBytes"))
	return rv
}
// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Executable] class.
var ExecutableClass = _ExecutableClass{objc.GetClass("MPSGraphExecutable")}

type _ExecutableClass struct {
	objc.Class
}

// An interface definition for the [Executable] class.
type IExecutable interface {
	objc.IObject
	RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.PCommandQueue, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData
	RunWithMTLCommandQueueObjectInputsArrayResultsArrayExecutionDescriptor(commandQueueObject objc.IObject, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData
	RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.PCommandQueue, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData
	RunAsyncWithMTLCommandQueueObjectInputsArrayResultsArrayExecutionDescriptor(commandQueueObject objc.IObject, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData
	SpecializeWithDeviceInputTypesCompilationDescriptor(device IDevice, inputTypes []IType, compilationDescriptor ICompilationDescriptor)
	EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer mps.ICommandBuffer, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData
	Options() Options
	SetOptions(value Options)
	FeedTensors() []Tensor
	TargetTensors() []Tensor
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable?language=objc
type Executable struct {
	objc.Object
}

func ExecutableFrom(ptr unsafe.Pointer) Executable {
	return Executable{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExecutableClass) Alloc() Executable {
	rv := objc.Call[Executable](ec, objc.Sel("alloc"))
	return rv
}

func Executable_Alloc() Executable {
	return ExecutableClass.Alloc()
}

func (ec _ExecutableClass) New() Executable {
	rv := objc.Call[Executable](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExecutable() Executable {
	return ExecutableClass.New()
}

func (e_ Executable) Init() Executable {
	rv := objc.Call[Executable](e_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787595-runwithmtlcommandqueue?language=objc
func (e_ Executable) RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.PCommandQueue, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	rv := objc.Call[[]TensorData](e_, objc.Sel("runWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), po0, inputsArray, resultsArray, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787595-runwithmtlcommandqueue?language=objc
func (e_ Executable) RunWithMTLCommandQueueObjectInputsArrayResultsArrayExecutionDescriptor(commandQueueObject objc.IObject, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData {
	rv := objc.Call[[]TensorData](e_, objc.Sel("runWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), objc.Ptr(commandQueueObject), inputsArray, resultsArray, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787594-runasyncwithmtlcommandqueue?language=objc
func (e_ Executable) RunAsyncWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(commandQueue metal.PCommandQueue, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData {
	po0 := objc.WrapAsProtocol("MTLCommandQueue", commandQueue)
	rv := objc.Call[[]TensorData](e_, objc.Sel("runAsyncWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), po0, inputsArray, resultsArray, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787594-runasyncwithmtlcommandqueue?language=objc
func (e_ Executable) RunAsyncWithMTLCommandQueueObjectInputsArrayResultsArrayExecutionDescriptor(commandQueueObject objc.IObject, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData {
	rv := objc.Call[[]TensorData](e_, objc.Sel("runAsyncWithMTLCommandQueue:inputsArray:resultsArray:executionDescriptor:"), objc.Ptr(commandQueueObject), inputsArray, resultsArray, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787596-specializewithdevice?language=objc
func (e_ Executable) SpecializeWithDeviceInputTypesCompilationDescriptor(device IDevice, inputTypes []IType, compilationDescriptor ICompilationDescriptor) {
	objc.Call[objc.Void](e_, objc.Sel("specializeWithDevice:inputTypes:compilationDescriptor:"), objc.Ptr(device), inputTypes, objc.Ptr(compilationDescriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787591-encodetocommandbuffer?language=objc
func (e_ Executable) EncodeToCommandBufferInputsArrayResultsArrayExecutionDescriptor(commandBuffer mps.ICommandBuffer, inputsArray []ITensorData, resultsArray []ITensorData, executionDescriptor IExecutableExecutionDescriptor) []TensorData {
	rv := objc.Call[[]TensorData](e_, objc.Sel("encodeToCommandBuffer:inputsArray:resultsArray:executionDescriptor:"), objc.Ptr(commandBuffer), inputsArray, resultsArray, objc.Ptr(executionDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787593-options?language=objc
func (e_ Executable) Options() Options {
	rv := objc.Call[Options](e_, objc.Sel("options"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787593-options?language=objc
func (e_ Executable) SetOptions(value Options) {
	objc.Call[objc.Void](e_, objc.Sel("setOptions:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787592-feedtensors?language=objc
func (e_ Executable) FeedTensors() []Tensor {
	rv := objc.Call[[]Tensor](e_, objc.Sel("feedTensors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutable/3787597-targettensors?language=objc
func (e_ Executable) TargetTensors() []Tensor {
	rv := objc.Call[[]Tensor](e_, objc.Sel("targetTensors"))
	return rv
}
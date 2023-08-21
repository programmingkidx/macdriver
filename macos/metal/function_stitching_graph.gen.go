// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FunctionStitchingGraph] class.
var FunctionStitchingGraphClass = _FunctionStitchingGraphClass{objc.GetClass("MTLFunctionStitchingGraph")}

type _FunctionStitchingGraphClass struct {
	objc.Class
}

// An interface definition for the [FunctionStitchingGraph] class.
type IFunctionStitchingGraph interface {
	objc.IObject
	FunctionName() string
	SetFunctionName(value string)
	OutputNode() FunctionStitchingFunctionNode
	SetOutputNode(value IFunctionStitchingFunctionNode)
	Nodes() []FunctionStitchingFunctionNode
	SetNodes(value []IFunctionStitchingFunctionNode)
	Attributes() []FunctionStitchingAttributeWrapper
	SetAttributes(value []PFunctionStitchingAttribute)
}

// A description of a new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph?language=objc
type FunctionStitchingGraph struct {
	objc.Object
}

func FunctionStitchingGraphFrom(ptr unsafe.Pointer) FunctionStitchingGraph {
	return FunctionStitchingGraph{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FunctionStitchingGraph) InitWithFunctionNameNodesOutputNodeAttributes(functionName string, nodes []IFunctionStitchingFunctionNode, outputNode IFunctionStitchingFunctionNode, attributes []PFunctionStitchingAttribute) FunctionStitchingGraph {
	rv := objc.Call[FunctionStitchingGraph](f_, objc.Sel("initWithFunctionName:nodes:outputNode:attributes:"), functionName, nodes, objc.Ptr(outputNode), attributes)
	return rv
}

// Creates a description of a new function call graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750541-initwithfunctionname?language=objc
func NewFunctionStitchingGraphWithFunctionNameNodesOutputNodeAttributes(functionName string, nodes []IFunctionStitchingFunctionNode, outputNode IFunctionStitchingFunctionNode, attributes []PFunctionStitchingAttribute) FunctionStitchingGraph {
	instance := FunctionStitchingGraphClass.Alloc().InitWithFunctionNameNodesOutputNodeAttributes(functionName, nodes, outputNode, attributes)
	instance.Autorelease()
	return instance
}

func (fc _FunctionStitchingGraphClass) Alloc() FunctionStitchingGraph {
	rv := objc.Call[FunctionStitchingGraph](fc, objc.Sel("alloc"))
	return rv
}

func FunctionStitchingGraph_Alloc() FunctionStitchingGraph {
	return FunctionStitchingGraphClass.Alloc()
}

func (fc _FunctionStitchingGraphClass) New() FunctionStitchingGraph {
	rv := objc.Call[FunctionStitchingGraph](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFunctionStitchingGraph() FunctionStitchingGraph {
	return FunctionStitchingGraphClass.New()
}

func (f_ FunctionStitchingGraph) Init() FunctionStitchingGraph {
	rv := objc.Call[FunctionStitchingGraph](f_, objc.Sel("init"))
	return rv
}

// The name of the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750540-functionname?language=objc
func (f_ FunctionStitchingGraph) FunctionName() string {
	rv := objc.Call[string](f_, objc.Sel("functionName"))
	return rv
}

// The name of the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750540-functionname?language=objc
func (f_ FunctionStitchingGraph) SetFunctionName(value string) {
	objc.Call[objc.Void](f_, objc.Sel("setFunctionName:"), value)
}

// The node with the output that’s the output of the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750543-outputnode?language=objc
func (f_ FunctionStitchingGraph) OutputNode() FunctionStitchingFunctionNode {
	rv := objc.Call[FunctionStitchingFunctionNode](f_, objc.Sel("outputNode"))
	return rv
}

// The node with the output that’s the output of the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750543-outputnode?language=objc
func (f_ FunctionStitchingGraph) SetOutputNode(value IFunctionStitchingFunctionNode) {
	objc.Call[objc.Void](f_, objc.Sel("setOutputNode:"), objc.Ptr(value))
}

// The nodes in the function’s call graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750542-nodes?language=objc
func (f_ FunctionStitchingGraph) Nodes() []FunctionStitchingFunctionNode {
	rv := objc.Call[[]FunctionStitchingFunctionNode](f_, objc.Sel("nodes"))
	return rv
}

// The nodes in the function’s call graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750542-nodes?language=objc
func (f_ FunctionStitchingGraph) SetNodes(value []IFunctionStitchingFunctionNode) {
	objc.Call[objc.Void](f_, objc.Sel("setNodes:"), value)
}

// A list of attributes to configure how the Metal device object generates the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750539-attributes?language=objc
func (f_ FunctionStitchingGraph) Attributes() []FunctionStitchingAttributeWrapper {
	rv := objc.Call[[]FunctionStitchingAttributeWrapper](f_, objc.Sel("attributes"))
	return rv
}

// A list of attributes to configure how the Metal device object generates the new stitched function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionstitchinggraph/3750539-attributes?language=objc
func (f_ FunctionStitchingGraph) SetAttributes(value []PFunctionStitchingAttribute) {
	objc.Call[objc.Void](f_, objc.Sel("setAttributes:"), value)
}
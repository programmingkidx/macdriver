// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingMaxNode] class.
var CNNPoolingMaxNodeClass = _CNNPoolingMaxNodeClass{objc.GetClass("MPSCNNPoolingMaxNode")}

type _CNNPoolingMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingMaxNode] class.
type ICNNPoolingMaxNode interface {
	ICNNPoolingNode
}

// A representation of a max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxnode?language=objc
type CNNPoolingMaxNode struct {
	CNNPoolingNode
}

func CNNPoolingMaxNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxNode {
	return CNNPoolingMaxNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}

func (cc _CNNPoolingMaxNodeClass) Alloc() CNNPoolingMaxNode {
	rv := objc.Call[CNNPoolingMaxNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingMaxNode_Alloc() CNNPoolingMaxNode {
	return CNNPoolingMaxNodeClass.Alloc()
}

func (cc _CNNPoolingMaxNodeClass) New() CNNPoolingMaxNode {
	rv := objc.Call[CNNPoolingMaxNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingMaxNode() CNNPoolingMaxNode {
	return CNNPoolingMaxNodeClass.New()
}

func (c_ CNNPoolingMaxNode) Init() CNNPoolingMaxNode {
	rv := objc.Call[CNNPoolingMaxNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNPoolingMaxNodeClass) NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingMaxNode {
	rv := objc.Call[CNNPoolingMaxNode](cc, objc.Sel("nodeWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource?language=objc
func CNNPoolingMaxNode_NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingMaxNode {
	return CNNPoolingMaxNodeClass.NodeWithSourceFilterSize(sourceNode, size)
}

func (c_ CNNPoolingMaxNode) InitWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingMaxNode {
	rv := objc.Call[CNNPoolingMaxNode](c_, objc.Sel("initWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource?language=objc
func NewCNNPoolingMaxNodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingMaxNode {
	instance := CNNPoolingMaxNodeClass.Alloc().InitWithSourceFilterSize(sourceNode, size)
	instance.Autorelease()
	return instance
}
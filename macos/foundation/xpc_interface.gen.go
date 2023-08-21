// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XPCInterface] class.
var XPCInterfaceClass = _XPCInterfaceClass{objc.GetClass("NSXPCInterface")}

type _XPCInterfaceClass struct {
	objc.Class
}

// An interface definition for the [XPCInterface] class.
type IXPCInterface interface {
	objc.IObject
	SetClassesForSelectorArgumentIndexOfReply(classes ISet, sel objc.Selector, arg uint, ofReply bool)
	ClassesForSelectorArgumentIndexOfReply(sel objc.Selector, arg uint, ofReply bool) Set
	InterfaceForSelectorArgumentIndexOfReply(sel objc.Selector, arg uint, ofReply bool) XPCInterface
	SetInterfaceForSelectorArgumentIndexOfReply(ifc IXPCInterface, sel objc.Selector, arg uint, ofReply bool)
	Protocol() objc.Protocol
	SetProtocol(value objc.IProtocol)
}

// An interface that may be sent to an exported object or remote object proxy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface?language=objc
type XPCInterface struct {
	objc.Object
}

func XPCInterfaceFrom(ptr unsafe.Pointer) XPCInterface {
	return XPCInterface{
		Object: objc.ObjectFrom(ptr),
	}
}

func (xc _XPCInterfaceClass) Alloc() XPCInterface {
	rv := objc.Call[XPCInterface](xc, objc.Sel("alloc"))
	return rv
}

func XPCInterface_Alloc() XPCInterface {
	return XPCInterfaceClass.Alloc()
}

func (xc _XPCInterfaceClass) New() XPCInterface {
	rv := objc.Call[XPCInterface](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXPCInterface() XPCInterface {
	return XPCInterfaceClass.New()
}

func (x_ XPCInterface) Init() XPCInterface {
	rv := objc.Call[XPCInterface](x_, objc.Sel("init"))
	return rv
}

// Sets the classes that can appear within the (numerically) specified collection object argument to the specified method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1415555-setclasses?language=objc
func (x_ XPCInterface) SetClassesForSelectorArgumentIndexOfReply(classes ISet, sel objc.Selector, arg uint, ofReply bool) {
	objc.Call[objc.Void](x_, objc.Sel("setClasses:forSelector:argumentIndex:ofReply:"), objc.Ptr(classes), sel, arg, ofReply)
}

// Returns an NSXPCInterface instance for a given protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1410202-interfacewithprotocol?language=objc
func (xc _XPCInterfaceClass) InterfaceWithProtocol(protocol objc.IProtocol) XPCInterface {
	rv := objc.Call[XPCInterface](xc, objc.Sel("interfaceWithProtocol:"), objc.Ptr(protocol))
	return rv
}

// Returns an NSXPCInterface instance for a given protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1410202-interfacewithprotocol?language=objc
func XPCInterface_InterfaceWithProtocol(protocol objc.IProtocol) XPCInterface {
	return XPCInterfaceClass.InterfaceWithProtocol(protocol)
}

// Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1418323-classesforselector?language=objc
func (x_ XPCInterface) ClassesForSelectorArgumentIndexOfReply(sel objc.Selector, arg uint, ofReply bool) Set {
	rv := objc.Call[Set](x_, objc.Sel("classesForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}

// Returns the interface previously set for the specified selector and parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1409127-interfaceforselector?language=objc
func (x_ XPCInterface) InterfaceForSelectorArgumentIndexOfReply(sel objc.Selector, arg uint, ofReply bool) XPCInterface {
	rv := objc.Call[XPCInterface](x_, objc.Sel("interfaceForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}

// Configures a specific parameter of a method to be sent as a proxy object instead of copied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1414293-setinterface?language=objc
func (x_ XPCInterface) SetInterfaceForSelectorArgumentIndexOfReply(ifc IXPCInterface, sel objc.Selector, arg uint, ofReply bool) {
	objc.Call[objc.Void](x_, objc.Sel("setInterface:forSelector:argumentIndex:ofReply:"), objc.Ptr(ifc), sel, arg, ofReply)
}

// The Objective-C protocol that this interface is based on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1416353-protocol?language=objc
func (x_ XPCInterface) Protocol() objc.Protocol {
	rv := objc.Call[objc.Protocol](x_, objc.Sel("protocol"))
	return rv
}

// The Objective-C protocol that this interface is based on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcinterface/1416353-protocol?language=objc
func (x_ XPCInterface) SetProtocol(value objc.IProtocol) {
	objc.Call[objc.Void](x_, objc.Sel("setProtocol:"), objc.Ptr(value))
}
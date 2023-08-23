// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionPickerPanel] class.
var CompositionPickerPanelClass = _CompositionPickerPanelClass{objc.GetClass("QCCompositionPickerPanel")}

type _CompositionPickerPanelClass struct {
	objc.Class
}

// An interface definition for the [CompositionPickerPanel] class.
type ICompositionPickerPanel interface {
	appkit.IPanel
}

// The QCCompositionPickerPanel class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The QCCompositionPickerPanel class cannot be subclassed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccompositionpickerpanel?language=objc
type CompositionPickerPanel struct {
	appkit.Panel
}

func CompositionPickerPanelFrom(ptr unsafe.Pointer) CompositionPickerPanel {
	return CompositionPickerPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

func (cc _CompositionPickerPanelClass) Alloc() CompositionPickerPanel {
	rv := objc.Call[CompositionPickerPanel](cc, objc.Sel("alloc"))
	return rv
}

func CompositionPickerPanel_Alloc() CompositionPickerPanel {
	return CompositionPickerPanelClass.Alloc()
}

func (cc _CompositionPickerPanelClass) New() CompositionPickerPanel {
	rv := objc.Call[CompositionPickerPanel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionPickerPanel() CompositionPickerPanel {
	return CompositionPickerPanelClass.New()
}

func (c_ CompositionPickerPanel) Init() CompositionPickerPanel {
	rv := objc.Call[CompositionPickerPanel](c_, objc.Sel("init"))
	return rv
}

func (cc _CompositionPickerPanelClass) WindowWithContentViewController(contentViewController appkit.IViewController) CompositionPickerPanel {
	rv := objc.Call[CompositionPickerPanel](cc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func CompositionPickerPanel_WindowWithContentViewController(contentViewController appkit.IViewController) CompositionPickerPanel {
	return CompositionPickerPanelClass.WindowWithContentViewController(contentViewController)
}

func (c_ CompositionPickerPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) CompositionPickerPanel {
	rv := objc.Call[CompositionPickerPanel](c_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewCompositionPickerPanelWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) CompositionPickerPanel {
	instance := CompositionPickerPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}
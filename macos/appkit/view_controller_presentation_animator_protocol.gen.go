// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that let you define animations to play when transitioning between two view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator?language=objc
type PViewControllerPresentationAnimator interface {
	// optional
	AnimateDismissalOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
	HasAnimateDismissalOfViewControllerFromViewController() bool

	// optional
	AnimatePresentationOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController)
	HasAnimatePresentationOfViewControllerFromViewController() bool
}

// ensure impl type implements protocol interface
var _ PViewControllerPresentationAnimator = (*ViewControllerPresentationAnimatorObject)(nil)

// A concrete type for the [PViewControllerPresentationAnimator] protocol.
type ViewControllerPresentationAnimatorObject struct {
	objc.Object
}

func (v_ ViewControllerPresentationAnimatorObject) HasAnimateDismissalOfViewControllerFromViewController() bool {
	return v_.RespondsToSelector(objc.Sel("animateDismissalOfViewController:fromViewController:"))
}

// Called when a previously-presented view controller is about to be dismissed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator/1434458-animatedismissalofviewcontroller?language=objc
func (v_ ViewControllerPresentationAnimatorObject) AnimateDismissalOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController) {
	objc.Call[objc.Void](v_, objc.Sel("animateDismissalOfViewController:fromViewController:"), objc.Ptr(viewController), objc.Ptr(fromViewController))
}

func (v_ ViewControllerPresentationAnimatorObject) HasAnimatePresentationOfViewControllerFromViewController() bool {
	return v_.RespondsToSelector(objc.Sel("animatePresentationOfViewController:fromViewController:"))
}

// Called when the specified view controller is about to be presented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontrollerpresentationanimator/1434396-animatepresentationofviewcontrol?language=objc
func (v_ ViewControllerPresentationAnimatorObject) AnimatePresentationOfViewControllerFromViewController(viewController ViewController, fromViewController ViewController) {
	objc.Call[objc.Void](v_, objc.Sel("animatePresentationOfViewController:fromViewController:"), objc.Ptr(viewController), objc.Ptr(fromViewController))
}
// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	. "github.com/progrium/macdriver/pkg/ns/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSBundle struct {
	objc.Object
}

func NSMainBundle() NSBundle {
	return NSBundle{objc.GetClass("NSBundle").SendMsg("mainBundle")}
}

func (bundle NSBundle) InfoDictionary() NSDictionary {
	return NSDictionary{bundle.SendMsg("infoDictionary")}
}

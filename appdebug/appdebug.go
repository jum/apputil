//
//  appdebug.go
//
//  Created by Jens-Uwe Mager on 22.04.15.
//  Copyright Best Search Infobrokerage, Inc 2015. All rights reserved.
//

// Package appdebug is intended to enable and disable the debug logging in an
// appengine environment with a boolean type DebugVar. To conditionalize
// debug output in a module, use it like this:
//	var dbg appdebug.DebugVar = true
//	dbg.Debugf(c, "Hello World")
// The context arg is kept for now to be compatible with older go
// appengine runtimes but is not used.
package appdebug

import (
	"log"

	"golang.org/x/net/context"
)

type DebugVar bool

func (d DebugVar) Debugf(c context.Context, format string, a ...interface{}) {
	if d {
		log.Printf(format, a...)
	}
}

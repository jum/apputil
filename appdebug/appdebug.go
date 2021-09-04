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
// The context arg used when running in go111 runtime mode with the old
// appengine/log import to enable grouping our output under the request
// in the google cloud console log.
package appdebug

import (
	"golang.org/x/net/context"

	aplog "google.golang.org/appengine/v2/log"
)

type DebugVar bool

func (d DebugVar) Debugf(c context.Context, format string, a ...interface{}) {
	if d {
		aplog.Debugf(c, format, a...)
	}
}

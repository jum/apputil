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
	"log"
	"os"
	"sync"

	"golang.org/x/net/context"

	aplog "google.golang.org/appengine/log"
)

type DebugVar bool

var isgo111 = false
var isgo111Check sync.Once

func (d DebugVar) Debugf(c context.Context, format string, a ...interface{}) {
	if d {
		isgo111Check.Do(func() {
			isgo111 = os.Getenv("GAE_RUNTIME") == "go111"
		})
		if isgo111 {
			aplog.Debugf(c, format, a...)
		} else {
			log.Printf(format, a...)
		}
	}
}

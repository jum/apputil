//
//  appredir.go
//
//  Created by Jens-Uwe Mager on 26.04.15.
//  Copyright Best Search Infobrokerage, Inc 2015. All rights reserved.
//

// Package appredir makes it easy to register a redirect handler for one or
// multiple pages to make it easier to restructure web pages and keep
// compatibility. Use like this:
//
//	var redirects = []appredir.RedirEntry{
//		{Path: "/oldpageA.html", Dest: "/path/newpageA.html"},
//		{Path: "/oldpageB.html", Dest: "/path/newpageB.html"},
//	}
//	appredir.RegisterRedirects(redirects)
package appredir

import (
	"net/http"
)

type RedirEntry struct {
	Path string // the local path of the old document
	Dest string // the path to the new document to use instead
}

func (re *RedirEntry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, re.Dest, http.StatusFound)
}

func RegisterRedirect(re *RedirEntry) {
	http.Handle(re.Path, re)
}

func RegisterRedirects(re []RedirEntry) {
	for i := range re {
		RegisterRedirect(&re[i])
	}
}

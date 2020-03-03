//
//  appstart.go
//
//  Created by Jens-Uwe Mager on 22.04.15.
//  Copyright Best Search Infobrokerage, Inc 2015. All rights reserved.
//

//Process any incoming requests looking like:
//
//	http://server/start?lang=xx&vers=1.1
//
//and redirect to the apropriate index.html for that version and
//language. Use via an import for side effect:
//
//	import _ "github.com/jum/apputil/appstart"
//
//Parsing the version and acting upon it is not yet implemented.
package appstart

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// The available languages, en is default and should not be in the list
var AvailLang = map[string][]string{
	"": []string{},
}

func init() {
	http.HandleFunc("/start", start)
}

type acceptLang struct {
	lang       string
	langPrefix string
	prio       float64
}

type acceptLangArray []acceptLang

func (p acceptLangArray) Len() int           { return len(p) }
func (p acceptLangArray) Less(i, j int) bool { return p[i].prio >= p[j].prio }
func (p acceptLangArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func parseAcceptLang(w http.ResponseWriter, r *http.Request) acceptLangArray {
	ret := make(acceptLangArray, 0)
	al := strings.Split(r.Header.Get("Accept-Language"), ",")
	for _, i := range al {
		l := strings.Split(i, ";")
		prio := 1.0
		if len(l) == 2 {
			attr := strings.Split(l[1], "=")
			if attr[0] == "q" {
				v, err := strconv.ParseFloat(attr[1], 64)
				if err != nil {
				} else {
					prio = v
				}
			}
		}
		p := new(acceptLang)
		p.lang = l[0]
		p.prio = prio
		l = strings.Split(p.lang, "-")
		p.langPrefix = l[0]
		ret = append(ret, *p)
	}
	sort.Sort(ret)
	return ret
}

func start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("ParseForm: %v", err)
	}
	lang := r.FormValue("lang")
	version := r.FormValue("vers")
	root := r.FormValue("root")
	file := r.FormValue("file")
	suffix := ""
	if len(lang) != 0 {
		for _, l := range AvailLang[root] {
			if l == lang {
				suffix = "." + l
				break
			}
		}
	} else {
		// The client did not specify an explicit language, so attempt
		// some educated guesses according to his browser locale.
		al := parseAcceptLang(w, r)
	outer:
		for _, a := range al {
			for _, l := range AvailLang[root] {
				if l == a.lang {
					suffix = "." + l
					break outer
				}
			}
		}
		if len(suffix) == 0 {
			// did not find the complete lang-locale pair, try the
			// language alone.
		outer1:
			for _, a := range al {
				for _, l := range AvailLang[root] {
					if l == a.langPrefix {
						suffix = "." + l
						break outer1
					}
				}
			}
		}
	}
	query := ""
	if version != "" {
		if query == "" {
			query = "?"
		}
		query += "vers=" + version
	}
	if len(file) != 0 {
		http.Redirect(w, r, file+suffix+".html"+query, http.StatusFound)
	} else {
		http.Redirect(w, r, root+"/index"+suffix+".html"+query, http.StatusFound)
	}
}

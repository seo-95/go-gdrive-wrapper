package drive

import "net/url"

// pointers to url.URL are used to make literal addressable (to call .String())
var apis = map[string]*url.URL{
	"info": &url.URL{
		Scheme: "https",
		Host:   "www.googleapis.com",
		Path:   "drive/v3/about",
		//Query:  "fields=",
	},
}

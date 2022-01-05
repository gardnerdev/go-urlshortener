package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HanlderFunc (which also
// implements http.Handler) that will attempt to map paths
// to thier coresponding URL
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if we can match the path...
		// redirect to it
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok { //ok is saying, if find key it will be True
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

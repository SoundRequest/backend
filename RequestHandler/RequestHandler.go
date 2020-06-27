package RequestHandler

import "net/http"

func Redirect(destination string, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, destination, 301)
}

package handlers

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	setNoCacheHeaders(w)
	ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

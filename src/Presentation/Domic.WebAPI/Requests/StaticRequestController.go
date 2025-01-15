package WebAPIRequest

import "net/http"

func HandleStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "Static/chat.css")
}

func HandleScript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, "Static/chat.js")
}

func HandlePublicChatPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Static/chat.html")
}

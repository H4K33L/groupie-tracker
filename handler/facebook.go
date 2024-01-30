package handler

import (
	"net/http"
)

func FacebookHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://www.facebook.com/", http.StatusFound)

	err := home.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

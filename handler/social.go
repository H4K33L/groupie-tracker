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

func TwitterHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://twitter.com/", http.StatusFound)

	err := home.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SnapHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://www.snapchat.com/", http.StatusFound)

	err := home.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InstaHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://www.instagram.com", http.StatusFound)

	err := home.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}




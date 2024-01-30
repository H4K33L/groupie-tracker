package handler

import (
	"fmt"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
	fmt.Fprint(w, "Cherche ton artiste favoris !")

}

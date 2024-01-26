package handler

import (
	"fmt"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Cherche ton artiste favoris !")
}
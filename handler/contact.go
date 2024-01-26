package handler

import (
	"fmt"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contactez-nous")
}

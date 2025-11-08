package handler

import (
	"fmt"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui")
}

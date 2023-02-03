package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBarang(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var barang models.Barang

	if result := h.DB.First(&barang, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	h.DB.Delete(&barang)

	w.Header().Add("conten-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}

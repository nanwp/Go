package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetBarang(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])


	var barang models.Barang

	if result := h.DB.Table("tbl_barang").First(&barang, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(barang)

}

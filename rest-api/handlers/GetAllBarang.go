package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/models"
)

func (h handler) GetAllBarang(w http.ResponseWriter, r *http.Request) {

	var barang [] models.Barang

	h.DB.Table("tbl_barang")

	if result := h.DB.Table("tbl_barang").Find(&barang); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(barang)

}

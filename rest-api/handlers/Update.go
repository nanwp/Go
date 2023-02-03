package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBarang(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBarang models.Barang
	json.Unmarshal(body, &updateBarang)

	var barang models.Barang

	if result := h.DB.First(&barang, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	barang.Nama = updateBarang.Nama
	barang.Harga = updateBarang.Harga
	barang.Satuan = updateBarang.Satuan
	barang.Stock = updateBarang.Stock

	h.DB.Save(&barang)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}

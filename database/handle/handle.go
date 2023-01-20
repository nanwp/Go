package handle

import (
	"database/koneksi"
	"fmt"
	"net/http"
	"strconv"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	name, age := koneksi.Query(idNumb)

	if name == "" {
		fmt.Fprintln(w, "Data Kosong")
		return
	}

	fmt.Fprintf(w, "id : %d \nName : %s \nUmur : %d",idNumb, name, age)

}

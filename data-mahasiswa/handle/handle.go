package handle

import (
	"data-mahasiswa/entity"
	"data-mahasiswa/koneksi"
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "ERROR CUY", http.StatusInternalServerError)
			return
		}

		//var name, nrp string
		
		data, err := koneksi.Connect().Query("SELECT nrp, nama FROM tbl_mahasiswa ORDER BY nrp ASC")

		if err != nil {
			log.Println("DATA KOSONG")
		}

		//defer data.Close()

		kumpulanMahasiswa := make([]*entity.Mahasiswa, 0)

		for data.Next() {
			mahasis := new(entity.Mahasiswa)
			err := data.Scan(&mahasis.NRP, &mahasis.Name)
			if err != nil {
				log.Println("EROR NJIRT")
			}
			kumpulanMahasiswa = append(kumpulanMahasiswa, mahasis)
		}

		err = tmpl.Execute(w, kumpulanMahasiswa)

		if err != nil {
			log.Println(err)
			http.Error(w, "ERROR CUY", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "ERROR BAD REQ", http.StatusBadRequest)
}

func DetailHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "detailmahasiswa.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "ERROR LAGI", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println()
			http.Error(w, "EROR KOSONG", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "ERROR CUY", http.StatusBadRequest)
}





package handle

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"website/entity"
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
			http.Error(w, "error disini ya diks", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, entity.DataMahasiswa)

		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr lagi kan", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happingggg", http.StatusBadRequest)
}

func AddMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "addmahasiswa.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "ERRORR ATUH KAK", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println()
			http.Error(w, "EROR LAGI KAN", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "error mulu dah", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "ERROR LAGI DAH", http.StatusInternalServerError)
			return
		}

		intNim, err := strconv.Atoi(r.Form.Get("nim"))

		newData := entity.Mahasiswa{
			NIM:   intNim,
			Name:  r.Form.Get("name"),
			Email: r.Form.Get("email"),
		}
		entity.DataMahasiswa = append(entity.DataMahasiswa, newData)
		
		tmpl, err := template.ParseFiles(path.Join("views", "resoult.html"), path.Join("views", "layout.html"))

		if err != nil{
			log.Println(err)
			http.Error(w, "AMAN KAH? EROR DONG", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil{
			log.Println(err)
			http.Error(w, "error lagi ajah", http.StatusInternalServerError)
		}

		return
		
	}

	http.Error(w, "ERORRR ANJIRT", http.StatusBadRequest)
}

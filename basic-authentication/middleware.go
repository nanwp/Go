package main

import "net/http"

const USERNAME = "admin"
const PASSWORD = "pw123"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`Ada yang salah`))
		return false
	}
	isValid := (username == USERNAME ) && (password == PASSWORD)
	if !isValid{
		w.Write([]byte(`salah pw`))
		return false 
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request)bool{
	if r.Method != "GET" {
		w.Write([]byte("HANYA GET"))
		return false
	}

	return true
}
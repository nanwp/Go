package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"name"`
	Age      int    `json:"umur"`
}

func main() {

	jsonString := `{"name":"Nanda Wijaya Putra", "umur":20}`
	jsonData := []byte(jsonString)

	var data User

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("data : ",data.Fullname)
	fmt.Println("data :",data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("data 1:",data1["name"])
	fmt.Println("data 1:",data1["umur"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	decodeData := data2.(map[string]interface{})

	fmt.Println("data 2:",decodeData["name"])
	fmt.Println("data 2:",decodeData["umur"])

	object := []User{{"Salman AL", 27}, {"Dany FD", 16}}

	jsonHasil, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString = string(jsonHasil)
	fmt.Println(jsonString)
}

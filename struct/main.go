package main

import (
	"struct/management"
)


func main(){

	user1 := management.User{01, "Nanda Wijaya Putra", "nanda@gmail.com"}
	user2 := management.User{02, "Zilong Man", "zilong@gmail.com"}
	user3 := management.User{03, "Aluchart", "alu@gmail.com"}
	user4 := management.User{04, "Snow Man", "Snow@gmail.com"}

	pahlawan := management.Group{"Grup Pahlawan", user1, []management.User{user2, user3, user4}}

	pahlawan.Display()

}


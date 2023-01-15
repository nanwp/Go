package main

import "fmt"

type User struct{
	ID int
	Name string
	Email string
}

type Group struct{
	GroupName string
	Lead User
	Member []User
}

func (user User) display(){
	fmt.Println("Id User :", user.ID, "\nNama :", user.Name, "\nEmail :", user.Email)
}

func (group Group) display(){
	fmt.Println("Group \t:", group.GroupName, "\nKetua \t:", group.Lead.Name, "\nMember \t:")
	
	for _,members := range group.Member{
		fmt.Println("------->",members.Name)
	}
}


func main(){

	user1 := User{01, "Nanda Wijaya Putra", "nanda@gmail.com"}
	user2 := User{02, "Zilong Man", "zilong@gmail.com"}
	user3 := User{03, "Aluchart", "alu@gmail.com"}
	user4 := User{04, "Snow Man", "Snow@gmail.com"}

	pahlawan := Group{"Grup Pahlawan", user1, []User{user2, user3, user4}}

	pahlawan.display()
	//user1.display()


}


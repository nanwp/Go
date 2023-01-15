package management

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

func (user User) Display(){
	fmt.Println("Id User :", user.ID, "\nNama :", user.Name, "\nEmail :", user.Email)
}

func (group Group) Display(){
	fmt.Println("Group \t:", group.GroupName, "\nKetua \t:", group.Lead.Name, "\nMember \t:")
	
	for _,members := range group.Member{
		fmt.Println("------->",members.Name)
	}
}

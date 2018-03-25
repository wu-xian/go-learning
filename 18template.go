package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Id   int
	Name string
	//RelatedUsers []User
}

func main() {
	fmt.Println("18 template")
	t := template.New("name123")
	t, err := t.ParseFiles("tmpl/template-learning.tmpl")

	if err != nil {
		fmt.Println(err)
	}
	// user3 := User{
	// 	Id:   3,
	// 	Name: "wuxian",
	// 	//RelatedUsers: nil,
	// }

	// user4 := User{
	// 	Id:   4,
	// 	Name: "wuxian",
	// 	//RelatedUsers: nil,
	// }
	user := User{
		Id:   1,
		Name: "wuxian",
		//RelatedUsers: []User{user3, user4},
	}
	if err = t.ExecuteTemplate(os.Stdout, "name123", user); err != nil {
		fmt.Println(err)
	}
}

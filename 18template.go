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
	tmpName := "tmpl/template-learning.tmpl"
	fmt.Println("18 template")
	t, err := template.New(tmpName).ParseFiles("tmpl/template-learning.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	user := User{
		Id:   1,
		Name: "123",
		//RelatedUsers: []User{user3, user4},
	}
	// if err = t.ExecuteTemplate(os.Stdout, tmpName, user); err != nil {
	// 	fmt.Println(err)
	// }

	err = t.Execute(os.Stdout, user)
	if err != nil {
		fmt.Println(err)
	}
}

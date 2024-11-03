package main

import (
	"errors"
	"fmt"
)

type user struct {
	ID   int64
	Name string
	Age  int64
}

var users []user

func main() {

	fmt.Println("hello world")
	s := "Hello"
	fmt.Println(s)

	for index := range 10 {
		_, err := createUser(int64(index+1), fmt.Sprintf("Thiraphat%d", index+1), 21)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, user := range users {
		fmt.Println(user)
	}

	users[0].updateUser("ThiraphatUpdate", 22)
	fmt.Println("User update : ", users[0])
}

func createUser(id int64, name string, age int64) (*user, error) {
	var u user
	for _, v := range users {
		if v.ID == id {
			return nil, errors.New("this user id has already taken")
		}
	}

	u.ID = id
	u.Name = name
	u.Age = age

	users = append(users, u)

	return &u, nil
}

func (u *user) updateUser(name string, age int64) {
	u.Name = name
	u.Age = age
}

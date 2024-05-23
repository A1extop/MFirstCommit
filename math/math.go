package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) UserInfo() error {
	_, err := fmt.Println(u.Name)
	if err != nil {
		return err
	}
	_, err = fmt.Println(u.Age)
	if err != nil {
		return err
	}
	return nil
}

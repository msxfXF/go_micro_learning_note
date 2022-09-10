package main

import "strconv"

type User struct {
	ID   int
	Name string
}

func GetUserList(num int) (u []*User) {
	u = make([]*User, num)
	for i, _ := range u {
		u[i] = &User{
			ID:   i,
			Name: "user_" + strconv.Itoa(i),
		}
	}
	return
}

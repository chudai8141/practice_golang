package main

import "fmt"

type T struct {
	// 変数名と型名が同じ場合，変数名のみで型名は省略することができる．
	User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	// 型名を省略したときのみ，↓の書き方をすることができる．
	// fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)

}

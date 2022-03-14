package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Persist data of user %s in database \n", u.name)
}

func (u user) olderOfAge() bool {
	return u.age >= 18
}

func (u *user) haveABirthday() {
	u.age++
}

func main() {
	user1 := user{"Baby Yoda", 50}
	fmt.Println(user1)
	user1.save()
	fmt.Println("User is older age: ", user1.olderOfAge())
	user1.haveABirthday()
	fmt.Println("User have a birthday: ", user1.age)
}

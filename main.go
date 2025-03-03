package main

import (
	"fmt"

	hm "github.com/arodrigow/go-challenge-v1/hashMap"
)

func main() {
	var test hm.Application
	test.Data = make(map[hm.Id]hm.User)

	var user = hm.User{FirstName: "Andre", LastName: "Alves", Biography: "Estou perdido"}
	var userUpdate = hm.User{FirstName: "Rodrigo", LastName: "Maciel", Biography: "Talvez?"}

	fmt.Println(test.FindAll()) //Testing empty list

	userResp := test.Insert(user)
	fmt.Println(userResp) //Testing return of Insert User

	test.Data = test.FindAll()
	fmt.Println(test.Data) //Testing list with new User

	fmt.Println(test.FindById(userResp.Id)) //Testing finding user through ID

	fmt.Println(test.Update(userResp.Id, userUpdate)) //Testing Update User
	fmt.Println(test.Data)                            //Testing list with updated User

	fmt.Println(test.Delete(userResp.Id)) //Testing Delete User
	fmt.Println(test.Data)                //Testing list with deleted User
}

// func run() error {
// }

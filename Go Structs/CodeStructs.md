# Structs & Pointers

package main
import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM /DD/YYYY): ")
	
	var appUser user 

	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	// ... do something awesome with that gathered data!

	 outputUserDetails(&appUser)
}

func outputUserDetails(u *user){
	fmt.Println(u.firstName, u.lastName , u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
} 

# Structs Final with two files

# structs.go
package main

import (
	"fmt"

	"github.com/dimcodes09/Go-Lang/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM /DD/YYYY): ")
	
	var appUser *user.User 

	appUser , err := user.New(userFirstName, userLastName , userBirthdate)

	if err != nil{
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	// ... do something awesome with that gathered data!

	 appUser.OutputUserDetails()
	 appUser.ClearUserName()
	 appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
} 

# USer code File
package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct{
	email string
	password string
	User 
}

func (u *User) OutputUserDetails(){
	fmt.Println(u.firstName, u.lastName , u.birthdate)
}

func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email , password string) Admin {
	return  Admin{
		email:email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "___",
			createdAt: time.Now(),
		},
	}
}
func New(firstName,lastName, birthdate string) (*User , error) {
	if firstName == "" || lastName == "" || birthdate == ""{
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	},nil
}




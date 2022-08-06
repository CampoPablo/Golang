package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Pablo",
		LastName:    "Campo",
		PhoneNumber: "+54 11 65465140",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate, "Phone number:", user.PhoneNumber)
}

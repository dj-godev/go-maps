// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Name struct {
	FirstName string
	LastName  string
}
type Address struct {
	Line1 string
	Line2 string
	Line3 string
}
type Customer struct {
	Name    Name
	Email   string
	Address Address
	DOB     time.Time
}

func main() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	john := Customer{
		Name: Name{FirstName: "John",
			LastName: "Smith",
		},
		Email: "johnsmith@example.com",
		Address: Address{
			Line1: "The White House",
			Line2: "1600 Pennsylvania Avenue NW",
			Line3: "Washington, DC 20500",
		},
		DOB: dob,
	}
	//johnJSON, err := json.Marshal(john)
	johnJSON, err := json.MarshalIndent(john, " ", " ")
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}

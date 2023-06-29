// You can edit this code!
// Click here and start typing.
//Encoding interfaces to JSON
/* Sometimes you don’t really want to fix the number of fields in your structs.
Instead, you want to be able to add additional data as and when you need to. You
can do that using empty interfaces, as in the following example:
type Customer map[string]interface{}
type Name map[string]interface{}
type Address map[string]interface{} */
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Customer map[string]interface{}
type Name map[string]interface{}
type Address map[string]interface{}

func main() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	john := Customer{
		"Name": Name{
			"FirstName": "John",
			"LastName":  "Smith",
		},
		"Email": "johnsmith@example.com",
		"Address": Address{
			"Line1": "The White House",
			"Line2": "1600 Pennsylvania Avenue NW",
			"Line3": "Washington, DC 20500",
		},
		"DOB": dob,
	}
	johnJSON, err := json.MarshalIndent(john, "", " ")
	if err == nil {
		fmt.Println(string(johnJSON))
	} else {
		fmt.Println(err)
	}
}

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type course struct {
	name  string
	desc  string
	price int
}

type user struct {
	name   string
	email  string
	course course
}

var members map[int]user

func main() {

	members = make(map[int]user)

	members[1] = user{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		course: course{
			name:  "Prelims",
			desc:  "This is prelims desc",
			price: 1990,
		},
	}
	members[2] = user{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		course: course{
			name:  "Mains",
			desc:  "This is Mains desc",
			price: 1990,
		},
	}

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.course)
	}
}

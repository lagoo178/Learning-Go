package main
import (
	"fmt"
	"sort"
)

// Using Structs and Maps in Go

// Creating a map of structs
type dob struct {
	day int
	month int
	year int
}
type people struct {
	name string
	email string
	dob dob
}
var members map[int]people
func main() {
	members = make(map[int]people)
	members[1] = people{
		name: "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
		day: 17,
		month: 3,
		year: 1990,
		},
	}
	members[2] = people{
		name: "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day: 9,
			month: 12,
			year: 1988,
		},
	}
	members[3] = people{
		name: "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day: 1,
			month: 12,
			year: 1988,
		},
	}
	members[4] = people{
		name: "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day: 19,
			month: 8,
			year: 2001,
		},
	}

	// for k, v := range members {
	// 	fmt.Println(k, v.name, v.email, v.dob)
	// }
	
	// Sorting a map of structs

	// get all the keys in members
	var keys []int
	for k := range members {
		keys = append(keys, k)
	}
	// sort the keys in ascending order
	sort.Ints(keys)
	// copy the value in members to the slice
	var sliceMembers []people
	for _, k := range keys {
		sliceMembers = append(sliceMembers, members[k])
	}
	
	// for k, v := range sliceMembers {
	// 	fmt.Println(k, v.name, v.email, v.dob)
	// }
	/*
	[
		{Mary Smith marysmith@example.com {17 3 1990}}
		{John Smith johnsmith@example.com {9 12 1988}}
		{Janet Doe janetdoe@example.com {1 12 1988}}
		{Adam Jones adamjones@example.com {19 8 2001}}
	]
	*/

	// sort all the members in the sliceMembers based on their date-of-birth (dob) from oldest to youngest
	sort.SliceStable(sliceMembers, func(i, j int) bool {
		// compare year
		if sliceMembers[i].dob.year !=
		sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year <
			sliceMembers[j].dob.year
		}
		// compare month
		if sliceMembers[i].dob.month !=
		sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month <
			sliceMembers[j].dob.month
		}
		// compare day and sort based on day
		// return sliceMembers[i].dob.day <
		// sliceMembers[j].dob.day

		// sort based on year
		return sliceMembers[i].dob.year <
 		sliceMembers[j].dob.year
	})
	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}

}
package main

import (
	"fmt"
	"sort"
)

// User is ...
type User struct {
	Name string
	Age  int
}

// UserSlice is ...
type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Firas", 20},
		{"Luthfi", 30},
		{"Dwiyansyah", 40},
		{"Putri", 19},
		{"Indah", 25},
		{"Lestari", 23},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}

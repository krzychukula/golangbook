package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//new type to slice of what I want to sort
type ByName []Person

//implement sort.Interface on this new type
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	//cast the list onto new type
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

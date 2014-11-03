package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Un"])

}

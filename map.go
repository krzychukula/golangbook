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

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("no Un")
	}

	complex := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}

	if el, ok := complex["He"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}

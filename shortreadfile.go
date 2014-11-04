package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("shortreadfile.go")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

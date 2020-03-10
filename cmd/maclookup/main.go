package main

import (
	"fmt"
	"maclookup"
	"os"
)

func main() {
	mac := os.Args[1]
	v, err := maclookup.Run(mac)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

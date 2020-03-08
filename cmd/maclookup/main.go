package main

import (
	"fmt"
	"maclookup"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1]
	mac := strings.ToLower(argsWithoutProg)
	v, err := maclookup.Run(mac)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

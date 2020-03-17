package main

import (
	"fmt"
	"maclookup"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		mac := os.Args[1]
		v, err := maclookup.GetVendorName(mac)
		if err != nil {
			panic(err)
		}
		fmt.Println(v)
	} else {
		fmt.Println("err: please give MAC address")
	}

}

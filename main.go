package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// read the mac
	fmt.Print("mac address: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	mac := strings.ToLower(text)

	// check the mac is of one of the supported formats
	pattern1 := regexp.MustCompile(`(?:[0-9a-f]{2}[-:\.]{1}){5}[0-9a-f]{2}`)
	pattern2 := regexp.MustCompile(`(?:[0-9a-f]{4}[-:\.]{1}){2}[0-9a-f]{4}`)
	pattern3 := regexp.MustCompile(`[0-9a-f]{12}`)
	if !pattern1.MatchString(mac) && !pattern2.MatchString(mac) && !pattern3.MatchString(mac) {
		panic("not a mac address")
	}

	// api call to get the vendor
	url := "https://api.macvendors.com/" + strings.Trim(mac, "\n")
	client := http.DefaultClient
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// and the vendor iiiiiis....
	fmt.Println(string(b))
}

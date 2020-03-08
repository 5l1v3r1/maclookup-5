package maclookup

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const URL = "https://api.macvendors.com/"

func verification(mac string) error {
	// check the mac is of one of the supported formats
	pattern1 := regexp.MustCompile(`(?:[0-9a-f]{2}[-:\.]{1}){5}[0-9a-f]{2}`)
	pattern2 := regexp.MustCompile(`(?:[0-9a-f]{4}[-:\.]{1}){2}[0-9a-f]{4}`)
	pattern3 := regexp.MustCompile(`[0-9a-f]{12}`)
	if !pattern1.MatchString(mac) && !pattern2.MatchString(mac) && !pattern3.MatchString(mac) {
		return errors.New("wrong mac type")
	}
	return nil
}

func request(mac string) (string, error) {
	url := URL + strings.Trim(mac, "\n")
	client := http.DefaultClient
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(b), nil
}

// Run main program function
func Run(mac string) (string, error) {
	err := verification(mac)
	if err != nil {
		return "", err
	}
	vendor, err := request(mac)
	if err != nil {
		return "", err
	}
	return vendor, err
}

package maclookup

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const URL = "https://api.macvendors.com/"

type ErrJSON struct {
	Errors struct {
		Detail string `json:"detail"`
	} `json:"errors"`
}

func normalizing(mac string) string {
	return strings.Trim(strings.ToLower(mac), "\n")
}

func verification(mac string) error {
	// check the mac is of one of the supported formats
	pattern1 := regexp.MustCompile(`(?:[0-9a-f]{2}[-:\.]{1}){5}[0-9a-f]{2}`)
	pattern2 := regexp.MustCompile(`(?:[0-9a-f]{4}[-:\.]{1}){2}[0-9a-f]{4}`)
	pattern3 := regexp.MustCompile(`[0-9a-f]{12}`)
	if !pattern1.MatchString(mac) && !pattern2.MatchString(mac) && !pattern3.MatchString(mac) {
		return errors.New("wrong mac address format")
	}
	return nil
}

func request(mac string) (string, error) {
	url := URL + mac
	client := http.DefaultClient
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 404 {
		var e ErrJSON
		err = json.Unmarshal(b, &e)
		if err != nil {
			return "", errors.New("can't decode JSON")
		}
		return "", errors.New(e.Errors.Detail)
	}
	return string(b), nil
}

// GetVendorName main program function
func GetVendorName(mac string) (string, error) {
	mac = normalizing(mac)
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

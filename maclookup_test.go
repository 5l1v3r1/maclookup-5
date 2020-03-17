package maclookup

import "testing"

const orgName = "XEROX CORPORATION"
const errMAC = "wrong mac address format"
const errJSON = "Not Found"

func TestGetVendorName(t *testing.T) {
	vendor, err := GetVendorName("00:00:00:00:00:00")
	if err != nil {
		t.Error(err)
	}
	if vendor != orgName {
		t.Errorf("Program return wrong Orgname, got:%s, want:%s", vendor, orgName)
	}
}

func TestGetVendorName2(t *testing.T) {
	_, err := GetVendorName("XX:XX:")
	if err.Error() != errMAC {
		t.Errorf("Program not return right error, got:%s, want:%s", err.Error(), errMAC)
	}
}

func TestGetVendorName3(t *testing.T) {
	_, err := GetVendorName("FF:FF:FF:FF:FF:FF")
	if err.Error() != errJSON {
		t.Errorf("Program not return right error, got:%s, want:%s", err.Error(), errJSON)
	}
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Zone struct {
	Account          string `json:"account"`
	APIRectify       bool   `json:"api_rectify"`
	Dnssec           bool   `json:"dnssec"`
	EditedSerial     int    `json:"edited_serial"`
	ID               string `json:"id"`
	Kind             string `json:"kind"`
	LastCheck        int    `json:"last_check"`
	MasterTsigKeyIds []any  `json:"master_tsig_key_ids"`
	Masters          []any  `json:"masters"`
	Name             string `json:"name"`
	NotifiedSerial   int    `json:"notified_serial"`
	Nsec3Narrow      bool   `json:"nsec3narrow"`
	Nsec3Param       string `json:"nsec3param"`
	Rrsets           []struct {
		Comments []any  `json:"comments"`
		Name     string `json:"name"`
		Records  []struct {
			Content  string `json:"content"`
			Disabled bool   `json:"disabled"`
		} `json:"records"`
		TTL  int    `json:"ttl"`
		Type string `json:"type"`
	} `json:"rrsets"`
	Serial          int    `json:"serial"`
	SlaveTsigKeyIds []any  `json:"slave_tsig_key_ids"`
	SoaEdit         string `json:"soa_edit"`
	SoaEditAPI      string `json:"soa_edit_api"`
	URL             string `json:"url"`
}

type Zones []struct {
	Account        string `json:"account"`
	Dnssec         bool   `json:"dnssec"`
	EditedSerial   int    `json:"edited_serial"`
	ID             string `json:"id"`
	Kind           string `json:"kind"`
	LastCheck      int    `json:"last_check"`
	Masters        []any  `json:"masters"`
	Name           string `json:"name"`
	NotifiedSerial int    `json:"notified_serial"`
	Serial         int    `json:"serial"`
	URL            string `json:"url"`
}

var KEY string = "XPS2jM2XX91DTL7PJTzzGM1vv97hwK" // Insert your own PDNS API Key here this is just a sample for local dev environment

func GetZones() Zones {
	var URL string = "http://localhost:8081/api/v1/servers/localhost/zones"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("X-API-Key", KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var result Zones
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func GetZone() {
	var URL string = "http://localhost:8081/api/v1/servers/localhost/zones/emmatest.se."
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("X-API-Key", KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

var URL string = "http://localhost:8081/api/v1/servers/localhost/zones"
var KEY string = "XPS2jM2XX91DTL7PJTzzGM1vv97hwK" // Insert your own PDNS API Key here this is just a sample for local dev environment

func GetZones() Zones {
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

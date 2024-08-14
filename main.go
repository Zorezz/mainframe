package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url string = "http://localhost:8081/api/v1/servers/localhost/zones"
var key string = "XPS2jM2XX91DTL7PJTzzGM1vv97hwK" // Insert your own PDNS API Key here this is just a sample for local dev environment

type zones []struct {
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

func GetZones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", key)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var result zones
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result {
		io.WriteString(w, "<a href='/zones/"+rec.Name+"'>"+rec.Name+"</a>\n")
	}
}

func GetZone(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("GET /", GetZones)
	http.HandleFunc("GET /zones/{name}", GetZone)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("HTTP Error: ", err)
	}
}

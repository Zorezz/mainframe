package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mainframe/handlers"
	"mainframe/views"
	"net/http"
	"strconv"
	"strings"
)

var URL string = "http://localhost:8081/api/v1/servers/localhost/zones"
var KEY string = "XPS2jM2XX91DTL7PJTzzGM1vv97hwK" // Insert your own PDNS API Key here this is just a sample for local dev environment

func ZonesHandler(ctx echo.Context) error {
	domains := handlers.GetZones()

	return views.ZonesView(domains).Render(ctx.Request().Context(), ctx.Response())
}

func ZoneHandler(ctx echo.Context) error {
	domain := ctx.Param("domain")
	records := handlers.GetZone(domain)

	return views.ZoneView(records, domain).Render(ctx.Request().Context(), ctx.Response())
}

func ZoneEditHandler(ctx echo.Context) error {
	domain := ctx.Param("domain")
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	records := handlers.GetZone(domain)

	return views.ZoneEdit(records, intId, domain).Render(ctx.Request().Context(), ctx.Response())
}

func ZonePutHandler(ctx echo.Context) error {
	domain := ctx.Param("domain")

	reqBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println(err)
	}

	parsedKeys := map[string]string{}
	for _, pair := range strings.Split(string(reqBody), "&") {
		kv := strings.Split(pair, "=")
		parsedKeys[kv[0]] = kv[1]
	}

	ttlInt, err := strconv.Atoi(parsedKeys["TTL"])
	if err != nil {
		fmt.Println(err)
	}

	content := handlers.Content{
		handlers.Rrsets{{
			Name:       parsedKeys["name"],
			Type:       parsedKeys["type"],
			TTL:        ttlInt,
			Changetype: "REPLACE",
			Records: handlers.Records{{
				Content:  parsedKeys["content"],
				Disabled: false,
			}},
		}},
	}

	contentJson, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Failed to Marshal JSON")
	}

	var URL string = "http://localhost:8081/api/v1/servers/localhost/zones/" + domain
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", URL, bytes.NewBuffer(contentJson))
	fmt.Printf("%s\n", contentJson)
	req.Header.Set("X-API-Key", KEY)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status: ", resp.Status)
	fmt.Println("response Header: ", resp.Header)

	return nil
}

func RecordCreateHandler(ctx echo.Context) error {
	domain := ctx.Param("domain")

	return views.RecordCreate(domain).Render(ctx.Request().Context(), ctx.Response())
}

func ZoneCreateHandler(ctx echo.Context) error {
	domain := ctx.Param("domain")

	handlers.CreateZoneFunc(domain)

	return nil
}

func main() {
	app := echo.New()

	app.GET("/zones", func(c echo.Context) error {
		return ZonesHandler(c)
	})
	app.GET("/zones/:domain", func(c echo.Context) error {
		return ZoneHandler(c)
	})
	app.GET("/zones/:domain/:id/edit", func(c echo.Context) error {
		return ZoneEditHandler(c)
	})
	app.PUT("/zones/:domain/:id", func(c echo.Context) error {
		return ZonePutHandler(c)
	})
	app.GET("/zones/:domain/record/create", func(c echo.Context) error {
		return RecordCreateHandler(c)
	})
	app.PUT("/zones/:domain/record/create", func(c echo.Context) error {
		return ZonePutHandler(c)
	})
	app.GET("/zones/:domain/create", func(c echo.Context) error {
		return ZoneCreateHandler(c)
	})

	app.Start(":8088")
}

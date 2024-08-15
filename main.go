package main

import (
	"github.com/labstack/echo/v4"
	"mainframe/handlers"
	"mainframe/views"
)

var URL string = "http://localhost:8081/api/v1/servers/localhost/zones"
var KEY string = "XPS2jM2XX91DTL7PJTzzGM1vv97hwK" // Insert your own PDNS API Key here this is just a sample for local dev environment

func ZonesHandler(ctx echo.Context) error {
	domains := handlers.GetZones()

	return views.ZonesView(domains).Render(ctx.Request().Context(), ctx.Response())
}

func ZoneHandler(ctx echo.Context, domainName string) error {
	records := handlers.GetZone(domainName)

	return views.ZoneView(records).Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	app := echo.New()

	app.GET("/zones", func(c echo.Context) error {
		return ZonesHandler(c)
	})
	app.GET("/zones/:domain", func(c echo.Context) error {
		return ZoneHandler(c, "/zones/:domain")
	})

	app.Start(":8088")
}

run:
	@templ generate
	@npm run build-css
	@go run main.go

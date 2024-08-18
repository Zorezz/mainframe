run:
	@set TEMPL_EXPERIMENT=rawgo
	@templ generate
	@npm run build-css
	@go run main.go

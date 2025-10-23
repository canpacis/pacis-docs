module github.com/canpacis/pacis-docs

go 1.24.5

replace github.com/canpacis/pacis => ../pacis

replace components => ./src/components

require (
	components v0.0.0-00010101000000-000000000000
	github.com/canpacis/pacis v0.6.5-0.20251018140617-51d4c3378d40
	github.com/joho/godotenv v1.5.1
	github.com/sivukhin/godjot v1.0.6
)

require (
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/Oudwins/tailwind-merge-go v0.2.1 // indirect
	github.com/alecthomas/chroma/v2 v2.20.0 // indirect
	github.com/canpacis/http-payload v0.3.1 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.6.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

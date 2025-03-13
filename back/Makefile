.PHONY: build 

build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/bootstrap main.go 
	chmod +x bin/bootstrap
	zip -j bin/api.zip bin/bootstrap

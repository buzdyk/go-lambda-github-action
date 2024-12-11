build-and-zip:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o build/bootstrap
	cd build && zip bootstrap.zip bootstrap


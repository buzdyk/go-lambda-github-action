build-hellogo:
	GOOS=linux go build -o bootstrap
	cp ./bootstrap $(ARTIFACTS_DIR)/.

build-and-zip:
	CGO_ENABLED=0 GOOS=linux go build -o bootstrap
	zip bootstrap.zip bootstrap
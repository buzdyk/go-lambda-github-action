build-hellogo:
	GOOS=linux go build -o bootstrap
	cp ./bootstrap $(ARTIFACTS_DIR)/.
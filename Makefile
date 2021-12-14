
.PHONY: test
test:
	go test -coverprofile=coverage.out
	gocov convert coverage.out > coverage.json
	cat coverage.json | gocov-html > coverage.html
	cat coverage.json | gocov-xml > coverage.xml

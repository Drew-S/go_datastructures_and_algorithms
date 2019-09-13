test:
	go test -timeout 10s -coverprofile cover.out ./...
	go tool cover -html=cover.out -o coverage.html
	rm cover.out
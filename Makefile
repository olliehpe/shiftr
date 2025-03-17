run:
	go run ./cmd/shiftr/*.go

test:
	go test -v --cover ./...

release:
	@echo "Enter the release version (format vx.x.x).."; \
	read VERSION; \
	git tag -a $$VERSION -m "Releasing "$$VERSION; \
	git push origin $$VERSION

upload:
	GOOS=linux GOARCH=amd64 go build ./cmd/shiftr
	scp ./shiftr ollie@10.32.20.37:/tmp/shiftr
APP=auth

generate-doc:
	swag init \
		-g cmd/api/main.go \
		--parseDependency \
		--parseInternal

build: generate-doc
	go build -o bin/$(APP) ./cmd/api

run: generate-doc
	go run ./cmd/api/main.go

clean:
	rm -rf bin docs
build:
	@ printf "Mod download.. "
	@ go mod download
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o osauamsapi \
		.cmd/app/
	@ echo "done"
run: go run ./cmd/app/main.go
debug: go run ./cmd/main.go
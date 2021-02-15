compile:
    mkdir -p bin/
    @echo "Compiling for every OS and Platform"
    @echo "🐧 Compile for Linux"
    GOOS=linux GOARCH=amd64 go build -o ./bin/gme-cli-linux-amd64 ./cmd/gme-shortener/main.go
    GOOS=linux GOARCH=386 go build -o ./bin/gme-cli-linux-386 ./cmd/gme-shortener/main.go  
    GOOS=linux GOARCH=arm go build -o ./bin/gme-cli-linux-arm ./cmd/gme-shortener/main.go 
    GOOS=linux GOARCH=arm64 go build -o ./bin/gme-cli-linux-arm64 ./cmd/gme-shortener/main.go
    @echo "🍏 Compile for Apple"
    GOOS=darwin GOARCH=amd64 go build -o ./bin/gme-cli-darwin-amd64 ./cmd/gme-shortener/main.go
    @echo "🪟 Compile for Windows"
    GOOS=windows GOARCH=amd64 go build -o ./bin/gme-cli-windows-amd64 ./cmd/gme-shortener/main.go
    GOOS=windows GOARCH=386 go build -o ./bin/gme-cli-windows-386 ./cmd/gme-shortener/main.go
    @echo "🐡 Compile for FreeBSD"
    GOOS=freebsd GOARCH=amd64 go build -o ./bin/gme-cli-freebsd-amd64 ./cmd/gme-shortener/main.go
    GOOS=freebsd GOARCH=386 go build -o ./bin/gme-cli-freebsd-386 ./cmd/gme-shortener/main.go
    GOOS=freebsd GOARCH=arm go build -o ./bin/gme-cli-freebsd-arm ./cmd/gme-shortener/main.go
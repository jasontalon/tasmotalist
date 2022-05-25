build/mac/m1:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/mac/m1/tasmotalist

build/win/amd64:
	GOOS=windows GOARCH=amd64 go build -o ./bin/win/amd64/tasmotalist.exe

build/win/arm64:
	GOOS=windows GOARCH=arm64 go build -o ./bin/win/arm64/tasmotalist.exe

build:
	rimraf bin
	make build/mac/m1
	make build/win/amd64
	make build/win/arm64
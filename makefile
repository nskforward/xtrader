compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./build/xtrader_amd64 ./cmd/*

run:
	go build -o ./build/xtrader_darwin ./cmd/* && ./build/xtrader_darwin
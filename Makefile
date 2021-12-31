all:
	go build

armv6:
	GOOS=linux GOARCH=arm GOARM=6 go build

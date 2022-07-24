package main

import (
	"fmt"
	"io"
)

func stream(address, username, password string, out io.Writer) {
	url := getURL(
		address,
		fmt.Sprintf("flv?port=1935&app=bcs&stream=channel0_main.bcs&user=%s&password=%s", username, password),
	)

	getRequest(url, out)
}

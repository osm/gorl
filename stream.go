package main

import (
	"fmt"
	"io"
)

func stream(address, username, password string, out io.Writer) {
	token, err := getToken(address, username, password)
	if err != nil {
		exitf("%v\n", err)
	}

	url := getURL(
		address,
		fmt.Sprintf("flv?port=1935&app=bcs&stream=channel0_main.bcs&token=%s", token),
	)

	getRequest(url, out)
}

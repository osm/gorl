package main

import (
	"fmt"
	"io"
)

func dl(address, username, password, file string, out io.Writer) {
	token, err := getToken(address, username, password)
	if err != nil {
		exitf("%v\n", err)
	}

	url := getURL(
		address,
		fmt.Sprintf("/cgi-bin/api.cgi?cmd=Download&source=%s&output=%s&token=%s",
			file, file, token,
		),
	)
	getRequest(url, out)
}

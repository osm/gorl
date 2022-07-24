package main

import (
	"fmt"
	"io"
)

func dl(address, username, password, file string, out io.Writer) {
	url := getURL(
		address,
		fmt.Sprintf("/cgi-bin/api.cgi?cmd=Download&source=%s&output=%s&user=%s&password=%s",
			file, file, username, password,
		),
	)
	getRequest(url, out)
}

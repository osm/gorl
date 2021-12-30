package main

import (
	"fmt"
	"io"
)

func snap(address, username, password string, out io.Writer) {
	token, err := getToken(address, username, password)
	if err != nil {
		exitf("%v\n", err)
	}

	url := getURL(address, fmt.Sprintf("/cgi-bin/api.cgi?cmd=Snap&token=%s", token))
	getRequest(url, out)
}

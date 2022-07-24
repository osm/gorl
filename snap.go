package main

import (
	"fmt"
	"io"
)

func snap(address, username, password string, out io.Writer) {
	url := getURL(address, fmt.Sprintf("/cgi-bin/api.cgi?cmd=Snap&user=%s&password=%s", username, password))
	getRequest(url, out)
}

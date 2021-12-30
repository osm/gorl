package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

var cacheFile string

func init() {
	dir, _ := os.UserHomeDir()
	cacheFile = filepath.Join(dir, ".gorl")
}

type Cache struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func getToken(address, username, password string) (string, error) {
	cacheToken, _ := getCachedToken()
	if cacheToken != "" {
		return cacheToken, nil
	}

	reqBody, _ := json.Marshal([]Request{
		{
			Command: "Login",
			Action:  0,
			Param: Param{
				User: &User{
					Username: username,
					Password: password,
				},
			},
		},
	})

	resp, err := postRequest(
		getURL(address, "/cgi-bin/api.cgi?cmd=Login"),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}

	token := resp.Value.Token.Name
	leaseTime := resp.Value.Token.LeaseTime
	cache, _ := json.Marshal(Cache{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Second*leaseTime - 1*time.Minute),
	})
	if err := os.WriteFile(cacheFile, cache, 0600); err != nil {
		return "", err
	}

	return token, nil
}

func getCachedToken() (string, error) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", err
	}

	var cache Cache
	if err := json.Unmarshal(data, &cache); err != nil {
		return "", err
	}

	if time.Now().After(cache.ExpiresAt) {
		return "", errors.New("token has expired")
	}

	return cache.Token, nil
}

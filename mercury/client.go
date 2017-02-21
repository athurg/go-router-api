package router

import (
	"fmt"
	"net"
)

type Client struct {
	Username string
	Password string
	baseUri  string
}

const (
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

func New(username, password, host string, port int) (*Client, error) {
	client := Client{Username: username, Password: password}

	if username == "" {
		client.Username = DefaultUsername
	}

	if password == "" {
		client.Password = DefaultPassword
	}

	ips, err := net.LookupIP(host)
	if err != nil || len(ips) == 0 {
		return nil, fmt.Errorf("主机解析失败:%s", err)
	}

	client.baseUri = fmt.Sprintf("http://%s:%d", ips[0], port)

	return &client, nil
}

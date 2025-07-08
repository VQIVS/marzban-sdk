package client

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient   *http.Client
	BaseURL      string
	Token        string
	ClientID     string
	ClientSecret string
}

type ClientOption func(*Client)

func NewClient(baseURL string, options ...ClientOption) *Client {
	client := &Client{
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
		},
		BaseURL: baseURL,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HttpClient.Timeout = timeout
	}
}

func WithClientID(clientID string, clientSecret string) ClientOption {
	return func(c *Client) {
		c.ClientID = clientID
		c.ClientSecret = clientSecret
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HttpClient = httpClient
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.Token = token
	}
}

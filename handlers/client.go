package handlers

import (
	"github.com/VQIVS/marzban-sdk/internal/client"
)

type MarzbanClient struct {
	*client.Client
}

func NewMarzbanClient(baseURL string, options ...client.ClientOption) *MarzbanClient {
	c := client.NewClient(baseURL, options...)
	return &MarzbanClient{Client: c}
}

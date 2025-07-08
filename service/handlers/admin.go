package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/VQIVS/marzban-sdk/internal/client"
	"github.com/VQIVS/marzban-sdk/internal/models"
)

func (mc *MarzbanClient) LoginWithUsername(req models.UserLoginReq) (*models.UserLoginResponse, error) {
	var reqBody models.UserLoginReq
	reqBody.Username = req.Username
	reqBody.Password = req.Password

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := mc.Client.HttpClient.Post(mc.Client.BaseURL+client.EndpointAdminToken, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to login, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	var loginResponse models.UserLoginResponse
	if err := json.Unmarshal(responseBody, &loginResponse); err != nil {
		return nil, err
	}
	mc.Client.Token = loginResponse.Token
	return &loginResponse, nil

}

func (mc *MarzbanClient) LoginWithClientID(clientID, clientSecret string) (*models.UserLoginResponse, error) {
	var reqBody models.UserLoginReq
	reqBody.Username = clientID
	reqBody.Password = clientSecret

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := mc.Client.HttpClient.Post(mc.Client.BaseURL+client.EndpointAdminToken, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to login, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	var loginResponse models.UserLoginResponse
	if err := json.Unmarshal(responseBody, &loginResponse); err != nil {
		return nil, err
	}
	mc.Client.Token = loginResponse.Token
	return &loginResponse, nil
}

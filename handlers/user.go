package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/VQIVS/marzban-sdk/internal/client"
	"github.com/VQIVS/marzban-sdk/internal/models"
	"github.com/VQIVS/marzban-sdk/utils"
)

func (mc *MarzbanClient) CreateUser(user models.User) (*models.User, error) {
	reqBody := &user

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := mc.Client.HttpClient.Post(
		mc.Client.BaseURL+client.EndpointUser,
		"application/json",
		bytes.NewBuffer(reqBodyBytes),
	)
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
			Detail:  "Failed to create user, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	var createdUser models.User
	if err := json.Unmarshal(responseBody, &createdUser); err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (mc *MarzbanClient) GetUserByUsername(username string) (*models.User, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	resp, err := mc.Client.HttpClient.Get(mc.Client.BaseURL + endpoint)
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
			Detail:  "Failed to get user by username, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	var user models.User
	if err := json.Unmarshal(responseBody, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// TODO: Double function check for improvements
func (mc *MarzbanClient) UpdateUser(user models.User) (*models.User, error) {

	reqBody := &user

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	endpoint := client.GetUserByUsernameEndpoint(user.Username)

	//TODO: Use StringToURL function in utils/common.go
	fullURL, err := url.Parse(mc.Client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}

	resp, err := mc.Client.HttpClient.Do(&http.Request{
		Method: http.MethodPut,
		URL:    fullURL,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBuffer(reqBodyBytes)),
	})

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
			Detail:  "Failed to update user, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	var updatedUser models.User
	if err := json.Unmarshal(responseBody, &updatedUser); err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (mc *MarzbanClient) DeleteUserByUsername(username string) error {
	endpoint := client.GetUserByUsernameEndpoint(username)

	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, fullURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := mc.Client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to delete user, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	return nil
}
func (mc *MarzbanClient) GetUserSubURL(username string) (string, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return "", err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to get user sub URL, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		SubURL string `json:"subscription_url"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", err
	}
	if response.SubURL == "" {
		return "", &models.ErrorResponse{
			Message: "No subscription URL found",
			Detail:  "The user does not have a subscription URL.",
		}
	}
	return response.SubURL, nil
}

func (mc *MarzbanClient) GetUserInbounds(username string) ([]string, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
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
			Detail:  "Failed to get user inbounds, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		Inbounds []string `json:"inbounds"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response.Inbounds, nil
}

func (mc *MarzbanClient) GetUserProxies(username string) ([]string, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
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
			Detail:  "Failed to get user proxies, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		Proxies []string `json:"proxies"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response.Proxies, nil
}
func (mc *MarzbanClient) GetUserUsage(username string) (int64, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return 0, err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to get user usage, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		Usage int64 `json:"usage"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return 0, err
	}

	return response.Usage, nil
}
func (mc *MarzbanClient) GetUserStatus(username string) (string, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return "", err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to get user status, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", err
	}

	return response.Status, nil
}

func (mc *MarzbanClient) GetUserExpire(username string) (int64, error) {
	endpoint := client.GetUserByUsernameEndpoint(username)
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return 0, err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to get user expire time, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var response struct {
		Expire int64 `json:"expire"`
	}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return 0, err
	}

	return response.Expire, nil
}

func (mc *MarzbanClient) ResetUserUsage(username string) error {
	endpoint := client.GetUserByUsernameEndpoint(username) + client.EndpointUserReset

	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fullURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := mc.Client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to reset user usage, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	return nil
}

//TODO: fix user sub revokes

func (mc *MarzbanClient) RevokeUserSub(username string) error {
	endpoint := client.GetUserByUsernameEndpoint(username) + client.EndpointUserRevokeSubsription

	fullURL, err := utils.StringToURL(mc.Client.BaseURL + endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fullURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := mc.Client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return &models.ErrorResponse{
			Message: "HTTP " + resp.Status,
			Detail:  "Failed to revoke user subscription, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}
	return nil
}

func (mc *MarzbanClient) GetExpiredUsers() ([]models.User, error) {
	fullURL, err := utils.StringToURL(mc.Client.BaseURL + client.EndpointUsersExpired)
	if err != nil {
		return nil, err
	}
	resp, err := mc.Client.HttpClient.Get(fullURL.String())
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
			Detail:  "Failed to get expired users, status code: " + resp.Status + ", body: " + string(responseBody),
		}
	}

	var users []models.User
	if err := json.Unmarshal(responseBody, &users); err != nil {
		return nil, err
	}

	return users, nil
}

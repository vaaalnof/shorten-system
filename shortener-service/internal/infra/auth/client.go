package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(
	baseURL string,
	timeout time.Duration,
) *Client {

	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ValidateToken(
	ctx context.Context,
	token string,
) (
	userID string,
	sessionID string,
	emailVerified string,
	err error,
) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+"/api/v1/validate-token",
		nil,
	)

	if err != nil {
		return "", "", "", err
	}

	req.Header.Set(
		"Authorization",
		token,
	)

	resp, err := c.client.Do(
		req,
	)

	if err != nil {
		return "", "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", errors.New(
			"invalid token",
		)
	}

	var result struct {
		Message string `json:"message"`

		Data struct {
			UserID        string `json:"user_id"`
			SessionID     string `json:"session_id"`
			EmailVerified string `json:"email_verified"`
		} `json:"data"`
	}

	if err := json.NewDecoder(
		resp.Body,
	).Decode(
		&result,
	); err != nil {

		return "", "", "", err
	}

	return result.Data.UserID,
		result.Data.SessionID,
		result.Data.EmailVerified,
		nil
}

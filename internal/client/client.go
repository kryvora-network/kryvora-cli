package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type StatusResponse struct {
	Status    string    `json:"status"`
	Version   string    `json:"version"`
	UptimeSec int64     `json:"uptime_sec"`
	Peers     int       `json:"peers"`
	SyncState string    `json:"sync_state"`
	Timestamp time.Time `json:"timestamp"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) *Client {
	if baseURL == "" {
		baseURL = "http://127.0.0.1:4177"
	}
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) FetchStatus() (*StatusResponse, error) {
	url := fmt.Sprintf("%s/status", c.baseURL)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("connect to node daemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	var status StatusResponse
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("unmarshal status json: %w", err)
	}

	return &status, nil
}

func (c *Client) FetchHealth() (bool, error) {
	url := fmt.Sprintf("%s/health", c.baseURL)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return false, fmt.Errorf("connect to node daemon: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientFetchHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	c := New(server.URL)
	ok, err := c.FetchHealth()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected health true")
	}
}

func TestClientFetchStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/status" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"active","version":"0.2.0","uptime_sec":120,"peers":4,"sync_state":"synced"}`))
	}))
	defer server.Close()

	c := New(server.URL)
	status, err := c.FetchStatus()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status.Status != "active" {
		t.Errorf("expected active status, got: %s", status.Status)
	}
	if status.Peers != 4 {
		t.Errorf("expected 4 peers, got: %d", status.Peers)
	}
}

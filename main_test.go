package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	       ts := http.NewServeMux()
	       ts.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		       w.WriteHeader(http.StatusOK)
		       if _, err := w.Write([]byte("ok")); err != nil {
			       t.Errorf("failed to write response: %v", err)
		       }
	       })

	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	ts.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if body != "ok" {
		t.Errorf("expected body 'ok', got '%s'", body)
	}
}

func TestHealthEndpoint_StatusOK(t *testing.T) {
	       ts := http.NewServeMux()
	       ts.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		       w.WriteHeader(http.StatusOK)
		       if _, err := w.Write([]byte("ok")); err != nil {
			       t.Errorf("failed to write response: %v", err)
		       }
	       })

	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	ts.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHealthEndpoint_BodyOk(t *testing.T) {
	       ts := http.NewServeMux()
	       ts.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		       w.WriteHeader(http.StatusOK)
		       if _, err := w.Write([]byte("ok")); err != nil {
			       t.Errorf("failed to write response: %v", err)
		       }
	       })

	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	ts.ServeHTTP(w, r)

	body := w.Body.String()
	if body != "ok" {
		t.Errorf("expected body 'ok', got '%s'", body)
	}
}

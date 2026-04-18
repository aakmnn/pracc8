package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRate_AllCases(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"rate":1.5}`))
		}))
		defer server.Close()

		s := NewExchangeService(server.URL)
		r, err := s.GetRate()

		if err != nil || r != 1.5 {
			t.Errorf("fail")
		}
	})

	t.Run("api error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(400)
			w.Write([]byte(`{"error":"bad"}`))
		}))
		defer server.Close()

		s := NewExchangeService(server.URL)
		_, err := s.GetRate()

		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("bad json", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`invalid`))
		}))
		defer server.Close()

		s := NewExchangeService(server.URL)
		_, err := s.GetRate()

		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("empty body", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer server.Close()

		s := NewExchangeService(server.URL)
		_, err := s.GetRate()

		if err == nil {
			t.Errorf("expected error")
		}
	})
}

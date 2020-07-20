package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	rootHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	t.Run(fmt.Sprintf("returns %v", http.StatusOK), func(t *testing.T) {
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected: %v; Got: %v", http.StatusOK, res.StatusCode)
		}
	})

	t.Run(fmt.Sprintf("returns '%s'", "application/json"), func(t *testing.T) {
		if expected, got := "application/json", res.Header.Get("Content-Type"); got != expected {
			t.Errorf("Expected: %v; Got: %v", expected, got)
		}
	})

	t.Run("returns expected JSON", func(t *testing.T) {
		expected := Message{"hello"}
		var got Message
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("could not read json: %v", err)
		}

		if got.Message != expected.Message {
			t.Errorf("Expected: {'message': '%s'}; Got: {'message': '%s'}", expected.Message, got.Message)
		}
	})
}

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    bool
	}{
		{"when true", "foo", true},
		{"when false", "bar", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := handler(tt.given)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			h.ServeHTTP(w, r)

			b, err := ioutil.ReadAll(w.Body)
			if err != nil {
				t.Fatal(err)
			}

			if s := string(bytes.TrimSpace(b)); s != tt.expected {
				t.Errorf("got %v, want %v", s, tt.expected)
			}
		})
	}
}

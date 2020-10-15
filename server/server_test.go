package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_WelcomeHandler(t *testing.T) {
	tests := []struct {
		name       string
		wantStatus int
		want       *bytes.Buffer
	}{
		{
			name:       "ok",
			wantStatus: 200,
			want:       bytes.NewBufferString("ok"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()

			WelcomeHandler(rec, req)

			buff := new(bytes.Buffer)
			if _, err := buff.ReadFrom(rec.Body); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.wantStatus, rec.Code)
			assert.Equal(t, tt.want, buff)
		})
	}
}

func TestHandler_Middleware(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	m := Middleware(http.HandlerFunc(WelcomeHandler))
	m.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	buff := new(bytes.Buffer)
	if _, err := buff.ReadFrom(rec.Body); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, bytes.NewBufferString("ok"), buff)

	k := "Middleware"
	assert.Equal(t, "Middleware", rec.Header().Get(k))
}

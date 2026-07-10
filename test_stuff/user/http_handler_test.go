package user

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("What a boring life"))
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(body)
}

func TestHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	Handler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "What a boring life", rr.Body.String())
}

func TestEchoHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	// reqBody := bytes.NewBufferString("This is the request body")
	reqBody := bytes.NewBuffer([]byte("This is the request body"))
	r := httptest.NewRequest("GET", "/", reqBody)

	EchoHandler(rr, r)
	assert.Equal(t, "This is the request body", rr.Body.String())
}
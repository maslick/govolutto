package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maslick/govolutto/src"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = src.SetupRouter(src.CreateService())
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealthEndpoint(t *testing.T) {
	w := performRequest(router, "GET", "/health", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckBalance(t *testing.T) {
	body := gin.H{
		"username": "daisy",
		"balance":  "100",
	}
	w := performRequest(router, "GET", "/v1/daisy/balance", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	usernameVal, usernameExists := response["username"]
	balanceVal, balanceExists := response["balance"]

	assert.Nil(t, err)
	assert.True(t, usernameExists)
	assert.True(t, balanceExists)
	assert.Equal(t, body["username"], usernameVal)
	assert.Equal(t, body["balance"], balanceVal)
}

func TestTransactionEndpoint(t *testing.T) {
	reqBody := gin.H{
		"from":   "scrooge",
		"to":     "daisy",
		"amount": 10000,
	}

	respBody := gin.H{
		"from":    "scrooge",
		"to":      "daisy",
		"amount":  "10000",
		"success": "true",
	}

	gson, ok := json.Marshal(reqBody)
	if ok != nil {
		panic("could not serialize body")
	}
	var buf = bytes.Buffer{}
	buf.Write(gson)

	w := performRequest(router, "POST", "/v1/transfer", &buf)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	if err != nil {
		panic("could not deserialize body")
	}
	fromVal, fromExists := response["from"]
	toVal, toExists := response["to"]
	amountVal, amountExists := response["amount"]
	successVal, successExists := response["success"]

	assert.Nil(t, err)
	assert.True(t, fromExists)
	assert.True(t, toExists)
	assert.True(t, amountExists)
	assert.True(t, successExists)
	assert.Equal(t, respBody["from"], fromVal)
	assert.Equal(t, respBody["to"], toVal)
	assert.Equal(t, respBody["amount"], amountVal)
	assert.Equal(t, respBody["success"], successVal)

	w = performRequest(router, "GET", "/v1/daisy/balance", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	daisyBalance, _ := response["balance"]

	assert.Equal(t, "10100", daisyBalance)
}

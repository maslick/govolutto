package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maslick/govolutto/src"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
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

	w := performRequest(router, "POST", "/v1/transfer", str2buf(reqBody))
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

func TestBadRequesDuringTransaction(t *testing.T) {
	w := performRequest(router, "POST", "/v1/transfer", &bytes.Buffer{})
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConcurrentTransactions(t *testing.T) {
	var wg sync.WaitGroup
	var count = 1000
	var amount = decimal.NewFromFloat(float64(10000.0 / float32(count)))
	wg.Add(count)

	reqBody := gin.H{
		"from":   "gyro",
		"to":     "donald",
		"amount": amount,
	}

	for i := 0; i < count; i++ {
		go func() {
			w := performRequest(router, "POST", "/v1/transfer", str2buf(reqBody))
			assert.Equal(t, http.StatusOK, w.Code)
			wg.Done()
		}()
	}

	wg.Wait()
	w1 := performRequest(router, "GET", "/v1/gyro/balance", nil)
	w2 := performRequest(router, "GET", "/v1/donald/balance", nil)

	var response1 map[string]string
	var response2 map[string]string

	_ = json.Unmarshal([]byte(w1.Body.String()), &response1)
	_ = json.Unmarshal([]byte(w2.Body.String()), &response2)

	balanceFrom, _ := response1["balance"]
	balanceTo, _ := response2["balance"]

	assert.Equal(t, "0", balanceFrom)
	assert.Equal(t, "10000", balanceTo)
}

func str2buf(reqBody gin.H) *bytes.Buffer {
	gson, ok := json.Marshal(reqBody)
	if ok != nil {
		panic("could not serialize body")
	}
	var buf = bytes.Buffer{}
	buf.Write(gson)
	return &buf
}

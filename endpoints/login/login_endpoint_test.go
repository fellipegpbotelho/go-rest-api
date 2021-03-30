package login

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/login", LoginEndpoint)

	t.Run("return status code 200 if login ok", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		body, _ := json.Marshal(gin.H{
			"email":    "clark@dc.com",
			"password": "iamsuperman",
		})
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
	})

	t.Run("return token if login ok", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		body, _ := json.Marshal(gin.H{
			"email":    "clark@dc.com",
			"password": "iamsuperman",
		})
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		router.ServeHTTP(responseRecorder, request)

		var got gin.H
		json.Unmarshal(responseRecorder.Body.Bytes(), &got)
		assert.NotNil(t, got["token"])
	})

	t.Run("return status code 401 if credentials do not exists", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		body, _ := json.Marshal(gin.H{
			"email":    "clark@dc.com",
			"password": "wrong_password",
		})
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		router.ServeHTTP(responseRecorder, request)
		assert.Equal(t, http.StatusUnauthorized, responseRecorder.Code)
	})

	t.Run("return unauthorized message if credentials do not exists", func(t *testing.T) {
		responseRecorder := httptest.NewRecorder()
		body, _ := json.Marshal(gin.H{
			"email":    "clark@dc.com",
			"password": "wrong_password",
		})
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		router.ServeHTTP(responseRecorder, request)

		var got gin.H
		json.Unmarshal(responseRecorder.Body.Bytes(), &got)
		assert.Equal(t, got["message"], "Unauthorized")
	})
}

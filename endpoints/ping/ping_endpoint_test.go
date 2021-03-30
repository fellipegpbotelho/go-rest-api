package ping

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpointReturnStatusCode200(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)
	PingEndpoint(context)
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "status code should be 200")
}

func TestPingEndpoint(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)
	PingEndpoint(context)

	var got gin.H
	json.Unmarshal(responseRecorder.Body.Bytes(), &got)

	want := gin.H{
		"message": "pong",
	}
	assert.Equal(t, want, got)
}

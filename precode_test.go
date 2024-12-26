package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4

    req := httptest.NewRequest("GET", "/cafe?count=11&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusOK, responseRecorder.Code, "The response code is not 200 OK")

    body := responseRecorder.Body.String()
    require.NotEmpty(t, body, "The response body is empty")

    list := strings.Split(body, ",")
    assert.Len(t, list, totalCount, "The answer does not contain all cities")

}

func TestMainHandlerWhenBodyEmpty(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusOK, responseRecorder.Code, "The response code is not 200 OK")

    body := responseRecorder.Body.String()
    assert.NotEmpty(t, body, "The response body is empty")
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=4&city=london", nil)
    responseRecorder := httptest.NewRecorder()

    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "The response code is not 400 Bad Request")

    assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

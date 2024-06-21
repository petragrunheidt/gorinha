package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGET(t *testing.T) {
	t.Run("returns OK", func(t *testing.T) {
		routes := SetupRoutes()

		request, _ := http.NewRequest(http.MethodGet, "/transacoes", nil)
		response := httptest.NewRecorder()

		routes.ServeHTTP(response, request)

		result := response.Result()
    defer result.Body.Close()

    got := result.StatusCode
    want := http.StatusOK

		if got != want {
			t.Errorf("got status %d, want %d", got, want)
	}
	})
}
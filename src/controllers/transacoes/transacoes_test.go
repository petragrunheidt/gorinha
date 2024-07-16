package transacoes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"gorinha/src/helpers"
	"github.com/gin-gonic/gin"
)

func TestGET(t *testing.T) {
	t.Run("returns OK", func(t *testing.T) {
		helpers.InitTestDB()

		router := gin.Default()
		router.GET("/balances/1", Get)

		req, _ := http.NewRequest(http.MethodGet, "/balances/1", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
}

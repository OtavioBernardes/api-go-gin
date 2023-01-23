package StudentController

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()
	return routes
}

func Test_provider_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		wantStatusCode string
	}{
		{
			name:           "Success to get all students",
			wantStatusCode: "200",
		},
	}
	studentController := New(*database.New())
	r := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r.GET("/student", studentController.GetAll)

			req, _ := http.NewRequest("GET", "/student", nil)
			response := httptest.NewRecorder()
			r.ServeHTTP(response, req)
			if response.Code != http.StatusAccepted {
				t.Fatalf("Invalid Status Code - Expected: %v | Received: %v ", tt.wantStatusCode, response.Code)
			}
		})
	}
}

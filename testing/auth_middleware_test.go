
package testing

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "your_project/middleware" 
)


func TestAuthMiddleware(t *testing.T) {
    req := httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer valid_token") 
    
    w := httptest.NewRecorder()

   
    middleware.AuthMiddleware(w, req, func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK) 
    })

    if w.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
    }
}

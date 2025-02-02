// user_handler_test.go
package testing

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "your_project/web/handlers" 
)


func TestCreateUserHandler(t *testing.T) {
    
    reqBody := []byte(`{"name":"John Doe", "email":"john@example.com"}`)
    req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    
    w := httptest.NewRecorder()

    
    handlers.CreateUserHandler(w, req)

    
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
    }
}

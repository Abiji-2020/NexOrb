
package testing

import (
    "testing"
    "your_project/database" 
)


func TestCreateUser(t *testing.T) {
    user, err := database.CreateUser("John Doe", "john@example.com")
    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
    if user.Name != "John Doe" || user.Email != "john@example.com" {
        t.Errorf("Expected user with Name: 'John Doe' and Email: 'john@example.com', got: %+v", user)
    }
}

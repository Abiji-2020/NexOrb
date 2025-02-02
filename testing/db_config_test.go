// db_config_test.go
package testing

import (
    "testing"
    "your_project/config" 
)


func TestDBConnection(t *testing.T) {
    err := config.SetupDB() 
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }
}


package testing

import (
    "your_project/models"
)


func CreateMockUser() models.User {
    return models.User{
        Name:  "John Doe",
        Email: "john@example.com",
    }
}

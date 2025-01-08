package mock 

import(
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"github.com/Abiji-2020/NexOrb/pkg/models"
	"github.com/Abiji-2020/NexOrb/database"
)

func InitTestDB() *database.Database {
	db, err := gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return &database.Database{Instance: db}
}

func MigrateTestDB(db *gorm.DB){
	err:= db.AutoMigrate(&models.ApiKeys{}, &models.Users{}, &models.Organizations{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
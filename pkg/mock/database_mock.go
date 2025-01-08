package mock 

import(
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/Abiji-2020/NexOrb/pkg/models"
)

func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}

func MigrateTestDB(db *gorm.DB){
	err:= db.AutoMigrate(&models.ApiKeys{}, &models.Users{}, &models.Organizations{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
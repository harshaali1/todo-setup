package database

import (
	"fmt"

	"tms-backend/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Supabase connection string
	dsn := `host=localhost
	        user=test1
			password=password
			port=5432
			sslmode=disable
			TimeZone=Asia/Shanghai`
	//dsn := "user=postgres password=BflfYfDOrV/+51wmM357RiQt3cDzwsufFuY3Al63bPQXUMD6oyn2WH7pnu4vVIAuUNJvvtBs7dE7VkOyVM545g== host=db.xmhltpudbjzuxyttuvho.supabase.co port=5432 dbname=postgres sslmode=verify-full"

	config := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(config), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	return nil
}

func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database connection not initialized")
	}
	return DB.AutoMigrate(&models.User{}, &models.Todo{})
}

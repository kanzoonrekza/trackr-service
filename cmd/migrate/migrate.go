package main

import (
	"fmt"
	"log"
	"trackr-service/internal/initialize"
	"trackr-service/internal/models"
)

func init() {
	initialize.EnvironmentVariables()
	initialize.ConnectDatabase()
}

func main() {
	initialize.DB.AutoMigrate(&models.Trackr{}, &models.User{})

	log.Println("Database migrated")
}

func DropColumn(table string, column string) {
	query := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", table, column)
	if err := initialize.DB.Exec(query).Error; err != nil {
		log.Fatalf("Failed to drop column: %v", err)
	}
	log.Printf("Column '%v' dropped successfully", column)
}

func DropTable(table string) {
	query := fmt.Sprintf("DROP TABLE %s", table)
	if err := initialize.DB.Exec(query).Error; err != nil {
		log.Fatalf("Failed to drop column: %v", err)
	}
	log.Printf("Table '%v' dropped successfully", table)
}

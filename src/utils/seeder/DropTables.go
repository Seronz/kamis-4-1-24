package seeder

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func DropAllTable(db *gorm.DB) error {
	var tables []string
	result := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';").Scan(&tables)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	for _, table := range tables {
		result := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", table))
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
	return nil
}

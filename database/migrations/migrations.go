package migrations

import (
	"fmt"
	"go-quiz/database"
	"go-quiz/models/entity"
	"log"
)

func RunMigration() {
	db := database.DatabaseInit()

	//db.Migrator().DropTable(&entity.User{}, &entity.Question{}, &entity.Answer{})
	//fmt.Println("Database Freshed")

	err := db.AutoMigrate(&entity.User{}, &entity.Question{}, &entity.Answer{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}

package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"timertracker/internel/config"
	"timertracker/internel/repository"
	"timertracker/internel/server"
	"timertracker/internel/service"
	"timertracker/internel/utils"
)

func runMigrations(c config.Config, db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("error migrating '%T' reason: %s", model, err)
		}
		if c.Debug {
			log.Println(fmt.Sprintf("DEBUG: Migration model to database %s", model))
		}
	}
	return nil
}

// @title TimerTracker API
// @version 1.0
// @description Это API для отслеживания времени работы с использованием таймеров
// @host localhost:8083
// @BasePath /v1
func main() {
	//КОНФИГУРАЦИЯ
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file %s", err))
	}

	c := config.NewConfig()

	//БАЗА ДАННЫХ
	db, err := gorm.Open(postgres.Open(c.DbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Error connect to database reason: %s", err))
	}
	if c.Debug {
		log.Println(fmt.Sprintf("DEBUG: Connect to database %s", err))
	}

	//МИГРАЦИЯ
	models := []interface{}{
		&repository.User{},
		&repository.Task{},
		&repository.Work{},
	}

	if err := runMigrations(*c, db, models...); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	//ВАЛИДАТОР
	valider := validator.New()
	valider.RegisterValidation("passport", utils.ValidatePassport)

	//СЕРВЕС ЛОГИКА
	serviceContext := service.NewServiceContext(*c, db, valider)

	//ИНИЦИАЛИЗАЦИЯ ФАЙБЕР
	srv := server.NewServer(*serviceContext)
	err = srv.StartWebServer()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error start server %s", err))
	}

}

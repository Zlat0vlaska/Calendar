package main

import (
	"github.com/Zlat0vlaska/Calendar/internal/handler"
	"github.com/Zlat0vlaska/Calendar/internal/repository"
	"github.com/Zlat0vlaska/Calendar/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	repo := repository.NewPostgresEventRepository(db)
	svc := service.NewEventService(repo)
	h := handler.NewEventHandler(svc)

	r := gin.Default()

	// Настройте маршруты
	r.POST("/events", h.CreateEvent)
	r.PUT("/events/:id", h.UpdateEvent)
	r.DELETE("/events/:id", h.DeleteEvent)
	r.GET("/events/:id", h.GetEvent)

	r.Run(":8080")
}

func setupDatabase() *gorm.DB {
	// Настройте подключение к базе данных
}

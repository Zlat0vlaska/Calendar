package repository

import (
	"github.com/Zlat0vlaska/Calendar/internal/domain"
	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *domain.Event) error
	Update(event *domain.Event) error
	Delete(id int64) error
	GetByID(id int64) (*domain.Event, error)
	GetByUID(uid string) (*domain.Event, error)
}

type PostgresEventRepository struct {
	db *gorm.DB
}

func NewPostgresEventRepository(db *gorm.DB) *PostgresEventRepository {
	return &PostgresEventRepository{db: db}
}

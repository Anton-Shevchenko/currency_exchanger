package repo

import (
	"exchange_mailer/internal"
	"exchange_mailer/internal/models"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) internal.RepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(form *models.Subscriber) (*models.Subscriber, error) {
	err := repo.db.Create(&form).Error

	return form, err
}

func (repo *Repo) GetAll() []models.Subscriber {
	var subscribes []models.Subscriber
	repo.db.Find(&subscribes)

	return subscribes
}

func (repo *Repo) Exists(email string) bool {
	var sub models.Subscriber

	r := repo.db.
		Where("email = ?", email).
		Limit(1).
		Find(&sub)

	return r.RowsAffected > 0
}

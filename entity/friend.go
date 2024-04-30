package entity

import (
	"time"

	"github.com/greetinc/greet-auth-srv/entity"
	"gorm.io/gorm"
)

type Friend struct {
	ID            string         `gorm:"primary_key" json:"id"`
	StatusAccount bool           `gorm:"status_account" json:"status_account"`
	LikeID        string         `gorm:"type:varchar(36);index" json:"like_id"`
	UserID        string         `gorm:"type:varchar(36);index" json:"user_id"`
	User          entity.User    `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt     time.Time      `json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

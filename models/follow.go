package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowerID uint `gorm:"not null" json:"follower_id"`
	FolloweeID uint `gorm:"not null" json:"followee_id"`
}

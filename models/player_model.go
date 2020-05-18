package models

import (
	"time"
)

// Player :
type Player struct {
	ID             *string    `json:"id" gorm:"column:id" structs:"id"`
	UserName       *string    `json:"userName" gorm:"column:user_name" structs:"user_name"`
	PassHash       *string    `json:"-" gorm:"column:pass_hash" structs:"-"`
	EmailAddress   *string    `json:"emailAddress" gorm:"column:email" structs:"email"`
	DisplayName    *string    `json:"displayName" gorm:"column:display_name" structs:"display_name"`
	AvatarURL      *string    `json:"avatarUrl" gorm:"column:avatar_url" structs:"avatar_url"`
	AvatarBlurHash *string    `json:"avatarBlurHash" gorm:"column:avatar_blur_hash" structs:"avatar_blur_hash"`
	CreatedAt      *time.Time `json:"createdAt" gorm:"column:created_at" structs:"created_at"`
	UpdatedAt      *time.Time `json:"updatedAt" gorm:"column:updated_at" structs:"updated_at"`
	DeletedAt      *time.Time `json:"deletedAt" gorm:"column:deleted_at" structs:"deleted_at"`
}

// PlayerAuthentication :
type PlayerAuthentication struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// PlayerRegistration :
type PlayerRegistration struct {
	UserName        string `json:"userName,omitempty" binding:"required,min=3,max=25"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
}

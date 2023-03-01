package middleware

import "gorm.io/gorm"

type Middleware struct {
	db *gorm.DB
}

func New() *Middleware {
	return &Middleware{}
}

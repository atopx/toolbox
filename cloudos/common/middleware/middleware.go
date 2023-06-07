package middleware

import "gorm.io/gorm"

type Middleware struct {
	db *gorm.DB
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) SetDatabase(db *gorm.DB) {
	m.db = db
}

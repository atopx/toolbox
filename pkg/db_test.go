package pkg

import (
	"testing"
)

func TestNewDbClient(t *testing.T) {
	db, err := NewDbClient("test.db", nil)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate()
}

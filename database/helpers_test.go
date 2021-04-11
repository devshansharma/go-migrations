package database_test

import (
	"testing"
	"saketsharma0805/migrator/database"
)


func TestGetStructFields (t *testing.T) {
	database.GetStructFields(&User{})
	t.Errorf("unknown")
}
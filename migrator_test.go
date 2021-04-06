package migrator_test

import (
	"saketsharma0805/migrator"
	"testing"
)



type inMemoryDB struct {}

func (db *inMemoryDB) Query () {}

// User for verifying user table migration
type User struct {
	migrator.Model
	Name string
	IsActive bool
}

func (u *User) TableName () string {
	return "auth_user"
}


func TestAddTable (t *testing.T) {

}

func TestAddFields (t *testing.T) {
	
}


// TestMigrator for validating Model Migration
func TestMigrator (t *testing.T) {
	t.Run("check if common fields are migrating", func(t *testing.T){
		db := &inMemoryDB{}
		migrate := migrator.NewMigrate(db)

		migrate.AutoMigrate(&User{})


	})
}
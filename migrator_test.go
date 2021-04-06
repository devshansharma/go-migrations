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


func TestParseToColumn (t *testing.T) {
	testCases := []struct{
		Key string
		Expected string
	}{
		{Key : "ID", Expected: "id"},
		{Key : "UUID", Expected: "uuid"},
		{Key : "FKUser", Expected: "fk_user"},
		{Key : "UserID", Expected: "user_id"},
		{Key : "IsActive", Expected: "is_active"},
		{Key : "CreatedAt", Expected: "created_at"},
	}

	for _, tc := range testCases {
		t.Run("case", func(t *testing.T){
			ret := migrator.ParseToColumn(tc.Key)
			if ret != tc.Expected {
				t.Errorf("Expected %q, Got %q", tc.Expected, ret)
			}
		})
	}
}


func TestMigrator (t *testing.T) {
	db := &inMemoryDB{}
	m := migrator.NewMigrate(db)

	user := &User{}

	table, err := m.AddTable(user)
	if err != nil {
		t.Error(err)
	}
	t.Run("test table creation", func (t *testing.T){
		if table.GetTableName() != user.TableName() {
			t.Errorf("[Table Name] Expected: %q, Got: %q\n", user.TableName(), table.GetTableName())
		}	

		if table.GetModelName() != "User" {
			t.Errorf("[Model Name] Expected: %q, Got: %q\n", "User", table.GetModelName())
		}

		fields := table.GetFields()
		if len(fields) != 0 {
			t.Errorf("[Fields] Expected: %d, Got: %d\n", 0, len(fields))
		}
	})


	// table.AddFields(user)
}



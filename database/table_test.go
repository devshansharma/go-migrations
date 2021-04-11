package database_test

import (
	"saketsharma0805/migrator/database"
	"testing"
	"strings"
	"time"
)

type User struct {
	ID        uint64
	Name      string `migrator:"varchar(30),notNull"`
	Email     string `migrator:"unique,index:,notNull"`
	IsActive  bool   `migrator:"default(0)"`
	CompanyID uint64 `migrator:"foreignKey(company,id)"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt time.Time `migrator:"index:idx_users_deleted_at"`
}

func (u *User) TableName() string {
	return "auth_user"
}

func TestTable(t *testing.T) {
	testCases := []struct {
		model interface{}
		expected string
	}{
		{model:&User{},expected:"expected"},
	}

	for _, tc := range testCases {
		t.Run("case", func(t *testing.T) {
			tbl := database.NewTable(tc.model)

			got := tbl.String()
			got = strings.ReplaceAll(got, "\n", "")
			got = strings.ReplaceAll(got, "\t", "")

			if got != tc.expected {
				t.Errorf("Expected: %q, Got: %q",tc.expected,got)
			}
			
		})
	}
}

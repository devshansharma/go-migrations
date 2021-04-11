package database_test

import (
	"saketsharma0805/migrator/database"
	"strings"
	"testing"
)

func TestForeignKey(t *testing.T) {

	testCases := []struct {
		tableName  string
		col        string
		fTableName string
		fCol       string
		onUpdate   database.ConstraintType
		onDelete   database.ConstraintType
		expected   string
	}{
		{tableName: "users", col: "company_id", fTableName: "company", fCol: "id",
			onUpdate: database.ConstraintRestrict,
			onDelete: database.ConstraintCascade,
			expected: `constraint fk_company_id foreign key (company_id) references company(id) on update restrict on delete cascade`,
		},
	}

	for _, tc := range testCases {
		t.Run("case", func(t *testing.T) {
			fk := database.NewForeignKey(tc.col, tc.fTableName, tc.fCol)
			fk.SetOnUpdate(tc.onUpdate)
			fk.SetOnDelete(tc.onDelete)

			got := strings.ToLower(fk.String())
			if got != tc.expected {
				t.Errorf("Expected %q, Got %q", tc.expected, got)
			}

		})
	}
}

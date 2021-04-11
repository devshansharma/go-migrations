package database

import (
	"fmt"
)

type ConstraintType string

func (c ConstraintType) String() string {
	return string(c)
}

var (
	ConstraintRestrict ConstraintType = "RESTRICT"
	ConstraintNoAction ConstraintType = "NO ACTION"
	ConstraintCascade  ConstraintType = "CASCADE"
	ConstraintSetNull  ConstraintType = "SET NULL"
)

type ForeignKey struct {
	tableName        string
	keyName          string
	keyColumn        string
	keyForeignTable  string
	keyForeignColumn string
	onUpdate         ConstraintType
	onDelete         ConstraintType
}

// NewForeignKey will take current table name, current column name,
// foreign table name and foreign column as params
// and return ForeignKey object
func NewForeignKey(col, fTable, fColumn string) *ForeignKey {
	keyName := fmt.Sprintf("fk_%s_%s", fTable, fColumn)
	return &ForeignKey{
		keyName:          keyName,
		// tableName:        tableName,
		keyColumn:        col,
		keyForeignTable:  fTable,
		keyForeignColumn: fColumn,
		onDelete:         ConstraintRestrict,
		onUpdate:         ConstraintRestrict,
	}
}

func (f *ForeignKey) SetOnUpdate(c ConstraintType) {
	f.onUpdate = c
}

func (f *ForeignKey) SetOnDelete(c ConstraintType) {
	f.onDelete = c
}

// String will return string representation of foreign key to be used in SQL query
func (f *ForeignKey) String() string {
	return fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s) ON UPDATE %s ON DELETE %s",
		f.keyName,
		f.keyColumn,
		f.keyForeignTable,
		f.keyForeignColumn,
		f.onUpdate,
		f.onDelete,
	)
}

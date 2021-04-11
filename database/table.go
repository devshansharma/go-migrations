package database

import (
	// "errors"
	"fmt"
	"reflect"
	// "regexp"
	"strings"
)

// Table for creating a table
type Table struct {
	migrationName string
	tableName     string
	model         interface{}
	columns       []*Column
	primaryKey    string
	foreignKey    []*ForeignKey
}

func NewTable(model interface{}) *Table {
	t := &Table{
		model: model,
	}

	t.SetTableName().
		AddColumns()

	t.migrationName = fmt.Sprintf("create_%s_table", t.tableName)
	return t
}

func (t *Table) GetTableName() string {
	return t.tableName
}

func (t *Table) SetTableName() *Table {
	var tableName string

	ft := reflect.TypeOf(t.model)
	if ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
	}

	if ft.Kind() == reflect.Struct {
		tableName = ft.Name()
		tMethod, ok := t.model.(interface{ TableName() string })
		if ok {
			tableName = tMethod.TableName()
		}
	}

	t.tableName = strings.ToLower(tableName)
	return t
}

func (t *Table) AddColumns() *Table {
	cols := GetStructFields(t.model, t.tableName)
	t.columns = append(t.columns, cols...)
	return t
}

func (t *Table) String() string {
	columnSlice := make([]string, 0)

	for _, c := range t.columns {
		columnSlice = append(columnSlice, c.String())
	}

	colString := strings.Join(columnSlice, ", ")
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s) 
		ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 DEFAULT COLLATE=utf8mb4_unicode_ci`,
		t.tableName,
		colString,
	)
	return query
}

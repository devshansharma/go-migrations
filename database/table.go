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
		AddColumns().
		SetPrimaryKey().
		AddForeignKeys()

	t.migrationName = fmt.Sprintf("create_%s_table", t.tableName)
	return t
}

func (t *Table) GetTableName() string {
	return t.tableName
}

func (t *Table) GetPrimaryKeyString() string {
	return fmt.Sprintf("PRIMARY KEY (%s)", t.primaryKey)
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
	cols := GetStructFields(t.model)
	t.columns = append(t.columns, cols...)
	return t
}

func (t *Table) SetPrimaryKey() *Table {
	// t.primaryKey = name
	return t
}

func (t *Table) AddForeignKeys() *Table {
	// f := NewForeignKey(t.tableName)
	// t.foreignKey = append(t.foreignKey, f)
	return t
}

func (t *Table) String() string {
	columnSlice := make([]string, 0)

	for _, c := range t.columns {
		columnSlice = append(columnSlice, c.String())
	}
	columnSlice = append(columnSlice, t.GetPrimaryKeyString())
	for _, f := range t.foreignKey {
		columnSlice = append(columnSlice, f.String())
	}

	colString := strings.Join(columnSlice, ", ")
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s) 
		ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 DEFAULT COLLATE=utf8mb4_unicode_ci`,
		t.tableName,
		colString,
	)
	return query
}

// const (
// 	foreignTag string = "foreignKey\\(([a-zA-Z_-]+),([a-zA-Z_-]+)\\)"
// 	DeleteTag  string = "onDelete\\(([a-zA-Z]+)\\)"
// 	varcharTag string = "varchar\\(([0-9]+)\\)"

// 	migrationTableName string = "migration"
// 	createTableStr     string = "CREATE TABLE IF NOT EXISTS %s (%s) "
// )

// var (
// 	ErrNotAStruct error = errors.New("Model is not a struct")

// 	regFKTag      = regexp.MustCompile(foreignTag)
// 	regOnDelTag   = regexp.MustCompile(DeleteTag)
// 	regVarcharTag = regexp.MustCompile(varcharTag)
// )

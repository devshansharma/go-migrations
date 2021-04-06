package migrator

import (
	"reflect"
	"time"
	"fmt"
	"errors"
	"strings"
)

/*
	migrator use case

	// migrator.Model definition
	type Model struct {
		ID        uint64 		`migrator:"primaryKey"`
		CreatedAt time.time
		UpdatedAt time.Time
		DeletedAt time.Time 	`migrator:"index"`
	}

	// for user model
	type User struct {
		migrator.Model
		Name         string
		Email        *string
		Age          uint8
		Birthday     *time.Time
		MemberNumber sql.NullString
		ActivatedAt  sql.NullTime
	}

	type Post struct {
		migrator.Model
		UserID uint64 			`migrator:"foreignKey"`
		Title string 			`migrator:"varchar(120)"`
		Content sql.NullString 	`migrator:"varchar(4000)"`
	}

	type Tag struct {
		migrator.Model
		Title string
		Slug string
	}

	type PostTag struct {
		TagID uint64
		PostID uint64
	}

	migrator.AutoMigrate(&User{})

	Migrator will create a database table with the name
	migrations, to record the changes done and queries executed.
	It will be helpful in debugging.

	A simple struct will be created, that can be inherited by
	models for the purpose of some common fields:
		CreatedAt, ModifiedAt, DeletedAt, ID

	Separate naming convention:
		// more suitable, can be used by other functions
		func (u *User) TableName () string {
			return "auth_user"
		}


	for every object, it will first check if the table is created,
	if not, create table

	if table is already created, look for changes in column definition
	if there are changes, do changes

	initial phase, table names will not be mutable, it will work only
	for mysql.

	foreign keys will be represented by '{columnName}ID' syntax,
	set null and delete will effect as per tags

*/

var (
	ErrNotAStruct error = errors.New("Model is not a struct")
)

type Idb interface {
	Query()
}

type field struct {
	name string
	column string
	attribute string
}

func (f *field) String () {
	fmt.Printf("[Field] Name: %s, Column: %s\n", f.name, f.column)
}

type table struct {
	tableName string
	model interface{}
	modelName string
	fields []*field
}

func (t *table) String () {
	fmt.Printf("Table Name: %q\n", t.tableName)
	fmt.Printf("Model Name: %q\n", t.modelName)
	for i := 0; i < len(t.fields); i++ {
		t.fields[i].String()
	}	
}


func parseToColumn (fieldName string) string {
	newName := ""

	for i, v := range fieldName {
		// uppercase 
		if v >= 65 && v <= 90 {

			// for i > 0, if the last character was in lower case, put _ 
			// PostID
			if i > 0 && (fieldName[i-1] >= 97 && fieldName[i-1] <= 123) {
				newName += fmt.Sprintf("_%c", v+32)	
				continue
			}

			// if next character in lower case, put _
			// FKPost
			if i > 0 && len(fieldName) > i+1 && fieldName[i+1] >= 97 && fieldName[i+1] <= 123 {
				newName += fmt.Sprintf("_%c", v+32)	
				continue
			}

			// first character
			newName += fmt.Sprintf("%c", v+32)	
			
		} else {
			newName += fmt.Sprintf("%c", v)
		}
	}

	return newName
}

func (t *table) addFields (iface interface{}) {
	// iface type
	ift := reflect.TypeOf(iface)
	if ift.Kind() == reflect.Ptr {
		ift = ift.Elem()
	}

	// iface value, 
	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}

	// loop through all fields
	for i :=0; i < ift.NumField(); i++ {
		ft := ift.Field(i)

		// if embedded struct
		if ft.Anonymous {
			t.addFields(ifv.Field(i).Interface())
		} else {
			t.fields = append(t.fields, &field{name:ft.Name, column:parseToColumn(ft.Name)})
		}

	}

	
}

type migrate struct {
	db Idb
	table []*table
}

func NewMigrate (db Idb) *migrate {
	return &migrate{
		db : db,
	}
}


func (m *migrate) AutoMigrate (model interface{}) {
	table, err := m.addTable(model)
	if err != nil {
		panic(err.Error())
	}

	table.addFields(model)
	table.String()
}

func (m *migrate) addTable (model interface{}) (*table, error) {
	var tableName string

	ft := reflect.TypeOf(model)
	if ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
	}

	if ft.Kind() == reflect.Struct {
		tableName = strings.ToLower(ft.Name())

		// check if tableName method is defined
		tMethod, ok := model.(interface{TableName()string})
		if ok {
			tableName = tMethod.TableName()
		}

		t := &table{
			tableName : tableName,
			model : model,
			modelName: ft.Name(),
			fields : make([]*field, 0),
		}
		m.table = append(m.table, t)

		return t, nil
	}

	return nil, ErrNotAStruct	
}



// Model for common fields
type Model struct {
	ID         uint64 `migrator:"primaryKey"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time `migrator:"index"`
}

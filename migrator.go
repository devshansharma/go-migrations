package migrator

// import (
// 	"reflect"
// 	"sync"
// 	"fmt"
// 	"strings"
// 	"saketsharma0805/migrator/database"
// )

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


// func NewMigrate (db database.Idb) *migrate {
// 	return &migrate{
// 		db : db,
// 		table : make([]*database.Table,0),
// 		migrateTableExists: false,
// 		mu : sync.RWMutex{},
// 	}
// }


// type migrate struct {
// 	db database.Idb
// 	migrateTableExists bool
// 	table []*database.Table
// 	mu sync.RWMutex
// }

// func (m *migrate) createMigrationTable () {
// 	m.migrateTableExists = true
// 	sql := CreateTableStr + DefaultEngine + DefaultCharset + DefaultCollation
// 	columnStr := `id int(64) PRIMARY KEY AUTO_INCREMENT,
// 		migration_name varchar(190) not null,
// 		table_name varchar(60) not null,
// 		created_at datetime default current_timestamp
// 	`
// 	sql = fmt.Sprintf(sql, migrationTableName, columnStr)
// 	m.db.Exec(sql)
// }
 
// func (m *migrate) String () {
// 	for _, t := range m.table {
// 		t.String()
// 	}
// }

// func (m *migrate) AutoMigrate (model interface{}) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	if !m.migrateTableExists {
// 		m.createMigrationTable()
// 	}

// 	table, err := m.AddTable(model)
// 	if err != nil {
// 		panic(err.Error())
// 	}
	
// 	table.AddFields(model)
// 	table.ParseTags()
// 	table.CreateQuery()
// 	m.db.Exec(table.createdQuery)
// }

// func (m *migrate) AddTable (model interface{}) (*table, error) {
// 	var tableName string

// 	ft := reflect.TypeOf(model)
// 	if ft.Kind() == reflect.Ptr {
// 		ft = ft.Elem()
// 	}

// 	if ft.Kind() == reflect.Struct {
// 		tableName = strings.ToLower(ft.Name())

// 		// check if tableName method is defined
// 		tMethod, ok := model.(interface{TableName()string})
// 		if ok {
// 			tableName = tMethod.TableName()
// 		}

// 		t := &table{
// 			tableName : tableName,
// 			model : model,
// 			modelName: ft.Name(),
// 			fields : make([]*field, 0),
// 		}
// 		m.table = append(m.table, t)

// 		return t, nil
// 	}

// 	return nil, ErrNotAStruct	
// }


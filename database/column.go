package database

import (
	"fmt"
	"reflect"
	"strings"
	"regexp"
)

var (
	DeleteTag  string = "onDelete\\(([a-zA-Z]+)\\)"
	varcharTag string = "varchar\\(([0-9]+)\\)"	
	foreignTag string = "foreignKey\\(([a-zA-Z_-]+),([a-zA-Z_-]+)\\)"
	indexTag string = "index:([a-zA-Z_]+)"

	regFKTag      = regexp.MustCompile(foreignTag)	
	regOnDelTag   = regexp.MustCompile(DeleteTag)
	regVarcharTag = regexp.MustCompile(varcharTag)
	regIndexTag = regexp.MustCompile(indexTag)
)


type Column struct {
	tableName string
	name string
	tags string
	colType string
	attributes string

	isPrimary bool

	// foreign key
	isForeign bool
	fTable string
	fCol string

	varcharLength string
	index string
	unique bool
}

func NewColumn (name string, tableName string) *Column {
	return &Column{
		name : ParseColumnNames(name),
		tableName : tableName,
		varcharLength:"255",
	}
}


func (c *Column) SetType (tags string, t reflect.Type) {
	colType := t.String()
	c.tags = tags

	c.parseTags(tags)

	switch colType {
	case "uint64":
		c.attributes = "BIGINT"
		if c.name == "id" {
			c.isPrimary = true
			c.attributes += " AUTO_INCREMENT"
		}
	case "string":
		vlen := c.varcharLength
		c.attributes = "VARCHAR("+vlen+")"
	case "bool":
		c.attributes = "TINYINT(1)"
	case "time.Time":
		c.attributes = "DATETIME"
		if c.name == "created_at" {
			c.attributes += " DEFAULT CURRENT_TIMESTAMP"
		} else if c.name == "updated_at" {
			c.attributes += " DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"
		} else if c.name == "deleted_at" {
			c.attributes += " DEFAULT null"
		}
	}

	c.colType = colType
}

func (c *Column) parseTags (tags string) {
	if strings.Contains(tags, "foreignKey") {
		fks := regFKTag.FindStringSubmatch(tags)
		c.isForeign = true
		c.fTable = fks[1]
		c.fCol = fks[2]
	}

	if strings.Contains(tags, "varchar") {
		varChar := regVarcharTag.FindStringSubmatch(tags)
		c.varcharLength = varChar[1]
	}

	if strings.Contains(tags, "index") {
		index := regIndexTag.FindStringSubmatch(tags)
		idx_name := "idx_" + c.name
		if len(index) > 1 {
			idx_name = index[1]
		}
		c.index = idx_name
	}

	if strings.Contains(tags, "unique") {
		c.unique = true
	}	
}	

func (c *Column) String () string {
	return fmt.Sprintf("%s %s", c.name, c.attributes)
}
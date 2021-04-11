package database

// import (
// 	"reflect"
// 	"fmt"
// 	"strings"
// )


// type field struct {
// 	name string
// 	column string
// 	fieldType reflect.Type
// 	tag string

// 	isPrimary bool
// 	isUnique bool

// 	isAutoIncrement bool
// 	isForeign bool
// 	foreignKeyTable string
// 	foreignKeyField string	
// 	isIndex bool
// 	onDelete string
// 	length string
// }

// func (f *field) createColumnStr () string {
// 	attributes := ""

// 	if f.isPrimary && f.fieldType.String() == "uint64" {
// 		attributes = "BIGINT AUTO_INCREMENT"
// 	}

// 	if f.fieldType.String() == "time.Time" {
// 		if f.column == "created_at"{
// 			attributes = "DATETIME DEFAULT CURRENT_TIMESTAMP"	
// 		} else if f.column == "modified_at" {
// 			attributes = "DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"	
// 		} else if f.column == "deleted_at" {
// 			attributes = "DATETIME NULL DEFAULT NULL"	
// 		}
// 	}

// 	col := fmt.Sprintf("%s %s", f.column, attributes)
// 	return col
// }

// func (f *field) IsPrimary () bool {
// 	return strings.Contains(f.tag, "primaryKey")
// }

// func (f *field) IsUnique () bool {
// 	return strings.Contains(f.tag, "unique")
// }


// func (f *field) Parse () {
// 	if f.tag == "" {
// 		return
// 	}
// 	tag := f.tag

// 	f.isPrimary = strings.Contains(tag, "primaryKey")
// 	f.isUnique = strings.Contains(tag, "unique")
// 	f.isForeign = strings.Contains(tag, "foreignKey")
// 	f.isIndex = strings.Contains(tag, "index")

// 	// foreignKey regexp
// 	if f.isForeign {
// 		ans := regFKTag.FindStringSubmatch(tag)
// 		if len(ans) == 3{
// 			f.foreignKeyTable, f.foreignKeyField = ans[1], ans[2]	
// 		} else {
// 			panic("foreign key table and column not provided")
// 		}

// 		// remove foreign key from tag string
// 		tag = regFKTag.ReplaceAllString(tag,"")
// 	}

// 	// delete regexp
// 	onDel := regOnDelTag.FindStringSubmatch(tag)
// 	if len(onDel) == 2 {
// 		f.onDelete = onDel[1]
// 		tag = regOnDelTag.ReplaceAllString(tag,"")
// 	}

// 	// varchar regexp
// 	varChar := regVarcharTag.FindStringSubmatch(tag)
// 	if len(varChar) == 2 {
// 		f.length = varChar[1]
// 		tag = regVarcharTag.ReplaceAllString(tag,"")
// 	}

// 	// remove all used tags, and see if there are unknown tags
// 	var tagsToRemove = []string{"primaryKey", "unique", "index", ","}
// 	for _, v := range tagsToRemove {
// 		tag = strings.ReplaceAll(tag, v, "")
// 	}

// 	if tag != "" {
// 		panic("unknown tag found: " + tag)
// 	}	
// }


// func (f *field) GetName () string {
// 	return f.name
// }

// func (f *field) GetColumnName () string {
// 	return f.column
// }

// func (f *field) GetFieldType () reflect.Type {
// 	return f.fieldType
// }

// func (f *field) GetTag () string {
// 	return f.tag
// }


// func (f *field) String () {
// 	sign := ""
// 	if f.isPrimary {
// 		sign = "PK"
// 	} else if f.isForeign {
// 		sign = "FK"
// 	} else if f.isUnique && f.isIndex {
// 		sign = "UI"
// 	} else if f.isUnique {
// 		sign = "UN"
// 	} else if f.isIndex {
// 		sign = "IX"
// 	}

// 	fmt.Printf("[%2s] Name: %-12s Column: %-15s Type: %-12s Tag: %s\n", 
// 		sign,
// 		f.GetName(), 
// 		f.GetColumnName(), 
// 		f.GetFieldType().String(),
// 		f.GetTag(),
// 	)
// }



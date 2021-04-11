package database

import (
	"reflect"
	"fmt"
)


func ParseColumnNames (fieldName string) string {
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


func GetStructFields (iface interface{}) []*Column {
	columnSlice := make([]*Column,0)

	ift := reflect.TypeOf(iface)
	if ift.Kind() == reflect.Ptr {
		ift = ift.Elem()
	}

	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}

	for i := 0; i < ift.NumField(); i++ {
		ft := ift.Field(i)

		// if embedded struct
		if ft.Anonymous {
			cols := GetStructFields(ifv.Field(i).Interface())
			columnSlice = append(columnSlice, cols...)
		} else {
			tag, _ := ft.Tag.Lookup("migrator")
			col := NewColumn(ft.Name)
			col.SetType(tag, ft.Type)
			columnSlice = append(columnSlice, col)
		}

	}	

	return columnSlice
}
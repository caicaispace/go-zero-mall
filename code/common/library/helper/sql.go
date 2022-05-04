package helper

import "strings"

func SqlFieldWithTableName(fields, tableName, filter string) string {
	strSlice := strings.Split(fields, ",")
	newSlice := make([]string, 0)
	for _, v := range strSlice {
		if strings.Trim(v, "`") == filter {
			continue
		}
		newSlice = append(newSlice, tableName+"."+v)
	}
	return strings.Join(newSlice, ",")
}

func sliceRemoveByIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

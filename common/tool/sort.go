package tool

import (
	"fmt"
	"sort"
	"strings"
)

// MapSortKeyAZString
//
//	@Description: Map key按照字段A-Z排序
//	@param data
//	@return string
func MapSortKeyAZString(data map[string]string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sBuild strings.Builder
	for _, v := range keys {
		value, _ := data[v]
		if sBuild.Len() > 0 {
			sBuild.WriteString("&")
		}
		sBuild.WriteString(fmt.Sprintf("%v=%v", v, value))
	}
	return sBuild.String()
}

func MapSortKeyAZValueString(data map[string]string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sBuild strings.Builder
	//var sBuild string
	for _, v := range keys {
		value, _ := data[v]

		//sBuild += strings.Trim(value, " ")
		sBuild.WriteString(value)
		//fmt.Println(sBuild.String())
	}
	//fmt.Println(sBuild.String())
	return sBuild.String()
	//return strings.Replace(sBuild, " ", "", -1)
}

// MapKeyLowToUpSortNotZero
//
//	@Description: Map key按照字段A-Z排序, value 为空不参与排序
//	@param data
//	@return string
func MapKeyLowToUpSortNotZero(data map[string]string) string {
	var keys []string
	for k, v := range data {
		if v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sBuild strings.Builder
	for _, v := range keys {
		value, _ := data[v]
		if sBuild.Len() > 0 {
			sBuild.WriteString("&")
		}
		sBuild.WriteString(fmt.Sprintf("%v=%v", v, value))
	}
	return sBuild.String()
}

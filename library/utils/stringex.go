package utils

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"reflect"
	"regexp"
)

//通过正则表达式进行数据过滤,过虑掉指定的内容
func NormFormat(str, filter string) string {
	if filter != "" {
		tmpTxt, err := gregex.ReplaceString(filter, "", str)
		if err != nil {
			glog.Error(err)
		}
		return tmpTxt
	}
	return str
}

//FindAndReplace 查找并替换
func FindAndReplace(docString, findString, replaceString string) string {
	reg := regexp.MustCompile(findString)
	return reg.ReplaceAllString(docString, replaceString)
}

// IsContains 查找值val是否在数组array中存在
func IsContains(val interface{}, array interface{}) bool {
	if array == nil {
		return false
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

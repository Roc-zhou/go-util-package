package structUtil

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// 获取 struct 字段名称
func GetStructKey(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	fmt.Println(t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		fmt.Println(t)
	}
	if t.Kind() != reflect.Struct {
		panic("Check You Struct!")
		return nil
	}
	// 字段数量
	num := t.NumField()
	result := make([]string, 0, num)
	for i := 0; i < num; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

// 获取 struct tag
func GetStructTag(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	num := t.NumField()
	result := make([]string, 0, num)
	for i := 0; i < num; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}

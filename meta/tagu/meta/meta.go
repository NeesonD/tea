package meta

import (
	"fmt"
	"github.com/fatih/structtag"
	"reflect"
	"runtime/debug"
)

const (
	defaultTag = "json"
	tableName  = "tableName"
)

type Meta struct {
	info      map[uintptr]Key
	tableName Key
}

func (m *Meta) Init(i interface{}) {
	m.InitByTag(i, defaultTag)
}

func (m *Meta) InitByTag(p interface{}, tagName string) {
	m.info = make(map[uintptr]Key)

	if reflect.TypeOf(p).Elem().Kind() != reflect.Struct &&
		reflect.TypeOf(p).Kind() != reflect.Ptr {
		panic("NEED A PTR TO STRUCT ")
	}

	name, ok := reflect.TypeOf(p).Elem().FieldByName(tableName)
	if !ok {
		panic("NO tableName FIELD")
	}

	tags, err := structtag.Parse(string(name.Tag))
	if err != nil {
		panic(fmt.Errorf("parse tag err: %+v", err))
	}

	tag, err := tags.Get(tagName)
	if err != nil {
		panic(fmt.Errorf("get tag err: %+v", err))
	}
	m.tableName = Key(tag.Name)
	m.recursive(p, tagName)

}

func (m *Meta) recursive(p interface{}, tagName string) {
	rt := reflect.TypeOf(p).Elem()
	rv := reflect.ValueOf(p).Elem()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		if field.Type.Kind() == reflect.Struct {
			if value.Addr().CanInterface() {
				m.recursive(value.Addr().Interface(), tagName)
			}
		}

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			continue
		}

		tag, err := tags.Get(tagName)
		if err != nil {
			continue
		}

		m.info[value.UnsafeAddr()] = Key(tag.Name)
	}
}

func (m *Meta) TableName() string {
	return string(m.tableName)
}

func (m *Meta) Tag(p interface{}) Key {
	if reflect.TypeOf(p).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("NEED PTR:%s", string(debug.Stack())))
		return ""
	}
	addr := reflect.ValueOf(p).Elem().UnsafeAddr()
	result, ok := m.info[addr]
	if !ok {
		panic(fmt.Sprintf("NO FIELD:%s", string(debug.Stack())))
		return ""
	}
	return result
}

func (m *Meta) TagName(p interface{}) string {
	result := m.Tag(p)
	return result.V()
}

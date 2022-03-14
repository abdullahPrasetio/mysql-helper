package mysql_helper

import (
	"fmt"
	"strings"
)

type Model struct {
	Wheres      []WhereModel
	WhereString string
	WhereValues []interface{}
}

type WhereModel struct {
	Field    string
	Value    interface{}
	Operator string
	Logic    string
}

func Create() *Model {
	return &Model{}
}

func (m *Model) Where(field string, operator string, value interface{}, logic string) *Model {
	model := WhereModel{
		Field:    field,
		Value:    value,
		Operator: operator,
		Logic:    logic,
	}

	m.Wheres = append(m.Wheres, model)
	return m
}

func (m *Model) GenerateWhere() *Model {
	var values []interface{}
	var where []string
	for _, v := range m.Wheres {
		values = append(values, v.Value)
		where = append(where, fmt.Sprintf("%s %s ? %s", v.Field, v.Operator, v.Logic))
	}
	m.WhereString = "WHERE " + strings.Join(where, " ")
	m.WhereValues = values
	return m
}

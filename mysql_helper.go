package mysql_helper

import (
	"fmt"
	"strings"
)

type Model struct {
	Wheres        []WhereModel
	WhereString   string
	WhereValues   []interface{}
	QueryType     string
	SelectColumn  string
	NameTable     string
	LimitQuery    string
	OrderByStatus bool
	OrderByColumn string
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
	m.generateWheres()
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

func (m *Model) generateWheres() *Model {
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

func (m *Model) From(column string) *Model {
	from := ""
	if m.QueryType == "SELECT" {
		from = "FROM "
	}
	m.NameTable = from + column
	return m
}

func (m *Model) Select(column string) *Model {
	m.QueryType = "SELECT"
	m.SelectColumn = column

	return m
}

func (m *Model) Limit(limit string) *Model {
	m.LimitQuery = "Limit " + limit
	return m
}

// func (m *Model) OrderBy(column string, condition string) *Model {
// 	m.OrderByColumn = append(m.OrderByColumn, column)
// 	m.OrderByCondition = append(m.OrderByCondition, condition)
// 	return m
// }

// func (m *Model) generateOrderBy() string {
// 	query := "ORDER BY "
// 	for i,v := range m.OrderByColumn {
// 		query+=
// 	}
// 	return strings.Join("")
// }

func (m *Model) OrderBy(orderBy string) *Model {
	m.OrderByColumn = "ORDER BY " + orderBy
	return m
}

func (m *Model) Generate() string {
	var query string

	if m.QueryType == "SELECT" {
		query = addString(query, m.QueryType)
		query = addString(query, m.SelectColumn)
		query = addString(query, m.NameTable)
		query = addString(query, m.WhereString)
		query = addString(query, m.OrderByColumn)
		query = addString(query, m.LimitQuery)
	}
	return query
}

func addString(query string, column string) string {
	if len(column) > 0 {
		return fmt.Sprintf("%s %s", query, column)
	}
	return query
}

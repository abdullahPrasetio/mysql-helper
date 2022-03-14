package mysql_helper

import (
	"fmt"
	"testing"
)

func TestWhere(t *testing.T) {
	where := Create().Where("column1", "=", "1", "AND").GenerateWhere()
	fmt.Println(where.WhereString)
	if where.WhereString != "WHERE column1 = ? AND" {
		t.Fatal("Result must be 'WHERE column1 = ? AND'")
	}
	fmt.Println("Test Where test Done")
}

func TestSelect(t *testing.T) {
	selectQuery := Create().Select("name,address").From("merchant_data").Generate()
	fmt.Println(selectQuery)
	if selectQuery != "SELECT name,address FROM merchant_data" {
		t.Fatal("Result must be 'SELECT name,address FROM merchant_data'")
	}
	fmt.Println("Test Select Done")
}

func TestSelectWithWhere(t *testing.T) {
	selectQuery := Create().Select("name,address").From("merchant_data").OrderBy("name ASC,address DESC").Where("name", "=", "Waluyo", "And").Where("name", "!=", "Adi", "").Limit("1").Generate()
	fmt.Println(selectQuery)
	if selectQuery != "SELECT name,address FROM merchant_data" {
		t.Fatal("Result must be 'SELECT name,address FROM merchant_data'")
	}
	fmt.Println("Test Select Done")
}

package main

import (
	orm "SimpleGoORM"
	"SimpleGoORM/cmd_test/dialect"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := orm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	dialect, _ := dialect.GetDialect("sqlite3")
	s := engine.NewSession(dialect)
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}

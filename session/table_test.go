package session

import (
	"SimpleGoORM/cmd_test/dialect"
	"SimpleGoORM/log"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	dialect, _ := dialect.GetDialect("sqlite3")
	s := NewSession(db, dialect).Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

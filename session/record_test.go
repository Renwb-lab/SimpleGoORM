package session

import (
	"SimpleGoORM/cmd_test/dialect"
	"SimpleGoORM/log"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Error(err)
		return nil
	}

	dialect, _ := dialect.GetDialect("sqlite3")
	s := NewSession(db, dialect).Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	defer func() {
		_ = s.db.Close()
	}()
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	defer func() {
		_ = s.db.Close()
	}()
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	defer func() {
		_ = s.db.Close()
	}()
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	defer func() {
		_ = s.db.Close()
	}()
	affected, _ := s.Where("Name = ?", "Tom").Update("Age", 30)
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 30 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	defer func() {
		_ = s.db.Close()
	}()
	affected, _ := s.Where("Name = ?", "Tom").Delete()
	count, _ := s.Count()

	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}

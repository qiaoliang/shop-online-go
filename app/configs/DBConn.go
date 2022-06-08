package configs

import (
	"fmt"

	"gorm.io/gorm"
)

type DBConn struct {
	db    *gorm.DB
	Error error
}

func NewConn(db *gorm.DB) *DBConn {
	return &DBConn{db, nil}
}
func (conn *DBConn) Find(dest interface{}, conds ...interface{}) (tx *DBConn) {
	conn.db = conn.db.Find(dest, conds...)
	conn.Error = conn.db.Error
	return conn
}

func (conn *DBConn) Create(value interface{}) (tx *DBConn) {
	conn.db = conn.db.Create(value)
	fmt.Println(conn.db.Error)
	conn.Error = conn.db.Error
	return conn
}

func (conn *DBConn) Where(query interface{}, args ...interface{}) (tx *DBConn) {
	conn.db = conn.db.Where(query, args)
	conn.Error = conn.db.Error
	return conn
}

func (conn *DBConn) First(dest interface{}, conds ...interface{}) (tx *DBConn) {
	conn.db = conn.db.First(dest, conds)
	conn.Error = conn.db.Error
	return conn
}
func (conn *DBConn) Errors() error {
	return conn.Error
}
func (conn *DBConn) Delete(value interface{}, conds ...interface{}) (tx *DBConn) {
	conn.db = conn.db.First(value, conds)
	conn.Error = conn.db.Error
	return conn

}

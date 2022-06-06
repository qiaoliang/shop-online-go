package configs

import "gorm.io/gorm"

type DBConn struct {
	db *gorm.DB
}

func NewConn(db *gorm.DB) *DBConn {
	return &DBConn{db}
}
func (this *DBConn) Find(dest interface{}, conds ...interface{}) (tx *DBConn) {
	this.db = this.db.Find(dest, conds...)
	return this
}

package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	User   string
	Passwd string
	Addr   string
	Port   int
	DBName string
}

var DB *gorm.DB
var err error

func DbMigrate() {
	dsn := getDbURI() + "?multiStatements=true"

	db, _ := sql.Open("mysql", dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://dbscripts",
		"mysql",
		driver,
	)
	if err != nil {
		// **I get error here!!**
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	fmt.Println("migration completed!")
}

func StaticPicURI() string {
	return viper.Get("STATIC_PIC_URI").(string)
}
func getDbURI() string {
	dbcfg := config{
		User:   viper.Get("MYSQL.DB_USERNAME").(string),
		Passwd: viper.Get("MYSQL.DB_PASSWORD").(string),
		Addr:   viper.Get("MYSQL.BASE_URL").(string),
		Port:   viper.Get("MYSQL.DB_PORT").(int),
		DBName: viper.Get("MYSQL.DB_NAME").(string),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbcfg.User, dbcfg.Passwd, dbcfg.Addr, dbcfg.Port, dbcfg.DBName)
	return dsn
}
func InitMysqlDB() {

	dsn := getDbURI() + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}

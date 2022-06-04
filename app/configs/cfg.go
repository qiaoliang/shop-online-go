package configs

import (
	"bookstore/app/utils"
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

type Config struct {
	User      string
	Passwd    string
	Addr      string
	Port      int
	DBName    string
	StaticPic string
}

func GetConfigInstance(cfgfile string) *Config {
	if !utils.IsPathExist(cfgfile) {
		fmt.Println("config file " + cfgfile + " is NOT existed")
		panic(cfgfile + "is NOT existed.")
	}
	viper.SetConfigFile(cfgfile)
	viper.ReadInConfig()
	Cfg = Config{
		User:      viper.Get("MYSQL.DB_USERNAME").(string),
		Passwd:    viper.Get("MYSQL.DB_PASSWORD").(string),
		Addr:      viper.Get("MYSQL.BASE_URL").(string),
		Port:      viper.Get("MYSQL.DB_PORT").(int),
		DBName:    viper.Get("MYSQL.DB_NAME").(string),
		StaticPic: viper.Get("MYSQL.STATIC_PIC_URI").(string),
	}
	return &Cfg
}

var DB *gorm.DB
var err error

func (cfg *Config) DbMigrate() {
	dsn := cfg.getDbURI() + "?multiStatements=true"

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

var Cfg Config

func (cfg *Config) StaticPicPrefix() string {
	return cfg.StaticPic
}
func (cfg *Config) BannerPicPrefix() string {
	return fmt.Sprintf("%v/banners/", cfg.StaticPicPrefix())
}
func (cfg *Config) GoodsPicPrefix() string {
	return fmt.Sprintf("%v/goods/", cfg.StaticPicPrefix())
}
func (cfg *Config) AvatarPicPrefix() string {
	return fmt.Sprintf("%v/avatar/", cfg.StaticPicPrefix())

}

func (cfg *Config) getDbURI() string {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Passwd, cfg.Addr, cfg.Port, cfg.DBName)
	return dsn
}
func (cfg *Config) InitMysqlDB() {

	dsn := cfg.getDbURI() + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}

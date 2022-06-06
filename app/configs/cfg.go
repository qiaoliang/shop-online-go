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
	DBUser          string
	DBPasswd        string
	DBAddr          string
	DBPort          int
	DBName          string
	DBMigrationPath string

	StaticPic  string
	BannerPath string
	GoodsPath  string
	AvatarPath string
}

func GetConfigInstance(cfgfile string) *Config {
	if !utils.IsPathExist(cfgfile) {
		fmt.Println("config file " + cfgfile + " is NOT existed")
		panic(cfgfile + "is NOT existed.")
	}
	viper.SetConfigFile(cfgfile)
	viper.ReadInConfig()
	Cfg = Config{
		DBUser:          viper.Get("MYSQL.DB_USERNAME").(string),
		DBPasswd:        viper.Get("MYSQL.DB_PASSWORD").(string),
		DBAddr:          viper.Get("MYSQL.BASE_URL").(string),
		DBPort:          viper.Get("MYSQL.DB_PORT").(int),
		DBName:          viper.Get("MYSQL.DB_NAME").(string),
		DBMigrationPath: viper.Get("MYSQL.DB_SCRIPTS").(string),

		StaticPic:  viper.Get("RESOURCES.STATIC_PIC_URI").(string),
		GoodsPath:  viper.Get("RESOURCES.GOODS_RELETIVE_PATH").(string),
		BannerPath: viper.Get("RESOURCES.BANNERS_RELETIVE_PATH").(string),
		AvatarPath: viper.Get("RESOURCES.AVARAE_RELETIVE_PATH").(string),
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
		cfg.DBMigrationPath,
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
	return fmt.Sprintf("%s/%s", cfg.StaticPicPrefix(), cfg.BannerPath)
}
func (cfg *Config) GoodsPicPrefix() string {
	return fmt.Sprintf("%s/%s", cfg.StaticPicPrefix(), cfg.GoodsPath)
}
func (cfg *Config) AvatarPicPrefix() string {
	return fmt.Sprintf("%s/%s", cfg.StaticPicPrefix(), cfg.AvatarPath)

}

func (cfg *Config) getDbURI() string {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.DBUser, cfg.DBPasswd, cfg.DBAddr, cfg.DBPort, cfg.DBName)
	return dsn
}
func (cfg *Config) InitMysqlDB() {

	dsn := cfg.getDbURI() + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}

package configs

import (
	"bookstore/app/utils"
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Persistence bool
	cfgDir      string
	DBConn      *gorm.DB
	Host        string
	Port        int

	DBUser          string
	DBPasswd        string
	DBAddr          string
	DBPort          int
	DBName          string
	DBMigrationPath string
	DBMigrateProto  string
	DBMigrateDir    string
	m               *migrate.Migrate

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

		Persistence:    viper.GetBool("PERSISTANCE"),
		Host:           viper.Get("HOST").(string),
		Port:           viper.Get("PORT").(int),
		DBUser:         viper.Get("MYSQL.DB_USERNAME").(string),
		DBPasswd:       viper.Get("MYSQL.DB_PASSWORD").(string),
		DBAddr:         viper.Get("MYSQL.BASE_URL").(string),
		DBPort:         viper.Get("MYSQL.DB_PORT").(int),
		DBName:         viper.Get("MYSQL.DB_NAME").(string),
		DBMigrateProto: viper.Get("MYSQL.DB_MIG_PROTO").(string),
		DBMigrateDir:   viper.Get("MYSQL.DB_MIG_DIR").(string),

		StaticPic:  viper.Get("RESOURCES.STATIC_PIC_URI").(string),
		GoodsPath:  viper.Get("RESOURCES.GOODS_RELETIVE_PATH").(string),
		BannerPath: viper.Get("RESOURCES.BANNERS_RELETIVE_PATH").(string),
		AvatarPath: viper.Get("RESOURCES.AVARAE_RELETIVE_PATH").(string),
	}
	Cfg.getAbsDir(cfgfile)
	return &Cfg
}

var DB *gorm.DB
var err error

func (cfg *Config) getAbsDir(filename string) string {
	fp, _ := filepath.Abs(filename)
	dp, _ := filepath.Split(fp)
	cfg.cfgDir = dp
	return cfg.cfgDir
}
func (cfg *Config) getMigretionPath() string {
	return fmt.Sprintf("%s/%s/%s", cfg.DBMigrateProto, cfg.cfgDir, cfg.DBMigrateDir)
}
func (cfg *Config) prepareMigration() {
	dsn := cfg.getDbURI() + "?multiStatements=true"

	db, _ := sql.Open("mysql", dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		cfg.getMigretionPath(),
		"mysql",
		driver,
	)

	if err != nil {
		// **I get error here!!**
		panic(err)
	}
	cfg.m = m
}
func (cfg *Config) Downgrade() {
	if cfg.m == nil {
		cfg.prepareMigration()
	}

	if err := cfg.m.Down(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf(" downgrade error:\n %v\n", err)
	}

}
func (cfg *Config) Upgrade() {
	if cfg.m == nil {
		cfg.prepareMigration()
	}

	if err := cfg.m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
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
func (cfg *Config) MysqlDBConn() *gorm.DB {
	if cfg.DBConn == nil {
		dsn := cfg.getDbURI() + "?charset=utf8mb4&parseTime=True&loc=Local"

		DB, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Failed to connect database")
		}
		cfg.DBConn = DB
	}
	return cfg.DBConn
}
func (cfg *Config) DBConnection() *DBConn {
	db := cfg.MysqlDBConn()
	return NewConn(db)
}

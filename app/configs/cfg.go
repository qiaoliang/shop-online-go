package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"bookstore/app/migrations"
)

type Config struct {
	cfgDir      string
	dbConn      *gorm.DB
	Host        string
	Port        int

	SQLiteDBFile   string
	SQLiteDBMemory string
	MigrationDir   string

	StaticPic  string
	BannerPath string
	GoodsPath  string
	AvatarPath string
}

func IsPathExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func GetConfigInstance(cfgfile string) *Config {
	if !IsPathExist(cfgfile) {
		fmt.Println("config file " + cfgfile + " is NOT existed")
		panic(cfgfile + "is NOT existed.")
	}
	viper.SetConfigFile(cfgfile)
	viper.ReadInConfig()
	migrationDir := viper.GetString("MIGRATION_DIR")
	if migrationDir == "" {
		migrationDir = "dbscripts" // 默认值
	}
	Cfg = Config{
		Host:           viper.Get("HOST").(string),
		Port:           viper.Get("PORT").(int),
		SQLiteDBFile:   viper.GetString("SQLITE.DB_FILE"),
		SQLiteDBMemory: viper.GetString("SQLITE.DB_MEMORY"),
		MigrationDir:   migrationDir,
		StaticPic:      viper.Get("RESOURCES.STATIC_PIC_URI").(string),
		GoodsPath:      viper.Get("RESOURCES.GOODS_RELETIVE_PATH").(string),
		BannerPath:     viper.Get("RESOURCES.BANNERS_RELETIVE_PATH").(string),
		AvatarPath:     viper.Get("RESOURCES.AVARAE_RELETIVE_PATH").(string),
	}
	Cfg.getAbsDir(cfgfile)
	return &Cfg
}

var err error

func (cfg *Config) getAbsDir(filename string) string {
	fp, _ := filepath.Abs(filename)
	dp, _ := filepath.Split(fp)
	cfg.cfgDir = dp
	return cfg.cfgDir
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

func (cfg *Config) runMigrations() {
	err := migrations.MigrateUp(cfg.SQLiteDBMemory, cfg.MigrationDir)
	if err != nil {
		panic("Database migration failed: " + err.Error())
	}
}

func (cfg *Config) DBConnection() *gorm.DB {
	if cfg.dbConn == nil {
		cfg.dbConn, err = gorm.Open(sqlite.Open(cfg.SQLiteDBMemory), &gorm.Config{})
		if err != nil {
			panic("Failed to connect database: " + err.Error())
		}
		cfg.runMigrations()
	}
	return cfg.dbConn
}
func (cfg *Config) DBDisconnect() {
	if cfg.dbConn != nil {
		cfg.dbConn = nil
	}
}

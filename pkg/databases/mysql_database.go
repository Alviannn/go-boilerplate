package databases

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	DbmateMigrationDir = "./migrations"
)

func NewMySQLDB() (db *gorm.DB, err error) {
	var (
		cfg      = configs.Default()
		logLevel = gormLogger.Warn
	)

	if cfg.Environment != constants.EnvProduction {
		logLevel = gormLogger.Info
	}

	return gorm.Open(
		mysql.Open(cfg.MySQL.ToGormURL()),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(logLevel),
		},
	)
}

func MigrateMySQL() (err error) {
	rawDBUrl := configs.Default().MySQL.ToDbmateURL()
	dbUrl, err := url.Parse(rawDBUrl)
	if err != nil {
		return
	}

	db := dbmate.New(dbUrl)
	db.MigrationsDir = []string{DbmateMigrationDir}
	db.AutoDumpSchema = false
	db.Verbose = false

	err = db.Migrate()
	return
}

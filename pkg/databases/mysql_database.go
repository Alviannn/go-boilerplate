package databases

import (
	"fmt"
	"go-boilerplate/internal/configs"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	GormMySQLURLFmt    = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC"
	DbmateMySQLURLFmt  = "mysql://%s:%s@%s:%s/%s"
	DbmateMigrationDir = "./migrations"
)

func NewMySQLDB() (db *gorm.DB, err error) {
	mysqlConfig := configs.Default().MySQL
	dsn := fmt.Sprintf(GormMySQLURLFmt,
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Name,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
}

func MigrateMySQL() (err error) {
	mysqlConfig := configs.Default().MySQL
	rawDBUrl := fmt.Sprintf(DbmateMySQLURLFmt,
		mysqlConfig.Username,
		url.QueryEscape(mysqlConfig.Password),
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Name,
	)

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

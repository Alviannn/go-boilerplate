package databases

import (
	"fmt"
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_TIMEZONE"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
}

func MigratePosgres() (err error) {
	rawDBUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		url.QueryEscape(os.Getenv("DB_PASS")),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	dbUrl, err := url.Parse(rawDBUrl)
	if err != nil {
		return
	}

	db := dbmate.New(dbUrl)
	db.MigrationsDir = []string{"./migrations"}
	db.AutoDumpSchema = false
	db.Verbose = false

	err = db.Migrate()
	return
}

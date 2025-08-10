package main

import (
	"flag"
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/databases"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
)

func main() {
	var (
		actionFlag = flag.String("action", "", "action")
		nameFlag   = flag.String("name", "", "name")
	)

	flag.Parse()

	switch *actionFlag {
	case "new":
		if *nameFlag == "" {
			panic("name is required")
		}

		db, err := getDB(false)
		if err != nil {
			panic(err)
		}

		err = db.NewMigration(*nameFlag)
		if err != nil {
			panic(err)
		}
	case "up":
		db, err := getDB(true)
		if err != nil {
			panic(err)
		}

		err = db.CreateAndMigrate()
		if err != nil {
			panic(err)
		}
	case "down":
		db, err := getDB(true)
		if err != nil {
			panic(err)
		}

		err = db.Rollback()
		if err != nil {
			panic(err)
		}
	case "status":
		db, err := getDB(true)
		if err != nil {
			panic(err)
		}

		_, err = db.Status(false)
		if err != nil {
			panic(err)
		}
	}
}

func getDB(isNeedConnection bool) (db *dbmate.DB, err error) {
	if isNeedConnection {
		if err = configs.Load(); err != nil {
			return
		}

		var (
			rawMysqlUrl = configs.Default().MySQL.ToDbmateURL()
			mysqlUrl    *url.URL
		)

		mysqlUrl, err = url.Parse(rawMysqlUrl)
		if err != nil {
			return
		}

		db = dbmate.New(mysqlUrl)
	} else {
		db = dbmate.New(nil)
	}

	db.Verbose = true
	db.MigrationsDir = []string{databases.DbmateMigrationDir}
	return
}

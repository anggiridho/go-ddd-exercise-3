package main

import (
	"errors"
	"os"

	"github.com/rwese/go_example_3/pkg/app"
	"github.com/sirupsen/logrus"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	feedsRepoMySQL "github.com/rwese/go_example_3/pkg/app/feed/repository/mysql"
	feedsRepoSqLite3 "github.com/rwese/go_example_3/pkg/app/feed/repository/sqlite"
)

const (
	dbEngine = "Sqlite3"
	sqliteDb = "./feeds.db"
)

func main() {
	var (
		rf app.FeedRepository
	)
	if dbEngine == "Sqlite3" {

		defer os.Remove(sqliteDb)
		dbConn, err := sql.Open("sqlite3", sqliteDb)
		rf = feedsRepoSqLite3.NewFeedRepository(dbConn)
		if err != nil {
			logrus.Fatal(err)
		}
	} else if dbEngine == "MySQL" {
		mysqldsn := "TODO" // TODO Read from Config
		dbConn, err := sql.Open("mysql", mysqldsn)
		if err != nil {
			logrus.Fatal(err)
		}

		rf = feedsRepoMySQL.NewFeedRepository(dbConn)

	} else {
		panic(errors.New("dbEngine not set or not supported")) // TODO handle errors properly
	}

	storeTestData(rf)
}

func storeTestData(r app.FeedRepository) {

	f := app.Feed{
		Name: "test",
		URL:  "testurl",
	}
	if err := r.CreateFeed(f); err != nil {
		panic(err)
	}

}

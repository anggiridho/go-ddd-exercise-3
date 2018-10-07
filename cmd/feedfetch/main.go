package main

import (
	"errors"

	"github.com/rwese/go_example_3/pkg/app"

	feedsRepoMySQL "github.com/rwese/go_example_3/pkg/app/feed/repository/mysql"
	feedsRepoSqLite3 "github.com/rwese/go_example_3/pkg/app/feed/repository/sqlite"
)

const (
	dbEngine = "Sqlite3"
)

func main() {
	var (
		feeds app.FeedRepository
	)
	if dbEngine == "Sqlite3" {
		feeds = feedsRepoSqLite3.NewFeedRepository()
	} else if dbEngine == "MySQL" {
		feeds = feedsRepoMySQL.NewFeedRepository()

	} else {
		panic(errors.New("nothing found"))
	}
}

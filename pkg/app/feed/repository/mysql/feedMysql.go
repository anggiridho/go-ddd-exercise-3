package mysql

import (
	"database/sql"
	"errors"
	"log"

	// mysql is used as the backend
	_ "github.com/go-sql-driver/mysql"
	"github.com/rwese/go_example_3/pkg/app"
	"github.com/sirupsen/logrus"
)

type feedRepository struct {
	dbConn *sql.DB
}

// NewFeedRepository returns FeedRepo
func NewFeedRepository(dbConn *sql.DB) app.FeedRepository {
	r := &feedRepository{
		dbConn: dbConn,
	}

	return r

}

func (r *feedRepository) CreateFeed(f app.Feed) error {

	tx, err := r.dbConn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into feeds(Name, URL) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(f.Name, f.URL)
	if err != nil {
		log.Fatal(err)
	}
	// insertID, _ := res.LastInsertId()
	tx.Commit()
	return nil
}

func (r *feedRepository) FindByID(id int64) ([]app.Feed, error) {

	rows, err := r.dbConn.Query("SELECT id, name, url FROM feeds WHERE id = $1", id)

	if err != nil {
		logrus.Fatal("Unable to query foo table:", err)
	}
	defer rows.Close()

	var f []app.Feed

	for rows.Next() {
		var feed app.Feed

		if err := rows.Scan(&feed.FeedID, &feed.Name, &feed.URL); err != nil {
			logrus.Error("Unable to scan results:", err)
			continue
		}
		f = append(f, feed)
	}
	return f, errors.New("nothing found")
}
func (r *feedRepository) FindAll() ([]app.Feed, error) {

	rows, err := r.dbConn.Query("SELECT id, name, url FROM feeds")

	if err != nil {
		logrus.Fatal("Unable to query foo table:", err)
	}
	defer rows.Close()

	var f []app.Feed

	for rows.Next() {
		var feed app.Feed

		if err := rows.Scan(&feed.FeedID, &feed.Name, &feed.URL); err != nil {
			logrus.Error("Unable to scan results:", err)
			continue
		}
		f = append(f, feed)
	}
	return f, errors.New("nothing found")
}

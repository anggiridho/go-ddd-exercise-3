package mysql

import (
	"database/sql"
	"errors"
	"log"

	// sqlite is used as the backend
	_ "github.com/mattn/go-sqlite3"
	"github.com/rwese/go_example_3/pkg/app"
)

type feedRepository struct {
	dbConn *sql.DB
}

// NewFeedRepository returns a sqlite repository
func NewFeedRepository(dbConn *sql.DB) (feedRepository, error) {
	r := feedRepository{
		dbConn: dbConn,
	}

	return r, nil

}

func (r *feedRepository) Find(id int) (*app.Feed, error) {

	rows, err := r.dbConn.Query("SELECT id, name, url FROM feeds WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	f := app.Feed{}

	return &f, errors.New("nothing found")
}

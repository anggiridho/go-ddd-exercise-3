package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	// mysql is used as the backend
	_ "github.com/go-sql-driver/mysql"
	"github.com/rwese/go_example_3/pkg/app"
)

type feedRepository struct {
	dbConn *sql.DB
}

// NewFeedRepository returns FeedRepo
func NewFeedRepository(dsn string, dbConn *sql.DB) (feedRepository, error) {
	r := feedRepository{
		dbConn: dbConn,
	}

	return r, nil

}

func (r *feedRepository) Find(id int) (*app.Feed, error) {

	rows, err := r.dbConn.Query("SELECT name FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, id)
		f := &app.Feed{Name: name}
		return f, nil
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &app.Feed{}, errors.New("nothing found")
}

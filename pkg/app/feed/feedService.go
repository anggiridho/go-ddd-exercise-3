package feed

import (
	"errors"

	"github.com/rwese/go_example_3/pkg/app"
)

var ErrInvalidArgument = errors.New("invalid argument")

type service struct {
	feeds app.FeedRepository
}
type Service interface {
	// CreateFeed method
	CreateFeed(Name string, URL string) error
	FindByID(id int64) ([]app.Feed, error)
	FindAll() ([]app.Feed, error)

	// Feeds method
	Feeds() []Feed
}

type Feed struct {
	Name string `json: "Name"`
	URL  string `json: "URL"`
}

func NewService(feeds app.FeedRepository) Service {
	return &service{
		feeds: feeds,
	}
}

func (s *service) CreateFeed(Name string, URL string) error {
	err := s.CreateFeed(Name, URL)
	return err
}
func (s *service) FindByID(id int64) ([]app.Feed, error) {
	f, err := s.FindByID(id)
	return f, err
}
func (s *service) FindAll() ([]app.Feed, error) {
	f, err := s.FindAll()
	return f, err
}

func (s *service) Feeds() []Feed {
	var result []Feed

	return result
}

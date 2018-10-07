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
	CreateFeed(Name string, URL string) (app.FeedID, error)

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

func (s *service) CreateFeed(Name string, URL string) (app.FeedID, error) {
	f := app.CreateFeed(Name, URL)
	return f.FeedID, nil
}

func (s *service) Feeds() []Feed {
	var result []Feed

	return result
}

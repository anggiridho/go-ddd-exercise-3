package feed

import "github.com/rwese/go_example_3/pkg/app"

type service struct {
	feeds app.FeedRepository
}
type Service interface {
	Feeds() []Feed
	CreateFeed(Name string, URL string) (app.FeedID, error)
}

type Feed struct {
	Name string
	URL  string
}

func NewService(feeds app.FeedRepository) Service {
	return &service{
		feeds: feeds,
	}
}

func (s *service) CreateFeed(Name string, URL string) (app.FeedID, error) {
	f := app.NewFeed(Name, URL)
	return f.FeedID, nil
}

func (s *service) Feeds() []Feed {
	var result []Feed

	return result
}

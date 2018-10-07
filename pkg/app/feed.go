package app

type FeedID int

type Feed struct {
	FeedID    FeedID
	Name, URL string
}

type FeedRepository interface {
	Find(id int) (*Feed, error)
}

func CreateFeed(Name string, URL string) *Feed {

	return &Feed{
		Name: Name,
		URL:  URL,
	}
}

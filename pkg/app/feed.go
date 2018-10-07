package app

type Feed struct {
	FeedID    int64
	Name, URL string
}

type FeedRepository interface {
	CreateFeed(Feed) error
	FindByID(id int64) ([]Feed, error)
	FindAll() ([]Feed, error)
}

// func CreateFeed(Name string, URL string) *Feed {

// 	return &Feed{
// 		Name: Name,
// 		URL:  URL,
// 	}
// }

// func CreateFeed(Name string, URL string) *Feed {

// 	return &Feed{
// 		Name: Name,
// 		URL:  URL,
// 	}
// }

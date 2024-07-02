package gofeed

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"
)

type Feed struct {
	*gofeed.Feed
}

func New(url string) (*Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = userAgent

	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting and parsing rss feed")
	}

	return &Feed{Feed: feed}, nil
}

func (f *Feed) GetGUIDs() (guids []string) {
	for _, item := range f.Items {
		guids = append(guids, item.GUID)
	}
	return
}

func (f *Feed) GetAuthorsNameToEmail() map[string]string {
	nameToAuthor := map[string]string{}
	for _, item := range f.Items {
		for _, author := range item.Authors {
			nameToAuthor[author.Name] = author.Email
		}
	}

	return nameToAuthor
}

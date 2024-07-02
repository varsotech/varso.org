package rank

import (
	"time"
)

type ItemRanker interface {
	Rank(*RankedItem) int64
}

type RankedItem struct {
	Rank        int64
	PublishDate time.Time
	FeedRank    float64
}

type Ranker struct {
	rankers     []ItemRanker
	sortedItems []*RankedItem
}

func NewRanker() *Ranker {
	return &Ranker{
		rankers: []ItemRanker{
			NewPublishDateRanker(),
			NewFeedRanker(),
		},
		sortedItems: []*RankedItem{},
	}
}

func (r *Ranker) Rank(publishDate time.Time, feedRank float64) int64 {
	rankedItem := &RankedItem{
		Rank:        1,
		PublishDate: publishDate,
		FeedRank:    feedRank,
	}

	for _, ranker := range r.rankers {
		rankedItem.Rank += ranker.Rank(rankedItem)
	}

	return rankedItem.Rank
}

package rank

type FeedRanker struct {
}

func NewFeedRanker() *FeedRanker {
	return &FeedRanker{}
}

func (r *FeedRanker) Rank(item *RankedItem) int64 {
	return int64(item.FeedRank)
}
